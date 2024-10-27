package interactions

import (
	"fmt"
	"hello-go/animals"
)

func ProcessAnimal(animal animals.Animal) {
	fmt.Printf("Животное: %T\n", animal)
	fmt.Println("Звук: ", animal.Sound())
	fmt.Println("Передвижение: ", animal.Move())
	fmt.Println("Еда: ", animal.Feed())
	fmt.Println("Сон: ", animal.Sleep())
	fmt.Println("Особенность: ", animal.Special())

	if swimmer, ok := animal.(animals.Swimmer); ok {
		fmt.Println("Умеет ли плавать:", swimmer.IsSwimming())
		fmt.Println("Способ плавания:", swimmer.Swim())
	}

	fmt.Println()
}
