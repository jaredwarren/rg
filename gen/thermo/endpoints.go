// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// thermo endpoints
//
// Command:
// $ goa gen github.com/jaredwarren/rg/design

package thermo

import (
	goa "goa.design/goa"
)

// Endpoints wraps the "thermo" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "thermo" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "thermo" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}
