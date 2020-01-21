package main

import "fmt"

func main() {
	//var x, y int
	//p :=&x
	//p1 :=&x
	//fmt.Println(&x == &x, &x == &y, p  == p1)
	//fmt.Println(p)

	//freq := rand.Float64() * 3.0
	//fmt.Println(freq)

	var x,y = 1,3
	fmt.Println(&x == &x, &x == &y, &x == nil)
	fmt.Println(&x)
	fmt.Println(&y)
}
