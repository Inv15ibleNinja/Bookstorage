package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	db.Find(&books)
	json.NewEncoder(w).Encode(&books)
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
	}
)

func main() {

	router := mux.NewRouter()
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=Book sslmode=disable password=postgres")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Publisher{})
	db.AutoMigrate(&Book{})

	for index := range author {
		db.Create(&author[index])
	}

	for index := range book {
		db.Create(&book[index])
	}

	for index := range publisher {
		db.Create(&publisher[index])
	}
	router.HandleFunc("/books", GetBooks).Methods("GET")
	//   router.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	//   router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")
	//   router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
	os.Exit(0)
}
