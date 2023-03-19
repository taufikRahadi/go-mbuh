package routes

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
	"quiz-1/infrastructure/application/presenters"
	"quiz-1/infrastructure/use_case/company"
)

func InitCompanyRoute(db *sql.DB, api *mux.Router) {
	repo := company.NewRepo(db)
	service := company.NewService(&repo)
	presenter := presenters.NewCompanyPresenter(&service)

	api.HandleFunc("/company", presenter.ListCompany).Methods(http.MethodGet)
	api.HandleFunc("/company", presenter.StoreCompany).Methods(http.MethodPost)
	api.HandleFunc("/company/{id}", presenter.GetCompanyById).Methods(http.MethodGet)
	api.HandleFunc("/company/{id}", presenter.UpdateCompany).Methods(http.MethodPut)
	api.HandleFunc("/company/{id}", presenter.DeleteCompany).Methods(http.MethodDelete)
}
