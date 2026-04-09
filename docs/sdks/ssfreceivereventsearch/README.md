# SSFReceiverEventSearch

## Overview

### Available Operations

* [Search](#search) - Search

## Search

Invokes the c1.api.ssf_receiver.v1.SSFReceiverEventSearchService.Search method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ssf_receiver.v1.SSFReceiverEventSearchService.Search" method="post" path="/api/v1/search/ssf-receiver-events" -->
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

    res, err := s.SSFReceiverEventSearch.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SSFReceiverEventSearchServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [shared.SSFReceiverEventSearchServiceSearchRequest](../../pkg/models/shared/ssfreceivereventsearchservicesearchrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.C1APISSFReceiverV1SSFReceiverEventSearchServiceSearchResponse](../../pkg/models/operations/c1apissfreceiverv1ssfreceivereventsearchservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |