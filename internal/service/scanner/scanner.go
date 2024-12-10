package scanner

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type Service struct {
	db      *gorm.DB
	scraper *Scraper
}

func NewScannerService(db *gorm.DB) *Service {
	return &Service{
		db:      db,
		scraper: NewScraper(),
	}
}

func (s *Service) Scan(ctx context.Context, symbol string) ([]ShareHolder, error) {
	fs, err := s.fetchSymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("/tr/sirket-bilgileri/genel/%s", fs.MemberOrFundOid)

	sr := s.scraper.Fetch(ctx, http.MethodGet, url, nil, nil)
	if sr.error != nil {
		return nil, sr.error
	}

	doc, err := goquery.NewDocumentFromReader(sr.body)
	if err != nil {
		return nil, err
	}

	selector := ".exportClass > div:contains('Ortağın Adı')"
	list := doc.Find(selector).Parent()

	res := make([]ShareHolder, 0)
	list.Each(func(i int, s *goquery.Selection) {
		s.Find(".w-clearfix.w-inline-block.a-table-row.infoRow").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			res = append(res, ShareHolder{
				Name:            strings.TrimSpace(s.Find("div:nth-child(1)").Text()),
				CapitalByAmount: strings.TrimSpace(s.Find("div:nth-child(2)").Text()),
				CapitalByVolume: strings.TrimSpace(s.Find("div:nth-child(3)").Text()),
			})
		})
	})

	return res, nil
}

func (s *Service) fetchSymbol(ctx context.Context, symbol string) (*SymbolResultItem, error) {
	req := SymbolRequest{
		Keyword:   symbol,
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
		sr := s.scraper.Fetch(ctx, http.MethodPost, url, req, &res)
		if sr.error != nil {
			return nil, sr.error
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
