package scanner

import (
	"context"
	"errors"
	"github.com/imroc/req/v3"
	"io"
)

const (
	baseUrl = "https://www.kap.org.tr"
)

type Scraper struct {
	client *req.Client
}

type SR struct {
	error error
	body  io.Reader
}

func NewScraper() *Scraper {
	client := req.NewClient().
		//DevMode().
		SetBaseURL(baseUrl).
		EnableForceHTTP1()
	return &Scraper{
		client: client,
	}
}

func NewSR(error error, body io.Reader) *SR {
	return &SR{
		error: error,
		body:  body,
	}
}

func (s *Scraper) Fetch(ctx context.Context, method, url string, reqBody, resBody any) *SR {
	res, err := s.client.R().
		SetContext(ctx).
		SetContentType("application/json").
		SetBody(reqBody).
		SetSuccessResult(resBody).
		Send(method, url)
	if err != nil {
		return NewSR(err, nil)
	}
	if res.StatusCode != 200 {
		return NewSR(errors.New(res.Status), nil)
	}

	return NewSR(nil, res.Body)
}
