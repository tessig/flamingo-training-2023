helloworld: {
	greeting: "Hello from config"
}

core: {
	gotemplate: engine: {
		templates: basepath: "templates"
		layout: dir:         "layouts"
	}
}

openweather: {
	apiURL:      "http://api.openweathermap.org/data/2.5/"
	apiKey:      flamingo.os.env.OPENWEATHER_API_KEY
}
