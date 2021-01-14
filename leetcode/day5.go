package main

import "fmt"

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var newHead *ListNode = nil
	var arr []int

	for head != nil {
		node := head.Next
		head.Next = newHead
		newHead = head
		head = node
	}

	for newHead != nil {
		arr = append(arr, newHead.Val)
		newHead = newHead.Next
	}

	return arr
}

func reversrList2(head *ListNode) []int {
	cur := head
	var arr []int
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	}

	for pre != nil {
		arr = append(arr, pre.Val)
		pre = pre.Next
	}

	return arr
}

func reversePrint3(head *ListNode) []int {
	if head == nil {
		return nil
	}

	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}

func main() {
	head := &ListNode{}
	head.Val = 1
	ln1 := &ListNode{}
	ln1.Val = 3
	ln2 := &ListNode{}
	ln2.Val = 2

	head.Next = ln1
	ln1.Next = ln2

	arr := reversePrint(head)
	fmt.Println(arr)
}
