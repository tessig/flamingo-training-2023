core: {
	auth: web: {
		broker: [
			core.auth.oidc & {
				broker:             "keycloak"
				clientID:           "client1"
				clientSecret:       "client1"
				endpoint:           "http://localhost:8080/realms/Realm1"
				enableOfflineToken: true
				claims: {
					idToken: core.auth.oidc.claims.idToken & {
						address: "address"
					}
				}
			},
		]
	}
	gotemplate: engine: {
		templates: basepath: "templates"
		layout: dir:         "layouts"
	}
}

flamingo: {
	opencensus: jaeger: enable: true
}

openweather: {
	defaultCity: "Berlin"
	apiURL:      "http://api.openweathermap.org/data/2.5/"
	apiKey:      flamingo.os.env.OPENWEATHER_API_KEY
}

graphql: {
	introspectionEnabled: true
}
