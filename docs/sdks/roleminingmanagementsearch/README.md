# RoleMiningManagementSearch

## Overview

### Available Operations

* [Search](#search) - Search

## Search

Search role mining suggestions by name, description, or cohort filter values with optional state and type filters.
 Each suggestion row is large (cohort filters, entitlements, insights, profile matches) — request a small page_size (≤10) to keep responses small.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.role_mining_management.v1.RoleMiningManagementSearchService.Search" method="post" path="/api/v1/search/role-mining/suggestions" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RoleMiningManagementSearch.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RoleMiningSearchSuggestionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [shared.RoleMiningSearchSuggestionsRequest](../../pkg/models/shared/roleminingsearchsuggestionsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.C1APIRoleMiningManagementV1RoleMiningManagementSearchServiceSearchResponse](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementsearchservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |