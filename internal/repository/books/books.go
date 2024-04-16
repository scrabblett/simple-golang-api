package books

import (
	"context"
	"database/sql"
	"simple-golang-api/internal/repository/books/model"
)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (repo *BookRepo) InsertBook(ctx context.Context, book *model.Book) error {
	tx, err := repo.db.Begin()

	if err != nil {
		return err
	}

	err = tx.QueryRowContext(ctx, `
		INSERT INTO book (title, description, age_group, publishing_date) VALUES ($1, $2, $3, $4)
		RETURNING id
	`, book.Title, book.Description, book.AgeGroup, book.PublishingDate).Scan(&book.Id)

	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
