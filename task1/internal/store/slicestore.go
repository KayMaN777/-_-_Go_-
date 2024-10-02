package store

import (
	"task1/internal/idgenerator"
	"task1/internal/model"
)

type SliceStore struct {
	books []model.Book
}

func NewSliceStore() *SliceStore {
	return &SliceStore{
		books: []model.Book{},
	}
}

func (store *SliceStore) Add(id uint32, book model.Book) {
	if _, ok := store.Search(id); ok {
		return
	}
	book.Id = id
	store.books = append(store.books, book)
}

func (store *SliceStore) Search(id uint32) (model.Book, bool) {
	for _, book := range store.books {
		if book.Id == id {
			return book, true
		}
	}
	return model.Book{}, false
}

func (store *SliceStore) Remove(id uint32) {
	for i, book := range store.books {
		if book.Id == id {
			store.books = append(store.books[:i], store.books[i+1:]...)
			break
		}
	}
}

func (store *SliceStore) Regenerate(generator idgenerator.BookIdGenerator) {
	for i, book := range store.books {
		store.books[i].Id = generator.GenerateID(book.Title)
	}
}
