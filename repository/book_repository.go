package repository

import (
	"context"
	"golang_crud/model/domain"
)

type BookRepository interface {
	Save(ctx context.Context, book domain.Book) (domain.Book, error)
	Update(ctx context.Context, book domain.Book) (domain.Book, error)
	Delete(ctx context.Context, bookId int) error
	FindById(ctx context.Context, bookId int) (domain.Book, error)
	FindAll(ctx context.Context) ([]domain.Book, error)
}
