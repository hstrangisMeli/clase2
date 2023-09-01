package main

import (
	"fmt"
)

func CalcularImpuesto(sueldo float64) float64 {
	if sueldo > 150000 {
		return sueldo * 0.27
	} else if sueldo > 50000 {
		return sueldo * 0.17
	} else {
		return 0
	}
}

func main() {
	fmt.Println(CalcularImpuesto(50001))
}
