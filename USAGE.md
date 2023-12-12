<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"log"
)

func main() {
	s := conductoronesdkgo.New(
		conductoronesdkgo.WithSecurity(shared.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "Bearer <YOUR_ACCESS_TOKEN_HERE>",
		}),
	)

	ctx := context.Background()
	res, err := s.Apps.Create(ctx, &shared.CreateAppRequest{
		Owners: []string{
			"string",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateAppResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->