# RequestCatalogSearch

### Available Operations

* [SearchEntitlements](#searchentitlements) - Search Entitlements

## SearchEntitlements

Search request catalogs based on filters specified in the request body.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogSearch.SearchEntitlements(ctx, shared.RequestCatalogSearchServiceSearchEntitlementsRequest{
        AppEntitlementExpandMask: &shared.AppEntitlementExpandMask{
            Paths: []string{
                "id",
            },
        },
        AppDisplayName: conductoroneapi.String("quis"),
        EntitlementAlias: conductoroneapi.String("reprehenderit"),
        GrantedStatus: shared.RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusGranted.ToPointer(),
        IncludeDeleted: conductoroneapi.Bool(false),
        PageSize: conductoroneapi.Float64(764.86),
        PageToken: conductoroneapi.String("corporis"),
        Query: conductoroneapi.String("quidem"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogSearchServiceSearchEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [shared.RequestCatalogSearchServiceSearchEntitlementsRequest](../../models/shared/requestcatalogsearchservicesearchentitlementsrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogsearchservicesearchentitlementsresponse.md), error**

