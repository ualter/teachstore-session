// Ualter Otoni Pereira
// ualter.junior@gmail.com
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	logrus "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
	"github.com/ualter/teachstore/session/service"
	"github.com/ualter/teachstore/tracing"
	"github.com/ualter/teachstore/utils"
)

var (
	outputLog      = os.Stdout
	bindAddress    string
	closerJaeger   io.Closer
	jaegerEndpoint = "192.168.1.42"
)

func init() {
	// Logging
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetOutput(outputLog)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	// External Configuration
	externalConfiguration()
	// Distributed Tracing
	startOpenTracingAPI()

	r := mux.NewRouter()
	// Service
	addSessionServiceHandlers(r)
	// Tracing
	r.Use(tracing.Middleware)

	logServer := log.New(outputLog, "", 0)
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	s := http.Server{
		Addr:         (":" + bindAddress),
		Handler:      ch(r),             // set the default handler
		ErrorLog:     logServer,         // set the logger for the server
		ReadTimeout:  10 * time.Second,  // max time to read request from the client
		WriteTimeout: 30 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func() {
		fmt.Printf("Server started at %s \n", bindAddress)
		err := s.ListenAndServe()
		if err != nil {
			if err.Error() != "http: Server closed" {
				logrus.Error("Error attempting start the server %s", err.Error())
				os.Exit(1)
			}
		}
	}()

	// shutdown the server if the O.S. says so
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	fmt.Printf("I am listening...\n")
	<-c
	//sig := <-c
	fmt.Printf("Requesting stop the server! \n")
	//logrus.Info("Requesting stop the server!", sig)

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	s.Shutdown(ctx)
}

func startOpenTracingAPI() {
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
	//defer closerJaeger.Close()
}

func addSessionServiceHandlers(r *mux.Router) {
	sessionSvc := service.NewService()

	getR := r.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/session", sessionSvc.ListAll)
}

func externalConfiguration() {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "develop"
	}

	viper.SetConfigName(environment)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("Error %s", err.Error())
		panic(err.Error())
	}

	viper.SetDefault("port", "9999")
	bindAddress = viper.GetString("port")
	jaegerEndpoint = utils.ReplaceEnvInConfig(viper.Get("opentracing.jaeger.http-sender.url").(string))

	logrus.Debugf("Jaeger Endpoint configured to: %s", utils.ReplaceEnvInConfig(viper.Get("opentracing.jaeger.http-sender.url").(string)))
}
