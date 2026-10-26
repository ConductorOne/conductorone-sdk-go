# APIKeySearch

## Overview

### Available Operations

* [Search](#search) - NOTE: Searches personal and service principal API keys

## Search

Search finds personal client credentials and service principal credentials, with optional filtering by query text or key type.
 Request a modest page_size (≤25) to keep responses small.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.api_key.v1.APIKeySearchService.Search" method="post" path="/api/v1/search/iam/api_keys" -->
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

    res, err := s.APIKeySearch.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.APIKeySearchServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [shared.APIKeySearchServiceSearchRequest](../../pkg/models/shared/apikeysearchservicesearchrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.C1APIAPIKeyV1APIKeySearchServiceSearchResponse](../../pkg/models/operations/c1apiapikeyv1apikeysearchservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |