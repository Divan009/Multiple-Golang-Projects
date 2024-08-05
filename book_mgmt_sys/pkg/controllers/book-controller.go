package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/divan009/Multiple-Golang-Projects/pkg/models"
	"github.com/divan009/Multiple-Golang-Projects/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	// newBooks val converted to JSON below
	resp, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("err while parsing ID")
	}
	newBook, _ := models.GetBookById(ID)
	resp, _ := json.Marshal(newBook)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}

	// Parse the incoming JSON into the newBook struct interface
	err := utils.ParseBody(r, newBook)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Create the book record in the database
	createdBook := newBook.CreateBook()

	// Marshal the created book into JSON
	resp, err := json.Marshal(createdBook)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	// Set the header and write the response
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, _ := strconv.ParseInt(bookId, 0, 0)

	deleteBook := models.DeleteBook(ID)

	resp, err := json.Marshal(deleteBook)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	updatedBook := &models.Book{}
	err := utils.ParseBody(r, updatedBook)
	if err != nil {
		return
	}

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("parsing bad")
	}

	bookDetails, db := models.GetBookById(ID)
	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}

	err = db.Save(&bookDetails).Error // Use a clean or appropriately scoped `db` instance
	if err != nil {
		http.Error(w, "Failed to update book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
