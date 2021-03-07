package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

//GetBooks returns list of all books in db
func GetBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBooks request")
	var books []Book
	db.Find(&books)
	json.NewEncoder(w).Encode(&books)
}

//GetBook returns book by book id
func GetBook(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBook request")
	params := mux.Vars(r)
	var book Book
	db.First(&book, params["id"])
	json.NewEncoder(w).Encode(&book)
}

var (
	author = []Author{

		{Name: "Stephen King"},
		{Name: "John Ronald Reuel Tolkien"},
		{Name: "George Raymond Richard Martin"},
		{Name: "Александр Сергеевич Пушкин"},
		{Name: "Сергей Васильевич Лукьяненко"},
		{Name: "Аркадий Натанович Стругаций"},
		{Name: "Борис Натанович Стругаций"},
		{Name: "Сергей Васильевич Лукьяненко"},
	}

	publisher = []Publisher{

		{Name: "Росмэн"},
		{Name: "Экспоненента"},
		{Name: "МИФ"},
		{Name: "Guardian"},
		{Name: "Special"},
		{Name: "Букля"},
	}

	book = []Book{
		{Title: "Тёмная башня"},
		{Title: "Пикник на обочине"},
		{Title: "Оно"},
		{Title: "Танец с драконами"},
		{Title: "Песнь льда и пламени"},
		{Title: "Хоббит"},
		{Title: "Властелин колец"},
		{Title: "Руслан и Людмила"},
		{Title: "Дубровский"},
		{Title: "Ночной дозор"},
		{Title: "Спектр"},
		{Title: "Еще одна книга"},
		{Title: "Книга без автора 1"},
		{Title: "Книга без автора 2"},
		{Title: "Книга без издательства 1"},
		{Title: "Книга без издательства 2"},
		{Title: "Еще одна книга без автора"},
		{Title: "Еще одна книга без издательства"},
		{Title: "Еще одна книга без автора 2"},
	}

	book_author = []Book_author{

		{
			Book_id:   1,
			Author_id: 1,
		},
		{
			Book_id:   2,
			Author_id: 6,
		},
		{
			Book_id:   2,
			Author_id: 7,
		},
		{
			Book_id:   3,
			Author_id: 1,
		},
		{
			Book_id:   4,
			Author_id: 3,
		},
		{
			Book_id:   5,
			Author_id: 3,
		},
		{
			Book_id:   6,
			Author_id: 2,
		},
		{
			Book_id:   7,
			Author_id: 2,
		},
		{
			Book_id:   8,
			Author_id: 4,
		},
		{
			Book_id:   9,
			Author_id: 4,
		},
		{
			Book_id:   10,
			Author_id: 10,
		},
		{
			Book_id:   11,
			Author_id: 10,
		},
		{
			Book_id:   12,
			Author_id: 1,
		},
		{
			Book_id:   12,
			Author_id: 5,
		},
		{
			Book_id:   12,
			Author_id: 3,
		},
	}

	book_publisher = []Book_publisher{

		{
			Book_id:      1,
			Publisher_id: 1,
		},
		{
			Book_id:      2,
			Publisher_id: 1,
		},
		{
			Book_id:      3,
			Publisher_id: 1,
		},
		{
			Book_id:      2,
			Publisher_id: 1,
		},
		{
			Book_id:      2,
			Publisher_id: 3,
		},
		{
			Book_id:      3,
			Publisher_id: 2,
		},

		{
			Book_id:      4,
			Publisher_id: 4,
		},
		{
			Book_id:      4,
			Publisher_id: 5,
		},
		{
			Book_id:      5,
			Publisher_id: 6,
		},
		{
			Book_id:      6,
			Publisher_id: 1,
		},
		{
			Book_id:      7,
			Publisher_id: 2,
		},
		{
			Book_id:      8,
			Publisher_id: 4,
		},
	}
	err error
)

func main() {
	fmt.Println("start") //debug
	//коннектимся к бд
	router := mux.NewRouter()
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=Book sslmode=disable password=postgres")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	//добавляем таблицы в бд
	//db.AutoMigrate(&Author{})
	//db.AutoMigrate(&Publisher{})
	//db.AutoMigrate(&Book{})
	db.AutoMigrate(&Book_author{})
	db.AutoMigrate(&Book_publisher{})
	//заполняем тестовыми данными
	for i := range book {
		db.Create(&book[i])
	}

	for i := range book_author {
		db.Create(&book_author[i])
	}
	for i := range book_publisher {
		db.Create(&book_publisher[i])
	}

	// for i := range author {
	// 	db.Create(&author[i])
	// }

	// for i := range publisher {
	// 	db.Create(&publisher[i])
	// }

	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", GetBook).Methods("GET")
	//   router.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	//   router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")
	//   router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
	//os.Exit(0)
}
