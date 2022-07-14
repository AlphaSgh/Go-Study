package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished")
}
func main() {
	go slowFunc() //goroutine的使用只需要在相应的函数或方法前面加关键字go即可,goroutine立即返回

	fmt.Println("i am not shown until slowFunc() completes")
	time.Sleep(time.Second * 3)
}
