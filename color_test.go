package rg

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/boltdb/bolt"
	color "github.com/jaredwarren/rg/gen/color"
	"github.com/jaredwarren/rg/pi"
)

func TestColor(t *testing.T) {

	// Setup logger and goa log adapter. Replace logger with your own using
	// your log package of choice. The goa.design/middleware/logging/...
	// packages define log adapters for common log packages.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[rg_test] ", log.Ltime)
	}

	// Initialize service dependencies such as databases.
	var (
		bdb *bolt.DB
	)
	{
		os.Remove("rg_test.db")
		var err error
		bdb, err = bolt.Open("rg_test.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer bdb.Close()
	}
	var err error

	rpi := pi.NewPi()

	// Create the structs that implement the services.
	var colorSvc color.Service
	colorSvc = NewColor(rpi, logger)
	if colorSvc == nil {
		t.Errorf("service nil: %s", err)
	}

	fakeCtx := context.TODO()

	//
	// Get/Update color
	//

	// initial color should be "off"
	c, err := colorSvc.Color(fakeCtx)
	if err != nil {
		logger.Fatalf("error schedule.Color: %s", err)
	}
	if c.Color != "off" {
		logger.Fatalf("error initial color failed: %s", c.Color)
	}

	err = colorSvc.Update(fakeCtx, &color.UpdatePayload{
		Color: "yellow",
	})
	if err != nil {
		logger.Fatalf("error schedule.Update: %s", err)
	}

	c, err = colorSvc.Color(fakeCtx)
	if err != nil {
		logger.Fatalf("error schedule.Color: %s", err)
	}
	if c.Color != "yellow" {
		logger.Fatalf("error update color failed: %s", c.Color)
	}

}
