package main

import (
	"container/list"
	"testing"
)

func TestNewDeque(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("init len of deque should not be 0")
			}
		}()
		NewDeque(0)
	}()

	for i := 1; i < 3; i++ {
		d := NewDeque(i)
		if d.count != 0 {
			t.Fatal("d.count is not 0")
		}
		if d.data == nil || len(d.data) != i {
			t.Fatalf("d.data is nil or len(d.data) is not %d", i)
		}
		if d.tail != 0 {
			t.Fatal("d.tail is not 0")
		}
		if d.head != 0 {
			t.Fatal("d.head is not 0")
		}
	}
}

func TestDeque_Enqueue(t *testing.T) {
	d := NewDeque(1)
	if !d.Enqueue(1) {
		t.Fatal("enqueue fail")
	}
	if d.count != 1 {
		t.Fatal("d.count is not right")
	}
	if d.tail != 1 {
		t.Fatal("d.tail is not right")
	}
	if d.head != 0 {
		t.Fatal("d.head is not right")
	}
	d = NewDeque(2)
	d.Enqueue(1)
	if d.Empty() {
		t.Fatal("deque should not be empty")
	}
	if d.Full() {
		t.Fatal("deque should not be full")
	}
	if d.tail != 1 {
		t.Fatal("d.tail is not right")
	}
	if d.head != 0 {
		t.Fatal("d.head is not right")
	}
}

func TestDeque_Dequeue(t *testing.T) {
	d := NewDeque(3)
	d.Enqueue(1)
	d.Enqueue(2)
	d.Enqueue(3)
	item, ok := d.Dequeue()
	if item == nil || !ok {
		t.Fatal("dequeue fail")
	}
	d.Dequeue()
	d.Enqueue(4)
	if d.tail != 4 {
		t.Fatal("d.tail is not right")
	}
	if d.head != 2 {
		t.Fatal("d.head is not right")
	}
	if d.count != 2 {
		t.Fatal("d.count is not right")
	}
}

func TestDeque_grow_shrink(t *testing.T) {
	d := NewDeque(1)
	d.Enqueue(1)
	if d.head != 0 || d.tail != 1 || len(d.data) != 2 {
		t.Fatal("deque not grow properly")
	}
	d.Enqueue(2)
	if d.head != 0 || d.tail != 2 || len(d.data) != 4 {
		t.Fatal("deque not grow properly")
	}
	d.Enqueue(3)
	d.Enqueue(4)
	if d.head != 0 || d.tail != 4 || len(d.data) != 8 {
		t.Fatal("deque not grow properly")
	}
	d.Dequeue()
	d.Dequeue()
	if d.head != 0 || d.tail != 2 || len(d.data) != 4 {
		t.Fatal("deque not shrink properly")
	}
}

func BenchmarkDeque_Enqueue(b *testing.B) {
	d := NewDeque(100)
	for n := 0; n < b.N; n++ {
		d.Enqueue(d)
	}
}

func BenchmarkDequeSlice_Enqueue(b *testing.B) {
	d := NewDequeSlice(100)
	for n := 0; n < b.N; n++ {
		d.Enqueue(d)
	}
}

func BenchmarkLinkedList_Enqueue(b *testing.B) {
	d := list.New()
	for n := 0; n < b.N; n++ {
		d.PushBack(d)
	}
}

func BenchmarkDeque_Dequeue(b *testing.B) {
	d := NewDeque(10000)
	for i := 0; i < 1000; i++ {
		d.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		d.Dequeue()
	}
}

func BenchmarkDequeSlice_Dequeue(b *testing.B) {
	d := NewDequeSlice(10000)
	for i := 0; i < 1000; i++ {
		d.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		d.Dequeue()
	}
}

func BenchmarkLinkedList_Dequeue(b *testing.B) {
	d := list.New()
	for i := 0; i < 1000; i++ {
		d.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		d.Remove(d.Front())
	}
}

func BenchmarkDeque_Enqueue_Dequeue(b *testing.B) {
	d := NewDeque(1)
	for i := 0; i < 1000; i++ {
		//1 2 4 8 16 ... 2^n
		for j := 0; j < 50; j++ {
			d.Enqueue(j)
		}
		for k := 0; k < 50; k++ {
			d.Dequeue()
		}
	}
}

func BenchmarkDequeSlice_Enqueue_Dequeue(b *testing.B) {
	d := NewDequeSlice(1)
	for i := 0; i < 1000; i++ {
		//1 2 4 8 16 ... 2^n
		for j := 0; j < 50; j++ {
			d.Enqueue(j)
		}
		for k := 0; k < 50; k++ {
			d.Dequeue()
		}
	}
}

func BenchmarkLinkedList_Enqueue_Dequeue(b *testing.B) {
	d := list.New()
	for i := 0; i < 1000; i++ {
		//1 2 4 8 16 ... 2^n
		for j := 0; j < 50; j++ {
			d.PushBack(j)
		}
		for k := 0; k < 50; k++ {
			d.Remove(d.Front())
		}
	}
}



func BenchmarkDeque_Enqueue_Then_Dequeue(b *testing.B) {
	d := NewDeque(8)
	for i := 0; i < 6; i++ {
		d.Enqueue(i)
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		if j%2 == 0 {
			d.Enqueue(j)
		} else {
			d.Dequeue()
		}
	}
}

func BenchmarkDequeSlice_Enqueue_Then_Dequeue(b *testing.B) {
	d := NewDequeSlice(8)
	for i := 0; i < 6; i++ {
		d.Enqueue(i)
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		if j%2 == 0 {
			d.Enqueue(j)
		} else {
			d.Dequeue()
		}
	}
}

func BenchmarkLinkedList_Enqueue_Then_Dequeue(b *testing.B) {
	d := list.New()
	for i := 0; i < 6; i++ {
		d.PushBack(i)
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		if j%2 == 0 {
			d.PushBack(j)
		} else {
			d.Remove(d.Front())
		}
	}
}


