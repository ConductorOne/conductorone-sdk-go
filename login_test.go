package conductoroneapi

import (
	"context"
	"testing"
)

func TestLoginFlow(t *testing.T) {
	t.Run("login flow", func(t *testing.T) {
		creds, err := LoginFlow(context.Background(), "c1dev.logan.dev.ductone.com:2443", "2RCzHlak5q7CY14SdBc8HoZEJRf", "Created With Cone")

		if err != nil {
			t.Errorf("LoginFlow() error = %v", err)
			return
		}
		if creds == nil || creds.ClientID == "" || creds.ClientSecret == "" {
			t.Errorf("LoginFlow() creds = %v", creds)
			return
		}
	})
}
