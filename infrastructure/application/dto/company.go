package dto

type (
	StoreCompany struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Problem     string `json:"problem"`
	}

	UpdateCompany struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Problem     string `json:"problem"`
	}
)

func (sc *StoreCompany) Validate() (bool, error) {
	return true, nil
}

func (sc *UpdateCompany) Validate() (bool, error) {
	return true, nil
}
