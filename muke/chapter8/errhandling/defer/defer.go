package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer()  {
	defer fmt.Println(1)
	fmt.Println(2)
}

func writeFile(fileName string)  {
	file,err:=os.Create(fileName)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i:=0 ;i<20;i++{
		fmt.Fprintln(writer,i)
	}
}

func main() {
	tryDefer()
	writeFile("bufiotest")
}
