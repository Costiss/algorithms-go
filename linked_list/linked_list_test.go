package linkedlist_test

import (
	linkedlist "algos/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

var second = linkedlist.ListNode{
	Value: 2,
}

var first = linkedlist.ListNode{
	Value: 1,
	Next:  &second,
}

var head = linkedlist.ListNode{
	Value: 0,
	Next:  &first,
}

func TestGetLastNodet(t *testing.T) {
	last := linkedlist.GetLastNode(&head)

	assert.Equal(t, last, &second)
}

func TestFind(t *testing.T) {

	t.Run("should find existing value", func(t *testing.T) {

		found := linkedlist.Find(&head, 1)

		assert.Equal(t, &first.Value, found)
	})

	t.Run("should return nil if value not in the list", func(t *testing.T) {

		found := linkedlist.Find(&head, 9)

		assert.Nil(t, found)
	})

	t.Run("should return nil if head nil", func(t *testing.T) {

		found := linkedlist.Find(nil, 9)

		assert.Nil(t, found)
	})
}
