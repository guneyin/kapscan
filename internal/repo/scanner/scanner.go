package scanner

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/guneyin/gobist"
	"github.com/guneyin/kapscan/internal/dto"
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/scraper"
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
	list, err := bist.GetSymbolList()
	if err != nil {
		return nil, err
	}

	symbolList := make(entity.CompanyList, list.Count)
	for i, symbol := range list.Items {
		symbolList[i] = entity.Company{
			Code: symbol.Code,
			Name: symbol.Name,
			Icon: symbol.Icon,
		}
	}

	return symbolList, nil
}

func (r *Repo) GetCompany(ctx context.Context, symbolCode string) ([]dto.ShareHolder, error) {
	fs, err := r.fetchCompany(ctx, symbolCode)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("/tr/sirket-bilgileri/genel/%s", fs.MemberOrFundOid)

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

	res := make([]dto.ShareHolder, 0)
	list.Each(func(i int, s *goquery.Selection) {
		s.Find(".w-clearfix.w-inline-block.a-table-row.infoRow").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			res = append(res, dto.ShareHolder{
				Name:            strings.TrimSpace(s.Find("div:nth-child(1)").Text()),
				CapitalByAmount: strings.TrimSpace(s.Find("div:nth-child(2)").Text()),
				CapitalByVolume: strings.TrimSpace(s.Find("div:nth-child(3)").Text()),
			})
		})
	})

	return res, nil
}

func (r *Repo) fetchCompany(ctx context.Context, symbol string) (*dto.SymbolResultItem, error) {
	req := dto.SymbolRequest{
		Keyword:   symbol,
		DiscClass: "ALL",
		Lang:      "tr",
		Channel:   "WEB",
	}
	res := make([]dto.SymbolResponse, 0)

	sri := dto.SymbolResultItem{}

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
				if strings.ToUpper(result.CmpOrFundCode) == strings.ToUpper(symbol) {
					sri = result
					break loop
				}
			}
		}
	}

	if sri.MemberOrFundOid == "" {
		return nil, errors.New("symbol not found")
	}

	return &sri, nil
}
