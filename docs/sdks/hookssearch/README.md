# HooksSearch

## Overview

### Available Operations

* [Search](#search) - Search

## Search

Search returns hooks for the tenant, paginated. Setting query or refs
 returns UNIMPLEMENTED; filtering is not yet supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.hooks.v1.HooksSearch.Search" method="post" path="/api/v1/search/hooks" -->
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

    res, err := s.HooksSearch.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.HooksSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.HooksSearchRequest](../../pkg/models/shared/hookssearchrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../pkg/models/operations/option.md)               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.C1APIHooksV1HooksSearchSearchResponse](../../pkg/models/operations/c1apihooksv1hookssearchsearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |