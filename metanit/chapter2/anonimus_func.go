package main

import "fmt"



func main() {
  f := func(x, y int) int{return x + y}
  fmt.Println(f(3, 5))
  
  action(10,25, func(x, y int) int {return x + y})
  action(11,11, func(x,y int) int {return x * y})
  
  f1 := selectFn(1)
  fmt.Println(f1(3,2))
  f2 := selectFn(2)
  fmt.Println(f2(5,2))
  f4 := selectFn(4)
  fmt.Println(f4(11,11))
  
  fS := square()
  fmt.Println(fS())
}

// Анонимная функция как аргумент функции
func action(n1, n2 int, operation func(int, int) int) {
  result := operation(n1, n2)
  fmt.Println(result)
}

// Анонимная функция как результат функции
func selectFn(n int) (func(int, int) int) {
  if n == 1 {
    return func(x,y int) int {return x + y}
  } else if n == 2 {
    return func(x, y int) int {return x - y}
  } else {
    return func(x, y int) int {return x * y}
  }
}

// Доступ к окружению
func square() func() int {
  var x int = 2
  return func() int{
    x++
    return x * x
  }
}