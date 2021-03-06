package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = Service("schedule", func() {
	Description("The Alarm schedule service.")

	Method("list", func() {
		Description("List all stored bottles")
		Result(ArrayOf(Schedule))
		HTTP(func() {
			GET("/schedule")
			Response(StatusOK)
		})
	})

	Method("create", func() {
		Description("create new cron schedule")
		Result(Schedule)
		Payload(SchedulePayload)
		HTTP(func() {
			POST("/schedule")
			Response(StatusCreated)
		})
	})

	Method("remove", func() {
		Description("Remove cron schedule")
		Payload(func() {
			Attribute("id", String, "")
			Required("id")
		})
		Error("not_found", NotFound, "Bottle not found")
		HTTP(func() {
			DELETE("/schedule/{id}")
			Response(StatusNoContent)
		})
	})
})

// SchedulePayload describes a cron schedule payload.
var SchedulePayload = ResultType("application/vnd.rg.schedule", func() {
	Description("A SchedulePayload describes a cron schedule payload.")
	Reference(Schedule)
	TypeName("SchedulePayload")

	Attributes(func() {
		Attribute("name")
		Attribute("cron")
		Attribute("color")
		// TODO: add sound, camera ...
		Attribute("next")
	})
	View("default", func() {
		Attribute("name")
		Attribute("cron")
		Attribute("color")
		Attribute("next")
	})

	Required("name", "cron", "color", "next")
})

// Schedule describes a cron schedule.
var Schedule = Type("Schedule", func() {
	Description("Schedule describes a cron schedule.")
	Attribute("id", String, "ID is the unique id of the schedule.", func() {
		Example("schedule_21345")
	})
	Attribute("name", String, "Descriptive Name", func() {
		MaxLength(100)
		Default("")
		Example("Week Days at 6:30am")
	})
	Attribute("cron", String, "Valid cron string", func() {
		Example("30 6 * * 1-5") // Week Days at 6:30am
	})

	Attribute("color", String, "color to set", func() {
		Enum("red", "yellow", "green", "off")
	})

	// TODO: add sound, camera ...

	Attribute("next", String, "next time", func() {
		Example("") // Week Days at 6:30am
		Default("")
	})

	Required("id", "color", "cron")
})
