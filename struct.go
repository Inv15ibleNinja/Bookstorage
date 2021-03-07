package main

type Book struct {
	Id    int
	Title string
}

type Author struct {
	Id   int
	Name string
}

type Publisher struct {
	Id   int
	Name string
}

type Book_author struct {
	Book_id   int
	Author_id int
}

type Author_publisher struct {
	Author_id    int
	Publisher_id int
}

type Book_publisher struct {
	Book_id      int
	Publisher_id int
}
