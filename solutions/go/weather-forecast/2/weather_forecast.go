// Package weather declare.
package weather

var (
    // CurrentCondition shows the current weather condition.
	CurrentCondition string
    // CurrentLocation shows the current weather location.
	CurrentLocation string
)

// Forecast city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
