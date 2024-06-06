## SDK Example Usage

### Example

```go
package main

import (
	"context"
	sdk "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"log"
)

func main() {

	s, err := sdk.NewWithCredentials(ctx, &sdk.ClientCredentials{
		ClientID:     "",
		ClientSecret: "",
	})

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

## SDK Example Usage with Custom Server/Tenant

### Example

```go
package main

import (
	"context"
	sdk "github.com/conductorone/conductorone-sdk-go"
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
	opt, _ := sdk.WithTenantCustom("")
	opts = append(opts, opt)

	s, err := sdk.NewWithCredentials(ctx, &sdk.ClientCredentials{
		ClientID:     "",
		ClientSecret: "",
	}, opts...)

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
## SDK Example Usage with Pagination

### Example

```go
func main() {
	ctx := context.Background()

	s, err := sdk.NewWithCredentials(...)
	if err != nil {
		panic(err)
	}

	pageSize := 1.0
	abort := make(chan int)
	ch := Paginate[
		operations.C1APIAppV1AppsListRequest,
		*operations.C1APIAppV1AppsListResponse,
		shared.App,
	](ctx,
		operations.C1APIAppV1AppsListRequest{
			PageSize: &pageSize},
		func(ctx context.Context, req operations.C1APIAppV1AppsListRequest) (*operations.C1APIAppV1AppsListResponse, error) {
			return s.Apps.List(ctx, req)
		},
		func(resp *operations.C1APIAppV1AppsListResponse) (*string, []shared.App, error) {
			return resp.GetListAppsResponse().GetNextPageToken(), resp.GetListAppsResponse().GetList(), nil
		},
		abort,
	)

	rv := make([]shared.App, 0, 20)
	for v := range ch {
		if v.Error != nil {
			panic(v.Error)
		}
		rv = append(rv, v.Message...)
		if len(rv) == 20 {
			close(abort)
			// Be sure to break so no other values are read from the channel since the channel is buffered.
			break
		}
	}
}
```