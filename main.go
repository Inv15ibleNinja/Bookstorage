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

func GetBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBooks request")
	var books []Book
	db.Find(&books)
	json.NewEncoder(w).Encode(&books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBook request")
	params := mux.Vars(r)
	var book Book
	db.First(&book, params["id"])
	json.NewEncoder(w).Encode(&book)
}

var (
	author = []Author{

		{name: "Stephen King"},
		{name: "John Ronald Reuel Tolkien"},
		{name: "George Raymond Richard Martin"},
		{name: "Александр Сергеевич Пушкин"},
		{name: "Сергей Васильевич Лукьяненко"},
		{name: "Аркадий Натанович Стругаций"},
		{name: "Борис Натанович Стругаций"},
		{name: "Сергей Васильевич Лукьяненко"},
	}

	publisher = []Publisher{

		{name: "Росмэн"},
		{name: "Экспоненента"},
		{name: "МИФ"},
		{name: "Guardian"},
		{name: "Special"},
		{name: "Букля"},
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
	}
)
var err error

func main() {
	fmt.Println("start")
	router := mux.NewRouter()
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=Book sslmode=disable password=postgres")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	//db.AutoMigrate(&Author{})
	//db.AutoMigrate(&Publisher{})
	db.AutoMigrate(&Book{})
	for index := range book {
		db.Create(&book[index])
	}
	// for index := range author {
	// 	db.Create(&author[index])
	// }

	// for index := range publisher {
	// 	db.Create(&publisher[index])
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
