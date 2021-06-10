package resources

import (
	"../../utils/queue"
	"container/list"
)

type Orders struct {
	list              *queue.Queue
	ValueAllResources uint64 // Общее кол-во ресурсов на цене
}

func newOrders() *Orders {
	return &Orders{
		list:              queue.New(),
		ValueAllResources: 0,
	}
}

func (orders *Orders) Len() int {
	return orders.list.Len()
}

func (orders *Orders) Add(value interface{}) {
	orders.list.Add(value)
}

func (orders *Orders) GetAll() []interface{} {
	return orders.list.GetAll()
}

func (orders *Orders) Get() *list.Element {
	return orders.list.Get()
}

func (orders *Orders) Remove(element *list.Element)  {
	orders.list.Remove(element)
}