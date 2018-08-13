package design

// import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = Service("home", func() {
	Description("Home html/ui service")
	HTTP(func() {
		Path("/")
	})

	// HTML

	Files("/favicon.ico", "static/favicon.ico")
	Files("/static/{*filename}", "static/")
	Files("/home/", "static/index.html")

})
