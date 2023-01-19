helloworld: {
	greeting: "Hello from config"
}

core: {
	gotemplate: engine: {
		templates: basepath: "templates"
		layout: dir:         "layouts"
	}
	zap: {
		loglevel:   "Warn"
		json:       true
		colored:    false
		logsession: true
	}
	healthcheck: {
		checkSession: true
		checkAuth:    true
	}
}

flamingo: {
	opencensus: {
		jaeger: {
			enable: true
		}
	}
}

openweather: {
	apiURL: "http://api.openweathermap.org/data/2.5/"
	apiKey: flamingo.os.env.OPENWEATHER_API_KEY
}
