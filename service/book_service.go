package service

import (
	"context"

	"golang_crud/model/web"
)

type BookService interface {
	Save(ctx context.Context, request web.CreateBookRequest) (web.BookResponse, error)
	Update(ctx context.Context, request web.UpdateBookRequest) (web.BookResponse, error)
	Delete(ctx context.Context, bookId int) error
	FindById(ctx context.Context, bookId int) (web.BookResponse, error)
	FindAll(ctx context.Context) ([]web.BookResponse, error)
}
