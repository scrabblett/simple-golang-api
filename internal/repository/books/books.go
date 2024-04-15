package books

import (
	"awesomeProject/internal/repository/books/model"
	"context"
	"database/sql"
	"go.uber.org/zap"
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
		zap.L().Error("failed to create transaction", zap.Error(err))

		return err
	}

	err = tx.QueryRowContext(ctx, `
		INSERT INTO book (title, description, age_group, publishing_date) VALUES ($1, $2, $3, $4)
		RETURNING id
	`, book.Title, book.Description, book.AgeGroup, book.PublishingDate).Scan(&book.Id)

	if err != nil {
		zap.L().Error("failed to insert book", zap.Error(err))

		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
