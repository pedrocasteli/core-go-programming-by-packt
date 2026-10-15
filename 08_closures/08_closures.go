package main

import "fmt"

func createTemperatureAdjuster() (func(change float64) float64, float64) {
	var originalTemperature float64 = 90.0

	adjustTemperature := func(change float64) float64 {
		originalTemperature = originalTemperature + change
		return originalTemperature
	}

	return adjustTemperature, originalTemperature
}

func priceCalculator(taxRate float64) func(float64) float64 {
	return func(price float64) float64 {
		return price * (1 + taxRate)
	}
}

func main() {
	adjustTemp, originalTemp := createTemperatureAdjuster()
	fmt.Printf("Original temperature is %.2f\n", originalTemp)
	fmt.Printf("Adjusted temp +1.5: %.1f Celsius\n", adjustTemp(1.5))

	calculateNYPrice := priceCalculator(0.08)
	calculateCAPrice := priceCalculator(0.0725)

	product := 100.0

	fmt.Printf("Price in NY: %.2f\n", calculateNYPrice(product))
	fmt.Printf("Price in CA: %.2f\n", calculateCAPrice(product))
}
