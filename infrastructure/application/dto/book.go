package dto

type (
	StoreBook struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	UpdateBook struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)

func (s *StoreBook) Validate() (bool, error) {
	return true, nil
}

func (s *UpdateBook) Validate() (bool, error) {
	return true, nil
}
