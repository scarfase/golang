package main

import "fmt"

type contact struct {
  email string
  phone string
}

type  person2 struct {
  name string
  age int
  contactInfo contact
}

type newPerson struct {
  name string
  age int
  contact
}

type node struct {
  value int
  next *node
}

func printNodeValue(n *node) {
  fmt.Println(n.value)
  if n.next != nil {
    printNodeValue(n.next)
  }
}

func main() {
  var tom = person2 {
    name: "Tom",
    age: 24,
    contactInfo: contact{
      email: "tom@gmail.com",
      phone: "+77773437207",
    },
  }
  tom.contactInfo.email = "newMail@gmail.com"
  fmt.Println(tom.contactInfo.email)
  fmt.Println(tom.name)
  
  var cat = newPerson {
    name: "Mase",
    age: 23,
    contact: contact{
      email: "cat@gmail.com",
      phone: "+77773437207",
    },
  }
  
  fmt.Println(cat.email)
  fmt.Println(cat.name)
  
  first := node{value:4}
  second := node{value:5}
  third := node{value:6}
  
  first.next = &second
  second.next = &third
  
  var current *node = &first
  for current != nil {
    fmt.Println(current.value)
    current = current.next
  }
}
