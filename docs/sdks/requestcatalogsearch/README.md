# RequestCatalogSearch

### Available Operations

* [RequestCatalogSearchSvcSearchEntitlements](#requestcatalogsearchsvcsearchentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

## RequestCatalogSearchSvcSearchEntitlements

Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

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
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.RequestCatalogSearch.RequestCatalogSearchSvcSearchEntitlements(ctx, shared.RequestCatalogSearchServiceSearchEntitlementsRequest{
        EntitlementAlias: conductoroneapi.String("quasi"),
        ExpandMask: &shared.AppEntitlementExpandMask{
            Paths: []string{
                "temporibus",
                "laborum",
                "quasi",
            },
        },
        PageSize: conductoroneapi.Float64(9719.45),
        PageToken: conductoroneapi.String("voluptatibus"),
        Query: conductoroneapi.String("vero"),
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

