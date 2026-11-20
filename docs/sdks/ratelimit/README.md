# RateLimit

## Overview

### Available Operations

* [Get](#get) - Get
* [GetUsageReport](#getusagereport) - Get Usage Report

## Get

Get returns every API rate limit configured for the organization, the live usage of the organization-wide limits, and the number of requests throttled in the last 60 minutes.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ratelimit.v1.RateLimitService.Get" method="get" path="/api/v1/rate_limits" -->
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

    res, err := s.RateLimit.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.RateLimitServiceGetResponse != nil {
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

**[*operations.C1APIRatelimitV1RateLimitServiceGetResponse](../../pkg/models/operations/c1apiratelimitv1ratelimitservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUsageReport

GetUsageReport summarizes which users or API keys made API requests, by default in the last 60 minutes, or over a past range of up to 90 days when start_time is set. It returns at most `max_entries` entries, ranked over every matching consumer, together with totals over all of them. It is a bounded summary rather than a search, so it is not paginated: counts change between calls, and pages read at different moments would not join into a consistent list.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ratelimit.v1.RateLimitService.GetUsageReport" method="post" path="/api/v1/rate_limits/usage_report" -->
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

    res, err := s.RateLimit.GetUsageReport(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RateLimitServiceGetUsageReportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [shared.RateLimitServiceGetUsageReportRequest](../../pkg/models/shared/ratelimitservicegetusagereportrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.C1APIRatelimitV1RateLimitServiceGetUsageReportResponse](../../pkg/models/operations/c1apiratelimitv1ratelimitservicegetusagereportresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |