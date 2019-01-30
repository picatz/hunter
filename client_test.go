package hunter

import (
	"testing"
)

var client = New(UseDefaultEnvVariable, UseDefaultHTTPClient)

func TestClient(t *testing.T) {
	if client.Key == "" {
		t.Fatal("no api key found for the client using the HUNTER_API_KEY environment variable")
	}
}
