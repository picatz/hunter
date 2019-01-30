package hunter

import (
	"bytes"
	"context"
	"encoding/json"
)

// DomainSearchResult is returned by the DomainSearch function.
type DomainSearchResult struct {
	Data struct {
		Domain       string `json:"domain"`
		Disposable   bool   `json:"disposable"`
		Webmail      bool   `json:"webmail"`
		Pattern      string `json:"pattern"`
		Organization string `json:"organization"`
		Emails       []struct {
			Value      string `json:"value"`
			Type       string `json:"type"`
			Confidence int    `json:"confidence"`
			Sources    []struct {
				Domain      string `json:"domain"`
				URI         string `json:"uri"`
				ExtractedOn string `json:"extracted_on"`
				LastSeenOn  string `json:"last_seen_on"`
				StillOnPage bool   `json:"still_on_page"`
			} `json:"sources"`
			FirstName   string      `json:"first_name"`
			LastName    string      `json:"last_name"`
			Position    string      `json:"position"`
			Seniority   string      `json:"seniority"`
			Department  string      `json:"department"`
			Linkedin    interface{} `json:"linkedin"`
			Twitter     string      `json:"twitter"`
			PhoneNumber interface{} `json:"phone_number"`
		} `json:"emails"`
	} `json:"data"`
	Meta struct {
		Results int `json:"results"`
		Limit   int `json:"limit"`
		Offset  int `json:"offset"`
		Params  struct {
			Domain     string      `json:"domain"`
			Company    interface{} `json:"company"`
			Type       interface{} `json:"type"`
			Offset     int         `json:"offset"`
			Seniority  interface{} `json:"seniority"`
			Department interface{} `json:"department"`
		} `json:"params"`
	} `json:"meta"`
}

// DomainSearch searches a given domain. You give one domain name
// and it returns all the email addresses using this domain name
// found by https://hunter.io/ on the internet.
func (c *Client) DomainSearch(params Params) (*DomainSearchResult, error) {
	body, err := c.request(context.Background(), "https://api.hunter.io/v2/domain-search", params)
	result := new(DomainSearchResult)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DomainSearchWithContext searches a given domain. You give one domain name
// and it returns all the email addresses using this domain name
// found by https://hunter.io/ on the internet.
//
// This context-based version of the function would be more suitable for long-running
// applications like servers.
func (c *Client) DomainSearchWithContext(ctx context.Context, params Params) (*DomainSearchResult, error) {
	body, err := c.request(ctx, "https://api.hunter.io/v2/domain-search", params)
	result := new(DomainSearchResult)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
