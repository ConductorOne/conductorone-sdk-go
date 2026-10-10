# ShadowMcpOccurrence

## Overview

### Available Operations

* [Search](#search) - Search

## Search

Search returns the ungoverned, enabled observation rows behind one
 shadow-mcp finding -- the same present-set that produced its evidence
 counts. Authorized as VIEWER -- the same role required to read the
 finding itself.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.ShadowMcpOccurrenceService.Search" method="post" path="/api/v1/search/shadow_mcp_occurrences" -->
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

    res, err := s.ShadowMcpOccurrence.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ShadowMcpOccurrenceServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [shared.ShadowMcpOccurrenceServiceSearchRequest](../../pkg/models/shared/shadowmcpoccurrenceservicesearchrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIFindingV1ShadowMcpOccurrenceServiceSearchResponse](../../pkg/models/operations/c1apifindingv1shadowmcpoccurrenceservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |