package main

import (
  "fmt"
)

// func имя_функции (список_параметров) (типы_возвращаемых_значений) {
//    выполняемые_операторы
//    return
// }

func main() {
  hello()
  hello()
  add(4, 6)
  addArr(123,412,412)
  fmt.Println(sumTwoNumbers(45, 123))
  fmt.Println(sumThreeNumbers(1,2,3,))
  fmt.Println(returnTwoDiffValues(12,213, "Dima", "Aitkulov"))
  // Можем присвоить результат выполнения функции переменным
  // var age, name := returnTwoDiffValues(12,213, "Dima", "Aitkulov")
  // age, name := returnTwoDiffValues(12,213, "Dima", "Aitkulov")
}

func  hello(){
  fmt.Println("Hello World")
}

func add(x, y int) {
  var z int = x + y
  fmt.Println("x + y =", z)
}

func addArr(numbers ...int) {
  var sum = 0
  for _, number := range numbers{
    sum += number
  }
  fmt.Println("sum = ", sum)
}

func sumTwoNumbers(x,y int) int {
  return x + y
}

func sumThreeNumbers(x, y, z int) (sum int) {
  sum = x + y + z
  return
}

func returnTwoDiffValues(x, y int, firstName, lastName string) (int, string) {
  var z int = x + y
  var fullName = firstName + " " + lastName
  return z, fullName
}
