// Package weather declare.
package weather

var (
    // CurrentCondition string.
	CurrentCondition string
    // CurrentLocation string.
	CurrentLocation string
)

// Forecast city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
