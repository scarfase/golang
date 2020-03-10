package main
import "fmt"

func main() {
  fmt.Println(divide(15, 5))
  fmt.Println(divide(123, 0))
}

func divide(x, y float64) float64 {
  if y == 0 {
    panic("Division by zero!!!")
  }
  return x/y
}
