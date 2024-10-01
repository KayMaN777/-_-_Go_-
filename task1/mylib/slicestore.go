package mylib

type SliceStore struct {
	books []Book
}

func NewSliceStore() *SliceStore {
	return &SliceStore{
		books: []Book{},
	}
}

func (store *SliceStore) Add(id uint32, book Book) {
	if _, ok := store.Search(id); ok {
		return
	}
	book.Id = id
	store.books = append(store.books, book)
}

func (store *SliceStore) Search(id uint32) (Book, bool) {
	for _, book := range store.books {
		if book.Id == id {
			return book, true
		}
	}
	return Book{}, false
}

func (store *SliceStore) Remove(id uint32) {
	for i, book := range store.books {
		if book.Id == id {
			store.books = append(store.books[:i], store.books[i+1:]...)
			break
		}
	}
}

func (store *SliceStore) Regenerate(generator BookIdGenerator) {
	for i, book := range store.books {
		store.books[i].Id = generator.genID(book.Title)
	}
}
