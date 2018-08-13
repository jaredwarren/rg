package rg

import (
	"log"

	home "github.com/jaredwarren/rg/gen/home"
)

// home service example implementation.
// The example methods log the requests and return zero values.
type homeSvc struct {
	logger *log.Logger
}

// NewHome returns the home service implementation.
func NewHome(logger *log.Logger) home.Service {
	return &homeSvc{logger}
}
