package rg

import (
	"log"

	camera "github.com/jaredwarren/rg/gen/camera"
)

// camera service example implementation.
// The example methods log the requests and return zero values.
type cameraSvc struct {
	logger *log.Logger
}

// NewCamera returns the camera service implementation.
func NewCamera(logger *log.Logger) camera.Service {
	return &cameraSvc{logger}
}
