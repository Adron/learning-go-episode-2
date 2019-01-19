package main

import (
	"fmt"

	tempconv "github.com/Adron/awesomeProject/tempconv"
)

func main() {
	fmt.Println("Temperatures are hard to figure off the top of one's mind!")

	boilingF := tempconv.CelsiusToFahrenheit(BoilingC)

	fmt.Printf("%v\n", boilingF)
	fmt.Printf("%s\n", boilingF)
	fmt.Println(boilingF)
	fmt.Printf("%g\n", boilingF)
}
