package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/kapscan/internal/entity"
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
	grp.Get("/:code", cmp.GetByCode)

	return cmp
}

func (cmp *Company) GetList(c *fiber.Ctx) error {
	offset, limit := mw.GetPaginate(c)
	s := c.Query("search")

	cl, err := cmp.svc.Search(s).
		Offset(offset).
		Limit(limit).
		Do(c.Context())
	if err != nil {
		return err
	}

	companyList := entity.CompanyList{}
	err = cl.DataAs(companyList)
	if err != nil {
		return err
	}

	return c.JSON(companyList)
}

func (cmp *Company) GetByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	data, err := cmp.svc.GetByCode(c.Context(), code)
	if err != nil {
		return mw.Error(c, err)
	}

	return mw.OK(c, "company fetched", data)
}
