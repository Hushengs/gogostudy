package main

import (
	"fmt"
	"gogostudy/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5,nil,nil}
	root.Right.Left = new(tree.Node)
	root.Right.Left = tree.CreateNode(2)
	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	node := []tree.Node{
		{Value:3},
		{},
		{6,nil,&root},
	}
	fmt.Println(node)

	//地址可以取值或地址
	Proot :=&root
	Proot.Print()
	Proot.SetValue(200)
	Proot.Print()
	fmt.Println()
	root.Traverse()
}
