package pi

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

// Pi ...
type Pi struct {
	currentColor string
	leds         map[string]*gpio.LedDriver
}

// NewPi ...
func NewPi() Pi {
	rpi := raspi.NewAdaptor()
	// TODO: Find a better way to dynamically set leds
	red := gpio.NewLedDriver(rpi, fmt.Sprintf("%X", 17))
	yellow := gpio.NewLedDriver(rpi, fmt.Sprintf("%X", 27)) // make sure pins work
	green := gpio.NewLedDriver(rpi, fmt.Sprintf("%X", 22))

	work := func() {
		red.Off()
		yellow.Off()
		green.Off()
	}

	robot := gobot.NewRobot("rg_alarm",
		[]gobot.Connection{rpi},
		[]gobot.Device{red, yellow, green},
		work,
	)
	go robot.Start()
	return Pi{
		currentColor: "off",
		leds: map[string]*gpio.LedDriver{
			"red":    red,
			"yellow": yellow,
			"green":  green,
		},
	}
}

// SetColor ...
func (p *Pi) SetColor(color string) error {
	for _, led := range p.leds {
		led.Off()
	}
	if color != "off" {
		led, _ := p.leds[color]
		if led != nil {
			// TODO: for now just log error
			led.On()
			p.currentColor = color
		}
	}
	return nil
}

// GetColor ...
func (p *Pi) GetColor() string {
	return p.currentColor
}
