package main

import "fmt"

// type имя_структуры struct {
//   поля_структуры
//   имя тип
// }

type person struct {
  name string
  age int
}

var tom person = person{"Tom", 23}


func main() {
  bob := person{"bob", 23}
  fmt.Println(bob.name)
  fmt.Println(bob.age)
  tom.age = 234
  fmt.Println(tom.age)
  var tomPointer *person = &tom
  tomPointer.age = 324
  fmt.Println(tom.age)
  (*tomPointer).age = 21
  fmt.Println(tom.age)
  
  var dom *person = &person{name: "Dom", age:22}
  //var domEmpty *person = new(person)
  var domPointerAge *int = &dom.age
  //var domPointerName *string = &dom.name
  //var domEmptyPointerAge *int = &domEmpty.age
  //var domEmptyPointerName *string = &domEmpty.name
  *domPointerAge = 33
  fmt.Println(dom.age)
}
