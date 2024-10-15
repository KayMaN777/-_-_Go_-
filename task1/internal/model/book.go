package model

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
	Id     uint32
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}
