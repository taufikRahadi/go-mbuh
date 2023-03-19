package book

import (
	"quiz-1/business/book"
)

type service struct {
	repo book.IRepo
}

func NewService(repo *book.IRepo) book.IService {
	return &service{repo: *repo}
}

func (s *service) Store(payload book.IStoreBook) (book.IBook, error) {
	return s.repo.Store(payload)
}

func (s *service) FindAll(query string) ([]book.IBook, error) {
	return s.repo.FindAll(query)
}

func (s *service) FindOne(id int) (book.IBook, error) {
	return s.repo.FindOne(id)
}

func (s *service) Update(id int, payload book.IUpdateBook) (book.IBook, error) {
	return s.repo.Update(id, payload)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}
