package main

import "fmt"

type treeNode struct {
	value       int       //内容
	left, right *treeNode //左右指针
}

/**
ky可以自己定义工厂函数
 */
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

/**
打印NODE
nodexxx 是一个接收者
 */
func (nodexxx treeNode) printTreeNode() {
	fmt.Print(nodexxx.value, " ")
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

/**
遍历node
 */
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	//中式遍历：先左再中再右
	//先遍历子树，再遍历自己
	node.left.traverse()
	node.printTreeNode()
	node.right.traverse()
}

func main() {
	//var root treeNode
	root := treeNode{value: 3} //根是3
	root.left = &treeNode{}    //根的左边是0
	//按照顺序填充：right,left是指针，要取地址&
	root.right = &treeNode{5, nil, nil} //根的左边是0
	root.right.left = new(treeNode)     //根的右边的左边是0
	root.left.right = createNode(2)     //根的左边的右边是2

	//改值
	root.right.left.setValue(4)
	root.right.left.printTreeNode()
	root.printTreeNode()
	fmt.Println("-------------")
	root.traverse()
}
