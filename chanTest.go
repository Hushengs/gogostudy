package main

import "fmt"

func main() {
	var c = make(chan int)
	var a string

	go func() {
		a = "hello world"
		<-c
	}()

	c <- 0
	fmt.Println(a)
}