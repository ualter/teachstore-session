package tracing

import (
	"context"
	"fmt"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
	opentracingext "github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
)

// Middleware for Tracing
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		logrus.Debugf("%s", r.RequestURI)
		rw.Header().Add("Middleware", "true")

		//ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		//r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

// TraceRequest trace the request of a Service
func TraceRequest(function string, rw http.ResponseWriter, r *http.Request) {
	// Trace the Request
	traceLogger := logrus.WithField("httpRequest", &MyHTTPRequest{
		Request: r,
		Status:  http.StatusOK,
	})

	// Trace the Request (TraceId, SpanId, ParentId)
	var traced bool
	var span opentracing.Span
	traced, span = serializeFromTheWire(function, rw, r, traceLogger)
	if !traced {
		// Create a Parent Span (started from here)
		span = serializeToTheWire(function, rw, r, traceLogger)
	}

	// Add Tracing to the Log and Message (TraceId, SpanId, ParentId)
	var msg string
	if sc, ok := span.Context().(jaeger.SpanContext); ok {
		traceID := fmt.Sprintf("%016x", sc.TraceID().Low)
		spanID := fmt.Sprintf("%016x", uint64(sc.SpanID()))
		parentID := "0"
		if uint64(sc.ParentID()) > 0 {
			parentID = fmt.Sprintf("%016x", uint64(sc.ParentID()))
		}
		traceLogger = traceLogger.WithFields(logrus.Fields{
			"traceID":  traceID,
			"spanID":   spanID,
			"parentID": parentID,
		})
		msg = fmt.Sprintf("Span reported: %s:%s:%s:%x - %s", traceID, spanID, parentID, sc.Flags(), function)
	}
	if msg == "" {
		msg = fmt.Sprintf("%s", function)
	}
	traceLogger.Infof(msg)
}

func serializeToTheWire(function string, rw http.ResponseWriter, r *http.Request, traceLogger *logrus.Entry) opentracing.Span {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), function)
	defer span.Finish()
	_ = ctx
	injectToHeader(rw, span, traceLogger)
	return span
}

func serializeFromTheWire(function string, rw http.ResponseWriter, r *http.Request, traceLogger *logrus.Entry) (bool, opentracing.Span) {
	var serverSpan opentracing.Span
	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(rw.Header()))
	if err != nil {
		//traceLogger.Debugf(err.Error())
		return false, nil
	}
	serverSpan = opentracing.StartSpan(function, opentracingext.RPCServerOption(wireContext))
	defer serverSpan.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), serverSpan)
	_ = ctx

	injectToHeader(rw, serverSpan, traceLogger)
	return true, serverSpan
}

func injectToHeader(rw http.ResponseWriter, span opentracing.Span, traceLogger *logrus.Entry) {
	opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(rw.Header()))
}
