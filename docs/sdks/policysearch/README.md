# PolicySearch

### Available Operations

* [Search](#search) - Search

## Search

Search policies based on filters specified in the request body.

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
    res, err := s.PolicySearch.Search(ctx, shared.SearchPoliciesRequest{
        DisplayName: conductoroneapi.String("sapiente"),
        PageSize: conductoroneapi.Float64(8953.86),
        PageToken: conductoroneapi.String("illo"),
        PolicyTypes: []shared.SearchPoliciesRequestPolicyTypes{
            shared.SearchPoliciesRequestPolicyTypesPolicyTypeUnspecified,
            shared.SearchPoliciesRequestPolicyTypesPolicyTypeCertify,
            shared.SearchPoliciesRequestPolicyTypesPolicyTypeProvision,
            shared.SearchPoliciesRequestPolicyTypesPolicyTypeGrant,
        },
        Query: conductoroneapi.String("sed"),
        Refs: []shared.PolicyRef{
            shared.PolicyRef{
                ID: conductoroneapi.String("4e3698f4-47f6-403e-8b44-5e80ca55efd2"),
            },
            shared.PolicyRef{
                ID: conductoroneapi.String("0e457e18-58b6-4a89-bbe3-a5aa8e4824d0"),
            },
            shared.PolicyRef{
                ID: conductoroneapi.String("ab407508-8e51-4862-865e-904f3b1194b8"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.SearchPoliciesRequest](../../models/shared/searchpoliciesrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.C1APIPolicyV1PolicySearchSearchResponse](../../models/operations/c1apipolicyv1policysearchsearchresponse.md), error**

