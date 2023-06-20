package conductoroneapi

import (
	"context"
	"fmt"
	"testing"
)

func TestLoginFlow(t *testing.T) {
	t.Run("login flow", func(t *testing.T) {
		creds, err := LoginFlow(context.Background(), "insulator", "2RCzHlak5q7CY14SdBc8HoZEJRf", "Created With Cone", func(validateUrl string) error {
			fmt.Printf("Open: %s\n", validateUrl)
			return nil
		})

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
