package services

type CustomerResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	InService string `json:"inservice"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}
