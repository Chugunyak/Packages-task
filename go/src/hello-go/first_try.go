package main

import (
	"hello-go/animals"
	"hello-go/interactions"
)

func main() {
	animalsList := []animals.Animal{
		animals.Lion{},
		animals.Dolphin{},
		animals.Elephant{},
		animals.Frog{},
		animals.Penguin{},
	}

	for _, animal := range animalsList {
		interactions.ProcessAnimal(animal)
	}
}
