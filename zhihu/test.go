package main

import (
	"errors"
	"fmt"
)

func main() {
	i := 2
	if i > 1 {
		i, err := doDivision(i, 2)
		if err != nil {
			panic(err)
		}
		fmt.Println(i)
	}
	fmt.Println(i)
}

func doDivision(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("input is invalid")
	}
	return x / y, nil
}