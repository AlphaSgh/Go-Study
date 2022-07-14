package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "slowFunc() finished"
}

func receiver(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}
func pinger(c chan string) {
	t := time.NewTicker(time.Second * 1)
	for {
		c <- "ping"
		<-t.C
	}
}
func ping1(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "ping on channel1"
}
func ping2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "ping on channel2"
}

func sender(c chan string) {
	t := time.NewTicker(time.Second * 1)
	for {
		c <- "i am sending a message"
		<-t.C
	}
}
func main() {
	messages := make(chan string)
	stop := make(chan bool)
	go sender(messages)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("time is up")
		stop <- true
	}()
	for {
		select {
		case <-stop:
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}

	/* 	channel1 := make(chan string)
	   	channel2 := make(chan string)

	   	go ping1(channel1)
	   	go ping2(channel2)

	   	select {
	   	case msg1 := <-channel1:
	   		fmt.Println("received", msg1)
	   	case msg2 := <-channel2:
	   		fmt.Println("received", msg2)
	   	case <-time.After(time.Millisecond * 500):
	   		fmt.Println("no messages received giving up")
	   	} */

	/* 	messages := make(chan string)
	   	go pinger(messages)
	   	for {
	   		msg := <-messages
	   		fmt.Println(msg)
	   	} */

	/* 	messages := make(chan string, 2)
	   	messages <- "hello"
	   	messages <- "world"
	   	close(messages)
	   	fmt.Println("pushed two messages to channel with no receivers")
	   	time.Sleep(time.Second * 1)
	   	receiver(messages) */

	/* 	c := make(chan string)
	   	go slowFunc(c)

	   	msg := <-c
	   	fmt.Println(msg) */
}
