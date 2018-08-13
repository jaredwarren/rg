package design

// import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = Service("sound", func() {
	Description("Sound service")
	HTTP(func() {
		Path("/sound")
	})

	// TODO: add sound interface... volume, sound, play/pause, etc...

})
