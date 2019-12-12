package main

type DequeSlice struct {
	data []interface{}
}

func NewDequeSlice(c int) *DequeSlice {
	return &DequeSlice{
		data: make([]interface{}, 0, c),
	}
}

func (d *DequeSlice) Enqueue(item interface{}) {
	d.data = append(d.data, item)
}

func (d *DequeSlice) Dequeue() (interface{}, bool) {
	if len(d.data) == 0 {
		return nil, false
	}
	item := d.data[0]
	d.data = d.data[1:]
	return item, true
}
