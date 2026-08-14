package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CreateRequest struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type CreateResponse struct {
	IdBook      int64  `json:"id_book"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type Book struct {
	DB *sql.DB
}

func (book *Book) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 1. Ambil & ubah (Decode) JSON dari client ke struct CreateRequest
	req := CreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Simpan data req ke MySQL database
	SQL := "INSERT INTO book (name, author, publication) VALUES (?, ?, ?)"
	result, err := book.DB.ExecContext(r.Context(), SQL, req.Name, req.Author, req.Publication)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Ambil ID baru yang dihasilkan AUTO_INCREMENT MySQL
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Buat objek Response berisi data lengkap (termasuk id_book baru)
	response := CreateResponse{
		IdBook:      id,
		Name:        req.Name,
		Author:      req.Author,
		Publication: req.Publication,
	}

	// 5. Kirim balasan (Encode) sebagai JSON ke client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (book *Book) GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	SQL := "SELECT id_book, name, author, publication FROM book"
	_ = SQL
}

func (book *Book) GetBookById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	SQL := "SELECT id_book, name, author, publication FROM book WHERE id_book = ?"
	_ = SQL
}

func (book *Book) UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	SQL := "UPDATE book SET name = ?, author = ?, publication = ? WHERE id_book = ?"
	_ = SQL
}

func (book *Book) DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	SQL := "DELETE FROM book WHERE id_book = ?"
	_ = SQL
}
