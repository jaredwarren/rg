package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/boltdb/bolt"
	rg "github.com/jaredwarren/rg"
	"github.com/jaredwarren/rg/gen/color"
	colorsvr "github.com/jaredwarren/rg/gen/http/color/server"
	homesvr "github.com/jaredwarren/rg/gen/http/home/server"
	schedulesvr "github.com/jaredwarren/rg/gen/http/schedule/server"
	schedule "github.com/jaredwarren/rg/gen/schedule"
	"github.com/jaredwarren/rg/pi"
	goahttp "goa.design/goa/http"
	"goa.design/goa/http/middleware"
)

func main() {
	// Define command line flags, add any other flag required to configure
	// the service.
	var (
		addr = flag.String("listen", ":8080", "HTTP listen `address`")
		dbg  = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger and goa log adapter. Replace logger with your own using
	// your log package of choice. The goa.design/middleware/logging/...
	// packages define log adapters for common log packages.
	var (
		adapter middleware.Logger
		logger  *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[rg] ", log.Ltime)
		adapter = middleware.NewLogger(logger)
	}

	// Initialize service dependencies such as databases.
	var (
		db *bolt.DB
	)
	{
		var err error
		db, err = bolt.Open("alarm.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	}
	// initialize pi
	// Start PI
	rpi := pi.NewPi()

	// Create the structs that implement the services.
	var (
		// homeSvc     home.Service
		scheduleSvc schedule.Service
		colorSvc    color.Service
	)
	{
		var err error
		colorSvc = rg.NewColor(rpi, logger)
		// homeSvc = rg.NewHome(logger)

		scheduleSvc, err = rg.NewSchedule(db, rpi, logger)
		if err != nil {
			logger.Fatalf("error creating schedule service: %s", err)
		}
	}

	// Wrap the services in endpoints that can be invoked from other
	// services potentially running in different processes.
	var (
		// homeEndpoints     *home.Endpoints
		scheduleEndpoints *schedule.Endpoints
		colorEndpoints    *color.Endpoints
	)
	{
		// homeEndpoints = home.NewEndpoints(homeSvc)
		scheduleEndpoints = schedule.NewEndpoints(scheduleSvc)
		colorEndpoints = color.NewEndpoints(colorSvc)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		// homeServer     *homesvr.Server
		scheduleServer *schedulesvr.Server
		colorServer    *colorsvr.Server
	)
	{
		eh := ErrorHandler(logger)
		// homeServer = homesvr.New(homeEndpoints, mux, dec, enc, eh)
		scheduleServer = schedulesvr.New(scheduleEndpoints, mux, dec, enc, eh)
		colorServer = colorsvr.New(colorEndpoints, mux, dec, enc, eh)
	}

	// Configure the mux.
	homesvr.Mount(mux)
	schedulesvr.Mount(mux, scheduleServer)
	colorsvr.Mount(mux, colorServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		if *dbg {
			handler = middleware.Debug(mux, os.Stdout)(handler)
		}
		handler = middleware.Log(adapter)(handler)
		handler = middleware.RequestID()(handler)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the service to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: *addr, Handler: handler}
	go func() {
		for _, m := range scheduleServer.Mounts {
			logger.Printf("file %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
		}
		logger.Printf("listening on %s", *addr)
		errc <- srv.ListenAndServe()
	}()

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Shutdown gracefully with a 30s timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)

	logger.Println("exited")
}

// ErrorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func ErrorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
