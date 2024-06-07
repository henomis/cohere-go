package response

import (
	"encoding/json"
	"io"
)

type CheckAPIKey struct {
	Response
	Valid          bool   `json:"valid"`
	OrganizationID string `json:"organization_id"`
	OwnerID        string `json:"owner_id"`
}

func (r *CheckAPIKey) AcceptContentType() string {
	return ContentTypeJSON
}

func (r *CheckAPIKey) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}
