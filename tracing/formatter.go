package tracing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"
	"unicode/utf8"

	logrus "github.com/sirupsen/logrus"
)

// MyFormatter ...
type MyFormatter struct {
	// DisableTimestamp allows disabling automatic timestamps in output
	DisableTimestamp bool

	// TimestampFormat sets the format used for marshaling timestamps.
	TimestampFormat func(logrus.Fields, time.Time) error

	// SeverityMap allows for customizing the names for keys of the log level field.
	SeverityMap map[string]string

	// PrettyPrint will indent all json logs
	PrettyPrint bool
}

// MyDefaultFormat ...
func MyDefaultFormat(f *MyFormatter) error {
	f.SeverityMap = map[string]string{
		"panic":   "600",
		"fatal":   "600",
		"warning": "400",
		"debug":   "100",
		"error":   "500",
		"trace":   "100",
		"info":    "200",
	}
	f.TimestampFormat = func(fields logrus.Fields, now time.Time) error {
		fields["@timestamp"] = now.Format("2006-01-02T15:04:05.000-07:00")
		return nil
	}

	return nil
}

// MyFormat ...
type MyFormat func(*MyFormatter) error

// NewMyFormatter ...
func NewMyFormatter(opts ...MyFormat) *MyFormatter {
	f := MyFormatter{}
	if len(opts) == 0 {
		opts = append(opts, MyDefaultFormat)
	}
	for _, apply := range opts {
		if err := apply(&f); err != nil {
			panic(err)
		}
	}
	return &f
}

// Format ...
func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)

	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			// Special treatment for MyHTTPRequest (otherwise it will be Marshall as a JSON Object inside the main root JSON)
			if reflect.TypeOf(v) == reflect.TypeOf(&MyHTTPRequest{}) {
				httpRequest := v.(*MyHTTPRequest)

				u := *httpRequest.Request.URL
				u.Fragment = ""

				data["requestMethod"] = httpRequest.Request.Method
				data["requestUrl"] = fixUTF8(u.String())
				data["status"] = httpRequest.Status
				data["userAgent"] = httpRequest.Request.UserAgent()
				data["serverIP"] = httpRequest.LocalIP
				data["remoteIP"] = httpRequest.RemoteIP
				data["referer"] = httpRequest.Request.Referer()
				if httpRequest.RequestSize > 0 {
					data["requestSize"] = fmt.Sprintf("%d", httpRequest.RequestSize)
				}
				if httpRequest.ResponseSize > 0 {
					data["responseSize"] = fmt.Sprintf("%d", httpRequest.ResponseSize)
				}
			} else {
				data[k] = v
			}
		}
	}

	if !f.DisableTimestamp && f.TimestampFormat != nil {
		if err := f.TimestampFormat(data, entry.Time); err != nil {
			return nil, err
		}
	}

	if entry.Message != "" {
		data["message"] = entry.Message
	}

	if s, ok := f.SeverityMap[entry.Level.String()]; ok {
		data["level"] = strings.ToUpper(entry.Level.String())
		data["level_value"] = s
	} else {
		data["level"] = "DEBUG"
		data["level_value"] = f.SeverityMap["debug"]
	}

	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		if funcVal != "" {
			data[logrus.FieldKeyFunc] = funcVal
		}
		if fileVal != "" {
			data[logrus.FieldKeyFile] = fileVal
		}
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	if f.PrettyPrint {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	return b.Bytes(), nil
}

// MyHTTPRequest struct
type MyHTTPRequest struct {
	// Request is the http.Request passed to the handler.
	Request *http.Request

	// RequestSize is the size of the HTTP request message in bytes, including
	// the request headers and the request body.
	RequestSize int64

	// Status is the response code indicating the status of the response.
	// Examples: 200, 404.
	Status int

	// ResponseSize is the size of the HTTP response message sent back to the client, in bytes,
	// including the response headers and the response body.
	ResponseSize int64

	// Latency is the request processing latency on the server, from the time the request was
	// received until the response was sent.
	Latency time.Duration

	// LocalIP is the IP address (IPv4 or IPv6) of the origin server that the request
	// was sent to.
	LocalIP string

	// RemoteIP is the IP address (IPv4 or IPv6) of the client that issued the
	// HTTP request. Examples: "192.168.1.1", "FE80::0202:B3FF:FE1E:8329".
	RemoteIP string

	// CacheHit reports whether an entity was served from cache (with or without
	// validation).
	CacheHit bool

	// CacheValidatedWithOriginServer reports whether the response was
	// validated with the origin server before being served from cache. This
	// field is only meaningful if CacheHit is true.
	CacheValidatedWithOriginServer bool
}

// MarshalJSON ...
func (r MyHTTPRequest) MarshalJSON() ([]byte, error) {
	if r.Request == nil {
		return nil, nil
	}
	u := *r.Request.URL
	u.Fragment = ""
	e := &logEntry{
		RequestMethod:                  r.Request.Method,
		RequestURL:                     fixUTF8(u.String()),
		Status:                         r.Status,
		UserAgent:                      r.Request.UserAgent(),
		ServerIP:                       r.LocalIP,
		RemoteIP:                       r.RemoteIP,
		Referer:                        r.Request.Referer(),
		CacheHit:                       r.CacheHit,
		CacheValidatedWithOriginServer: r.CacheValidatedWithOriginServer,
	}
	if r.RequestSize > 0 {
		e.RequestSize = fmt.Sprintf("%d", r.RequestSize)
	}
	if r.ResponseSize > 0 {
		e.ResponseSize = fmt.Sprintf("%d", r.ResponseSize)
	}

	return json.Marshal(e)
}

type logEntry struct {
	RequestMethod                  string `json:"requestMethod,omitempty"`
	RequestURL                     string `json:"requestUrl,omitempty"`
	RequestSize                    string `json:"requestSize,omitempty"`
	Status                         int    `json:"status,omitempty"`
	ResponseSize                   string `json:"responseSize,omitempty"`
	UserAgent                      string `json:"userAgent,omitempty"`
	RemoteIP                       string `json:"remoteIp,omitempty"`
	ServerIP                       string `json:"serverIp,omitempty"`
	Referer                        string `json:"referer,omitempty"`
	CacheLookup                    bool   `json:"cacheLookup,omitempty"`
	CacheHit                       bool   `json:"cacheHit,omitempty"`
	CacheValidatedWithOriginServer bool   `json:"cacheValidatedWithOriginServer,omitempty"`
	CacheFillBytes                 string `json:"cacheFillBytes,omitempty"`
	Protocol                       string `json:"protocol,omitempty"`
}

// fixUTF8 is a helper that fixes an invalid UTF-8 string by replacing
// invalid UTF-8 runes with the Unicode replacement character (U+FFFD).
// See Issue https://github.com/googleapis/google-cloud-go/issues/1383.
func fixUTF8(s string) string {
	if utf8.ValidString(s) {
		return s
	}

	// Otherwise time to build the sequence.
	buf := new(bytes.Buffer)
	buf.Grow(len(s))
	for _, r := range s {
		if utf8.ValidRune(r) {
			buf.WriteRune(r)
		} else {
			buf.WriteRune('\uFFFD')
		}
	}
	return buf.String()
}
