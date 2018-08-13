package design

// import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = Service("thermo", func() {
	Description("Thermometer service")
	HTTP(func() {
		Path("/thermo")
	})

	// TODO: add thermometer interface...

})
