package main

import "fmt"

func main() {
	var x, y int
	p :=&x
	p1 :=&x
	fmt.Println(&x == &x, &x == &y, p  == p1)
	fmt.Println(p)
}
