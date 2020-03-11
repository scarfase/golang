package main

import "fmt"

func main() {
  d := 5
  fmt.Println("d before:", d)
  changeValue(&d)
  fmt.Println("d after:", d)
  p1 := createPointer(324)
  fmt.Println("p1:", *p1)
  p2 := createPointer(12938)
  fmt.Println("p2:", *p2)
  
}

func changeValue(x *int) {
  *x = (*x) * (*x)
}

func createPointer(x int) *int {
  p := new(int)
  *p = x
  return p
}
