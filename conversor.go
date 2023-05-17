package main

import "fmt"

func K2C(tempK float64) float64 {
	tempC := tempK - 273.15
	return tempC
}

func main() {
	var K float64
	fmt.Print("Digite a temperatura em Kelvin: ")
	fmt.Scanf("%f", &K)
	fmt.Printf("A temperatura %f Kelvins em Celsius é: %.2f", K, K2C(K))
}
