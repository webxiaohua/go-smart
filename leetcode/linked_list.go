package main

import "encoding/json"

// 单链表
type LinkedListNode struct {
	Next *LinkedListNode
	Val int
}

func (s *LinkedListNode) ToJson() string {
	b,_ :=json.Marshal(s)
	return string(b)
}