package main

/**
反转二叉树


反转前
      1
    /   \
  2       3
 / \    /  \
4   5  6    7
      / \
     8   9

反转后
      1
    /   \
  3       2
 / \    /  \
7   6  5    4
   / \
  9   8
*/

func main() {
	reverse(node1)
}

func reverse(node *Node) {
	if node == nil {
		return
	}

	tmp := node.left
	node.left = node.right
	node.right = tmp

	reverse(node.left)
	reverse(node.right)
}
