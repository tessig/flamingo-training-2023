package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/gotemplate"
	"flamingo.me/flamingo/v3/core/healthcheck"
	"flamingo.me/flamingo/v3/framework/opencensus"

	"flamingo.me/training/src/helloworld"
	"flamingo.me/training/src/openweather"
)

func main() {
	flamingo.App([]dingo.Module{
		new(helloworld.Module),
		new(gotemplate.Module),
		new(openweather.Module),
		new(healthcheck.Module),
		new(opencensus.Module),
	})
}
