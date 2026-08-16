package service

import (
	"context"

	"golang_crud/exception"
	"golang_crud/model/domain"
	"golang_crud/model/web"
	"golang_crud/repository"

	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	Repository repository.BookRepository
	Validate   *validator.Validate
}

func NewBookService(repository repository.BookRepository, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		Repository: repository,
		Validate:   validate,
	}
}

func (service *BookServiceImpl) Save(ctx context.Context, request web.CreateBookRequest) (web.BookResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.BookResponse{}, err
	}

	bookDomain := domain.Book{
		Name:        request.Name,
		Author:      request.Author,
		Publication: request.Publication,
	}
	book, err := service.Repository.Save(ctx, bookDomain)
	if err != nil {
		return web.BookResponse{}, err
	}

	return web.BookResponse{
		IdBook:      book.IdBook,
		Name:        book.Name,
		Author:      book.Author,
		Publication: book.Publication,
	}, nil
}

func (service *BookServiceImpl) Update(ctx context.Context, request web.UpdateBookRequest) (web.BookResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.BookResponse{}, err
	}

	_, err = service.Repository.FindById(ctx, request.IdBook)
	if err != nil {
		return web.BookResponse{}, &exception.NotFoundErr{ErrMessage: "404 Not Found"}
	}

	bookDomain := domain.Book{
		IdBook:      request.IdBook,
		Name:        request.Name,
		Author:      request.Author,
		Publication: request.Publication,
	}
	book, err := service.Repository.Update(ctx, bookDomain)
	if err != nil {
		return web.BookResponse{}, err
	}

	response := web.BookResponse{
		IdBook:      book.IdBook,
		Name:        book.Name,
		Author:      book.Author,
		Publication: book.Publication,
	}
	return response, nil
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId int) error {
	_, err := service.Repository.FindById(ctx, bookId)
	if err != nil {
		return &exception.NotFoundErr{ErrMessage: "404 Not Found"}
	}

	err = service.Repository.Delete(ctx, bookId)
	if err != nil {
		return err
	}

	return nil
}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId int) (web.BookResponse, error) {
	book, err := service.Repository.FindById(ctx, bookId)
	if err != nil {
		return web.BookResponse{}, &exception.NotFoundErr{ErrMessage: "404 Not Found"}
	}

	res := web.BookResponse{
		IdBook:      book.IdBook,
		Name:        book.Name,
		Author:      book.Author,
		Publication: book.Publication,
	}
	return res, nil
}

func (service *BookServiceImpl) FindAll(ctx context.Context) ([]web.BookResponse, error) {
	books, err := service.Repository.FindAll(ctx)
	if err != nil {
		return []web.BookResponse{}, err
	}

	bookResponse := []web.BookResponse{}
	for _, book := range books {
		res := web.BookResponse{
			IdBook:      book.IdBook,
			Name:        book.Name,
			Author:      book.Author,
			Publication: book.Publication,
		}
		bookResponse = append(bookResponse, res)
	}
	return bookResponse, nil
}
