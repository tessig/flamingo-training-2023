package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/gotemplate"

	"flamingo.me/training/src/helloworld"
)

func main() {
	flamingo.App([]dingo.Module{
		new(helloworld.Module),
		new(gotemplate.Module),
	})
}
