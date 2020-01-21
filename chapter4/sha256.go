package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("中"))
	fmt.Println(c1)
	fmt.Printf("%x",c1)
	}
