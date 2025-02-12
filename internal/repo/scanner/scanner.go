package scanner

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/guneyin/gobist"
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/scraper"
)

type Repo struct {
	scraper *scraper.Scraper
}

func NewRepo() *Repo {
	return &Repo{scraper: scraper.New()}
}

func (r *Repo) FetchCompanyList() (entity.CompanyList, error) {
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

func (r *Repo) SyncCompany(ctx context.Context, cmp *entity.Company) (*entity.Company, error) {
	err := r.syncCompany(ctx, cmp)
	if err != nil {
		return nil, err
	}

	err = r.syncCompanyDetail(ctx, cmp)
	if err != nil {
		return nil, err
	}

	err = r.syncCompanyShares(ctx, cmp)
	if err != nil {
		return nil, err
	}

	return cmp, nil
}

func (r *Repo) syncCompany(ctx context.Context, cmp *entity.Company) error {
	req := SymbolRequest{
		Keyword:   cmp.Code,
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
			return sr.Error
		}

		for _, cr := range res {
			for _, result := range cr.Results {
				sliced := strings.Split(result.CmpOrFundCode, ",")
				for _, symbol := range sliced {
					if strings.EqualFold(symbol, cmp.Code) {
						sri = result
						break loop
					}
				}
			}
		}
	}

	if sri.MemberOrFundOid == "" {
		return errors.New("company not found")
	}

	cmp.MemberID = sri.MemberOrFundOid

	return nil
}

func (r *Repo) syncCompanyDetail(ctx context.Context, cmp *entity.Company) error {
	url := fmt.Sprintf("/tr/sirket-bilgileri/ozet/%s", cmp.MemberID)

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

func (r *Repo) syncCompanyShares(ctx context.Context, cmp *entity.Company) error {
	url := fmt.Sprintf("tr/infoHistory/kpy41_acc5_sermayede_dogrudan/%s", cmp.MemberID)

	sr := r.scraper.Fetch(ctx, http.MethodGet, url, nil, nil)
	if sr.Error != nil {
		return sr.Error
	}

	doc, err := goquery.NewDocumentFromReader(sr.Body)
	if err != nil {
		return err
	}

	var shareDate time.Time
	selector := ".modal-info.my-modal-info > div > a"
	list := doc.Find(selector)
	list.Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}

		if dt, ok := isDate(s); ok {
			shareDate = *dt
			return
		}

		diff := time.Now().Year() - shareDate.Year()
		if diff > 2 {
			return
		}

		if cs, ok := parseLineAsCompanyShare(s); ok {
			cs.CompanyID = cmp.ID
			cs.Date = shareDate

			cmp.AddShare(*cs)
		}
	})

	return nil
}
