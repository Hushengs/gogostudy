package main

import "fmt"

func main() {
	ary := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(ary)
}
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
