package scanner

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/guneyin/gobist"
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/scraper"
	"github.com/guneyin/kapscan/util"
	"net/http"
	"strings"
)

type Repo struct {
	scraper *scraper.Scraper
}

func NewRepo() *Repo {
	return &Repo{scraper: scraper.New()}
}

func (r *Repo) GetCompanyList() (entity.CompanyList, error) {
	bist := gobist.New()
	symbolList, err := bist.GetSymbolList()
	if err != nil {
		return nil, err
	}

	cl := make(entity.CompanyList, symbolList.Count)
	for i, symbol := range symbolList.Items {
		cl[i] = entity.Company{
			Code: symbol.Code,
			Name: symbol.Name,
			Icon: symbol.Icon,
		}
	}

	return cl, nil
}

func (r *Repo) SyncCompany(ctx context.Context, cmp *entity.Company) error {
	fs, err := r.fetchCompany(ctx, cmp.Code)
	if err != nil {
		return err
	}
	cmp.MemberID = fs.MemberOrFundOid

	url := fmt.Sprintf("/tr/sirket-bilgileri/ozet/%s", fs.MemberOrFundOid)

	sr := r.scraper.Fetch(ctx, http.MethodGet, url, nil, nil)
	if sr.Error != nil {
		return sr.Error
	}

	doc, err := goquery.NewDocumentFromReader(sr.Body)
	if err != nil {
		return err
	}

	selector := ".w-clearfix.w-inline-block.a-table-row.infoRow"
	list := doc.Find(selector)
	list.Each(func(i int, s *goquery.Selection) {
		val := strings.TrimSpace(s.Find("div:nth-child(2)").Text())
		switch i {
		case 0:
			cmp.Address = val
		case 1:
			cmp.Email = val
		case 2:
			cmp.Website = val
		case 5:
			cmp.Index = val
		case 6:
			cmp.Sector = val
		case 7:
			cmp.Market = val
		}
	})

	return nil
}

func (r *Repo) GetCompanyShare(ctx context.Context, cmp entity.Company) ([]entity.CompanyShare, error) {
	url := fmt.Sprintf("/tr/sirket-bilgileri/genel/%s", cmp.MemberID)

	sr := r.scraper.Fetch(ctx, http.MethodGet, url, nil, nil)
	if sr.Error != nil {
		return nil, sr.Error
	}

	doc, err := goquery.NewDocumentFromReader(sr.Body)
	if err != nil {
		return nil, err
	}

	selector := ".exportClass > div:contains('Ortağın Adı')"
	list := doc.Find(selector).Parent()

	res := make([]entity.CompanyShare, 0)
	list.Each(func(i int, s *goquery.Selection) {
		s.Find(".w-clearfix.w-inline-block.a-table-row.infoRow").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			res = append(res, entity.CompanyShare{
				CompanyID:       cmp.ID,
				Title:           strings.TrimSpace(s.Find("div:nth-child(1)").Text()),
				CapitalByAmount: util.NewMoney(s.Find("div:nth-child(2)").Text()).Float64(),
				CapitalByVolume: util.NewMoney(s.Find("div:nth-child(3)").Text()).Float64(),
				VoteRight:       util.NewMoney(s.Find("div:nth-child(4)").Text()).Float64(),
			})
		})
	})

	return res, nil
}

func (r *Repo) fetchCompany(ctx context.Context, code string) (*SymbolResultItem, error) {
	req := SymbolRequest{
		Keyword:   code,
		DiscClass: "ALL",
		Lang:      "tr",
		Channel:   "WEB",
	}
	res := make([]SymbolResponse, 0)

	sri := SymbolResultItem{}

	keys := []string{"combined", "smart"}
loop:
	for _, key := range keys {
		url := fmt.Sprintf("/kapsrc/%s", key)
		sr := r.scraper.Fetch(ctx, http.MethodPost, url, req, &res)
		if sr.Error != nil {
			return nil, sr.Error
		}

		for _, cr := range res {
			for _, result := range cr.Results {
				sliced := strings.Split(result.CmpOrFundCode, ",")
				for _, symbol := range sliced {
					if strings.ToUpper(symbol) == strings.ToUpper(code) {
						sri = result
						break loop
					}
				}
			}
		}
	}

	if sri.MemberOrFundOid == "" {
		return nil, errors.New("company not found")
	}

	return &sri, nil
}
