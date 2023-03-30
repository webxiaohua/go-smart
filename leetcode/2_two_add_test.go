/*
2.两数相加
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/
package main

import (
	"fmt"
	"testing"
)

func TestTwoAdd(t *testing.T) {
	l1 := &LinkedListNode{
		Next: &LinkedListNode{
			Next: &LinkedListNode{
				Next: nil,
				Val:  3,
			},
			Val:  4,
		},
		Val:  2,
	}
	l2 := &LinkedListNode{
		Next: &LinkedListNode{
			Next: &LinkedListNode{
				Next: nil,
				Val:  4,
			},
			Val:  6,
		},
		Val:  5,
	}
	res := twoAdd(l1,l2)
	fmt.Println(res.ToJson())
}

/*
实现思路：
定义两个指针 p,q 分别指向两个链表的头节点
定义 cur 指向返回链表的当前值
定义 carry 用于进位
循环：
	若p,q同时为nil，准备退出，退出前处理 carry，若carry大于0，需要再进一位;
	计算当前值公式： sum = x + y + carry
	cur.Val = sum % 10
	carry = sum / 10
	p,q 指针继续前进
	cur 指针前进
*/
func twoAdd(l1,l2 *LinkedListNode) (res *LinkedListNode){
	res = &LinkedListNode{}
	cur := res
	carry := 0
	p := l1
	q := l2
	for {
		if p==nil && q==nil {
			if carry > 0 {
				cur.Next = &LinkedListNode{
					Next: nil,
					Val:  carry,
				}
			}
			break
		}
		x,y := 0,0
		if p!= nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		sum := x + y + carry
		cur.Next = &LinkedListNode{
			Next: nil,
			Val:  sum % 10,
		}
		cur = cur.Next
		carry = sum / 10
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	return res.Next
}