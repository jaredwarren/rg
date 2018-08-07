// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// schedule HTTP server
//
// Command:
// $ goa gen github.com/jaredwarren/rg/design

package server

import (
	"context"
	"net/http"

	schedule "github.com/jaredwarren/rg/gen/schedule"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Server lists the schedule service endpoint HTTP handlers.
type Server struct {
	Mounts   []*MountPoint
	Home     http.Handler
	List     http.Handler
	Schedule http.Handler
	Remove   http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the schedule service endpoints.
func New(
	e *schedule.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Home", "GET", "/"},
			{"List", "GET", "/"},
			{"Schedule", "POST", "/schedule"},
			{"Remove", "DELETE", "/{id}"},
			{"static/favicon.ico", "GET", "/favicon.ico"},
			{"static/", "GET", "/static/*filename"},
		},
		Home:     NewHomeHandler(e.Home, mux, dec, enc, eh),
		List:     NewListHandler(e.List, mux, dec, enc, eh),
		Schedule: NewScheduleHandler(e.Schedule, mux, dec, enc, eh),
		Remove:   NewRemoveHandler(e.Remove, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "schedule" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Home = m(s.Home)
	s.List = m(s.List)
	s.Schedule = m(s.Schedule)
	s.Remove = m(s.Remove)
}

// Mount configures the mux to serve the schedule endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountHomeHandler(mux, h.Home)
	MountListHandler(mux, h.List)
	MountScheduleHandler(mux, h.Schedule)
	MountRemoveHandler(mux, h.Remove)
	MountStaticFaviconIco(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	}))
	MountStatic(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/")
	}))
}

// MountHomeHandler configures the mux to serve the "schedule" service "home"
// endpoint.
func MountHomeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/", f)
}

// NewHomeHandler creates a HTTP handler which loads the HTTP request and calls
// the "schedule" service "home" endpoint.
func NewHomeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeHomeResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "home")
		ctx = context.WithValue(ctx, goa.ServiceKey, "schedule")

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountListHandler configures the mux to serve the "schedule" service "list"
// endpoint.
func MountListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/", f)
}

// NewListHandler creates a HTTP handler which loads the HTTP request and calls
// the "schedule" service "list" endpoint.
func NewListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeListResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list")
		ctx = context.WithValue(ctx, goa.ServiceKey, "schedule")

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountScheduleHandler configures the mux to serve the "schedule" service
// "schedule" endpoint.
func MountScheduleHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/schedule", f)
}

// NewScheduleHandler creates a HTTP handler which loads the HTTP request and
// calls the "schedule" service "schedule" endpoint.
func NewScheduleHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeScheduleRequest(mux, dec)
		encodeResponse = EncodeScheduleResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "schedule")
		ctx = context.WithValue(ctx, goa.ServiceKey, "schedule")
		payload, err := decodeRequest(r)
		if err != nil {
			eh(ctx, w, err)
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountRemoveHandler configures the mux to serve the "schedule" service
// "remove" endpoint.
func MountRemoveHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/{id}", f)
}

// NewRemoveHandler creates a HTTP handler which loads the HTTP request and
// calls the "schedule" service "remove" endpoint.
func NewRemoveHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeRemoveRequest(mux, dec)
		encodeResponse = EncodeRemoveResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "remove")
		ctx = context.WithValue(ctx, goa.ServiceKey, "schedule")
		payload, err := decodeRequest(r)
		if err != nil {
			eh(ctx, w, err)
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountStaticFaviconIco configures the mux to serve GET request made to
// "/favicon.ico".
func MountStaticFaviconIco(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/favicon.ico", h.ServeHTTP)
}

// MountStatic configures the mux to serve GET request made to
// "/static/*filename".
func MountStatic(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/static/*filename", h.ServeHTTP)
}
