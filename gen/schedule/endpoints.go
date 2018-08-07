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
	Home     goa.Endpoint
	List     goa.Endpoint
	Schedule goa.Endpoint
	Remove   goa.Endpoint
}

// NewEndpoints wraps the methods of the "schedule" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Home:     NewHomeEndpoint(s),
		List:     NewListEndpoint(s),
		Schedule: NewScheduleEndpoint(s),
		Remove:   NewRemoveEndpoint(s),
	}
}

// Use applies the given middleware to all the "schedule" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Home = m(e.Home)
	e.List = m(e.List)
	e.Schedule = m(e.Schedule)
	e.Remove = m(e.Remove)
}

// NewHomeEndpoint returns an endpoint function that calls the method "home" of
// service "schedule".
func NewHomeEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, s.Home(ctx)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "schedule".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.List(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSchedulePayloadCollection(res, "default")
		return vres, nil
	}
}

// NewScheduleEndpoint returns an endpoint function that calls the method
// "schedule" of service "schedule".
func NewScheduleEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SchedulePayload)
		return s.Schedule(ctx, p)
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
