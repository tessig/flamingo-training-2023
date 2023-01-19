package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/gotemplate"

	"flamingo.me/training/src/helloworld"
	"flamingo.me/training/src/openweather"
)

func main() {
	flamingo.App([]dingo.Module{
		new(gotemplate.Module),
		new(helloworld.Module),
		new(openweather.Module),
	})
}
