package graphql

import (
	"flamingo.me/graphql"

	"flamingo.me/training/src/openweather/domain"
)

type (
	Service struct{}
)

func (*Service) Schema() []byte {
	// language=graphql
	return []byte(`
type Openweather_Weather {
	MainCharacter:       String
	Description:         String
	IconCode:            String
	Temp:                Int
	Humidity:            Int
	TempMin:             Int
	TempMax:             Int
	WindSpeed:           Float
	Cloudiness:          Int
	LocationName:        String
	LocationCountryCode: String
}

extend type Query {
	Openweather_Weather(city: String!): Openweather_Weather!
}
`)
}

// Types mapping between graphql and go
func (*Service) Types(types *graphql.Types) {
	types.Map("Openweather_Weather", new(domain.Weather))
	types.Resolve("Query", "Openweather_Weather", WeatherResolver{}, "Openweather_Weather")
}
