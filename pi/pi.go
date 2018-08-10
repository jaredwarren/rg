package pi

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

// Pi ...
type Pi struct {
	Leds map[string]*gpio.LedDriver
}

func NewPi() Pi {
	// Raspberry Pi
	rpi := raspi.NewAdaptor()
	red := gpio.NewLedDriver(rpi, fmt.Sprintf("%X", 17))
	yellow := gpio.NewLedDriver(rpi, fmt.Sprintf("%X", 18)) // make sure pins work
	green := gpio.NewLedDriver(rpi, fmt.Sprintf("%X", 19))

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
		Leds: map[string]*gpio.LedDriver{
			"red":    red,
			"yellow": yellow,
			"green":  green,
		},
	}
}
