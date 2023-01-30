package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const BaseURL = "https://api2.isbndb.com/book/"

type Book struct {
	Book BookData `json:"book"`
}

type BookData struct {
	Title     string   `json:"title"`
	Language  string   `json:"language"`
	Authors   []string `json:"authors"`
	Publisher string   `json:"publisher"`
	Image     string   `json:"image"`
}

func main() {
	http.HandleFunc("/", getBookFromIsbndbApi)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getBookFromIsbndbApi(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Query().Get("isbn")
	book := fetchBookByIsbn(isbn)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func fetchBookByIsbn(isbn string) *BookData {
	fullUrl := BaseURL + isbn

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		fmt.Println("Error happened: ", err)
	}

	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}
	req.Header.Add("Authorization", os.Getenv("APIKey"))

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer response.Body.Close()

	bookData := getBookResponse(*response)
	return bookData
}

func getBookResponse(response http.Response) *BookData {
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		fmt.Println("Error happened: ", readErr)
	}

	var book Book
	json.Unmarshal([]byte(body), &book)
	return &book.Book
}
