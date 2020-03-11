package main

import "fmt"

// Объявление указателя
// var p *int
// Получение адреса у ранее созданной переменной
// p = &someVar
// Получение значения переменной или разыменование
// *p
// Используя указатель можно менять значение переменной
// *p = 12

func main() {
  var x int = 4
  var p *int = &x
  fmt.Println("Address:",p)
  fmt.Println("Value:", *p)
  *p = 25
  fmt.Println("New value:", *p)
  // Упрощенная форма записи
  f := 2.3
  pF := &f
  fmt.Println("Address:", pF)
  fmt.Println("Value:", *pF)
  var pEmpty *float64
  if pEmpty != nil {
    fmt.Println("VAlue:", *pEmpty)
  }
  // new() function
  pointerNew := new(int)
  fmt.Println("Value:", *pointerNew)
  *pointerNew = 8
  fmt.Println("Value:", *pointerNew)
  
}
