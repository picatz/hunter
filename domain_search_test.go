package hunter

import (
	"context"
	"testing"
	"time"
)

func TestClient_DomainSearch(t *testing.T) {
	results, err := client.DomainSearch(Params{"domain": "stripe.com"})
	if err != nil {
		t.Fatal(err)
	}
	if results.Meta.Results <= 0 {
		t.Error("got not results")
	}
}

func TestClient_DomainSearchWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	results, err := client.DomainSearchWithContext(ctx, Params{"domain": "stripe.com"})
	if err != nil {
		t.Fatal(err)
	}
	if results.Meta.Results <= 0 {
		t.Error("got not results")
	}
}
