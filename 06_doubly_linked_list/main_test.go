package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMiddleItem(t *testing.T) {
	list := List{}
	list.PushBack("middle")
	list.PushBack("last")
	list.PushFront("first")
	last := list.Last()
	mid1 := last.Previous().GetValue()
	first := list.First()
	mid2 := first.Next().GetValue()
	require.Equal(t, mid1, mid2, "test that middle is correct")
}

func TestNextOrPreviousOfSingleItem(t *testing.T) {
	list := List{}
	list.PushFront("single")
	last := list.Last()
	el := last.Next().GetValue()
	require.Nil(t, el, "test that next of last is nil")
	first := list.First()
	el = first.Previous().GetValue()
	require.Nil(t, el, "test that previous of first is nil")
}

func TestRemoveItem(t *testing.T) {
	list := List{}
	list.PushBack("middle")
	list.PushBack("last")
	list.PushFront("first")
	require.Equal(t, 3, list.Len(), "test that len is correct")
	last := list.Last()
	list.Remove(&last)
	newLast := list.Last()
	require.Equal(t, "middle", newLast.GetValue(), "test that middle is last after removal")
	require.Equal(t, "first", list.First().GetValue(), "test that first is still first")
	require.Equal(t, 2, list.Len(), "test that new len is correct")
}
