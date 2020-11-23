package main

import (
	"log"

	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	zipkingo "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	viper "github.com/spf13/viper"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-client-go/zipkin"
	"github.com/uber/jaeger-lib/metrics"
)

func StartOpenTracingWithZipkin() {
	reporter := zipkinhttp.NewReporter(zipkinEndpoint)
	//defer reporter.Close()

	endpoint, err := zipkingo.NewEndpoint(viper.GetString("name"), myIP+":"+bindAddress)
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}

	nativeTracer, err := zipkingo.NewTracer(reporter, zipkingo.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}

	// use zipkin-go-opentracing to wrap our tracer
	tracer := zipkinot.Wrap(nativeTracer)

	// optionally set as Global OpenTracing tracer instance
	opentracing.SetGlobalTracer(tracer)
}

func StartOpenTracingWithJaeger() {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: jaegerEndpoint,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Zipkin HTTP B3 compatible header propagation (enable or disable it here)
	zipkinHTTPB3CompitablePropagation := true
	if zipkinHTTPB3CompitablePropagation {
		zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
		closerJaeger, err := cfg.InitGlobalTracer(
			viper.GetString("name"),
			jaegercfg.Logger(jLogger),
			jaegercfg.Metrics(jMetricsFactory),
			jaegercfg.Injector(opentracing.HTTPHeaders, zipkinPropagator),
			jaegercfg.Extractor(opentracing.HTTPHeaders, zipkinPropagator),
			jaegercfg.ZipkinSharedRPCSpan(true),
		)
		if err != nil {
			log.Fatalf("%s", err.Error())
			panic(err.Error())
		}
		_ = closerJaeger
	} else {
		closerJaeger, err := cfg.InitGlobalTracer(
			viper.GetString("name"),
			jaegercfg.Logger(jLogger),
			jaegercfg.Metrics(jMetricsFactory),
		)
		if err != nil {
			log.Fatalf("%s", err.Error())
			panic(err.Error())
		}
		_ = closerJaeger
	}
}
