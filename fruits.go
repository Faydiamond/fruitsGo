package main

import (
	"fmt"
	"strings"
)

var fruits []string

func add_fruit() {
	var fruit string
	fmt.Println(" Digite una fruta")
	fmt.Scanln(&fruit)
	fruits = append(fruits, fruit)
	menu()
}

func show_fruits() {
	for _, value := range fruits {
		fmt.Println(value)
	}
}

func menu() {
	var question string
	fmt.Println(" desea agregar una fruta, digite si para agregar, digite no para ver la bolsa")
	fmt.Scanln(&question)

	question = strings.ToLower(question)
	if question == "si" {
		//fmt.Println("Agregar Fruta")
		add_fruit()
	} else if question == "no" {
		//fmt.Println("Mostrar Bolsa de frutas")
		show_fruits()
	} else {
		menu()
	}
}

func main() {

	menu()
}
