package main

import "project_google_gostudy/contaniner/tree"

func main() {
	var root tree.Node

	root = tree.Node{Value:3} //3
	root.Left = &tree.Node{} //0
	root.Right = &tree.Node{5, nil, nil} //5
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)


	root.Traverse()
	//root.right.left.print()
	//fmt.Println()
}
