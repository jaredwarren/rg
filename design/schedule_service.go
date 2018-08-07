package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = Service("schedule", func() {
	Description("The Alarm schedule service.")
	HTTP(func() {
		Path("/")
	})

	// HTML

	Method("home", func() {
		Description("Alarm Schedule Home")
		// Result(CollectionOf(Game))
		HTTP(func() {
			GET("/")
			// Response(StatusOK, "text/html")
			Response(StatusOK)
		})
	})

	Files("/favicon.ico", "static/favicon.ico")
	Files("/static/*filename", "static/")

	// JSON

	Method("list", func() {
		Description("List all stored bottles")
		Result(CollectionOf(SchedulePayload))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("schedule", func() {
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
			Attribute("id", String, "ID of bottle to remove")
			Required("id")
		})
		Error("not_found", NotFound, "Bottle not found")
		HTTP(func() {
			DELETE("/{id}")
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
		Attribute("id", String, "ID is the unique id of the schedule.", func() {
			Example("sched_21345")
		})
		Attribute("name")
		Attribute("cron")
		Attribute("color")
	})

	Required("id", "name", "cron", "color")
})

// Schedule describes a cron schedule.
var Schedule = Type("Schedule", func() {
	Description("Schedule describes a cron schedule.")
	Attribute("name", String, "Descriptive Name", func() {
		MaxLength(100)
		Example("Week Days at 6:30am")
	})
	Attribute("cron", String, "Valid cron string", func() {
		Example("30 6 * * 1-5") // Week Days at 6:30am
	})

	Attribute("color", String, "color to set", func() {
		Enum("red", "yellow", "green", "off")
	})

	Required("color", "cron")
})
