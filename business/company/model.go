package company

type ICompany struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Problem     string `json:"problem"`
}
