package main

import "fmt"

type kelvin float64
type celsius float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) toCelsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) String() string {
	f := float64(k)
	return fmt.Sprintf("%v degrees", f)
}

func main() {
	var k kelvin = 294.0
	var c celsius = kelvinToCelsius(k)
	c = k.toCelsius()
	fmt.Printf("%v, %v, %v, %v in celsius: %v\n", k1, k2, k3, k4, c)
}
