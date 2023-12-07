package repository

type Customer struct {
	Id   int
	Name string
}

type CustomerRepository interface {
	GetCustomers() ([]Customer, error)
	GetCustomer(int) (*Customer, error)
}
