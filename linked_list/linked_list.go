package linkedlist

type ListNode struct {
	Value int
	Prev  *ListNode
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

func Append(head *ListNode, newNode *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	last := GetLastNode(head)

	last.Next = newNode
	newNode.Prev = last

	return head
}

func Prepend(head *ListNode, newNode *ListNode) *ListNode {
	if head == nil {
		return newNode
	}
	if newNode == nil {
		return head
	}

	newNode.Next = head
	head.Prev = newNode

	return newNode
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
