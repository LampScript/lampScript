package main

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}



func main() {
	a := new(map[string]string)
	(*a)["a"] = "1"
	fmt.Println(a)
	//fmt.Println(levelOrder(&TreeNode{Val:3, Right:&TreeNode{Val:20, Left:&TreeNode{Val:15}, Right:&TreeNode{Val:7}},Left:&TreeNode{Val:9}}))
}

func levelOrder(root *TreeNode) [][]int {
	return Solve(root, make([][]int, 0, 0), 0)
}

func Solve(t *TreeNode, array [][]int, height int) [][]int {
	if t == nil {
		return array
	}
	if len(array) < height + 1 {
		array = append(array, make([]int,0,0))
	}
	tmp := array[height]
	tmp = append(tmp, t.Val)
	array[height] = tmp

	if t.Left != nil {
		array = Solve(t.Left, array, height+1)
	}

	if t.Right != nil {
		array = Solve(t.Right, array, height+1)
	}
	return array
}