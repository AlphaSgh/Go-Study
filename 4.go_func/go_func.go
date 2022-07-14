package main

import "fmt"

func isEven(i int) bool {
	return i%2 == 0
}
func getPrize() (int, string) {
	i := 2
	str := "hello"
	return i, str
}
func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func sayHi() (x, y string) {
	x = "hello"
	y = "world"
	return
}
func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten > 5 {
		fmt.Printf("i am full, i have eaten %d.\n", eaten)
		return eaten
	}
	fmt.Printf("i am still hungry! i have eaten %d.\n", eaten)
	return feedMe(portion, eaten)
}
func anotherFunction(f func() string) string {
	return f()
}
func main() {
	// fmt.Printf("%v\n", isEven(1))
	// fmt.Printf("%v\n", isEven(2))
	// i, str := getPrize()
	// fmt.Printf("you won %v %v\n", i, str)
	// fmt.Printf("1 2 3 4 = %v\n", sumNumbers(1, 2, 3, 4))
	// fmt.Println(sayHi())
	// feedMe(1, 0)
	fn := func() string {
		return "function called"
	}
	fmt.Println(anotherFunction(fn))
}
