// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// camera endpoints
//
// Command:
// $ goa gen github.com/jaredwarren/rg/design

package camera

import (
	goa "goa.design/goa"
)

// Endpoints wraps the "camera" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "camera" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "camera" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}