package main

import (
	"errors"
	"testing"
)

// Тестируемый класс Queue
type TestQueue struct {
	data []int
}

func (q *TestQueue) Push(value int) {
	q.data = append(q.data, value)
}

func (q *TestQueue) Pop() (int, error) {
	if len(q.data) == 0 {
		return 0, errors.New("очередь пустая")
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value, nil
}

// Тестируемый класс Heap
type TestHeap struct {
	data []int
}

func (h *TestHeap) Push(value int) {
	h.data = append(h.data, value)
}

func (h *TestHeap) Pop() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("куча пустая")
	}
	value := h.data[0]
	h.data = h.data[1:]
	return value, nil
}

// Тестируемый класс Binary
type TestBinary struct {
	data []int
}

func (bt *TestBinary) Push(value int) {
	bt.data = append(bt.data, value)
}

func (bt *TestBinary) Pop() (int, error) {
	if len(bt.data) == 0 {
		return 0, errors.New("двоичное дерево пусто")
	}
	value := bt.data[len(bt.data)-1]
	bt.data = bt.data[:len(bt.data)-1]
	return value, nil
}

func (bt *TestBinary) Search(value int) bool {
	for _, v := range bt.data {
		if v == value {
			return true
		}
	}
	return false
}

// Тесты для Queue
func TestQueue(t *testing.T) {
	var q Queue = &TestQueue{}

	q.Push(1)
	q.Push(2)

	val, err := q.Pop()
	if err != nil || val != 1 {
		t.Errorf("Ожидалось 1, получено %d", val)
	}

	val, err = q.Pop()
	if err != nil || val != 2 {
		t.Errorf("Ожидалось 2, получено %d", val)
	}

	_, err = q.Pop()
	if err == nil {
		t.Error("Ошибка, пустая очередь")
	}
}

// Тесты для Heap
func TestHeap(t *testing.T) {
	var h Heap = &TestHeap{}

	h.Push(3)
	h.Push(4)

	val, err := h.Pop()
	if err != nil || val != 3 {
		t.Errorf("Ожидалось 3, получено %d", val)
	}

	val, err = h.Pop()
	if err != nil || val != 4 {
		t.Errorf("Ожидалось 4, получено %d", val)
	}

	_, err = h.Pop()
	if err == nil {
		t.Error("Ошибка, пустая куча")
	}
}

// Тесты для Binary
func TestBinary(t *testing.T) {
	var bt Binary = &TestBinary{}

	bt.Push(5)
	bt.Push(6)

	val, err := bt.Pop()
	if err != nil || val != 6 {
		t.Errorf("Ожидалось 6, получено %d", val)
	}

	val, err = bt.Pop()
	if err != nil || val != 5 {
		t.Errorf("Ожидалось 5, получено %d", val)
	}

	if _, err := bt.Pop(); err == nil {
		t.Error("Ошибка, пустое двоичное дерево")
	}

	bt.Push(7)
	if !bt.Search(7) {
		t.Error("Ожидалось, что значение 7 будет найдено в бинарном дереве")
	}
	if bt.Search(8) {
		t.Error("Не ожидалось, что значение 8 будет найдено в бинарном дереве")
	}
}
