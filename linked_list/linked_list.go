package linkedlist

type DoublyLinkedList struct {
	Head   *ListNode
	Length int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head:   nil,
		Length: 0,
	}
}

type ListNode struct {
	Value any
	Prev  *ListNode
	Next  *ListNode
}

func (list *DoublyLinkedList) Find(value int) *ListNode {
	if list.Head == nil {
		return nil
	}

	current := list.Head

	for i := 0; i < list.Length; i++ {
		if current.Value == value {
			return current
		}

		if current.Next == nil {
			return nil
		}

		current = current.Next
	}

	return nil
}

func (list *DoublyLinkedList) Append(value any) *ListNode {
	if value == nil {
		return nil
	}
	node := &ListNode{Value: value}

	list.Length += 1

	if list.Head == nil {
		list.Head = node
		return node
	}

	last := list.GetLastNode()

	last.Next = node
	node.Prev = last

	return node
}

func (list *DoublyLinkedList) Prepend(value any) *ListNode {
	if value == nil {
		return nil
	}
	node := &ListNode{Value: value}

	list.Length += 1

	if list.Head == nil {
		list.Head = node
		return node
	}

	node.Next = list.Head
	list.Head.Prev = node
	list.Head = node

	return node
}

func (list *DoublyLinkedList) GetLastNode() *ListNode {
	if list.Head == nil {
		return nil
	}

	current := list.Head

	for i := 0; i < list.Length; i++ {
		if current.Next == nil {
			return current
		}

		current = current.Next
	}

	return nil
}
