package scanner

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestDocumentSelector(t *testing.T) {
	f, err := os.Open("testdata/8acae2c59145e00a01915a44a01d1179.html")
	if err != nil {
		t.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(f)
	assert.Nil(t, err)

	selector := ".exportClass > div:contains('Ortağın Adı')"
	doc.Find(selector).Parent().Each(func(i int, s *goquery.Selection) {
		s.Find(".w-clearfix.w-inline-block.a-table-row.infoRow").Each(func(i int, s *goquery.Selection) {
			name := strings.TrimSpace(s.Find("div:nth-child(1)").Text())
			shareByAmount := strings.TrimSpace(s.Find("div:nth-child(2)").Text())
			shareByRatio := strings.TrimSpace(s.Find("div:nth-child(3)").Text())
			votingRight := strings.TrimSpace(s.Find("div:nth-child(4)").Text())

			t.Logf("%50.50s %30.30s %30.30s %30.30s\n", name, shareByAmount, shareByRatio, votingRight)
		})
	})
}
