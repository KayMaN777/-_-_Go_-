package main

type BookStore interface {
	Search(id uint32) (Book, bool)
	Add(id uint32, b Book)
}

type Library struct {
	store     BookStore
	generator BookIdGenerator
}

func NewLibrary(store BookStore, generator BookIdGenerator) *Library {
	return &Library{
		store:     store,
		generator: generator,
	}
}

func (library *Library) AddBook(book Book) {
	library.store.Add(library.generator.genID(book.Title), book)
}

func (library *Library) FindBook(title string) (Book, bool) {
	return library.store.Search(library.generator.genID(title))
}
