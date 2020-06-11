package main

import (
	"fmt"
)

type people struct {
	name string
	age int64
}

func (p *people) toString(){
	fmt.Println(p.name)
	fmt.Println("AAA")
	fmt.Printf("地址 %p",p)
}

func (p people) tooString(){
	fmt.Println(p.name)
	fmt.Println("BBB")
	fmt.Printf("地址 %p",&p)
}

func main() {
	//p1 := people{"hushengs"}
	//p1.toString()
	//p1.tooString()
	//p2 :=&people{"hk"}
	//p2.toString()
	//p2.tooString()
	p3 := people{}
	fmt.Println(p3)
	p3.name="conan"
	fmt.Println(p3)
	p4 := new(people)
	fmt.Println(p4)
	p4.name = "hushengs"
	fmt.Println(p4)
}
