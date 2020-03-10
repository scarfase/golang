package main
import "fmt"

func main() {
  fmt.Println(factorial(4))
  fmt.Println(factorial(5))
  fmt.Println(factorial(8))
  fmt.Println(fibonachi(4))
  fmt.Println(fibonachi(28))
}

func factorial(n uint) uint {
  if n == 0 {
    return 1
  }
  return n * factorial(n - 1)
}

func fibonachi(n uint) uint {
  if n == 0 {
    return 0
  }
  if n == 1 {
    return 1
  }
  return fibonachi(n - 1) + fibonachi(n - 2)
}