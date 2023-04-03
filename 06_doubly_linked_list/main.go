package main

import (
	"fmt"
)

type Item struct {
	previous *Item
	next     *Item
	Value    interface{}
}

func (item Item) GetValue() interface{} {
	return item.Value
}

func (item Item) Previous() Item {
	previous := item.previous
	if previous == nil {
		return Item{}
	}
	return *previous
}

func (item Item) Next() Item {
	next := item.next
	if next == nil {
		return Item{}
	}
	return *next
}

type List struct {
	Items []Item
}

func (list *List) First() Item {
	if list.Len() == 0 {
		return Item{}
	}
	return list.Items[0]
}

func (list *List) Last() Item {
	if list.Len() == 0 {
		return Item{}
	}
	return list.Items[list.Len()-1]
}

func (list *List) Len() int {
	return len(list.Items)
}

func (list *List) PushFront(val interface{}) {
	var next *Item
	if list.Len() > 0 {
		next = &list.Items[0]
	}
	item := Item{
		nil, next, val,
	}
	if list.Len() > 0 {
		list.Items[0].previous = &item
	}
	list.Items = append([]Item{item}, list.Items...)
}

func (list *List) PushBack(val interface{}) {
	var prev *Item
	if list.Len() > 0 {
		prev = &list.Items[list.Len()-1]
	}
	item := Item{
		prev, nil, val,
	}
	if list.Len() > 0 {
		list.Items[list.Len()-1].next = &item
	}
	list.Items = append(list.Items, item)
}

func (list *List) Remove(item *Item) {
	for idx, it := range list.Items {
		if *item == it {
			newSlice := make([]Item, 0)
			newSlice = append(newSlice, list.Items[:idx]...)
			list.Items = append(newSlice, list.Items[idx+1:]...)
		}
	}
}

func main() {
	list := List{}
	list.PushBack("middle")
	list.PushBack("last")
	list.PushFront("first")
	last := list.Last()
	mid1 := last.Previous().GetValue()
	first := list.First()
	mid2 := first.Next().GetValue()
	fmt.Printf("%v %v\n", mid1, mid2)
}
