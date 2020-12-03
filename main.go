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

	logrus "github.com/sirupsen/logrus"
	config "github.com/ualter/teachstore-session/config"
	"github.com/ualter/teachstore-session/session/service"
	"github.com/ualter/teachstore-session/tracing"
)

var (
	outputLog    = os.Stdout
	closerJaeger io.Closer
)

func init() {
	// Logging
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(tracing.NewMyFormatter(func(f *tracing.MyFormatter) error {
		f.PrettyPrint = false
		return nil
	}))

	//logrus.SetReportCaller(true) // Add the Caller (file.go:line)
	logrus.SetOutput(outputLog)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	// Just For Debug Time
	//os.Setenv("IP_DOCKER_HOST", "192.168.1.42")

	// External Configuration
	config.LoadExternalConfiguration()

	// Distributed Tracing
	tracingEnable := config.GetString("opentracing.enable")
	if tracingEnable == "jaeger" {
		StartOpenTracingWithJaeger()
	} else if tracingEnable == "zipkin" {
		StartOpenTracingWithZipkin()
	}

	r := mux.NewRouter()
	// Service
	addSessionServiceHandlers(r)
	// Middleware (generic filter/chain for HTTP Requests)
	r.Use(tracing.Middleware)

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	port := config.GetString("port")
	s := http.Server{
		Addr:         (":" + port),
		Handler:      ch(r),                     // set the default handler
		ErrorLog:     log.New(outputLog, "", 0), // set the logger for the server
		ReadTimeout:  10 * time.Second,          // max time to read request from the client
		WriteTimeout: 30 * time.Second,          // max time to write response to the client
		IdleTimeout:  120 * time.Second,         // max time for connections using TCP Keep-Alive
	}

	go func() {
		fmt.Printf("Server started at %s \n", port)
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

func addSessionServiceHandlers(r *mux.Router) {
	sessionSvc := service.NewService()

	getR := r.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/health{_dummy:/|$}", sessionSvc.Ping)
	getR.HandleFunc("/sessions{_dummy:/|$}", sessionSvc.GetAll)
	getR.HandleFunc("/sessions/{id:[0-9]+}", sessionSvc.GetByID)

}
