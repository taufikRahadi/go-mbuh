package book

import (
	"database/sql"
	"fmt"
	"quiz-1/business/book"
)

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) book.IRepo {
	return &repository{db: db}
}

func (r *repository) Store(payload book.IStoreBook) (book.IBook, error) {
	b := book.IBook{}
	querySt := `INSERT INTO books (name, description) VALUES ($1, $2)`
	_, err := r.db.Exec(querySt, payload.Name, payload.Description)

	if err != nil {
		return b, err
	}

	return b, nil
}

func (r *repository) FindAll(query string) ([]book.IBook, error) {
	books := []book.IBook{}
	querySt := `
		SELECT * FROM books
	`

	if len(query) > 0 {
		querySt += " WHERE name ILIKE $1"
		rows, err := r.db.Query(querySt, "%"+query+"%")
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			var b book.IBook
			err := rows.Scan(&b.ID, &b.Name, &b.Description)
			if err != nil {
				return nil, err
			}
			books = append(books, b)
		}

		return books, nil
	} else {
		rows, err := r.db.Query(querySt)
		if err != nil {
			return nil, err
		}
		fmt.Println(err)

		defer rows.Close()

		for rows.Next() {
			var b book.IBook
			err := rows.Scan(&b.ID, &b.Name, &b.Description)
			if err != nil {
				return nil, err
			}
			books = append(books, b)
		}

		return books, nil
	}
}

func (r *repository) FindOne(id int) (book.IBook, error) {
	var b book.IBook
	querySt := `
		SELECT * FROM books WHERE id = $1 LIMIT 1
	`

	row := r.db.QueryRow(querySt, id)
	err := row.Scan(&b.ID, &b.Name, &b.Description)
	if err != nil {
		return b, err
	}

	return b, nil
}

func (r *repository) Update(id int, payload book.IUpdateBook) (book.IBook, error) {
	var b book.IBook
	querySt := `
		UPDATE books
		SET name = $1,
		    description = $2
		WHERE id = $3
	`
	_, err := r.db.Exec(querySt, payload.Name, payload.Description, id)
	if err != nil {
		return b, err
	}

	return b, nil
}

func (r *repository) Delete(id int) error {
	querySt := `
		DELETE FROM books
		WHERE id = $1
	`
	_, err := r.db.Exec(querySt, id)
	return err
}
