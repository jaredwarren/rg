package rg

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jaredwarren/rg/db"
	schedule "github.com/jaredwarren/rg/gen/schedule"
	"github.com/jaredwarren/rg/pi"
)

func TestRun(t *testing.T) {

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

	// Add item to db to cover more test
	tdb, err := db.NewScheduleDB(bdb)
	if err != nil {
		t.Errorf("db error: %s", err)
	}
	tdb.Save("SCHEDULE", "test_schedule_00", &schedule.Schedule{
		Name:  "test_schedule00",
		Cron:  "0 32 * * * *", // every hour on hh:30
		Color: "off",
	})

	rpi := pi.NewPi()

	// Create the structs that implement the services.
	var scheduleSvc schedule.Service
	scheduleSvc, err = NewSchedule(bdb, rpi, logger)
	if err != nil {
		logger.Fatalf("error creating service: %s", err)
	}
	if scheduleSvc == nil {
		t.Errorf("service nil: %s", err)
	}

	fakeCtx := context.TODO()

	//
	// Create
	//
	res1, err := scheduleSvc.Create(fakeCtx, &schedule.SchedulePayload{
		Name:  "test_schedule",
		Cron:  "0 30 * * * *", // every hour on hh:30
		Color: "red",
	})
	if err != nil {
		logger.Fatalf("error schedule.Create: %s", err)
	}
	if res1 == nil {
		t.Errorf("service response nil: %s", err)
	}

	tNext, err := time.Parse("2006-01-02 15:04:05 -0700 MST", res1.Next)
	if err != nil {
		t.Errorf("next time error: %s", err)
	}

	// could be as hight as 59 minutes away. i.e. if now is hh:30:01
	t1 := time.Now().Local().Add(time.Hour * 1)
	if !t1.After(tNext) {
		t.Error("next fail 1")
	}
	fmt.Println(t1, " - ", tNext, " - ", time.Now())

	t2 := time.Now()
	if !tNext.After(t2) {
		t.Error("next fail 2")
	}

	//
	// List
	//
	// add an extra one to cover more test
	scheduleSvc.Create(fakeCtx, &schedule.SchedulePayload{
		Name:  "test_schedule1",
		Cron:  "0 32 * * * *", // every hour on hh:30
		Color: "red",
	})
	rList, err := scheduleSvc.List(fakeCtx)
	if err != nil {
		logger.Fatalf("error schedule.Remove: %s", err)
	}
	if len(rList) != 3 {
		t.Error("List wrong length expected 3 got:", len(rList))
	}

	//
	// Remove
	//
	err = scheduleSvc.Remove(fakeCtx, &schedule.RemovePayload{
		ID: res1.ID,
	})
	if err != nil {
		logger.Fatalf("error schedule.Remove: %s", err)
	}
	// make sure removed from list
	rList, err = scheduleSvc.List(fakeCtx)
	if err != nil {
		logger.Fatalf("error schedule.Remove: %s", err)
	}
	if len(rList) != 2 {
		t.Error("List wrong length expected 2 got:", len(rList))
	}

	//
	// Get/Update color
	//

	// // initial color should be "off"
	// color, err := scheduleSvc.Color(fakeCtx)
	// if err != nil {
	// 	logger.Fatalf("error schedule.Color: %s", err)
	// }
	// if color.Color != "off" {
	// 	logger.Fatalf("error initial color failed: %s", color.Color)
	// }

	// err = scheduleSvc.Update(fakeCtx, &schedule.UpdatePayload{
	// 	Color: "yellow",
	// })
	// if err != nil {
	// 	logger.Fatalf("error schedule.Update: %s", err)
	// }

	// color, err = scheduleSvc.Color(fakeCtx)
	// if err != nil {
	// 	logger.Fatalf("error schedule.Color: %s", err)
	// }
	// if color.Color != "yellow" {
	// 	logger.Fatalf("error update color failed: %s", color.Color)
	// }

}
