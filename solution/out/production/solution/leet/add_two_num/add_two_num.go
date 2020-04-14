package main

// https://leetcode.com/problems/add-two-numbers/submissions/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	l1 := ListNode{0, nil}

	r2 := ListNode{6, nil}
	r1 := ListNode{4, &r2}

	ans := addTwoNumbers(&l1, &r1)
	fmt.Println(ans.Val)
	ans = ans.Next
	fmt.Println(ans.Val)
	ans = ans.Next

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := ListNode{}
	crNode := &ans
	var prevNode *ListNode
	var carry int
	var temp1 int
	var temp2 int

	for l1 != nil || l2 != nil {
		if prevNode == nil {
			crNode = &ans
		} else {
			prevNode.Next = &ListNode{}
			crNode = prevNode.Next
		}
		temp1 = 0
		temp2 = 0
		if l1 == nil {
			temp2 = l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			temp1 = l1.Val
			l1 = l1.Next
		} else {
			temp1 = l1.Val
			temp2 = l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}
		crNode.Val = (temp1 + temp2 + carry) % 10
		carry = (temp1 + temp2 + carry) / 10

		prevNode = crNode
	}
	if carry > 0 {
		prevNode.Next = &ListNode{}
		crNode = prevNode.Next
		crNode.Val = carry
	}

	return &ans
}
