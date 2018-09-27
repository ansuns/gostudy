package tree

import "fmt"

/***
结构是创建在栈上的，不是堆上
 */
type Node struct {
	Value int
	//*Node 是一个指针类型，有星号.
	//left right 是一个指针类型，指向左右两个Node
	Left, Right *Node
}

/*
返回结构的一个地址，是一个局部变量
 */
func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

/**
为结构定义方法，是写在外面，有一个接收者:(node Node)
 */

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}
func (node *Node) SetValue(Value int) {
	node.Value = Value
}

/**
遍历函数
 */
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	//中序遍历
	//先左再右
	node.Left.Traverse() //遍历左子树
	node.Print() // 遍历自己
	node.Right.Traverse() //遍历右子树
}
