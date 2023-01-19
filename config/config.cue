helloworld: {
	greeting: "Hello from config"
}

core: {
	auth: {
		web: {
			broker: [
				core.auth.oidc & {
					broker:       "keycloak"
					clientID:     "client1"
					clientSecret: "client1"
					endpoint:     "http://localhost:8080/realms/Realm1"
					enableOfflineToken: false
					scopes: ["profile", "email", "address"]
					claims: {
						idToken: core.auth.oidc.claims.idToken & {"address": "address"}
					}
				},
			]
		}
	}
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
	os: env: {
		OPENWEATHER_API_KEY: string | *""
	}
	opencensus: {
		jaeger: {
			enable: false
		}
	}
}

openweather: {
	apiURL: "http://api.openweathermap.org/data/2.5/"
	apiKey: flamingo.os.env.OPENWEATHER_API_KEY
}
