package company

type IRepo interface {
	Store(payload IStoreCompany) (ICompany, error)
	FindAll() ([]ICompany, error)
	FindOne(id int) (ICompany, error)
	Update(id int, payload IUpdateCompany) (ICompany, error)
	Delete(id int) error
}
