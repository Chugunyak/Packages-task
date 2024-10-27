package utils

import "fmt"

func PrintAnimalInfo(animal string, sound string, move string, feed string, sleep string, special string) {
	fmt.Printf("Животное: %s\n", animal)
	fmt.Println("Звук: ", sound)
	fmt.Println("Передвижение: ", move)
	fmt.Println("Еда: ", feed)
	fmt.Println("Сон: ", sleep)
	fmt.Println("Особенность: ", special)
}
