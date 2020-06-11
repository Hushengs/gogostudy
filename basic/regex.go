package main

import (
	"fmt"
	"regexp"
)

const text = `My email is yym@163.com
			  My email is hk@168.com.cn
			  My email is hushengs@168.com.cn
			  My email is conan@gmail.com.cn
			`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)
	//re,err := regexp.Compile("yym@163.com")
	//match := re.FindString(text)
	//match := re.FindAllString(text,-1)
	match := re.FindAllStringSubmatch(text,-1)
	fmt.Println(match)
}
