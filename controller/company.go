package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/kapscan/dto"
	"github.com/guneyin/kapscan/entity"
	"github.com/guneyin/kapscan/mw"
	"github.com/guneyin/kapscan/service/company"
	"github.com/guneyin/kapscan/util"
)

const companyControllerName = "company"

type Company struct {
	svc *company.Service
}

func newCompanyController() IController {
	svc := company.NewService()

	return &Company{svc}
}

func (cm *Company) name() string {
	return companyControllerName
}

func (cm *Company) setRoutes(router fiber.Router) IController {
	grp := router.Group(cm.name())

	grp.Get("/", cm.GetList)
	grp.Get("/:code", cm.GetByCode)

	return cm
}

func (cm *Company) GetList(c *fiber.Ctx) error {
	offset, limit := mw.GetPaginate(c)
	s := c.Query("search")

	cl, err := cm.svc.Search(s).
		Offset(offset).
		Limit(limit).
		Do(c.Context())
	if err != nil {
		return err
	}

	companyList := entity.CompanyList{}
	err = cl.DataAs(companyList)
	if err != nil {
		return mw.Error(c, err)
	}

	return c.JSON(companyList)
}

func (cm *Company) GetByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	cmp, err := cm.svc.GetByCode(c.Context(), code)
	if err != nil {
		return mw.Error(c, err)
	}

	data, err := util.Convert(cmp, &dto.Company{})
	if err != nil {
		return mw.Error(c, err)
	}

	return mw.OK(c, "company fetched", data)
}
