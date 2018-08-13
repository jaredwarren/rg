package design

// import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = Service("camera", func() {
	Description("Camera service")
	HTTP(func() {
		Path("/camera")
	})

	// TODO: add camera interface...

})
