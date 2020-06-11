package main

import "fmt"

func main() {
	//testSlice()
	testSliceCap()
}


func testSlice(){
	s := []int64{1,2,3,4}
	sc :=s[2:]
	fmt.Println(s)
	fmt.Println(sc)
	sc[0] = 6
	fmt.Println(s)
	fmt.Println(sc)
}

func testSliceCap(){
	s := make([]int64,6,6)
	s[5] = 6
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s = append(s,7,8,9)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func testCopy(){
	s1
	s := make([]int64,6,6)
	copy(s,s1)
}