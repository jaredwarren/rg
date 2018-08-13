package rg

import (
	"log"

	sound "github.com/jaredwarren/rg/gen/sound"
)

// sound service example implementation.
// The example methods log the requests and return zero values.
type soundSvc struct {
	logger *log.Logger
}

// NewSound returns the sound service implementation.
func NewSound(logger *log.Logger) sound.Service {
	return &soundSvc{logger}
}
