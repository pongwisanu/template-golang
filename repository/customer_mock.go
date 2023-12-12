package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {

	cus := []Customer{
		{Id: 1001, Name: "Test1"},
		{Id: 1002, Name: "Test2"},
	}

	return &customerRepositoryMock{customers: cus}
}

func (r customerRepositoryMock) GetCustomers() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetCustomer(id int) (*Customer, error) {

	for _, vm := range r.customers {
		if vm.Id == id {
			return &vm, nil
		}
	}

	return nil, errors.New("Customer Not found")
}
