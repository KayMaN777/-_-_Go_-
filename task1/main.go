package main

import (
	"fmt"
)

func main() {
	book1 := Book{Title: "Война и мир", Author: "Lev Tolstoi", Pages: 1225}
	book2 := Book{Title: "Анна Каренина", Author: "Lev Tolstoi"}

	library := NewLibrary(NewMapStore(), new(FnvGenerator))
	library.AddBook(book1)
	library.AddBook(book2)

	book, ok := library.FindBook("Анна Каренина")
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	book, ok = library.FindBook("Harry Potter")
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}
}
