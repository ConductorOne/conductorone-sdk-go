package conductoronesdkgo

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
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

func TestXFlow(t *testing.T) {
	t.Run("x flow", func(t *testing.T) {
		ctx := context.Background()

		/* Optional Override
		* Server URL will be extracted from client, optionally, you can
		* provide a server URL or a tenant domain (will create URL https://{tenant_domain}.conductor.one)
		 */
		opts := []CustomSDKOption{}
		opt, err := WithTenantCustom("https://c1dev.anthony.dev.ductone.com:2443")
		opts = append(opts, opt)

		s, err := NewWithCredentials(ctx, &ClientCredentials{
			ClientID:     "",
			ClientSecret: "",
		}, opts...)
		x := 10
		res, err := s.User.List(ctx, operations.C1APIUserV1UserServiceListRequest{PageSize: &x})
		fmt.Printf("User: %s\n", *res.UserServiceListResponse.NextPageToken)

		if err != nil {
			log.Fatal(err)
		}
		for {
			res, err = res.Next()
			if err != nil {
				t.Error(err)
			}
			if res == nil {
				t.Error("Nil res")
				break
			}
			fmt.Printf("User: %+v\n", res.UserServiceListResponse.List)
		}
	})
}
