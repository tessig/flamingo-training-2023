helloworld: {
	greeting: "Hello from the config file"
}

core: {
	gotemplate: engine: {
		layout: dir:         "layouts"
	}

	zap: {
		loglevel:   "Warn"
		json:       true
		colored:    false
		logsession: true
	}
}

flamingo: {
	opencensus: {
		jaeger: {
			enable: true
		}
	}
}
