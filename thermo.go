package rg

import (
	"log"

	thermo "github.com/jaredwarren/rg/gen/thermo"
)

// thermo service example implementation.
// The example methods log the requests and return zero values.
type thermoSvc struct {
	logger *log.Logger
}

// NewThermo returns the thermo service implementation.
func NewThermo(logger *log.Logger) thermo.Service {
	return &thermoSvc{logger}
}
