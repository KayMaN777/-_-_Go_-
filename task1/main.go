package main

import (
	"fmt"
)

func main() {
	books := []Book{
		{Title: "Война и мир", Author: "Лев Толстой", Pages: 1225},
		{Title: "Преступление и наказание", Author: "Фёдор Достоевский", Pages: 550},
		{Title: "Анна Каренина", Author: "Лев Толстой", Pages: 864},
		{Title: "Мастер и Маргарита", Author: "Михаил Булгаков", Pages: 400},
		{Title: "1984", Author: "Джордж Оруэлл", Pages: 328},
		{Title: "Убить пересмешника", Author: "Харпер Ли", Pages: 400},
		{Title: "Гордость и предубеждение", Author: "Джейн Остин", Pages: 432},
		{Title: "Старик и море", Author: "Эрнест Хемингуэй", Pages: 128},
		{Title: "Над пропастью во ржи", Author: "Джером Д. Сэлинджер", Pages: 224},
		{Title: "Великий Гэтсби", Author: "Фрэнсис Скотт Фицджеральд", Pages: 208},
		{Title: "Тихий Дон", Author: "Михаил Шолохов", Pages: 1280},
		{Title: "Унесённые ветром", Author: "Маргарет Митчелл", Pages: 1024},
		{Title: "Доктор Живаго", Author: "Борис Пастернак", Pages: 672},
		{Title: "Собачье сердце", Author: "Михаил Булгаков", Pages: 192},
		{Title: "Золото пылающих скал", Author: "Джон Стейнбек", Pages: 320},
	}
	// Создаем библиотеку
	library := NewLibrary(NewMapStore(), new(FnvGenerator))
	library.AddBook(books[0])
	library.AddBook(books[1])
	library.AddBook(books[2])

	// Добавляли книгу, проверяем что она есть
	book, ok := library.FindBook(books[0].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Добавляли книгу, проверяем что она есть
	book, ok = library.FindBook(books[2].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Книгу не добавляли, проверяем что ее нет
	book, ok = library.FindBook(books[7].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Удалили книгу, проверяем что ее нет
	library.RemoveBook(books[0].Title)
	book, ok = library.FindBook(books[0].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Меняем генератор id
	library.SetGenerator(new(AdlerGenerator))

	// Проверяем, что книги которые были добавлены находятся
	book, ok = library.FindBook(books[1].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}
	book, ok = library.FindBook(books[2].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Проверяем, что книга которой не было добавлена не находится
	book, ok = library.FindBook(books[5].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Меняем хранилище
	library.SetStore(NewSliceStore())
	library.AddBook(books[7])
	library.AddBook(books[8])
	library.AddBook(books[9])

	// Книгу не добавляли, проверяем что ее нет
	book, ok = library.FindBook(books[0].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Книгу добавляли, проверяем что она есть
	book, ok = library.FindBook(books[8].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Удаляем книгу и проверяем, что ее больше нет
	library.RemoveBook(books[8].Title)
	book, ok = library.FindBook(books[8].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}
}
