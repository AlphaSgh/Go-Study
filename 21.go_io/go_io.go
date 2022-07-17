package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func iofile() {
	fileBytes, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileBytes)
	fileString := string(fileBytes)
	fmt.Println(fileString)
}
func ioCreatFile() {
	// b := make([]byte, 0)
	b := []byte("hello, world")
	err := ioutil.WriteFile("example.txt", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func ioReadFile() {
	files, err := ioutil.ReadDir("../")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Mode(), file.Name())
	}
}
func ioCopy() {
	from, err := os.Open("./example05.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile("./example05.copy.txt", os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}
func ioRemove() {
	err := os.Remove("./example.txt")
	if err != nil {
		log.Fatal(err)
	}
}
func readCfgJson() {
	type Config struct {
		Name   string `json:"name"`
		Awake  bool   `json:"awake"`
		Hungry bool   `json:"hungry"`
	}
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	c := Config{}
	err = json.Unmarshal(f, &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", c)
}
func readCfgToml() {
	type Config struct {
		Name   string
		Awake  bool
		Hungry bool
	}
	c := Config{}
	_, err := toml.DecodeFile("config.toml", &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", c)

}
func main() {
	readCfgToml()
	// readCfgJson()
	// ioRemove()
	// ioCopy()
	// ioReadFile()
	// ioCreatFile()
	// iofile()
}
