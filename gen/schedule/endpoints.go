// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// schedule endpoints
//
// Command:
// $ goa gen github.com/jaredwarren/rg/design

package schedule

import (
	"context"

	goa "goa.design/goa"
)

// Endpoints wraps the "schedule" service endpoints.
type Endpoints struct {
	List   goa.Endpoint
	Create goa.Endpoint
	Remove goa.Endpoint
}

// NewEndpoints wraps the methods of the "schedule" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:   NewListEndpoint(s),
		Create: NewCreateEndpoint(s),
		Remove: NewRemoveEndpoint(s),
	}
}

// Use applies the given middleware to all the "schedule" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Create = m(e.Create)
	e.Remove = m(e.Remove)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "schedule".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.List(ctx)
	}
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "schedule".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SchedulePayload)
		return s.Create(ctx, p)
	}
}

// NewRemoveEndpoint returns an endpoint function that calls the method
// "remove" of service "schedule".
func NewRemoveEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RemovePayload)
		return nil, s.Remove(ctx, p)
	}
}
