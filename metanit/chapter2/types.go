package main

import "fmt"

var _int8 int8   // -128 128                     1B
var _int16 int16 // -32768 32767                 2B
var _int32 int32 // -2147483648 2147483647       4B
var _int64 int64 // –9 223 372 036 854 775 808   8B
//  9 223 372 036 854 775 807
var _uint8 uint8   // 0 255                        1B
var _uint16 uint16 // 0 65535                      2B
var _uint32 uint32 // 0 4294967295                 4B
var _uint64 uint64 // 18 446 744 073 709 551 615   8B

var _byte byte // синоним uint8 0 255          1B
var _rune rune // синоним int32                4B
// -2147483648
// 2147483647
var _int int  // в зависимости от платформы может занимать 4B 8B
var _uint int // без знаковое - в зависимости от платформы может занимать 4B 8B

var _float32 float32 // 1.4*10 -45 -- 3.4*10 38      4B
var _float64 float64 // 4.9*10 -324 -- 1.8*10 308    8B

var _complex64 complex64   // вещественная и мнимая части представляют числа float32
var _complex128 complex128 // вещественная и мнимая части представляют числа float64

var _bool bool // true || false

var _string string // Строка

// Константы
// const имя_константы тип_константы
const pi float64 = 3.14

func main() {
	fmt.Println("types")
}
