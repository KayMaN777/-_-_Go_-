package main

type MapStore struct {
	bookMap map[uint32]Book
}

func NewMapStore() *MapStore {
	return &MapStore{
		bookMap: make(map[uint32]Book),
	}
}

func (store *MapStore) Add(id uint32, book Book) {
	store.bookMap[id] = book
}

func (store *MapStore) Search(id uint32) (Book, bool) {
	book, found := store.bookMap[id]
	return book, found
}
