// Package weather, tool to convert.
package weather

// CurrentCondition current condition.
var CurrentCondition  string

// CurrentLocation current location.
var CurrentLocation string

// Forecast return current forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
