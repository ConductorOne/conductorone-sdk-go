# FindingSearch

## Overview

### Available Operations

* [Search](#search) - Search

## Search

Search findings using full-text query and filters for severity, state, type, and app.
 Each Finding row is large (risk factors, evidence, target, tags) — request a small page_size (≤10) to keep responses small.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingSearchService.Search" method="post" path="/api/v1/findings/search" -->
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

    res, err := s.FindingSearch.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FindingSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.FindingSearchRequest](../../pkg/models/shared/findingsearchrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.C1APIFindingV1FindingSearchServiceSearchResponse](../../pkg/models/operations/c1apifindingv1findingsearchservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |