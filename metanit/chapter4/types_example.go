package main

import "fmt"

// type собственный_тип примитивный_тип

type mile int
type kilometer int
type library [] string

func main() {
  
  var distance mile = 5
  fmt.Println(distance)
  distanceToEnemy(distance)
  var distance2 kilometer = 4
  fmt.Println(distance2)
  // distanceToEnemy(distance2) не скомпилиться потому что ожидается тим mile
  var myLibrary library = library{"Book1", "Book2", "Book3"}
  fmt.Println(myLibrary)
  
}

func distanceToEnemy(distance mile) {
  fmt.Println("Расстояние до противника:")
  fmt.Println(distance, "миль")
}

func printBooks(lib library)  {
  for _, value := range lib{
    fmt.Println(value)
  }
}