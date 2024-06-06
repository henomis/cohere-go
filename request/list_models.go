package request

import (
	"io"

	"github.com/henomis/restclientgo"
)

type ListModels struct {
	PageSize    int    `json:"-"`
	PageToken   string `json:"-"`
	Endpoint    string `json:"-"`
	DefaultOnly bool   `json:"-"`
}

func (c *ListModels) Path() (string, error) {
	urlValues := restclientgo.NewURLValues()

	if c.PageSize > 0 {
		urlValues.AddInt("page_size", &c.PageSize)
	}

	if len(c.PageToken) > 0 {
		urlValues.Add("page_token", &c.PageToken)
	}

	if len(c.Endpoint) > 0 {
		urlValues.Add("endpoint", &c.Endpoint)
	}

	if c.DefaultOnly {
		urlValues.AddBool("default_only", &c.DefaultOnly)
	}

	params := urlValues.Encode()

	return "/models?" + params, nil
}

func (c *ListModels) Encode() (io.Reader, error) {
	return nil, nil
}

func (c *ListModels) ContentType() string {
	return ""
}
