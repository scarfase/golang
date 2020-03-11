package main

import "fmt"

// func (имя_параметра тип_получателя) имя_метода (параметры) (типы_возвращаемых_результатов) {
//    тело_метода
// }

type methodLibrary []string

type methodPerson struct {
  name string
  age int
}

func (p methodPerson) print() {
  fmt.Println("Имя:", p.name)
  fmt.Println("Возраст:", p.age)
}

func (p methodPerson) eat(meal string) {
  fmt.Println(p.name, "ест", meal)
}

func (lib methodLibrary)print()  {
  for _, val := range lib {
    fmt.Println(val)
  }
}

func main() {
  var lib methodLibrary = methodLibrary{"123", "123", "end"}
  lib.print()
  
  var tom = methodPerson {name: "Tom", age: 23}
  tom.print()
  tom.eat("фыдлваодлфаодлф")
}
