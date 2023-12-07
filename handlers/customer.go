package handlers

import (
	"fmt"
	"nida/vmwareapi/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type customerHandle struct {
	cusSrv services.CustomerService
}

func NewCustomerHandler(cusSrv services.CustomerService) customerHandle {
	return customerHandle{cusSrv: cusSrv}
}

func (h customerHandle) GetCustomers(c *fiber.Ctx) error {
	cus, err := h.cusSrv.GetCustomers()

	if err != nil {
		fmt.Print(err)
		return c.SendStatus(500)
	}

	c.JSON(cus)
	return c.SendStatus(200)
}

func (h customerHandle) GetCustomer(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	cus, err := h.cusSrv.GetCustomer(id)

	if err != nil {
		fmt.Println(err)
		return c.SendStatus(404)
	}

	c.JSON(cus)
	return c.SendStatus(200)

}
