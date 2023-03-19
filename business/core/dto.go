package core

type IDto interface {
	Validate() (bool, error)
}
