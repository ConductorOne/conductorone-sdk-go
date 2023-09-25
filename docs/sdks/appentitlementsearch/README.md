# AppEntitlementSearch

### Available Operations

* [Search](#search) - Search

## Search

Search app entitlements based on filters specified in the request body.

### Example Usage

```go
package main

import(
	"context"
	"log"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlementSearch.Search(ctx, shared.AppEntitlementSearchServiceSearchRequest{
        AppEntitlementExpandMask: &shared.AppEntitlementExpandMask{
            Paths: []string{
                "molestiae",
            },
        },
        AccessReviewID: conductoronesdkgo.String("minus"),
        Alias: conductoronesdkgo.String("placeat"),
        AppIds: []string{
            "voluptatum",
        },
        AppUserIds: []string{
            "iusto",
        },
        ComplianceFrameworkIds: []string{
            "excepturi",
        },
        ExcludeAppIds: []string{
            "nisi",
        },
        ExcludeAppUserIds: []string{
            "recusandae",
        },
        IncludeDeleted: conductoronesdkgo.Bool(false),
        OnlyGetExpiring: conductoronesdkgo.Bool(false),
        PageSize: conductoronesdkgo.Float64(8360.79),
        PageToken: conductoronesdkgo.String("ab"),
        Query: conductoronesdkgo.String("quis"),
        ResourceTypeIds: []string{
            "veritatis",
        },
        RiskLevelIds: []string{
            "deserunt",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppEntitlementSearchServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [shared.AppEntitlementSearchServiceSearchRequest](../../models/shared/appentitlementsearchservicesearchrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchResponse](../../models/operations/c1apiappv1appentitlementsearchservicesearchresponse.md), error**

