# MyRateLimit

## Overview

### Available Operations

* [Get](#get) - Get

## Get

Get returns the per-user and per-API-key limits that apply to the caller, the live usage of the caller's per-user limits, the usage of each of the caller's API keys, and, when the request was made with an API key, that key's usage.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ratelimit.v1.MyRateLimitService.Get" method="get" path="/api/v1/rate_limits/me" -->
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

    res, err := s.MyRateLimit.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.MyRateLimitServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.C1APIRatelimitV1MyRateLimitServiceGetResponse](../../pkg/models/operations/c1apiratelimitv1myratelimitservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |