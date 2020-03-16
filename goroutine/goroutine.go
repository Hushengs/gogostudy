package main

import (
	"fmt"
	"runtime"
	"time"
)

//协程
func main() {
	var a [10]int
	for i:= 0;i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				//主动交出控制器
				runtime.Gosched()
				//fmt.Printf("Hello from"+"goroutine %d\n", i)
			}
		}(i)
	}

	time.Sleep(time.Microsecond)
	fmt.Println(a)
}
