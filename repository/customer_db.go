package repository

import (
	"github.com/jmoiron/sqlx"
)

type customerRepositoryDb struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDb(db *sqlx.DB) CustomerRepository {
	return &customerRepositoryDb{db: db}
}

func (r customerRepositoryDb) GetCustomers() ([]Customer, error) {

	customers := []Customer{}

	query := "select * from table"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r customerRepositoryDb) GetCustomer(id int) (*Customer, error) {

	customer := Customer{}
	query := "select * from table where id=?"
	err := r.db.Get(&customer, query, id)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}
