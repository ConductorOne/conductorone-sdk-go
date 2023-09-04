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
    res, err := s.AppEntitlementSearch.Search(ctx, shared.AppEntitlementSearchServiceSearchRequest{
        AppEntitlementExpandMask: &shared.AppEntitlementExpandMask{
            Paths: []string{
                "molestiae",
                "minus",
            },
        },
        AccessReviewID: conductoroneapi.String("placeat"),
        Alias: conductoroneapi.String("voluptatum"),
        AppIds: []string{
            "excepturi",
            "nisi",
        },
        AppUserIds: []string{
            "temporibus",
            "ab",
            "quis",
            "veritatis",
        },
        ComplianceFrameworkIds: []string{
            "perferendis",
            "ipsam",
            "repellendus",
        },
        ExcludeAppIds: []string{
            "quo",
            "odit",
            "at",
            "at",
        },
        ExcludeAppUserIds: []string{
            "molestiae",
            "quod",
            "quod",
            "esse",
        },
        OnlyGetExpiring: conductoroneapi.Bool(false),
        PageSize: conductoroneapi.Float64(5204.78),
        PageToken: conductoroneapi.String("porro"),
        Query: conductoroneapi.String("dolorum"),
        ResourceTypeIds: []string{
            "nam",
        },
        RiskLevelIds: []string{
            "occaecati",
            "fugit",
            "deleniti",
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

