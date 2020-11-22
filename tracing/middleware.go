package tracing

import (
	"context"
	"net/http"

	//"github.com/uber/jaeger-client-go/config"
	//"github.com/uber/jaeger-lib/metrics/prometheus"

	opentracing "github.com/opentracing/opentracing-go"
	opentracingext "github.com/opentracing/opentracing-go/ext"
	log "github.com/sirupsen/logrus"
)

// Midleware for Tracing
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		log.Debugf("%s", r.RequestURI)
		rw.Header().Add("Middleware", "true")

		//ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		//r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}

// Trace the request of a Service
func TraceRequest(function string, rw http.ResponseWriter, r *http.Request) {
	traced := serializeFromTheWire(function, rw, r)
	if !traced {
		serializeToTheWire(function, rw, r)
	}
}

func serializeToTheWire(function string, rw http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), function)
	defer span.Finish()
	_ = ctx
	opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(rw.Header()))
	log.Debugf("Starting Tracing Request %+v", span)
}

func serializeFromTheWire(function string, rw http.ResponseWriter, r *http.Request) bool {
	var serverSpan opentracing.Span

	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(rw.Header()))
	if err != nil {
		log.Debugf(err.Error())
		return false
	}

	serverSpan = opentracing.StartSpan(function, opentracingext.RPCServerOption(wireContext))
	defer serverSpan.Finish()

	log.Debugf("Received Tracing Request %+v", serverSpan)

	ctx := opentracing.ContextWithSpan(context.Background(), serverSpan)
	_ = ctx
	return true
}
