package alimento

const (
	perro     = "Perro"
	gato      = "Gato"
	hamster   = "Hamster"
	tarantula = "Tarantula"
)

var alimentoPorAnimal = map[string]int{"Perro": 10000, "Gato": 5000, "Hamster": 250, "Tarantula": 150}

func Alimento(nombreAnimal string) (cantidadAlimento func(int) float64, err string) {
	switch nombreAnimal {
	case perro:
		cantidadAlimento = cantidadAlimentoPerro
	case gato:
		cantidadAlimento = cantidadAlimentoGato
	case hamster:
		cantidadAlimento = cantidadAlimentoHamster
	case tarantula:
		cantidadAlimento = cantidadAlimentoTarantula
	default:
		err = "No existe el animal ingresado"
	}
	return
}

func cantidadAlimentoPerro(cantPerros int) float64 {
	return float64(alimentoPorAnimal[perro]) * float64(cantPerros)
}

func cantidadAlimentoGato(cantGatos int) float64 {
	return float64(alimentoPorAnimal[gato]) * float64(cantGatos)
}

func cantidadAlimentoHamster(cantHamsters int) float64 {
	return float64(alimentoPorAnimal[hamster]) * float64(cantHamsters)
}

func cantidadAlimentoTarantula(cantTarantulas int) float64 {
	return float64(alimentoPorAnimal[tarantula]) * float64(cantTarantulas)
}

/*
Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.
Tienen los siguientes datos:
Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.
Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne
una función y un mensaje (en caso que no exista el animal).
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
Ejemplo:


const (
	dog    = "dog"
	cat    = "cat"
)


...


animalDog, msg := animal(dog)
animalCat, msg := animal(cat)


...


var amount float64
amount += animalDog(10)
amount += animalCat(10)

*/
