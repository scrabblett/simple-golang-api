package books

import (
	"context"
	"database/sql"
	"errors"
	"simple-golang-api/internal/domain"
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

func (repo *BookRepo) GetBook(ctx context.Context, id int64) (*model.Book, error) {
	tx, err := repo.db.Begin()

	if err != nil {
		return &model.Book{}, err
	}

	script := `
		SELECT book.id, book.title, book.description, book.age_group, book.publishing_date FROM book WHERE id = $1
    `

	var bookModel model.Book

	err = tx.QueryRowContext(
		ctx, script, id,
	).Scan(&bookModel.Id, &bookModel.Title, &bookModel.Description, &bookModel.AgeGroup, &bookModel.PublishingDate)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &model.Book{}, domain.BookNotFound
		}

		return &model.Book{}, err
	}

	return &bookModel, nil
}

func (repo *BookRepo) UpdateBook(ctx context.Context, id int64, book *model.Book) error {
	tx, err := repo.db.Begin()

	if err != nil {
		return err
	}

	script := `UPDATE book SET title = $1, description = $2, publishing_date = $3, age_group = $4 WHERE id = $5`

	res, err := tx.ExecContext(ctx, script, book.Title, book.Description, book.PublishingDate, book.AgeGroup, id)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.BookNotFound
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (repo *BookRepo) DeleteBook(ctx context.Context, id int64) error {
	tx, err := repo.db.Begin()

	if err != nil {
		return err
	}

	script := `DELETE FROM book WHERE id = $1`

	res, err := tx.ExecContext(ctx, script, id)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.BookNotFound
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
