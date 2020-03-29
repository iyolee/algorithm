package main

// LNode 链表节点
type LNode struct {
	Data interface{}
	Next *LNode
}

// Reverse 带头节点的逆序
func Reverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var pre *LNode
	var cur *LNode
	next := node.Next
	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

func main() {
	head := &LNode{}
	CreateNode(head, 8)
	PrintNode("逆序前：", head)
	Reverse(head)
	PrintNode("逆序后：", head)
}
