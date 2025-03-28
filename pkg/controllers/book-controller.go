package controllers

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"strconv"

	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"

	"log"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("[GET] /book/ called")

	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	log.Printf("[GET] /book/%s called\n", bookId)

	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		log.Printf("Error parsing ID: %v\n", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	bookDetails, db := models.GetBookById(ID)
	if db.Error != nil {
		log.Printf("Book with ID %d not found\n", ID)
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("[POST] /book/ called")

	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	log.Printf("Creating book: %+v\n", CreateBook)

	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	log.Printf("[DELETE] /book/%s called\n", bookId)

	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		log.Printf("Error parsing ID: %v\n", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	log.Printf("[PUT] /book/%s called\n", bookId)

	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		log.Printf("Error parsing ID: %v\n", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	bookDetails, db := models.GetBookById(ID)
	if db.Error != nil {
		log.Printf("Book with ID %d not found\n", ID)
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
