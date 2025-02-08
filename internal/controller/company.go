package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/kapscan/internal/mw"
	"github.com/guneyin/kapscan/internal/service/company"
)

const companyControllerName = "company"

type Company struct {
	svc *company.Service
}

func newCompanyController() IController {
	svc := company.NewService()

	return &Company{svc}
}

func (cmp *Company) name() string {
	return companyControllerName
}

func (cmp *Company) setRoutes(router fiber.Router) IController {
	grp := router.Group(cmp.name())

	grp.Get("/", cmp.GetList)
	grp.Get("/:id", cmp.GetByID)

	return cmp
}

func (cmp *Company) GetList(c *fiber.Ctx) error {
	offset, limit := mw.GetPaginate(c)

	res, err := cmp.svc.GetCompanyList().Offset(offset).Limit(limit).Do()
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (cmp *Company) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := cmp.svc.GetCompany(id)
	if err != nil {
		return mw.Error(c, err)
	}

	return mw.OK(c, "company fetched", res)
}
