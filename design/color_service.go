package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = Service("color", func() {
	Description("Color service")
	HTTP(func() {
		Path("/")
	})

	Method("update", func() {
		Description("Set color")
		Payload(func() {
			Attribute("color", String, "color to set", func() {
				Enum("red", "yellow", "green", "off")
			})
			Required("color")
		})
		HTTP(func() {
			POST("/color")
			Response(StatusNoContent)
		})
	})

	Method("color", func() {
		Description("get current color")
		Result(Color)
		HTTP(func() {
			GET("/color")
			Response(StatusOK)
		})
	})
})

// Color current state.
var Color = Type("Color", func() {
	Description("Color current state.")

	Attribute("color", String, "color to set", func() {
		Enum("red", "yellow", "green", "off")
	})

	Required("color")
})
