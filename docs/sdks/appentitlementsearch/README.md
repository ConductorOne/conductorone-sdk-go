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
                "debitis",
                "ipsa",
            },
        },
        AccessReviewID: conductoroneapi.String("delectus"),
        Alias: conductoroneapi.String("tempora"),
        AppIds: []string{
            "molestiae",
            "minus",
        },
        AppUserIds: []string{
            "voluptatum",
            "iusto",
            "excepturi",
            "nisi",
        },
        ComplianceFrameworkIds: []string{
            "temporibus",
            "ab",
            "quis",
            "veritatis",
        },
        ExcludeAppIds: []string{
            "perferendis",
            "ipsam",
            "repellendus",
        },
        ExcludeAppUserIds: []string{
            "quo",
            "odit",
            "at",
            "at",
        },
        OnlyGetExpiring: conductoroneapi.Bool(false),
        PageSize: conductoroneapi.Float64(9786.19),
        PageToken: conductoroneapi.String("molestiae"),
        Query: conductoroneapi.String("quod"),
        ResourceTypeIds: []string{
            "esse",
            "totam",
            "porro",
            "dolorum",
        },
        RiskLevelIds: []string{
            "nam",
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

