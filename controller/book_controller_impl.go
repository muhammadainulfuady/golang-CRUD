package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang_crud/exception"
	"golang_crud/helper"
	"golang_crud/model/web"
	"golang_crud/service"

	"github.com/julienschmidt/httprouter"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) Save(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := web.CreateBookRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			web.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "fail",
				Message: "Json formatter err",
				Errors:  err.Error(),
				Data:    nil,
			},
		)
		return
	}

	bookResponse, err := controller.BookService.Save(r.Context(), request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		formatErrors := helper.FormatValidationError(err)
		json.NewEncoder(w).Encode(
			web.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "fail",
				Message: "Validation error",
				Errors:  formatErrors,
				Data:    nil,
			},
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		web.ApiResponse{
			Code:    http.StatusCreated,
			Status:  "OK",
			Message: "Berhasil menambahkan buku",
			Errors:  nil,
			Data:    bookResponse,
		},
	)
}

func (controller *BookControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idBookStr := params.ByName("idBook")
	idBook, err := strconv.Atoi(idBookStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			web.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "fail",
				Message: "ID buku harus berupa angka",
				Errors:  err.Error(),
				Data:    nil,
			},
		)
		return
	}

	request := web.UpdateBookRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			web.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "fail",
				Message: "Json formatter err",
				Errors:  err.Error(),
				Data:    nil,
			},
		)
		return
	}

	request.IdBook = idBook

	bookResponse, err := controller.BookService.Update(r.Context(), request)
	if err != nil {
		if notFoundErr, ok := err.(*exception.NotFoundErr); ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(
				web.ApiResponse{
					Code:    http.StatusNotFound,
					Status:  "fail",
					Message: "Buku tidak ditemukan",
					Errors:  notFoundErr.Error(),
					Data:    nil,
				},
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		formatErrors := helper.FormatValidationError(err)
		json.NewEncoder(w).Encode(
			web.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "fail",
				Message: "Validation error",
				Errors:  formatErrors,
				Data:    nil,
			},
		)
		return
	}
	webResponse := web.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Berhasil update buku",
		Errors:  nil,
		Data:    bookResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}

func (controller *BookControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idBookStr := params.ByName("idBook")
	idBook, err := strconv.Atoi(idBookStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			web.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "fail",
				Message: "ID buku harus berupa angka",
				Errors:  err.Error(),
				Data:    nil,
			},
		)
		return
	}

	err = controller.BookService.Delete(r.Context(), idBook)
	if notFoundErr, ok := err.(*exception.NotFoundErr); ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			web.ApiResponse{
				Code:    http.StatusNotFound,
				Status:  "fail",
				Message: "Buku tidak ditemukan",
				Errors:  notFoundErr.Error(),
				Data:    nil,
			},
		)
		return
	}

	webResponse := web.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Berhasil menghapus buku",
		Errors:  nil,
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}

func (controller *BookControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idBookStr := params.ByName("idBook")
	idBook, err := strconv.Atoi(idBookStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			web.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "fail",
				Message: "ID buku harus berupa angka",
				Errors:  err.Error(),
				Data:    nil,
			},
		)
		return
	}

	bookResponse, err := controller.BookService.FindById(r.Context(), idBook)
	if notFoundErr, ok := err.(*exception.NotFoundErr); ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			web.ApiResponse{
				Code:    http.StatusNotFound,
				Status:  "fail",
				Message: "Buku tidak ditemukan",
				Errors:  notFoundErr.Error(),
				Data:    nil,
			},
		)
		return
	}

	webResponse := web.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Berhasil mengambil buku",
		Errors:  nil,
		Data:    bookResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}

func (controller *BookControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bookResponse, err := controller.BookService.FindAll(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			web.ApiResponse{
				Code:    http.StatusInternalServerError,
				Status:  "fail",
				Message: "Gagal mengambil semua buku server error",
				Errors:  err.Error(),
				Data:    nil,
			},
		)
		return
	}

	webResponse := web.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Berhasil mengambil semua buku",
		Errors:  nil,
		Data:    bookResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}
