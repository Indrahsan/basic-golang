package main

import "fmt"

func main() {
	//How to print
	fmt.Println("Hello", "world!", "how", "are", "you")

	//How to use variable
	var firstName string = "john"

	lastName := "wick"

	//karakter %s di situ akan diganti dengan data string yang berada di parameter ke-2, ke-3, dan seterusnya.
	fmt.Printf("halo %s %s!\n", firstName, lastName)

	//Multi variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	//%s for string
	//%d for integer
	//%v for float
	//%t for boolean
	fmt.Printf("%s %s %s %s %s %s %s %s %s %d %t %v %s\n", first, second, third, fourth, fifth, sixth, seventh, eight, ninth, one, isFriday, twoPointTwo, say)

	_ = "belajar Golang"

	var message = `Nama saya "John Wick".
Salam kenal.
Mari belajar "Golang".`

	fmt.Println(message)

	//Use another function
	aritmatik()
	operatorLogika()
}

func aritmatik() {
	fmt.Println()
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)
}

func operatorLogika() {
	fmt.Println()
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
