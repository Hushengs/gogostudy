package tree

import "fmt"

type Node struct {
	Value int
	Left,Right *Node
}

//工厂函数
func CreateNode(value int ) *Node{
	return &Node{Value:value}
}

//输出
func (node Node) Print()  {
	fmt.Print(node.Value)
}

//func print(node treeNode)  {
////	fmt.Print(node.value)
////}

//传值
func (node *Node) SetValue(value int)  {
	if node == nil{
		fmt.Println("Setting Value to nil node.Ignored")
	}
	node.Value = value
}

//中序遍历
func (node *Node) Traverse(){
	if node == nil{
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}