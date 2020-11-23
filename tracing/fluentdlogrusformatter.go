package tracing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"unicode/utf8"

	durpb "github.com/golang/protobuf/ptypes/duration"

	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
	logtypepb "google.golang.org/genproto/googleapis/logging/type"
)

// Format configuration of the logrus formatter output.
type Format func(*Formatter) error

// DefaultFormat is Stackdriver.
var DefaultFormat = StackdriverFormat

// Formatter that is called on by logrus.
type Formatter struct {
	// DisableTimestamp allows disabling automatic timestamps in output
	DisableTimestamp bool

	// TimestampFormat sets the format used for marshaling timestamps.
	TimestampFormat func(logrus.Fields, time.Time) error

	// SeverityMap allows for customizing the names for keys of the log level field.
	SeverityMap map[string]string

	// PrettyPrint will indent all json logs
	PrettyPrint bool
}

// NewFormatter with optional options. Defaults to the Stackdriver option.
func NewFormatterFluentD(opts ...Format) *Formatter {
	f := Formatter{}
	if len(opts) == 0 {
		opts = append(opts, DefaultFormat)
	}
	for _, apply := range opts {
		if err := apply(&f); err != nil {
			panic(err)
		}
	}

	return &f
}

// Format the log entry. Implements logrus.Formatter.
/*
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/Sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	prefixFieldClashes(data, entry.HasCaller())

	if !f.DisableTimestamp && f.TimestampFormat != nil {
		// https://cloud.google.com/logging/docs/agent/configuration#timestamp-processing
		if err := f.TimestampFormat(data, entry.Time); err != nil {
			return nil, err
		}
	}

	if entry.Message != "" {
		data["message"] = entry.Message
	}

	if s, ok := f.SeverityMap[entry.Level.String()]; ok {
		data["severity"] = s
	} else {
		data["severity"] = f.SeverityMap["debug"]
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
*/
func prefixFieldClashes(data logrus.Fields, reportCaller bool) {
	if t, ok := data["time"]; ok {
		data["fields.time"] = t
		delete(data, "time")
	}

	if m, ok := data["msg"]; ok {
		data["fields.msg"] = m
		delete(data, "msg")
	}

	if l, ok := data["level"]; ok {
		data["fields.level"] = l
		delete(data, "level")
	}

	if m, ok := data["message"]; ok {
		data["fields.message"] = m
		delete(data, "message")
	}

	if l, ok := data["timestamp"]; ok {
		data["fields.timestamp"] = l
		delete(data, "timestamp")
	}

	if l, ok := data["severity"]; ok {
		data["fields.severity"] = l
		delete(data, "severity")
	}

	if reportCaller {
		if l, ok := data[logrus.FieldKeyFunc]; ok {
			data["fields."+logrus.FieldKeyFunc] = l
		}
		if l, ok := data[logrus.FieldKeyFile]; ok {
			data["fields."+logrus.FieldKeyFile] = l
		}
	}
}

func PrettyPrintFormat(f *Formatter) error {
	f.PrettyPrint = true
	return nil
}

func DisableTimestampFormat(f *Formatter) error {
	f.DisableTimestamp = true
	return nil
}

// StackdriverFormat maps values to be recognized by the Google Cloud Platform.
// https://cloud.google.com/logging/docs/agent/configuration#special-fields
func StackdriverFormat(f *Formatter) error {
	f.SeverityMap = map[string]string{
		"panic":   logtypepb.LogSeverity_CRITICAL.String(),
		"fatal":   logtypepb.LogSeverity_CRITICAL.String(),
		"warning": logtypepb.LogSeverity_WARNING.String(),
		"debug":   logtypepb.LogSeverity_DEBUG.String(),
		"error":   logtypepb.LogSeverity_ERROR.String(),
		"trace":   logtypepb.LogSeverity_DEBUG.String(),
		"info":    logtypepb.LogSeverity_INFO.String(),
	}
	f.TimestampFormat = func(fields logrus.Fields, now time.Time) error {
		// https://cloud.google.com/logging/docs/agent/configuration#timestamp-processing
		ts, err := ptypes.TimestampProto(now)
		if err != nil {
			return err
		}
		fields["timestamp"] = ts
		return nil
	}

	return nil
}

// HTTPRequest contains an http.Request as well as additional
// information about the request and its response.
// https://github.com/googleapis/google-cloud-go/blob/v0.39.0/logging/logging.go#L617
type HTTPRequest struct {
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

func (r HTTPRequest) MarshalJSON() ([]byte, error) {
	if r.Request == nil {
		return nil, nil
	}
	u := *r.Request.URL
	u.Fragment = ""
	// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#HttpRequest
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
	if r.Latency != 0 {
		e.Latency = ptypes.DurationProto(r.Latency)
	}

	return json.Marshal(e)
}

type logEntry struct {
	RequestMethod                  string          `json:"requestMethod,omitempty"`
	RequestURL                     string          `json:"requestUrl,omitempty"`
	RequestSize                    string          `json:"requestSize,omitempty"`
	Status                         int             `json:"status,omitempty"`
	ResponseSize                   string          `json:"responseSize,omitempty"`
	UserAgent                      string          `json:"userAgent,omitempty"`
	RemoteIP                       string          `json:"remoteIp,omitempty"`
	ServerIP                       string          `json:"serverIp,omitempty"`
	Referer                        string          `json:"referer,omitempty"`
	Latency                        *durpb.Duration `json:"latency,omitempty"`
	CacheLookup                    bool            `json:"cacheLookup,omitempty"`
	CacheHit                       bool            `json:"cacheHit,omitempty"`
	CacheValidatedWithOriginServer bool            `json:"cacheValidatedWithOriginServer,omitempty"`
	CacheFillBytes                 string          `json:"cacheFillBytes,omitempty"`
	Protocol                       string          `json:"protocol,omitempty"`
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
