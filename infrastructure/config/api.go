package config

import (
	"database/sql"
	"github.com/gorilla/mux"
	"quiz-1/infrastructure/application/routes"
)

func InitApi(db *sql.DB, api *mux.Router) {
	// path /company/*
	//routes.InitCompanyRoute(db, api)

	//	path /book/*
	routes.InitBookRoute(db, api)
}
