package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Circul Area Calculation")
	fmt.Print("enter a radus value:")

	var radius float64
	fmt.Scanf("%f", &radius)

	area := math.Pi * math.Pow(radius, 2)

	fmt.Printf("circule area with radius %.2f =%.2f\n", radius, area)

}
