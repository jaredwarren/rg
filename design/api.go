package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = API("rg", func() {
	Title("Red Green Alarm")
	Description("HTTP service for managing your recipes")
	Server("http://localhost:8080")
})

// NotFound ...
var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a bottle that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Metadata("struct:error:name")
		Example("bottle 1 not found")
	})
	Attribute("id", String, "ID of missing bottle")
	Required("message", "id")
})
