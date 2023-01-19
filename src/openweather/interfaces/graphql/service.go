package graphql

import (
	"flamingo.me/graphql"

	"flamingo.me/training/src/openweather/domain"
)

type (
	Service struct {
	}
)

func (s *Service) Schema() []byte {
	// language=graphql
	return []byte(`
"""
The current state of the weather
"""
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
	""" 
	Get the current weather for the city
	"""
	Openweather_Weather(city: String!): Openweather_Weather!
}`)
}

func (s *Service) Types(types *graphql.Types) {
	types.Map("Openweather_Weather", domain.Weather{})
	types.Resolve("Query", "Openweather_Weather", WeatherResolver{}, "Openweather_Weather")
}
