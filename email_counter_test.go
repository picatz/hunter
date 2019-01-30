package hunter

import (
	"context"
	"testing"
	"time"
)

func TestClient_CountEmails(t *testing.T) {
	results, err := client.CountEmails(Params{
		"domain": "stripe.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	if results.Data.Total <= 0 {
		t.Error("should have more than 0 email results for this query, got:", results)
	}
}

func TestClient_CountEmailsWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	results, err := client.CountEmailsWithContext(ctx, Params{
		"domain": "stripe.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	if results.Data.Total <= 0 {
		t.Error("should have more than 0 email results for this query, got:", results)
	}
}
