package main

import (
	"github.com/djq8888/goSingleList"
)

func addTwoNumbers(l1 *goSingleList.ListNode, l2 *goSingleList.ListNode) *goSingleList.ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	res := goSingleList.ListNode{(l1.Val + l2.Val) % 10, nil}
	prev := &res
	flag := (l1.Val + l2.Val) / 10
	l1 = l1.Next
	l2 = l2.Next
	for l1 != nil && l2 != nil {
		tmp := goSingleList.ListNode{(l1.Val + l2.Val + flag) % 10, nil}
		prev.Next = &tmp
		prev = &tmp
		flag = (l1.Val + l2.Val + flag) / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		tmp := goSingleList.ListNode{(l1.Val + flag) % 10, nil}
		prev.Next = &tmp
		prev = &tmp
		flag = (l1.Val + flag) / 10
		l1 = l1.Next
	}
	for l2 != nil {
		tmp := goSingleList.ListNode{(l2.Val + flag) % 10, nil}
		prev.Next = &tmp
		prev = &tmp
		flag = (l2.Val + flag) / 10
		l2 = l2.Next
	}
	if flag == 1 {
		prev.Next = &goSingleList.ListNode{1, nil}
	}
	return &res
}