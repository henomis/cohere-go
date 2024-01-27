package response

import (
	"io"
	"strconv"

	"github.com/henomis/cohere-go/model"
	"github.com/henomis/restclientgo"
)

const (
	ContentTypeJSON = "application/json"
)

type Response struct {
	Meta    model.Meta `json:"meta,omitempty"`
	Headers Headers    `json:"-"`
	Code    int        `json:"-"`
	RawBody *string    `json:"-"`
}

type Headers struct {
	EndpointMonthlyCallLimit   int `json:"-"`
	RatelimitLimit             int `json:"-"`
	TrialEndpointCallRemaining int `json:"-"`
	RatelimitRemaining         int `json:"-"`
	RatelimitReset             int `json:"-"`
	TrialEndpointCallLimit     int `json:"-"`
}

type Detail struct {
	TypeURL *string `json:"typeUrl,omitempty"`
	Value   *string `json:"value,omitempty"`
}

func (r *Response) IsSuccess() bool {
	return (r.Code >= 200 && r.Code < 300)
}

func (r *Response) SetStatusCode(code int) error {
	r.Code = code
	return nil
}

func (r *Response) SetBody(body io.Reader) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	s := string(b)
	r.RawBody = &s

	return nil
}

func (r *Response) SetHeaders(headers restclientgo.Headers) error {
	if header, ok := headers["X-Endpoint-Monthly-Call-Limit"]; ok {
		headerAsInt, err := strconv.Atoi(header[0])
		if err != nil {
			return err
		}

		r.Headers.EndpointMonthlyCallLimit = headerAsInt
	}

	if header, ok := headers["X-Ratelimit-Limit"]; ok {
		headerAsInt, err := strconv.Atoi(header[0])
		if err != nil {
			return err
		}

		r.Headers.RatelimitLimit = headerAsInt
	}

	if header, ok := headers["X-Trial-Endpoint-Call-Remaining"]; ok {
		headerAsInt, err := strconv.Atoi(header[0])
		if err != nil {
			return err
		}

		r.Headers.TrialEndpointCallRemaining = headerAsInt
	}

	if header, ok := headers["X-Ratelimit-Remaining"]; ok {
		headerAsInt, err := strconv.Atoi(header[0])
		if err != nil {
			return err
		}

		r.Headers.RatelimitRemaining = headerAsInt
	}

	if header, ok := headers["X-Ratelimit-Reset"]; ok {
		headerAsInt, err := strconv.Atoi(header[0])
		if err != nil {
			return err
		}

		r.Headers.RatelimitReset = headerAsInt
	}

	if header, ok := headers["X-Trial-Endpoint-Call-Limit"]; ok {
		headerAsInt, err := strconv.Atoi(header[0])
		if err != nil {
			return err
		}

		r.Headers.TrialEndpointCallLimit = headerAsInt
	}

	return nil
}
