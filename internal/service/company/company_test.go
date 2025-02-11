package company_test

import (
	"os"
	"strings"
	"testing"

	"github.com/guneyin/kapscan/internal/service/company"

	"github.com/PuerkitoBio/goquery"
	"github.com/guneyin/kapscan/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDocumentSelector(t *testing.T) {
	f, err := os.Open("testdata/8acae2c59145e00a01915a44a01d1179.html")
	if err != nil {
		t.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(f)
	require.NoError(t, err)

	selector := ".exportClass > div:contains('Ortağın Adı')"
	doc.Find(selector).Parent().Each(func(_ int, s *goquery.Selection) {
		s.Find(".w-clearfix.w-inline-block.a-table-row.infoRow").Each(func(_ int, s *goquery.Selection) {
			name := strings.TrimSpace(s.Find("div:nth-child(1)").Text())
			shareByAmount := strings.TrimSpace(s.Find("div:nth-child(2)").Text())
			shareByRatio := strings.TrimSpace(s.Find("div:nth-child(3)").Text())
			votingRight := strings.TrimSpace(s.Find("div:nth-child(4)").Text())

			t.Logf("%50.50s %30.30s %30.30s %30.30s\n", name, shareByAmount, shareByRatio, votingRight)
		})
	})
}

func TestService_FetchCompanyList(t *testing.T) {
	svc := company.NewService()

	cl, err := svc.Search("").Do()
	require.NoError(t, err)
	assert.NotNil(t, cl)

	companyList := dto.CompanyList{}
	err = cl.DataAs(&companyList)
	require.NoError(t, err)

	for _, cmp := range companyList {
		t.Logf("%s\n", cmp.Code)
	}
}
