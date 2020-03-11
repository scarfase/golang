package main

import "fmt"

type pointerPerson struct {
  name string
  age int
}

func (p *pointerPerson) updateAge(newAge int) {
  (*p).age = newAge
}

func main() {
  var tom = pointerPerson{
    name: "Tom",
    age:  3434,
  }
  
  tomPointer := &tom
  fmt.Println("before:", tom.age)
  tomPointer.updateAge(22)
  fmt.Println("after:", tom.age)
}
