package hunter

import (
	"context"
	"testing"
	"time"
)

func TestClient_Account(t *testing.T) {
	_, err := client.Account()
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_AccountWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := client.AccountWithContext(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
