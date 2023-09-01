package main

import (
	"clase2/alimento"
	"fmt"
)

func main() {
	getCantidad, err := alimento.Alimento("Perro")
	fmt.Println(getCantidad(2), err)
}
