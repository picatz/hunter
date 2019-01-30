package hunter

import (
	"context"
	"testing"
	"time"
)

func TestClient_VerifyEmail(t *testing.T) {
	results, err := client.VerifyEmail(Params{
		"email": "steli@close.io",
	})
	if err != nil {
		t.Fatal(err)
	}
	if results.Data.Email != "steli@close.io" {
		t.Error("unable to verify email", results)
	}
}

func TestClient_VerifyEmailWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	results, err := client.VerifyEmailWithContext(ctx, Params{
		"email": "steli@close.io",
	})
	if err != nil {
		t.Fatal(err)
	}
	if results.Data.Email != "steli@close.io" {
		t.Error("unable to verify email", results)
	}
}
