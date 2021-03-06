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
var err error

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
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Publisher{})
	db.AutoMigrate(&Book{})

	//заполняем тестовыми данными
	for index := range book {
		db.Create(&book[index])
	}

	for index := range author {
		db.Create(&author[index])
	}

	for index := range publisher {
		db.Create(&publisher[index])
	}

	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", GetBook).Methods("GET")
	//   router.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	//   router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")
	//   router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
	//os.Exit(0)
}
