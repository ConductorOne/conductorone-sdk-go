# AppEntitlementSearch
(*AppEntitlementSearch*)

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
                "transition",
            },
        },
        AccessReviewID: conductoronesdkgo.String("withdrawal Coordinator Ngultrum"),
        Alias: conductoronesdkgo.String("Russian Northeast yearningly"),
        AppIds: []string{
            "Southeast",
        },
        AppUserIds: []string{
            "Books",
        },
        ComplianceFrameworkIds: []string{
            "Other",
        },
        ExcludeAppIds: []string{
            "eligendi",
        },
        ExcludeAppUserIds: []string{
            "Tobago",
        },
        IncludeDeleted: conductoronesdkgo.Bool(false),
        OnlyGetExpiring: conductoronesdkgo.Bool(false),
        PageSize: conductoronesdkgo.Float64(7320.8),
        PageToken: conductoronesdkgo.String("Arsenic Gasoline Oklahoma"),
        Query: conductoronesdkgo.String("opposite"),
        ResourceTypeIds: []string{
            "Synchronised",
        },
        RiskLevelIds: []string{
            "huzzah",
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

