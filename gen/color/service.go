// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// color service
//
// Command:
// $ goa gen github.com/jaredwarren/rg/design

package color

import (
	"context"
)

// Color service
type Service interface {
	// Set color
	Update(context.Context, *UpdatePayload) (err error)
	// get current color
	Color(context.Context) (res *Color, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "color"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"update", "color"}

// UpdatePayload is the payload type of the color service update method.
type UpdatePayload struct {
	// color to set
	Color string
}

// Color is the result type of the color service color method.
type Color struct {
	// color to set
	Color string
}
