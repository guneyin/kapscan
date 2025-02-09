package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/mw"
	"github.com/guneyin/kapscan/internal/service/company"
	"github.com/vcraescu/go-paginator/v2/view"
	"strconv"
)

const (
	companyHandlerName = "company"

	layoutMain = "layouts/main"
)

type Company struct {
	svc *company.Service
}

func newCompanyWebHandler() IHandler {
	svc := company.NewService()

	return &Company{svc}
}

func (cmp *Company) name() string {
	return companyHandlerName
}

func (cmp *Company) setRoutes(router fiber.Router) IHandler {
	router.Get("/", cmp.Index)

	grp := router.Group(cmp.name())
	grp.Get("/", cmp.GetCompanyList)
	grp.Get("/:id", cmp.GetCompany)

	return cmp
}

func (cmp *Company) Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{}, layoutMain)
}

func (cmp *Company) GetCompanyList(c *fiber.Ctx) error {
	page, size := mw.GetPaginate(c)

	_, pageData, err := cmp.svc.GetCompanyList().Offset(page).Limit(size).Do()
	if err != nil {
		return err
	}

	vw := view.New(pageData)

	list := &entity.CompanyList{}
	err = pageData.Results(list)
	if err != nil {
		return err
	}

	pageNavData := NewPageNavData(vw)

	return c.Render("components/company_list", fiber.Map{"CompanyList": list, "PageNavData": pageNavData})
}

func (cmp *Company) GetCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := cmp.svc.GetCompany(id)
	if err != nil {
		return err
	}

	return c.Render("company", fiber.Map{"Company": data}, layoutMain)
}

type PageNavData struct {
	pages                     []int
	next, prev, last, current int
	Items                     []PageNavItem
}

type PageNavItem struct {
	id       int16
	label    string
	active   bool
	disabled bool
}

func (pni PageNavItem) URL() string {
	return fmt.Sprintf("/company?page=%d", pni.id)
}

func (pni PageNavItem) Label() string {
	switch pni.label {
	case "":
		return fmt.Sprintf("%d", pni.id)
	default:
		return pni.label
	}
}

func (pni PageNavItem) Class() string {
	switch pni.active {
	case true:
		return "contrast"
	default:
		return ""
	}
}

func (pni PageNavItem) Disabled() string {
	switch pni.disabled {
	case true:
		return "disabled"
	default:
		return ""
	}
}

func NewPageNavData(vw view.Viewer) PageNavData {
	pages, _ := vw.Pages()
	next, _ := vw.Next()
	prev, _ := vw.Prev()
	last, _ := vw.Last()
	current, _ := vw.Current()

	pnd := PageNavData{
		pages:   pages,
		next:    next,
		prev:    prev,
		last:    last,
		current: current,
		Items:   make([]PageNavItem, 0),
	}
	pnd.buildItems()
	return pnd
}

func (pnd *PageNavData) buildItems() {
	pnd.Items = append(pnd.Items, PageNavItem{
		id:       int16(pnd.prev),
		label:    "<<",
		active:   false,
		disabled: pnd.prev == 0,
	})

	for _, v := range pnd.pages {
		if v > pnd.last {
			continue
		}
		pnd.Items = append(pnd.Items, PageNavItem{
			id:       int16(v),
			label:    strconv.Itoa(v),
			active:   v == pnd.current,
			disabled: false,
		})
	}

	pnd.Items = append(pnd.Items, PageNavItem{
		id:       int16(pnd.next),
		label:    ">>",
		active:   false,
		disabled: pnd.current == pnd.last,
	})
}
