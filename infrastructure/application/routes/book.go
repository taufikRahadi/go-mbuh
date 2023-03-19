package routes

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
	"quiz-1/infrastructure/application/presenters"
	"quiz-1/infrastructure/use_case/book"
)

func InitBookRoute(db *sql.DB, api *mux.Router) {
	repo := book.NewRepo(db)
	service := book.NewService(&repo)
	presenter := presenters.NewBookPresenter(&service)

	// list books
	api.HandleFunc("/book", presenter.ListBooks).Methods(http.MethodGet)

	// store book
	api.HandleFunc("/book", presenter.StoreBook).Methods(http.MethodPost)

	//	get book by id
	api.HandleFunc("/book/{id}", presenter.GetBookById).Methods(http.MethodGet)

	// update book
	api.HandleFunc("/book/{id}", presenter.UpdateBook).Methods(http.MethodPut)

	// delete book
	api.HandleFunc("/book/{id}", presenter.DeleteBook).Methods(http.MethodDelete)
}
