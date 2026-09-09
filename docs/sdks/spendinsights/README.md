# SpendInsights

## Overview

### Available Operations

* [GetAttributionRollups](#getattributionrollups) - Get Attribution Rollups
* [GetDenial](#getdenial) - Get Denial
* [GetMySpendStatus](#getmyspendstatus) - Get My Spend Status
* [GetOverview](#getoverview) - Get Overview
* [ResolveEffectiveLimits](#resolveeffectivelimits) - Resolve Effective Limits
* [SearchDenials](#searchdenials) - Search Denials

## GetAttributionRollups

Get ranked settled-spend attribution groups for one dimension and time window.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.spendinsights.v1.SpendInsightsService.GetAttributionRollups" method="post" path="/api/v1/spend-insights/attribution/search" -->
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

    res, err := s.SpendInsights.GetAttributionRollups(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAttributionRollupsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [shared.GetAttributionRollupsRequest](../../pkg/models/shared/getattributionrollupsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.C1APISpendinsightsV1SpendInsightsServiceGetAttributionRollupsResponse](../../pkg/models/operations/c1apispendinsightsv1spendinsightsservicegetattributionrollupsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDenial

Get one budget-denial episode by its unique ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.spendinsights.v1.SpendInsightsService.GetDenial" method="get" path="/api/v1/spend-insights/denials/{block_id}" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
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

    res, err := s.SpendInsights.GetDenial(ctx, operations.C1APISpendinsightsV1SpendInsightsServiceGetDenialRequest{
        BlockID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDenialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.C1APISpendinsightsV1SpendInsightsServiceGetDenialRequest](../../pkg/models/operations/c1apispendinsightsv1spendinsightsservicegetdenialrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |
| `opts`                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                   | :heavy_minus_sign:                                                                                                                                             | The options for this request.                                                                                                                                  |

### Response

**[*operations.C1APISpendinsightsV1SpendInsightsServiceGetDenialResponse](../../pkg/models/operations/c1apispendinsightsv1spendinsightsservicegetdenialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMySpendStatus

Get the caller's budget limits and current usage. An optional denial ID
 includes that denial only when it is visible to the caller.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.spendinsights.v1.SpendInsightsService.GetMySpendStatus" method="get" path="/api/v1/spend-insights/my/status" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
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

    res, err := s.SpendInsights.GetMySpendStatus(ctx, operations.C1APISpendinsightsV1SpendInsightsServiceGetMySpendStatusRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.GetMySpendStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.C1APISpendinsightsV1SpendInsightsServiceGetMySpendStatusRequest](../../pkg/models/operations/c1apispendinsightsv1spendinsightsservicegetmyspendstatusrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |
| `opts`                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                           | The options for this request.                                                                                                                                                |

### Response

**[*operations.C1APISpendinsightsV1SpendInsightsServiceGetMySpendStatusResponse](../../pkg/models/operations/c1apispendinsightsv1spendinsightsservicegetmyspendstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetOverview

Get an organization-wide spend overview for a requested time window.
 Returns current budget headroom, trends, forecasts, usage coverage, and denials.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.spendinsights.v1.SpendInsightsService.GetOverview" method="post" path="/api/v1/spend-insights/overview" -->
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

    res, err := s.SpendInsights.GetOverview(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetOverviewResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.GetOverviewRequest](../../pkg/models/shared/getoverviewrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../pkg/models/operations/option.md)               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.C1APISpendinsightsV1SpendInsightsServiceGetOverviewResponse](../../pkg/models/operations/c1apispendinsightsv1spendinsightsservicegetoverviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ResolveEffectiveLimits

Resolve the budget limits and current usage that apply to a user.
 Optionally includes application-specific limits.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.spendinsights.v1.SpendInsightsService.ResolveEffectiveLimits" method="get" path="/api/v1/spend-insights/effective-limits/{user_id}" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
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

    res, err := s.SpendInsights.ResolveEffectiveLimits(ctx, operations.C1APISpendinsightsV1SpendInsightsServiceResolveEffectiveLimitsRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResolveEffectiveLimitsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                | Type                                                                                                                                                                                     | Required                                                                                                                                                                                 | Description                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                       | The context to use for the request.                                                                                                                                                      |
| `request`                                                                                                                                                                                | [operations.C1APISpendinsightsV1SpendInsightsServiceResolveEffectiveLimitsRequest](../../pkg/models/operations/c1apispendinsightsv1spendinsightsserviceresolveeffectivelimitsrequest.md) | :heavy_check_mark:                                                                                                                                                                       | The request object to use for the request.                                                                                                                                               |
| `opts`                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                       | The options for this request.                                                                                                                                                            |

### Response

**[*operations.C1APISpendinsightsV1SpendInsightsServiceResolveEffectiveLimitsResponse](../../pkg/models/operations/c1apispendinsightsv1spendinsightsserviceresolveeffectivelimitsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchDenials

Search budget-denial episodes for the organization within a time window.
 Requires the denial-read permission.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.spendinsights.v1.SpendInsightsService.SearchDenials" method="post" path="/api/v1/spend-insights/denials/search" -->
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

    res, err := s.SpendInsights.SearchDenials(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchDenialsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.SearchDenialsRequest](../../pkg/models/shared/searchdenialsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.C1APISpendinsightsV1SpendInsightsServiceSearchDenialsResponse](../../pkg/models/operations/c1apispendinsightsv1spendinsightsservicesearchdenialsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |