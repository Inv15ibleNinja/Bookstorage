package main

type Book struct {
	id    int
	title string
}

type Author struct {
	id   int
	name string
}

type Publisher struct {
	id   int
	name string
}

type Book_author struct {
	book_id   int
	author_id int
}

type Author_publisher struct {
	author_id    int
	publisher_id int
}

type Book_publisher struct {
	book_id      int
	publisher_id int
}
