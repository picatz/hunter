package hunter

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// EmailVerifierResult is returned by the VerifyEmail function.
type EmailVerifierResult struct {
	Data struct {
		Result     string `json:"result"`
		Score      int    `json:"score"`
		Email      string `json:"email"`
		Regexp     bool   `json:"regexp"`
		Gibberish  bool   `json:"gibberish"`
		Disposable bool   `json:"disposable"`
		Webmail    bool   `json:"webmail"`
		MxRecords  bool   `json:"mx_records"`
		SMTPServer bool   `json:"smtp_server"`
		SMTPCheck  bool   `json:"smtp_check"`
		AcceptAll  bool   `json:"accept_all"`
		Block      bool   `json:"block"`
		Sources    []struct {
			Domain      string `json:"domain"`
			URI         string `json:"uri"`
			ExtractedOn string `json:"extracted_on"`
			LastSeenOn  string `json:"last_seen_on"`
			StillOnPage bool   `json:"still_on_page"`
		} `json:"sources"`
	} `json:"data"`
	Meta struct {
		Params struct {
			Email string `json:"email"`
		} `json:"params"`
	} `json:"meta"`
}

// VerifyEmail  allows you to verify the deliverability of an email address.
func (c *Client) VerifyEmail(params Params) (*EmailVerifierResult, error) {
	body, err := c.request(context.Background(), http.MethodGet, "https://api.hunter.io/v2/email-verifier", params)
	if err != nil {
		return nil, err
	}
	result := new(EmailVerifierResult)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// VerifyEmailWithContext allows you to verify the deliverability of an email address.
//
// This context-based version of the function would be more suitable for long-running
// applications like servers.
func (c *Client) VerifyEmailWithContext(ctx context.Context, params Params) (*EmailVerifierResult, error) {
	body, err := c.request(ctx, http.MethodGet, "https://api.hunter.io/v2/email-verifier", params)
	if err != nil {
		return nil, err
	}
	result := new(EmailVerifierResult)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
