package main

import "fmt"

func reverse(s []int){
	for i,j:=0,len(s)-1;i<j;i,j=i+1,j-1{
		s[i],s[j] = s[j],s[i]
	}
}

func reversePointer(s *[5]int) {
	for i,j:=0,len(*s)-1;i<j;i,j=i+1,j-1{
		(*s)[i],(*s)[j] = (*s)[j],(*s)[i]
	}
}

func reverses(s *[5]int){
	i,j := 0,len(*s)-1
	for i<j{
		(*s)[i],(*s)[j]=(*s)[j],(*s)[i]
		i+=1
		j-=1
	}
}

func main() {
	//s :=[]int{1,2,3,4,5}
	//reverse(s)
	//fmt.Println(s)

	arr :=[5]int{1,2,3,4,5}
	//reverses(&arr)
	reversePointer(&arr)
	fmt.Println(arr)
}
