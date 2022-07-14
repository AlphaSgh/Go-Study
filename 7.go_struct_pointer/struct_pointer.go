package main

import (
	"fmt"
)

type Movie struct {
	Name string
	Age  int64
}
type Address struct {
	Number int
	Street string
	City   string
}
type SuperHero struct {
	Name    string
	Age     int
	Address Address
}

type Alarm struct {
	Time  string
	Sound string
}

func NewAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: "Klaxon",
	}
	return a
}
func main() {
	//以值引用的方式复制结构体
	/* 	a := Alarm{
	   		Time:  "8:00",
	   		Sound: "china",
	   	}
	   	b := a
	   	b.Sound = "china+china"
	   	fmt.Printf("%+v\n", a)
	   	fmt.Printf("%+v\n", b)
	   	fmt.Printf("%p\n", &a)
	   	fmt.Printf("%p\n", &b) */
	//以指针引用的方式复制结构体
	c := Alarm{
		Time:  "11:11",
		Sound: "alice",
	}
	d := &c
	d.Sound = "alice + alice"
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", *d)
	fmt.Printf("%p\n", &c)
	fmt.Printf("%p\n", d)

	// a := NewAlarm("08:11")
	// b := NewAlarm("08:11")
	// if a == b {
	// 	fmt.Println("a and b are the same")
	// }
	// fmt.Println(a)
	// fmt.Println(reflect.TypeOf(a))
	// fmt.Println(reflect.TypeOf(b))
	// alarm := NewAlarm("07:00")
	// fmt.Println(alarm)
	/* 	hero := SuperHero{
	   		Name: "Hero",
	   		Age:  18,
	   		Address: Address{
	   			Number: 10086,
	   			Street: "Mountain Drive",
	   			City:   "China",
	   		},
	   	}
	   	fmt.Println(hero)
	   	fmt.Println(hero.Address.City) */
	/* 	var m1 Movie
	   	fmt.Printf("%+v\n", m1)
	   	m1.Name = "bob"
	   	m1.Age = 23
	   	fmt.Printf("%+v\n", m1)
	   	m := Movie{
	   		Name: "alice",
	   		Age:  18,
	   	}
	   	fmt.Println(m.Name, m.Age) */
}
