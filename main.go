package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name: ")

	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay, you can play!")
	} else {
		fmt.Println("Sorry, you can't play!")
		return
	}

	correct := 0
	incorrect := 0

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if strings.ToUpper(answer)+" "+answer2 == "RTX 3090" {
		fmt.Println("Correct!")
		correct++
	} else {
		fmt.Println("Incorrect!")
		incorrect++
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		correct++
	} else {
		fmt.Println("Incorrect!")
		incorrect++
	}

	percent := (float64(correct) / (float64(correct) + float64(incorrect))) * 100
	fmt.Printf("You got %v out of %v questions correct! This is equal to %v%%.", correct, correct+incorrect, percent)

}
