package LinkedList

//链表反转 1->2->3 反转为 3->2->1
func (l *LinkedList) Reverse(node *LinkedListNode) *LinkedListNode {
	if node == nil {
		return node
	}
	var prev *LinkedListNode = nil

	for {
		if node == nil {
			break
		}
		next := node.next
		node.next = prev
		prev = node
		node = next
	}

	return prev
}
