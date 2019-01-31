package hunter

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// EmailCounterResult is returned by the CountEmails function.
type EmailCounterResult struct {
	Data struct {
		Total          int `json:"total"`
		PersonalEmails int `json:"personal_emails"`
		GenericEmails  int `json:"generic_emails"`
		Department     struct {
			Executive     int `json:"executive"`
			It            int `json:"it"`
			Finance       int `json:"finance"`
			Management    int `json:"management"`
			Sales         int `json:"sales"`
			Legal         int `json:"legal"`
			Support       int `json:"support"`
			Hr            int `json:"hr"`
			Marketing     int `json:"marketing"`
			Communication int `json:"communication"`
		} `json:"department"`
		Seniority struct {
			Junior    int `json:"junior"`
			Senior    int `json:"senior"`
			Executive int `json:"executive"`
		} `json:"seniority"`
	} `json:"data"`
	Meta struct {
		Params struct {
			Domain string      `json:"domain"`
			Type   interface{} `json:"type"`
		} `json:"params"`
	} `json:"meta"`
}

// CountEmails  allows you to verify the deliverability of an email address.
func (c *Client) CountEmails(params Params) (*EmailCounterResult, error) {
	body, err := c.request(context.Background(), http.MethodGet, "https://api.hunter.io/v2/email-count", params)
	if err != nil {
		return nil, err
	}
	result := new(EmailCounterResult)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CountEmailsWithContext allows you to verify the deliverability of an email address.
func (c *Client) CountEmailsWithContext(ctx context.Context, params Params) (*EmailCounterResult, error) {
	body, err := c.request(ctx, http.MethodGet, "https://api.hunter.io/v2/email-count", params)
	if err != nil {
		return nil, err
	}
	result := new(EmailCounterResult)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
