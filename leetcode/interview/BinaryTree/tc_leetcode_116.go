package main

/**
填充二叉树


填充前
      1
    /   \
  2       3
 / \    /  \
4   5  6    7
      / \
     8   9

填充后
      1 ——> nil
    /   \
  2  ——> 3 ——> nil
 / \    /  \
4——>5  6 ——>7 ——> nil
      / \
     8——>9 ——> nil
*/

type Tree struct {
	data  string
	left  *Tree
	right *Tree
	next  *Tree
}

var tree1 = &Tree{data: "1", left: tree2, right: tree3}
var tree2 = &Tree{data: "2", left: tree4, right: tree5}
var tree3 = &Tree{data: "3", left: tree6, right: tree7}
var tree4 = &Tree{data: "4", left: nil, right: nil}
var tree5 = &Tree{data: "5", left: nil, right: nil}
var tree6 = &Tree{data: "6", left: tree8, right: tree9}
var tree7 = &Tree{data: "7", left: nil, right: nil}
var tree8 = &Tree{data: "8", left: nil, right: nil}
var tree9 = &Tree{data: "9", left: nil, right: nil}

func main() {
	connect(tree1)
}

func connect(root *Tree) *Tree {
	if root == nil {
		return nil
	}
	connect_two_node(root.left, root.right)
	return root
}

func connect_two_node(no1, no2 *Tree) {
	if no1 == nil || no2 == nil {
		return
	}

	no1.next = no2
	connect_two_node(no1.left, no1.right)
	connect_two_node(no2.left, no2.right)
}
