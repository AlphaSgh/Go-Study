package main

import "fmt"

type Movie struct {
	Name string
	Age  int64
}

func main() {
	var m1 Movie
	fmt.Printf("%+v\n", m1)
	m1.Name = "bob"
	m1.Age = 23
	fmt.Printf("%+v\n", m1)
	m := Movie{
		Name: "alice",
		Age:  18,
	}
	fmt.Println(m.Name, m.Age)
}
