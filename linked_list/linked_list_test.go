package linkedlist_test

import (
	linkedlist "algos/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getHead() linkedlist.ListNode {
	var second = linkedlist.ListNode{
		Value: 2,
	}

	var first = linkedlist.ListNode{
		Value: 1,
		Next:  &second,
	}

	return linkedlist.ListNode{
		Value: 0,
		Next:  &first,
	}
}

func TestGetLastNodet(t *testing.T) {

	head := getHead()

	last := linkedlist.GetLastNode(&head)

	assert.Equal(t, last.Value, 2)
}

func TestFind(t *testing.T) {

	t.Run("should find existing value", func(t *testing.T) {
		head := getHead()

		found := linkedlist.Find(&head, 1)

		assert.Equal(t, 1, *found)
	})

	t.Run("should return nil if value not in the list", func(t *testing.T) {
		head := getHead()

		found := linkedlist.Find(&head, 9)

		assert.Nil(t, found)
	})

	t.Run("should return nil if head nil", func(t *testing.T) {
		found := linkedlist.Find(nil, 9)

		assert.Nil(t, found)
	})
}

func TestPrepend(t *testing.T) {
	t.Run("should add node to the beginning", func(t *testing.T) {
		head := getHead()

		node := linkedlist.ListNode{
			Value: 9,
		}

		newHead := linkedlist.Prepend(&head, &node)

		assert.Equal(t, &node, newHead)
		assert.Equal(t, &head, newHead.Next)
	})
}

func TestAppend(t *testing.T) {
	t.Run("should add node to the end", func(t *testing.T) {
		head := getHead()

		node := linkedlist.ListNode{
			Value: 9,
		}

		linkedlist.Append(&head, &node)

		last := linkedlist.GetLastNode(&head)

		assert.Equal(t, &node, last)
	})
}
