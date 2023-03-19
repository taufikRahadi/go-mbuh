package company

import "quiz-1/business/company"

type service struct {
	repo company.IRepo
}

func (s *service) Store(payload company.IStoreCompany) (company.ICompany, error) {
	return s.repo.Store(payload)
}

func (s *service) FindAll() ([]company.ICompany, error) {
	return s.repo.FindAll()
}

func (s *service) FindOne(id int) (company.ICompany, error) {
	return s.repo.FindOne(id)
}

func (s *service) Update(id int, payload company.IUpdateCompany) (company.ICompany, error) {
	return s.repo.Update(id, payload)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

func NewService(repo *company.IRepo) company.IService {
	return &service{repo: *repo}
}
