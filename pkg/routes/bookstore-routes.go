package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func RegisterBookStoreRoutes(router *mux.Router) {
    router.HandleFunc("/book/", CreateBookHandler).Methods("POST")
    router.HandleFunc("/book/", GetBookHandler).Methods("GET")
    router.HandleFunc("/book/{bookId}", GetBookByIdHandler).Methods("GET")
    router.HandleFunc("/book/{bookId}", UpdateBookHandler).Methods("PUT")
    router.HandleFunc("/book/{bookId}", DeleteBookHandler).Methods("DELETE")
}
// Estrutura de exemplo para representar um livro
type Book struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    Quantity int    `json:"quantity"`
}

var books []Book // Slice para armazenar livros fictícios

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
    var newBook Book
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&newBook); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Gere um ID único para o novo livro (simplificado)
    newBook.ID = len(books) + 1
    books = append(books, newBook)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newBook)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func GetBookByIdHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID, err := strconv.Atoi(vars["bookId"])
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    // Buscar o livro pelo ID (simplificado)
    for _, book := range books {
        if book.ID == bookID {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(book)
            return
        }
    }

    http.Error(w, "Book not found", http.StatusNotFound)
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID, err := strconv.Atoi(vars["bookId"])
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    var updatedBook Book
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&updatedBook); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Atualizar o livro pelo ID (simplificado)
    for i, book := range books {
        if book.ID == bookID {
            books[i] = updatedBook
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(updatedBook)
            return
        }
    }

    http.Error(w, "Book not found", http.StatusNotFound)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID, err := strconv.Atoi(vars["bookId"])
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    // Excluir o livro pelo ID (simplificado)
    for i, book := range books {
        if book.ID == bookID {
            books = append(books[:i], books[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }

    http.Error(w, "Book not found", http.StatusNotFound)
}
