package hunter

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// AccountInformation is returned by the Account function.
type AccountInformation struct {
	Data struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		PlanName  string `json:"plan_name"`
		PlanLevel int    `json:"plan_level"`
		ResetDate string `json:"reset_date"`
		TeamID    int    `json:"team_id"`
		Calls     struct {
			Used      int `json:"used"`
			Available int `json:"available"`
		} `json:"calls"`
	} `json:"data"`
}

// Account is a function to get information regarding your
// Hunter account at any time. This call is free.
//
// This context-based version of the function would be more suitable for long-running
// applications like servers.
func (c *Client) Account() (*AccountInformation, error) {
	body, err := c.request(context.Background(), http.MethodGet, "https://api.hunter.io/v2/account", nil)
	if err != nil {
		return nil, err
	}
	result := new(AccountInformation)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AccountWithContext is a function to get information regarding your
// Hunter account at any time. This call is free.
//
// This context-based version of the function would be more suitable for long-running
// applications like servers.
func (c *Client) AccountWithContext(ctx context.Context) (*AccountInformation, error) {
	body, err := c.request(ctx, http.MethodGet, "https://api.hunter.io/v2/account", nil)
	if err != nil {
		return nil, err
	}
	result := new(AccountInformation)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
