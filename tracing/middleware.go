package tracing

import (
	"context"
	"fmt"
	"net/http"

	//"github.com/uber/jaeger-client-go/config"
	//"github.com/uber/jaeger-lib/metrics/prometheus"

	opentracing "github.com/opentracing/opentracing-go"
	opentracingext "github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
)

// Midleware for Tracing
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		logrus.Debugf("%s", r.RequestURI)
		rw.Header().Add("Middleware", "true")

		//ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		//r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}

// Trace the request of a Service
func TraceRequest(function string, rw http.ResponseWriter, r *http.Request) {

	traceLogger := logrus.WithField("httpRequest", &HTTPRequest{
		Request: r,
		Status:  http.StatusOK,
		//ResponseSize: 31337,
		//Latency:      123 * time.Millisecond,
	})

	/*traceLogger := logrus.WithFields(logrus.Fields{
		"microservice": viper.Get("name"),
		"operation":    function,
	})
	traceLogger.Infof("Called %s", function)*/

	traced := serializeFromTheWire(function, rw, r, traceLogger)
	if !traced {
		// Create a Parent Span (started from here)
		serializeToTheWire(function, rw, r, traceLogger)
	}

	fmt.Printf("%+v\n\n", rw.Header())
}

func serializeToTheWire(function string, rw http.ResponseWriter, r *http.Request, traceLogger *logrus.Entry) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), function)
	defer span.Finish()
	_ = ctx
	injectToHeader(rw, span, traceLogger)
}

func serializeFromTheWire(function string, rw http.ResponseWriter, r *http.Request, traceLogger *logrus.Entry) bool {
	var serverSpan opentracing.Span
	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(rw.Header()))
	if err != nil {
		traceLogger.Debugf(err.Error())
		return false
	}
	serverSpan = opentracing.StartSpan(function, opentracingext.RPCServerOption(wireContext))
	defer serverSpan.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), serverSpan)
	_ = ctx

	injectToHeader(rw, serverSpan, traceLogger)

	return true
}

func injectToHeader(rw http.ResponseWriter, span opentracing.Span, traceLogger *logrus.Entry) {
	opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(rw.Header()))

	traceLogger.Debugf("Tracing Request %+v", span)
}
