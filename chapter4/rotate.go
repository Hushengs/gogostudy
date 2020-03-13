package main

import "fmt"

func rotate(s []int,r int)[]int{
	lens := len(s)
	//创建一个空的指定长度的slice
	res :=make([]int, lens)
	for i:=0 ;i<lens;i++{
		index := i+r
		if index>=lens{
			index=index-lens
		}
		res[i] = s[index]
	}
	return res
}

func main() {
	s := []int{1,3,5,7,9}
	fmt.Println(rotate(s,1))

}
