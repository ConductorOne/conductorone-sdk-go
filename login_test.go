package conductoronesdkgo

import (
	"context"
	"fmt"
	"testing"
)

func TestLoginFlow(t *testing.T) {
	t.Run("login flow", func(t *testing.T) {
		t.Skip("skip login flow test")

		creds, err := LoginFlow(context.Background(), "insulator", ClientIdGolangSDK, "Created With Cone", func(codeResp *DeviceCodeResponse) error {
			fmt.Printf("Open: %s\n", codeResp.VerificationURI)
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
