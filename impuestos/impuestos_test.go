package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularImpuestoSueldoAlto(t *testing.T) {
	//Inicializo
	sueldo := 600000.0
	esperado := sueldo * 0.27

	//Ejecuto
	obtenido := CalcularImpuesto(sueldo)

	//Assert
	assert.Equal(t, esperado, obtenido)
}

func TestCalcularImpuestoSueldoMedio(t *testing.T) {
	//Inicializo
	sueldo := 100000.0
	esperado := sueldo * 0.17

	//Ejecuto
	obtenido := CalcularImpuesto(sueldo)

	//Assert
	assert.Equal(t, esperado, obtenido)
}

func TestCalcularImpuestoSueldoBajo(t *testing.T) {
	//Inicializo
	sueldo := 100.0
	esperado := 0.0

	//Ejecuto
	obtenido := CalcularImpuesto(sueldo)

	//Assert
	assert.Equal(t, esperado, obtenido)
}
