// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// sound endpoints
//
// Command:
// $ goa gen github.com/jaredwarren/rg/design

package sound

import (
	goa "goa.design/goa"
)

// Endpoints wraps the "sound" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "sound" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "sound" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}