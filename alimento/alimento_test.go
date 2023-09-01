package alimento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularCantidadAlimento1Perro(t *testing.T) {
	t.Run("Calcular alimento para un perro", func(t *testing.T) {
		//Arrange
		nombreAnimal := "Perro"
		cantidadAnimales := 1
		cantidadEsperada := 10000.0

		//Act
		funcAlimento, err := Alimento(nombreAnimal)
		var cantidadObtenida float64 = funcAlimento(cantidadAnimales)

		//Assert
		assert.Equal(t, cantidadEsperada, cantidadObtenida)
		assert.Empty(t, err)
	})

	t.Run("Calcular alimento para 1 gato", func(t *testing.T) {
		//Arrange
		nombreAnimal := "Gato"
		cantidadAnimales := 1
		cantidadEsperada := 5000.0

		//Act
		funcAlimento, err := Alimento(nombreAnimal)
		var cantidadObtenida float64 = funcAlimento(cantidadAnimales)

		//Assert
		assert.Equal(t, cantidadEsperada, cantidadObtenida)
		assert.Empty(t, err)
	})

	t.Run("Calcular alimento para 1 hamster", func(t *testing.T) {
		//Arrange
		nombreAnimal := "Hamster"
		cantidadAnimales := 1
		cantidadEsperada := 250.0

		//Act
		funcAlimento, err := Alimento(nombreAnimal)
		var cantidadObtenida float64 = funcAlimento(cantidadAnimales)

		//Assert
		assert.Equal(t, cantidadEsperada, cantidadObtenida)
		assert.Empty(t, err)
	})

	t.Run("Calcular alimento para 4 Tarantulas", func(t *testing.T) {
		//Arrange
		nombreAnimal := "Tarantula"
		cantidadAnimales := 4
		cantidadEsperada := 600.0

		//Act
		funcAlimento, err := Alimento(nombreAnimal)
		var cantidadObtenida float64 = funcAlimento(cantidadAnimales)

		//Assert
		assert.Equal(t, cantidadEsperada, cantidadObtenida)
		assert.Empty(t, err)
	})

	t.Run("Calcular alimento para 2 Vacas", func(t *testing.T) {
		//Arrange
		nombreAnimal := "Vaca"

		//Act
		_, err := Alimento(nombreAnimal)

		//Assert
		assert.Equal(t, err, "No existe el animal ingresado")
	})
}
