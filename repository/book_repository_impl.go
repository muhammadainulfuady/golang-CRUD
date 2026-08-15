package repository

import (
	"context"
	"database/sql"

	"golang_crud/model/domain"
)

type BookRepositoryImpl struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &BookRepositoryImpl{DB: db}
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, book domain.Book) (domain.Book, error) {
	SQL := "INSERT INTO book(name, author, publication) VALUES (?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, SQL, book.Name, book.Author, book.Publication)
	if err != nil {
		return book, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return book, err
	}

	book.IdBook = int(id)
	return book, nil
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, book domain.Book) (domain.Book, error) {
	SQL := "UPDATE book SET name = ?, author = ?, publication = ? WHERE id_book = ?"
	_, err := repository.DB.ExecContext(ctx, SQL, book.Name, book.Author, book.Publication, book.IdBook)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, bookId int) error {
	SQL := "DELETE FROM book WHERE id_book = ?"
	_, err := repository.DB.ExecContext(ctx, SQL, bookId)
	if err != nil {
		return err
	}

	return nil
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, bookId int) (domain.Book, error) {
	book := domain.Book{}

	SQL := "SELECT id_book, name, author, publication FROM book WHERE id_book = ?"
	err := repository.DB.QueryRowContext(ctx, SQL, bookId).Scan(
		&book.IdBook,
		&book.Name,
		&book.Author,
		&book.Publication,
	)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context) ([]domain.Book, error) {
	SQL := "SELECT id_book, name, author, publication FROM book"

	rows, err := repository.DB.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []domain.Book{}
	for rows.Next() {
		book := domain.Book{}
		err := rows.Scan(
			&book.IdBook,
			&book.Name,
			&book.Author,
			&book.Publication,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
