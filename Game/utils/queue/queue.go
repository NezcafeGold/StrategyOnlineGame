package queue

import (
	"container/list"
)

type Queue struct {
	list *list.List
}

// New - new Queue.
func New() *Queue {
	return &Queue{
		list: list.New(),
	}
}

// Add - adding a new object to the queue.
func (queue *Queue) Add(value interface{}) {
	queue.list.PushBack(value)
}

// Pop - getting and remove object to the queue.
func (queue *Queue) Pop() *list.Element {
	element := queue.list.Front()
	queue.list.Remove(element)

	return element
}

// Remove - remove object to the queue.
func (queue *Queue) Remove(element *list.Element)  {
	queue.list.Remove(element)
}

// Get - getting object to the queue.
func (queue *Queue) Get() *list.Element {
	return queue.list.Front()
}

// Len - getting the queue length.
func (queue *Queue) Len() int {
	return queue.list.Len()
}

// GetAll - getting all object to the queue.
func (queue *Queue) GetAll() []interface{} {
	var arr []interface{}

	for i := 0; i < queue.Len(); i++ {
		element := queue.Get()
		arr = append(arr, element.Value)

		queue.list.MoveToBack(element)
	}

	return arr
}