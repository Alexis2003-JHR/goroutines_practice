package data

import (
	"fmt"
	"sync"
)

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{1, "Dune", false},
	{2, "El perfume", false},
	{3, "The world of", false},
	{4, "Teoría de la noche", false},
	{5, "Blanca Olmedo", false},
	{6, "El principito", false},
	{7, "100 años de soledad", false},
	{8, "El alquimista", false},
	{9, "El libro del cemeterio", false},
	{10, "Maze runner", false},
}

func findBook(id int, m *sync.RWMutex) (int, *Book) {
	index := -1
	var book *Book

	m.RLock()
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}
	}
	m.RUnlock()

	return index, book
}

func FinishBook(id int, m *sync.RWMutex) {
	i, book := findBook(id, m)
	if i < 0 {
		return
	}

	m.Lock()
	book.Finished = true
	books[i] = book
	m.Unlock()

	fmt.Printf("Finished Book: %s\n", book.Title)
}
