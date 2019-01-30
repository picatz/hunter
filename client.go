package hunter

import (
	"context"
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

func (c *Client) request(ctx context.Context, path string, params Params) ([]byte, error) {
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
	return ioutil.ReadAll(resp.Body)
}
