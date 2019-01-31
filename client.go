package hunter

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

// Client implements an object to interact with
// the https://hunter.io API v2
type Client struct {
	Key    string
	client *http.Client
}

var (
	// UseDefaultEnvVariable is a default variable to tell the New method to lookup
	// the HUNTER_API_KEY environment variable.
	UseDefaultEnvVariable = ""
	// UseDefaultHTTPClient is a default variable to tell the New method to use the
	// default net/http DefaultClient.
	UseDefaultHTTPClient = http.DefaultClient
)

// New returns a Client object.
func New(key string, client *http.Client) *Client {
	if client == nil {
		client = UseDefaultHTTPClient
	}
	if key == UseDefaultEnvVariable {
		key = os.Getenv("HUNTER_API_KEY")
	}
	return &Client{Key: key, client: client}
}

var (
	ErrNoContent                  = errors.New("the request was successful and no additional content was sent")
	ErrBadRequest                 = errors.New("your request was not valid")
	ErrUnauthorized               = errors.New("no valid API key was provided")
	ErrForbidden                  = errors.New("you have reached the global rate limit (150 requests per second)")
	ErrNotFound                   = errors.New("the requested resource does not exist")
	ErrUnprocessableEntity        = errors.New("your request is valid but the creation of the resource failed")
	ErrTooManyRequests            = errors.New("you have reached your usage limit. Upgrade your plan if necessary")
	ErrUnavailableForLegalReasons = errors.New("the person behind the requested resource asked directly or indirectly to stop the processing of this resource")
	ErrServerError                = errors.New("something went wrong on hunter's end")
)

func (c *Client) request(ctx context.Context, method, path string, params Params) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	q := req.URL.Query()
	q.Add("api_key", c.Key)
	for k, v := range params {
		// skip if value is empty
		if v == "" {
			continue
		}
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200, 201:
		return ioutil.ReadAll(resp.Body)
	case 204:
		return nil, ErrNoContent
	case 400:
		return nil, ErrBadRequest
	case 401:
		return nil, ErrUnauthorized
	case 403:
		return nil, ErrForbidden
	case 404:
		return nil, ErrNotFound
	case 422:
		return nil, ErrUnprocessableEntity
	case 429:
		return nil, ErrTooManyRequests
	case 451:
		return nil, ErrUnavailableForLegalReasons
	default: // 5XX
		return nil, ErrServerError
	}
}
