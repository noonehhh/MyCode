package main

import "fmt"

/**
二叉树转换为链表
转换前
      1
    /   \
  2       3
 / \    /  \
4   5  6    7
      / \
     8   9

转换后
1-2-4-5-3-6-8-9-7
*/

type TreeNode struct {
	data  string
	left  *TreeNode
	right *TreeNode
}

var treenode1 = &TreeNode{data: "1", left: treenode2, right: treenode3}
var treenode2 = &TreeNode{data: "2", left: treenode4, right: treenode5}
var treenode3 = &TreeNode{data: "3", left: treenode6, right: treenode7}
var treenode4 = &TreeNode{data: "4", left: nil, right: nil}
var treenode5 = &TreeNode{data: "5", left: nil, right: nil}
var treenode6 = &TreeNode{data: "6", left: treenode8, right: treenode9}
var treenode7 = &TreeNode{data: "7", left: nil, right: nil}
var treenode8 = &TreeNode{data: "8", left: nil, right: nil}
var treenode9 = &TreeNode{data: "9", left: nil, right: nil}

func main() {
	flatten2(treenode1)
	a := treenode1
	fmt.Println(a)
}

var arr []string

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	arr = append(arr, root.data)
	flatten(root.left)
	flatten(root.right)
	fmt.Println(arr)
}

func flatten2(root *TreeNode) {
	if root == nil {
		return
	}

	flatten2(root.left)
	flatten2(root.right)

	right := root.right
	root.left, root.right = nil, root.left

	p := root
	for p.right != nil {
		p = p.right
	}
	p.right = right
}
