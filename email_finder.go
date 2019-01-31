package hunter

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// EmailFinderResult is returned by the FindEmail function.
type EmailFinderResult struct {
	Data struct {
		FirstName   string      `json:"first_name"`
		LastName    string      `json:"last_name"`
		Email       string      `json:"email"`
		Score       int         `json:"score"`
		Domain      string      `json:"domain"`
		Position    string      `json:"position"`
		Twitter     string      `json:"twitter"`
		LinkedinURL string      `json:"linkedin_url"`
		PhoneNumber interface{} `json:"phone_number"`
		Company     string      `json:"company"`
		Sources     []struct {
			Domain      string `json:"domain"`
			URI         string `json:"uri"`
			ExtractedOn string `json:"extracted_on"`
			LastSeenOn  string `json:"last_seen_on"`
			StillOnPage bool   `json:"still_on_page"`
		} `json:"sources"`
	} `json:"data"`
	Meta struct {
		Params struct {
			FirstName string      `json:"first_name"`
			LastName  string      `json:"last_name"`
			FullName  interface{} `json:"full_name"`
			Domain    string      `json:"domain"`
			Company   interface{} `json:"company"`
		} `json:"params"`
	} `json:"meta"`
}

// FindEmail generates or retrieves the most likely
// email address from a domain name, a first name and a last name.
func (c *Client) FindEmail(params Params) (*EmailFinderResult, error) {
	body, err := c.request(context.Background(), http.MethodGet, "https://api.hunter.io/v2/email-finder", params)
	if err != nil {
		return nil, err
	}
	result := new(EmailFinderResult)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindEmailWithContext generates or retrieves the most likely
// email address from a domain name, a first name and a last name.
//
// This context-based version of the function would be more suitable for long-running
// applications like servers.
func (c *Client) FindEmailWithContext(ctx context.Context, params Params) (*EmailFinderResult, error) {
	body, err := c.request(ctx, http.MethodGet, "https://api.hunter.io/v2/email-finder", params)
	if err != nil {
		return nil, err
	}
	result := new(EmailFinderResult)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
