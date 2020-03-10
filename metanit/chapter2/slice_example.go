package main

import "fmt"

// var имя_среза []тип_среза





func main() {
  var users = []string{"Mike", "Tom", "Kate"}
  users22 := []string{"Tom", "Alice", "Kate"}
  var usersMake []string = make([]string, 5)
  usersMake = append(usersMake, "TEst")
  fmt.Println(users22[1])
  fmt.Println(users)
  fmt.Println(usersMake)
  
  initialUsers := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
  users1 := initialUsers[2:6]
  users2 := initialUsers[:4]
  users3 := initialUsers[3:]
  
  fmt.Println(users1)
  fmt.Println(users2)
  fmt.Println(users3)
  
  usersDel := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
  var n = 3
  usersDel = append(usersDel[:n], usersDel[n + 1 :]...)
  fmt.Println(usersDel)
}


