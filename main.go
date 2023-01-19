package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/auth"
	"flamingo.me/flamingo/v3/core/auth/oauth"
	"flamingo.me/flamingo/v3/core/gotemplate"
	"flamingo.me/flamingo/v3/core/healthcheck"
	"flamingo.me/flamingo/v3/framework/opencensus"
	flamingoGQL "flamingo.me/graphql"

	"flamingo.me/training/graphql"
	"flamingo.me/training/src/helloworld"
	"flamingo.me/training/src/openweather"
)

//go:generate rm -f graphql/generated.go
//go:generate go run -tags graphql main.go graphql

func main() {
	flamingo.App([]dingo.Module{
		new(gotemplate.Module),
		new(helloworld.Module),
		new(openweather.Module),
		new(healthcheck.Module),
		new(opencensus.Module),
		new(auth.WebModule),
		new(oauth.Module),
		new(flamingoGQL.Module),
		new(graphql.Module),
	})
}
