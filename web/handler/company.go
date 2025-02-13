package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/kapscan/internal/dto"
	"github.com/guneyin/kapscan/internal/mw"
	"github.com/guneyin/kapscan/internal/service/company"
	"github.com/vcraescu/go-paginator/v2/view"
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
	grp.Get("/", cmp.CompanySearch)
	grp.Get("/:code", cmp.CompanyDetail)

	return cmp
}

func (cmp *Company) Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{}, layoutMain)
}

func (cmp *Company) CompanySearch(c *fiber.Ctx) error {
	search := c.Query("search")
	page, size := mw.GetPaginate(c)

	cl, err := cmp.svc.Search(search).
		Offset(page).
		Limit(size).
		Do(c.Context())
	if err != nil {
		return err
	}

	companyList := &dto.CompanyList{}
	err = cl.DataAs(companyList)
	if err != nil {
		return err
	}

	vw := view.New(cl.PageData())
	pageNavData := NewPageNavData(vw)

	return c.Render("components/company_search", fiber.Map{"CompanyList": companyList, "PageNavData": pageNavData})
}

func (cmp *Company) CompanyDetail(c *fiber.Ctx) error {
	code := c.Params("code")
	data, err := cmp.svc.GetByCode(c.Context(), code)
	if err != nil {
		return err
	}

	return c.Render("pages/company/detail", fiber.Map{"Company": data}, layoutMain)
}

type PageNavData struct {
	pages                     []int
	next, prev, last, current int
	Items                     []PageNavItem
}

type PageNavItem struct {
	id       int
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
		return strconv.Itoa(pni.id)
	default:
		return pni.label
	}
}

func (pni PageNavItem) Active() string {
	switch pni.active {
	case true:
		return "active"
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
		id:       pnd.prev,
		label:    "<<",
		active:   false,
		disabled: pnd.prev == 0,
	})

	for _, v := range pnd.pages {
		if v > pnd.last {
			continue
		}
		pnd.Items = append(pnd.Items, PageNavItem{
			id:       v,
			label:    strconv.Itoa(v),
			active:   v == pnd.current,
			disabled: false,
		})
	}

	pnd.Items = append(pnd.Items, PageNavItem{
		id:       pnd.next,
		label:    ">>",
		active:   false,
		disabled: pnd.current == pnd.last,
	})
}
