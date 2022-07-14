package main

import (
	"bytes"
	"fmt"
	"strings"
)

func useStr() {
	s1 := "after a backslash, certain single character escapes represent\nn is a line feed or new line \n\t t is a tab"
	s2 := `after a backslash, certain single character escapes \n tepresent
special values
	t is a tab`
	fmt.Println(s1)
	fmt.Println(s2)
}

//使用加号拼接字符串，少量拼接使用+ / +=效果好，但是拼接次数增加后效率不高
//如果需要在循环中拼接字符串可以使用字节缓冲区效率更高
func useStr1() {
	var buffer bytes.Buffer
	for i := 0; i < 500; i++ {
		buffer.WriteString("z")
	}
	fmt.Println(buffer.String())
}
func stringFuncs() {
	str := " We Are Friend"
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.Index(str, "Are"))
	fmt.Println(strings.Index(str, "www"))
	fmt.Println(strings.TrimSpace(str))

}
func main() {
	stringFuncs()
	// useStr1()
	// useStr()
}
