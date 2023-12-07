package repository

import "errors"

type customerRepositoryConn struct {
	customers []Customer
}

func NewCustomerRepositoryConn() CustomerRepository {

	cus := []Customer{
		{Id: 1001, Name: "Test1"},
		{Id: 1002, Name: "Test2"},
	}

	return &customerRepositoryConn{customers: cus}
}

func (r customerRepositoryConn) GetCustomers() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryConn) GetCustomer(id int) (*Customer, error) {

	for _, vm := range r.customers {
		if vm.Id == id {
			return &vm, nil
		}
	}

	return nil, errors.New("Customer Not found")
}
