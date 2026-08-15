package web

type CreateBookRequest struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
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
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
