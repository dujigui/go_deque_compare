package main

type Deque struct {
	head  int
	tail  int
	data  []interface{}
	count int
	init  int
}

func NewDeque(c int) *Deque {
	if c <= 0 {
		panic("init len of deque can not be 0")
	}
	return &Deque{
		data: make([]interface{}, c),
		init: c,
	}
}

func (d *Deque) Enqueue(item interface{}) bool {
	if d.Full() {
		return false
	}
	d.data[d.tail] = item
	d.tail = (d.tail + 1) % len(d.data)
	d.count++
	d.grow()
	return true
}

func (d *Deque) Dequeue() (interface{}, bool) {
	if d.Empty() {
		return nil, false
	}
	item := d.data[d.head]
	d.data[d.head] = nil
	d.head = (d.head + 1) % len(d.data)
	d.count--
	d.shrink()
	return item, true
}

func (d *Deque) Head() (interface{}, bool) {
	if d.Empty() {
		return nil, false
	}
	return d.data[d.head], true
}

func (d *Deque) Tail() (interface{}, bool) {
	if d.Empty() {
		return nil, false
	}
	return d.data[d.tail], true
}

func (d *Deque) Empty() bool {
	return d.count == 0
}

func (d *Deque) Full() bool {
	return d.count == len(d.data)
}


func (d *Deque) grow() {
	if !d.Full() || d.Empty() {
		return
	}
	old := d.data
	data := make([]interface{}, len(d.data)*2)
	if d.head == 0 {
		copy(data, old)
	} else {
		n := copy(data, old[d.head:])
		copy(data[n:], old[:d.head])
	}
	d.data = data
	d.head = 0
	d.tail = d.count
}

func (d *Deque) shrink() {
	l := len(d.data) / 4
	if l < d.count {
		return
	}
	if l < d.init {
		return
	}

	old := d.data
	data := make([]interface{}, len(d.data)/2)
	if d.tail > d.head {
		copy(data, old[d.head:d.tail])
	} else {
		n := copy(data, old[d.head:])
		copy(data[n:], old[:d.tail])
	}
	d.data = data
	d.head = 0
	d.tail = d.count
}
