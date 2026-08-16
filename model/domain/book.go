package domain

type Book struct {
	IdBook      int    `json:"id_book"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}
