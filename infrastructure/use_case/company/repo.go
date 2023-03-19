package company

import (
	"database/sql"
	"quiz-1/business/company"
	"strconv"
)

type repository struct {
	db *sql.DB
}

func (r *repository) Store(payload company.IStoreCompany) (company.ICompany, error) {
	c := company.ICompany{}
	querySt := `INSERT INTO companies (name, description, problem) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(querySt, payload.Name, payload.Description, payload.Problem)

	if err != nil {
		return c, err
	}

	return c, nil
}

func (r *repository) FindAll() ([]company.ICompany, error) {
	companies := []company.ICompany{}
	querySt := `
		SELECT * FROM companies
	`
	rows, err := r.db.Query(querySt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var c company.ICompany
		err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.Problem)
		if err != nil {
			return nil, err
		}
		companies = append(companies, c)
	}

	return companies, nil
}

func (r *repository) FindOne(id int) (company.ICompany, error) {
	var c company.ICompany
	querySt := `
		SELECT * FROM companies WHERE id = $1 LIMIT 1
	`
	row := r.db.QueryRow(querySt, strconv.Itoa(id))
	err := row.Scan(&c.ID, &c.Name, &c.Description, &c.Problem)
	if err != nil {
		return company.ICompany{}, err
	}

	return c, nil
}

func (r *repository) Update(id int, payload company.IUpdateCompany) (company.ICompany, error) {
	var c company.ICompany
	querySt := `
		UPDATE companies
		SET name = $1,
		    description = $2,
		    problem = $3
		WHERE id = $4
	`

	_, err := r.db.Exec(querySt, payload.Name, payload.Description, payload.Problem, id)
	if err != nil {
		return c, err
	}

	data, _ := r.FindOne(id)

	return data, nil
}

func (r *repository) Delete(id int) error {
	querySt := `
		DELETE FROM companies
		WHERE id = $1
	`

	_, err := r.db.Exec(querySt, id)
	return err
}

func NewRepo(db *sql.DB) company.IRepo {
	return &repository{db: db}
}
