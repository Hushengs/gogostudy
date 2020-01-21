package main

import (
	"fmt"
	"gogostudy/tree"
)

//内嵌Embedding
type myTreeNode struct {
	*tree.Node
}

//组合包装
//type myTreeNode struct {
//	node *tree.Node
//}

//重命名 后续遍历
func (myNode * myTreeNode) postOrder(){
	if myNode == nil || myNode.Node==nil{
		return
	}
	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse()  {
	fmt.Println("this method is shadowed")
}

func main() {
	//var root tree.Node
	root := myTreeNode{&tree.Node{Value:3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5,nil,nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	root.Node.Traverse()
	fmt.Println()
	root.postOrder()
	fmt.Println()

}
