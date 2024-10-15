package bookstore

import (
	"task1/internal/idgenerator"
	"task1/internal/model"
)

type BookStore interface {
	Search(id uint32) (model.Book, bool)
	Add(id uint32, b model.Book)
	Remove(id uint32)
	Regenerate(generator idgenerator.BookIdGenerator)
}

type Library struct {
	store     BookStore
	generator idgenerator.BookIdGenerator
}

func NewLibrary(store BookStore, generator idgenerator.BookIdGenerator) *Library {
	return &Library{
		store:     store,
		generator: generator,
	}
}

func (library *Library) AddBook(book model.Book) {
	library.store.Add(library.generator.GenerateID(book.Title), book)
}

func (library *Library) FindBook(title string) (model.Book, bool) {
	return library.store.Search(library.generator.GenerateID(title))
}

func (library *Library) RemoveBook(title string) {
	library.store.Remove(library.generator.GenerateID(title))
}

func (library *Library) SetStore(store BookStore) {
	library.store = store
}

func (library *Library) SetGenerator(generator idgenerator.BookIdGenerator) {
	library.generator = generator
	library.store.Regenerate(generator)
}
