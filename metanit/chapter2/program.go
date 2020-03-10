package main // Первым делом идет объявление пакета

// Импорт пакетов
import "fmt"

// Объявление переменных
// var имя_переменной тип_переменной
var hello string

// Если хотим определить несколько переменных, то можно их обернуть в скобки
var (
	name string = "Tom"
	age  int    = 19
)

// Динамическое присвоение типа переменный в зависимости от значения
// Работает внутри функций
// dinamicVar := "var"

// Выполнение программы начинается с функции main
func main() {
	fmt.Println("Hello Go")

	var hello string = "Hello world"
	fmt.Println(hello)

	dinamicVar := "var"
	fmt.Println(dinamicVar)
}
