package scraper

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/imroc/req/v3"
)

const (
	baseURL = "https://www.kap.org.tr"
)

type Scraper struct {
	client *req.Client
}

type Result struct {
	Error error
	Body  io.Reader
}

func New() *Scraper {
	client := req.NewClient().
		// DevMode().
		SetBaseURL(baseURL).
		EnableForceHTTP1()
	return &Scraper{
		client: client,
	}
}

func newResult(err error, body io.Reader) *Result {
	return &Result{
		Error: err,
		Body:  body,
	}
}

func (s *Scraper) Fetch(ctx context.Context, method, url string, reqBody, resBody any) *Result {
	res, err := s.client.R().
		SetContext(ctx).
		SetContentType("application/json").
		SetBody(reqBody).
		SetSuccessResult(resBody).
		Send(method, url)
	if err != nil {
		return newResult(err, nil)
	}
	if res.StatusCode != http.StatusOK {
		return newResult(errors.New(res.Status), nil)
	}

	return newResult(nil, res.Body)
}
