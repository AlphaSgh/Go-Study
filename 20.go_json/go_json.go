package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

func json1() {
	hobbiles := []string{"Cycling", "Cheese", "Techno"}
	p := Person{
		Name:    "george",
		Age:     18,
		Hobbies: hobbiles,
	}
	fmt.Printf("%+v\n", p)
	jsonByteData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData := string(jsonByteData)
	fmt.Println(jsonStringData)
}

type Person2 struct {
	Name    string   `json:"name,omitempty"`
	Age     int      `json:"age,omitempty"`
	Hobbies []string `json:"hobbies,omitempty"`
}

func json2() {
	hobbies := []string{"Game", "Eat", "Sleep"}
	p := Person2{
		Name:    "Teemo",
		Age:     18,
		Hobbies: hobbies,
	}
	fmt.Printf("%+v\n", p)
	jsonByteData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData := string(jsonByteData)
	fmt.Printf("%+v\n", jsonStringData)
}
func json3() {
	jsonStringData := `{"name":"Teemo","age":18,"hobbies":["Game","Eat","Sleep"]}`
	jsonByteData := []byte(jsonStringData)
	P := Person2{}
	err := json.Unmarshal(jsonByteData, &P)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", P)
}
func main() {
	json3()
	// json2()
	// json1()
}
