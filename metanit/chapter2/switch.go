package main

import "fmt"

func main() {
	a := 0
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 10, 11, 12:
		fmt.Println("a = 10")
	default:
		fmt.Println("Значение переменной не определенно")
	}
}
