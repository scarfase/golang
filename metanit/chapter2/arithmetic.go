package main

import "fmt"


func main() {
	var a int = 32
	var b int = 22
	a++
	b--
	fmt.Println("Сложение: a + b =", a+b)
	fmt.Println("Вычетание: a + b =", a-b)
	fmt.Println("Умножение: a + b =", a*b)
	fmt.Println("Деление: a + b =", a/b)
	fmt.Println("Остаток от деления: a + b =", a%b)
	fmt.Println("Постфиксный инкремент: а++ =", a)
	fmt.Println("Постфиксный декремент: b-- =", b)
}
