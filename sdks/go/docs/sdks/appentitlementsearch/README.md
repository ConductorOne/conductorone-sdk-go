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
        AccessReviewID: conductoroneapi.String("suscipit"),
        Alias: conductoroneapi.String("molestiae"),
        AppIds: []string{
            "placeat",
            "voluptatum",
            "iusto",
            "excepturi",
        },
        AppUserIds: []string{
            "recusandae",
            "temporibus",
        },
        ComplianceFrameworkIds: []string{
            "quis",
        },
        ExcludeAppIds: []string{
            "deserunt",
        },
        ExcludeAppUserIds: []string{
            "ipsam",
        },
        ExpandMask: &shared.AppEntitlementExpandMask{
            Paths: []string{
                "sapiente",
                "quo",
                "odit",
                "at",
            },
        },
        OnlyGetExpiring: conductoroneapi.Bool(false),
        PageSize: conductoroneapi.Float64(8700.88),
        PageToken: conductoroneapi.String("maiores"),
        Query: conductoroneapi.String("molestiae"),
        ResourceTypeIds: []string{
            "quod",
            "esse",
            "totam",
            "porro",
        },
        RiskLevelIds: []string{
            "dicta",
            "nam",
            "officia",
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

