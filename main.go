package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f-32) * 5/9)
}

func main () {
	fmt.Println("Temperatures are hard to figure off the top of one's mind!")

	boilingF := CelsiusToFahrenheit(BoilingC)

	fmt.Printf("%v\n", boilingF)
	fmt.Printf("%s\n", boilingF)
	fmt.Println(boilingF)
	fmt.Printf("%g\n", boilingF)

}