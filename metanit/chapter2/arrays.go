package main

import "fmt"

// var имя_массива [число_элементов]тим_элементов

func main() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	number2 := [...]int{1, 2, 2, 2, 2}
	fmt.Println(number2)
	fmt.Println(numbers[2])
}
