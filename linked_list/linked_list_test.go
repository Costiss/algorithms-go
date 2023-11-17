package linkedlist_test

import (
	linkedlist "algos/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getList() *linkedlist.DoublyLinkedList {
	list := linkedlist.NewDoublyLinkedList()

	list.Append(0)

	list.Append(1)

	list.Append(2)

	return list
}

func TestAppend(t *testing.T) {
	t.Run("should add node to the end", func(t *testing.T) {
		list := getList()

		oldLast := list.GetLastNode()

		last := list.Append(9)

		assert.Equal(t, list.GetLastNode(), last)
		assert.Equal(t, oldLast, last.Prev)
	})

	t.Run("should not add length if node is nil", func(t *testing.T) {
		list := getList()
		listLength := list.Length

		node := list.Append(nil)

		assert.Nil(t, node)
		assert.EqualValues(t, listLength, list.Length)
	})

	t.Run("should replace head if list is empty", func(t *testing.T) {
		list := linkedlist.NewDoublyLinkedList()

		node := list.Append(0)

		assert.Equal(t, node, list.Head)
		assert.EqualValues(t, list.Length, 1)
	})
}

func TestGetLastNode(t *testing.T) {
	list := getList()

	last := list.GetLastNode()

	assert.EqualValues(t, last.Value, 2)
}

func TestFind(t *testing.T) {

	t.Run("should find existing value", func(t *testing.T) {
		list := getList()

		found := list.Find(1)

		assert.EqualValues(t, 1, found.Value)
	})

	t.Run("should return nil if value not in the list", func(t *testing.T) {
		list := getList()

		found := list.Find(9)

		assert.Nil(t, found)
	})

	t.Run("should return nil if head nil", func(t *testing.T) {
		list := linkedlist.NewDoublyLinkedList()

		found := list.Find(1)

		assert.Nil(t, found)
	})
}

func TestPrepend(t *testing.T) {
	t.Run("should add node to the beginning", func(t *testing.T) {
		list := getList()

		oldHead := list.Head

		node := list.Prepend(9)

		assert.Equal(t, node, list.Head)
		assert.Equal(t, oldHead, node.Next)
		assert.Equal(t, oldHead.Prev, node)
		assert.EqualValues(t, list.Length, 4)
	})

	t.Run("should replace head if list is empty", func(t *testing.T) {
		list := linkedlist.NewDoublyLinkedList()

		node := list.Prepend(0)

		assert.Equal(t, node, list.Head)
	})

	t.Run("should not add length if node is nil", func(t *testing.T) {
		list := getList()
		listLength := list.Length

		node := list.Prepend(nil)

		assert.Nil(t, node)
		assert.EqualValues(t, listLength, list.Length)
	})
}
