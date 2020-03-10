package main
import "fmt"

// for [инициализация_счетчика]; [условие]; [изменение_счетчика] {
// 		действия	
// }

func main() {
  for i := 1; i < 10; i++ {
    fmt.Println(i * i)
  }
  // Перебор цикла
  // for индекс, значение := range массив {
  //    действия
  // }
  var  users = [3]string{"Tom", "Alice", "Jhon"}
  for index, value := range users{
    fmt.Println(index, value)
  }
  fmt.Println("Без индексов")
  for _, value := range users{
    fmt.Println(value)
  }
  fmt.Println("Стандартный цикл")
  for i := 0; i < len(users); i++ {
    fmt.Println(users[i])
  }
  
  var numbers = [10]int{1, 2, -3, 4, 4, 4, 4, 4, 4, -1}
  var sum int = 0
  for _, value := range numbers{
    if value > 0   {
      break
    }
    sum += value
  }
  fmt.Println("Sum:", sum)
}