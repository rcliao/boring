package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s said %d", name, i)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
			c <- <-input1
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Eric"), boring("Gopher"))
	fmt.Println("I'm listening")
	timeout := time.After(5 * time.Second)
	for i := 0; ; i++ {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("5 seconds pass. I hope you have a great dream!")
			return
		}
	}
}
