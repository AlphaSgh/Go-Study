package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func sayHello(s string) string {
	return "hello " + s
}

func addition(x, y int) int {
	return x + y
}

func testBool() {
	var b bool
	b = true
	fmt.Println(b)
}

func testReflect() {
	var s string = "string"
	var i int = 10
	var f float32 = 3.14

	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
}

func testStrconv() {
	var s string = "true"
	b, err := strconv.ParseBool(s)
	fmt.Println(b)
	fmt.Println(err)
	var ss = strconv.FormatBool(true)
	fmt.Printf("ss: %v\n", ss)
}
func isEvn(x int) bool {
	return x%2 == 0
}
func getPrize() (int, string) {
	i := 2
	str := "hello"
	return i, str
}
func sumNums(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
func sayHi() (a, b string) {
	a := "hello"
	b := "world"
	return
}
func main() {
	// fmt.Println(sayHello(" sgh"))
	// fmt.Println(addition(3, 4))
	// testBool()
	// testReflect()
	// testStrconv()
	// fmt.Printf("%v\n%v\n", isEvn(1), isEvn(2))
	// num, str := getPrize()
	// fmt.Printf("you won %v %v\n", num, str)
	// fmt.Printf("the sum = %v\n", sumNums(1, 2, 3, 4))
	fmt.Printf(sayHi())
}
