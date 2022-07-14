package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func err1() {
	file, err := ioutil.ReadFile("foo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", file)
}
func err2() {
	err := errors.New("something was wrong")
	if err != nil {
		fmt.Println(err)
	}
}
func err3() {
	name, role := "richard jupp", "drummer"
	err := fmt.Errorf("the %v %v quit", role, name)
	if err != nil {
		fmt.Println(err)
	}
}
func Half(number int) (int, error) {
	if number%2 != 0 {
		return -1, fmt.Errorf("connot half %v", number)
	}
	return number / 2, nil
}
func main() {
	n, err := Half(14)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
	// err3()
	// err2()
	// err1()
}
