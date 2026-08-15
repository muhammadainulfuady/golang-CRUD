package web

type CreateBookRequest struct {
	Name        string `json:"name" validate:"required, min:4, max:150"`
	Author      string `json:"author" validate:"required, min:4, max:100"`
	Publication string `json:"publication" validate:"required, datetime:2026-07-15"`
}

type UpdateBookRequest struct {
	IdBook      int    `json:"id_book"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type BookResponse struct {
	IdBook      int    `json:"id_book"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type ApiResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
