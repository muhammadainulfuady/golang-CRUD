package web

type CreateBookRequest struct {
	Name        string `json:"name" validate:"required,min=4,max=150"`
	Author      string `json:"author" validate:"required,min=4,max=100"`
	Publication string `json:"publication" validate:"required,datetime=206-01-02"`
}

type UpdateBookRequest struct {
	IdBook      int    `json:"id_book"`
	Name        string `json:"name" validate:"required,min=4,max=150"`
	Author      string `json:"author" validate:"required,min=4,max=100"`
	Publication string `json:"publication" validate:"required,datetime=2006-01-02"`
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
	Errors  any    `json:"errors"`
	Data    any    `json:"data"`
}
