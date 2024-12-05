// Package weather gives functions regarding weather.
package weather

// CurrentCondition holds the current condition.
var CurrentCondition string

// CurrentLocation holds the current location.
var CurrentLocation string

// Forecast returns a forecast for the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
