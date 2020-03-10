package main

import "fmt"

// var название_map = map[key]value

func main() {
  var people = map[string]int {
    "tom": 1,
    "kot": 2,
    "mause": 3,
  }
  
  if val, ok := people["kot"]; ok{
    fmt.Println(val)
  }
  
  for key, value := range people{
    fmt.Println(key, value)
  }
  fmt.Println(people["mause"])
  
  mapClean := make(map[string]int)
  fmt.Println(mapClean)
  people["tom"] = 123
  fmt.Println(people)
  
  delete(people, "mause")
  fmt.Println(people)
  
}
