## SDK Example Usage

### Example

```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := NewWithCredentials(ctx, &ClientCredentials{
		ClientID:     "",
		ClientSecret: "",
	} )

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
<!-- No SDK Example Usage [usage] -->
## SDK Expander

### Example

```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
	// Create SDK ...
	req := shared.RequestCatalogSearchServiceSearchEntitlementsRequest{
		AppEntitlementExpandMask: shared.AppEntitlementExpandMask{
			Paths: []string{"*"},
		},
	}
	resp, err := s.sdk.RequestCatalogSearch.SearchEntitlements(ctx, &req)
	if err != nil {
		return nil, err
	}

	response := resp.RequestCatalogSearchServiceSearchEntitlementsResponse
	if response == nil {
		return nil, nil
	}

	getStructWithPaths := func(response shared.AppEntitlementView) *shared.AppEntitlementView {
		return &response
	}

	expandedResponse, err := GetExpandResponse(response.List, response.Expanded, getStructWithPaths, newMockEntitlement)
	if err != nil {
		return nil, err 
	}
	return expandedResponse, nil

}

type mockEntitlement struct {
	*shared.AppEntitlement
	Expanded map[string]*any
}

func newMockEntitlement(x shared.AppEntitlementView, expanded map[string]*any) *mockEntitlement {
	entitlement := x.GetAppEntitlement()

	return &mockEntitlement{
		AppEntitlement: entitlement,
		Expanded:       expanded,
	}
}

```
## SDK Example Usage with Custom Server/Tenant

### Example

```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	/* Optional Override 
	* Server URL will be extracted from client, optionally, you can
	* provide a server URL or a tenant domain (will create URL https://{tenant_domain}.conductor.one) 
	*/
	opts := []sdk.CustomSDKOption{}
	opt, _ := sdk.WithTenantCustom("Server URL or Tenant Domain")
	opts = append(opts, opt)

	s := NewWithCredentials(ctx, &ClientCredentials{
		ClientID:     "",
		ClientSecret: "",
	} opts...)

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