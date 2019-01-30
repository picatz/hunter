package hunter

import (
	"context"
	"testing"
	"time"
)

func TestClient_FindEmail(t *testing.T) {
	results, err := client.FindEmail(Params{
		"domain":     "asana.com",
		"first_name": "Dustin",
		"last_name":  "Moskovitz",
	})
	if err != nil {
		t.Fatal(err)
	}
	if results.Data.Email != "dustin@asana.com" {
		t.Error("unable to find known email", results)
	}
}

func TestClient_FindEmailWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	results, err := client.FindEmailWithContext(ctx, Params{
		"domain":     "asana.com",
		"first_name": "Dustin",
		"last_name":  "Moskovitz",
	})
	if err != nil {
		t.Fatal(err)
	}
	if results.Data.Email != "dustin@asana.com" {
		t.Error("unable to find known email", results)
	}
}
