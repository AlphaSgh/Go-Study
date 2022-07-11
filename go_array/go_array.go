package main

import (
	"fmt"
)

func testArray() {
	var cheeses [2]string
	cheeses[0] = "marry"
	cheeses[1] = "deffer"
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	fmt.Println(cheeses)
}
func testSlicing() {
	var cheeses = make([]string, 2)
	cheeses[0] = "alice"
	cheeses[1] = "bolb"
	cheeses = append(cheeses, "clice")
	cheeses = append(cheeses, "a", "b", "c", "d")
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
}
func testCopySlicing() {
	var cheeses = make([]string, 3)
	cheeses[0] = "A"
	cheeses[1] = "B"
	var smellyCheeses = make([]string, 2)
	// copy(smellyCheeses, cheeses)
	copy(smellyCheeses, cheeses[1:2])
	fmt.Println(smellyCheeses)
}
func testMap() {
	var players = make(map[string]int)
	players["cook"] = 31
	players["bairstorw"] = 43
	players["stokes"] = 23
	fmt.Println(players["cook"])
	fmt.Println(players)
	delete(players, "cook")
	fmt.Println(players)
}
func main() {
	// testArray()
	// testSlicing()
	// testCopySlicing()
	testMap()
}
