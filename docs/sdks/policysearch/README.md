# PolicySearch

### Available Operations

* [Search](#search) - Invokes the c1.api.policy.v1.PolicySearch.Search method.

## Search

Invokes the c1.api.policy.v1.PolicySearch.Search method.

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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PolicySearch.Search(ctx, shared.SearchPoliciesRequest{
        DisplayName: conductoroneapi.String("corrupti"),
        PageSize: conductoroneapi.Float64(8792.35),
        PageToken: conductoroneapi.String("tempora"),
        PolicyTypes: []shared.SearchPoliciesRequestPolicyTypes{
            shared.SearchPoliciesRequestPolicyTypesPolicyTypeUnspecified,
            shared.SearchPoliciesRequestPolicyTypesPolicyTypeGrant,
            shared.SearchPoliciesRequestPolicyTypesPolicyTypeProvision,
        },
        Query: conductoroneapi.String("voluptatem"),
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

