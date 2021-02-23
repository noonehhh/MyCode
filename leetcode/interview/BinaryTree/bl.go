package main

import "fmt"

/**
遍历二叉树
前序遍历  根左右  一直访问左子树，直到没有左子树，然后访问右子树
中序遍历  左根右
后续遍历  左右根

遍历二叉树
      1
    /   \
  2       3
 / \    /  \
4   5  6    7
      / \
     8   9

前序输出: 1 2 4 5 3 6 8 9 7
中序输出: 4 2 5 1 8 6 9 3 7
后序输出: 4 5 2 8 9 6 7 3 1
层序输出: 1 2 3 4 5 6 7 8 9
*/

type Node struct {
	data  string
	left  *Node
	right *Node
}

var node1 = &Node{data: "1", left: node2, right: node3}
var node2 = &Node{data: "2", left: node4, right: node5}
var node3 = &Node{data: "3", left: node6, right: node7}
var node4 = &Node{data: "4", left: nil, right: nil}
var node5 = &Node{data: "5", left: nil, right: nil}
var node6 = &Node{data: "6", left: node8, right: node9}
var node7 = &Node{data: "7", left: nil, right: nil}
var node8 = &Node{data: "8", left: nil, right: nil}
var node9 = &Node{data: "9", left: nil, right: nil}

func main() {
	reverse(node1)
	pre1(node1)
}

//前序 递归
func pre1(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	pre1(node.left)
	pre1(node.right)
}

//前序 非递归
func pre2(node *Node) {
	var arr []string
	var stack []*Node
	if node == nil {
		return
	}
	for node != nil || len(stack) != 0 {
		if node != nil {
			arr = append(arr, node.data)
			stack = append(stack, node)
			node = node.left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = node.right
		}
	}
	fmt.Println(arr)
}

//中序 递归
func mid1(node *Node) {
	if node == nil {
		return
	}
	mid1(node.left)
	fmt.Println(node.data)
	mid1(node.right)
}

//中序 非递归
func mid2(node *Node) {
	var arr []string
	var stack []*Node
	for node != nil || len(stack) != 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.left
		} else {
			node = stack[len(stack)-1]
			arr = append(arr, node.data)
			stack = stack[:len(stack)-1]
			node = node.right
		}
	}
	fmt.Println(arr)
}

//后序 递归
func last1(node *Node) {
	if node == nil {
		return
	}
	last1(node.left)
	last1(node.right)
	fmt.Println(node.data)

}

//后序 非递归
func last2(node *Node) {

}

//层序遍历
func level(node *Node) {

}
