package mylib

type MapStore struct {
	bookMap map[uint32]Book
}

func NewMapStore() *MapStore {
	return &MapStore{
		bookMap: make(map[uint32]Book),
	}
}

func (store *MapStore) Add(id uint32, book Book) {
	book.Id = id
	store.bookMap[id] = book
}

func (store *MapStore) Search(id uint32) (Book, bool) {
	book, found := store.bookMap[id]
	return book, found
}

func (store *MapStore) Remove(id uint32) {
	delete(store.bookMap, id)
}

func (store *MapStore) Regenerate(generator BookIdGenerator) {
	newMap := make(map[uint32]Book)
	for _, book := range store.bookMap {
		newMap[generator.genID(book.Title)] = book
	}
	store.bookMap = newMap
}
