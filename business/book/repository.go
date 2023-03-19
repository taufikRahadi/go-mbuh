package book

type IRepo interface {
	Store(payload IStoreBook) (IBook, error)
	FindAll(query string) ([]IBook, error)
	FindOne(id int) (IBook, error)
	Update(id int, payload IUpdateBook) (IBook, error)
	Delete(id int) error
}
