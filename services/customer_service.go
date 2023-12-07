package services

import (
	"log"
	"nida/vmwareapi/repository"
)

type customerService struct {
	cusRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return &customerService{cusRepo: customerRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {

	vms, err := s.cusRepo.GetCustomers()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	vmReponses := []CustomerResponse{}

	for _, vm := range vms {
		vmReponse := CustomerResponse{
			Id:        vm.Id,
			Name:      vm.Name,
			InService: "True",
		}
		vmReponses = append(vmReponses, vmReponse)
	}

	return vmReponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {

	vm, err := s.cusRepo.GetCustomer(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	vmResponse := CustomerResponse{
		Id:        vm.Id,
		Name:      vm.Name,
		InService: "True",
	}

	return &vmResponse, nil

}
