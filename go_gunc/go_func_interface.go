package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}

func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + "," + r
}

type Sphere struct {
	Radius float64
}

func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}
func (s *Sphere) Volume() float64 {
	return (float64(4) / float64(3)) * math.Pi * (math.Pow(s.Radius, 3))
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) changeBase(f float64) {
	t.base = f
	return
}
func (t *Triangle) changeBase1(f float64) {
	t.base = f
	return
}

type Robot interface {
	PowerOn() error
}
type T850 struct {
	Name string
}
type R2D2 struct {
	Broken bool
}

func (t *T850) PowerOn() error {
	return nil
}
func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}
func Boot(r Robot) error {
	return r.PowerOn()
}
func main() {
	t := T850{
		Name: "the terminator",
	}
	r := R2D2{
		Broken: true,
	}
	err := Boot(&r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("rebot is power on")
	}
	err = Boot(&t)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("robot is power on")
	}

	/* 	tri := Triangle{
	   		base:   5,
	   		height: 10,
	   	}
	   	tri.changeBase(4)
	   	fmt.Println(tri.base)
	   	tri.changeBase1(4)
	   	fmt.Println(tri.base) */

	/* 	s := Sphere{
	   		Radius: 5,
	   	}
	   	fmt.Println(s.SurfaceArea())
	   	fmt.Println(s.Volume()) */
	/* 	m := Movie{
	   		Name:   "Spiderman",
	   		Rating: 3.2,
	   	}
	   	fmt.Println(m.summary()) */
}
