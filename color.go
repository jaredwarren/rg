package rg

import (
	"context"
	"log"

	color "github.com/jaredwarren/rg/gen/color"
	"github.com/jaredwarren/rg/pi"
)

// ColorSvc service example implementation.
// The example methods log the requests and return zero values.
type colorSvc struct {
	pi     pi.Pi
	logger *log.Logger
}

// NewColor returns the color service implementation.
func NewColor(rpi pi.Pi, logger *log.Logger) color.Service {
	return &colorSvc{rpi, logger}
}

// Set color
func (s *colorSvc) Update(ctx context.Context, p *color.UpdatePayload) (err error) {
	s.logger.Print("color.update:", p.Color)
	return s.pi.SetColor(p.Color)
}

// get current color
func (s *colorSvc) Color(ctx context.Context) (res *color.Color, err error) {
	s.logger.Print("color.color")
	return &color.Color{
		Color: s.pi.GetColor(),
	}, nil
}

// TODO add led service? crud leds and/or other services
