package main

import "fmt"

func testSwitch() {
	i := 2
	switch i {
	case 2:
		fmt.Printf("TWO")
	case 3:
		fmt.Printf("THREE")
	case 4:
		fmt.Printf("FOUR")
	default:
		fmt.Printf("ERROR")
	}

}
func testFor() {
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Printf("i is %d\n", i)
	// }
	// for j := 0; j < 5; j++ {
	// 	fmt.Printf("j is %d\n", j)
	// }
	numbers := []int{3, 3, 4, 2}
	for i, n := range numbers {
		// fmt.Printf("the index of the loop is %d\n", i)
		// fmt.Printf("the value of the loop is %d\n", n)
		fmt.Printf("the value %d 's index is %d\n", n, i)
	}
}
func testDefer() {
	defer fmt.Println("i am the first defer")
	defer fmt.Println("i am the second defer")
	defer fmt.Println("i am the third defer")
	fmt.Println("testDefer")
}

func main() {
	testDefer()
	// testFor()
	// b := true
	// b = false
	// if b {
	// 	fmt.Printf("b is true")
	// } else {
	// 	fmt.Printf("b is false")
	// }
	// if !b {
	// 	fmt.Printf("b is false")
	// }
	// i := 2
	// if i == 3 {
	// 	fmt.Printf("i is 3")
	// } else if i == 2 {
	// 	fmt.Printf("i is 2")
	// }
	// testSwitch()

}
