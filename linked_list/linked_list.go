package linkedlist

type ListNode struct {
	Value int
	Next  *ListNode
}

func Find(head *ListNode, valueToFind int) (value *int) {
	if head == nil {
		return nil
	}

	if head.Value == valueToFind {
		return &head.Value
	}

	if head.Next == nil {
		return nil
	}

	head = head.Next

	return Find(head, valueToFind)

}

func AddNode(head *ListNode, newNode *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	last := GetLastNode(head)

	last.Next = newNode

	return newNode
}

func ChangeHead(head *ListNode, newHead *ListNode) *ListNode {
	if head == nil {
		return newHead
	}
	if newHead == nil {
		return head
	}

	newHead.Next = head

	return newHead
}

func GetLastNode(startingNode *ListNode) *ListNode {
	if startingNode == nil {
		return nil
	}

	if startingNode.Next == nil {
		return startingNode
	}

	startingNode = startingNode.Next

	return GetLastNode(startingNode)
}
