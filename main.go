package main

// Queue интерфейс для очереди
type Queue interface {
	Push(value int)
	Pop() (int, error)
}

// Heap интерфейс для кучи
type Heap interface {
	Push(value int)
	Pop() (int, error)
}

// Binary интерфейс для бинарного дерева
type Binary interface {
	Push(value int)
	Pop() (int, error)
	Search(value int) bool
}
