package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/auth/oauth"
	"flamingo.me/flamingo/v3/core/cache"
	"flamingo.me/flamingo/v3/core/gotemplate"
	"flamingo.me/flamingo/v3/framework/opencensus"
	flamingoGraphQL "flamingo.me/graphql"

	"flamingo.me/training/graphql"
	"flamingo.me/training/src/helloworld"
	"flamingo.me/training/src/openweather"
)

//go:generate rm -f graphql/generated.go
//go:generate go run -tags graphql main.go graphql

func main() {
	flamingo.App([]dingo.Module{
		new(oauth.Module),
		new(gotemplate.Module),
		new(helloworld.Module),
		new(openweather.Module),
		new(opencensus.Module),
		new(flamingoGraphQL.Module),
		new(graphql.Module),
		dingo.ModuleFunc(func(injector *dingo.Injector) {
			injector.Bind(new(cache.Backend)).ToInstance(cache.NewInMemoryCache())
		}),
	})
}
