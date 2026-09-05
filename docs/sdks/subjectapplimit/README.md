# SubjectAppLimit

## Overview

### Available Operations

* [Delete](#delete) - Delete
* [Get](#get) - Get
* [ListHistory](#listhistory) - List History
* [Search](#search) - Search
* [SetLimit](#setlimit) - Set Limit
* [Suspend](#suspend) - Suspend
* [Unsuspend](#unsuspend) - Unsuspend

## Delete

Remove the subject's limit on this app entirely. The app is then bounded
 only by the fund, by any tenant-wide cap, and by whatever the subject
 sets again themselves.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.SubjectAppLimitService.Delete" method="delete" path="/api/v1/funds/subject-app-limits/{user_id}/{app_id}" -->
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

    res, err := s.SubjectAppLimit.Delete(ctx, operations.C1APIFundsV1SubjectAppLimitServiceDeleteRequest{
        AppID: "<id>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubjectAppLimitServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIFundsV1SubjectAppLimitServiceDeleteRequest](../../pkg/models/operations/c1apifundsv1subjectapplimitservicedeleterequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APIFundsV1SubjectAppLimitServiceDeleteResponse](../../pkg/models/operations/c1apifundsv1subjectapplimitservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get returns one subject's limit row for one app, together with any
 suspension acting as their pause. A subject with no row for this app is
 not found, meaning nothing bounds the app beyond the layers above.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.SubjectAppLimitService.Get" method="get" path="/api/v1/funds/subject-app-limits/{user_id}/{app_id}" -->
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

    res, err := s.SubjectAppLimit.Get(ctx, operations.C1APIFundsV1SubjectAppLimitServiceGetRequest{
        AppID: "<id>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubjectAppLimitServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.C1APIFundsV1SubjectAppLimitServiceGetRequest](../../pkg/models/operations/c1apifundsv1subjectapplimitservicegetrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.C1APIFundsV1SubjectAppLimitServiceGetResponse](../../pkg/models/operations/c1apifundsv1subjectapplimitservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHistory

List the change history for one subject's limit on one app, newest
 first. Admin-tier per the object-history convention, and the same
 ObjectHistory stream the subject reads through MyFundLimitsService:
 admin-originated mutations are attributable here and self-service
 history shows them to the subject.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.SubjectAppLimitService.ListHistory" method="get" path="/api/v1/funds/subject-app-limits/{user_id}/{app_id}/history" -->
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

    res, err := s.SubjectAppLimit.ListHistory(ctx, operations.C1APIFundsV1SubjectAppLimitServiceListHistoryRequest{
        AppID: "<id>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubjectAppLimitServiceListHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APIFundsV1SubjectAppLimitServiceListHistoryRequest](../../pkg/models/operations/c1apifundsv1subjectapplimitservicelisthistoryrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APIFundsV1SubjectAppLimitServiceListHistoryResponse](../../pkg/models/operations/c1apifundsv1subjectapplimitservicelisthistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Search

Search the tenant's subject-app limit rows. Reads the Postgres mirror:
 the runtime row is keyed on (tenant, subject), so a tenant-wide listing
 has no runtime query at all. Filters narrow by subject, by app, and by
 state; empty filters return every live row in the tenant.

 POST rather than GET: the filters are repeated fields, which the gateway
 cannot carry as query parameters. Pagination is keyset-based over
 (pksk), the storage primary key, so a concurrent write to a filtered
 field cannot duplicate or omit a row across pages the way an offset page
 can.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.SubjectAppLimitService.Search" method="post" path="/api/v1/funds/subject-app-limits/search" -->
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

    res, err := s.SubjectAppLimit.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SubjectAppLimitServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [shared.SubjectAppLimitServiceSearchRequest](../../pkg/models/shared/subjectapplimitservicesearchrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.C1APIFundsV1SubjectAppLimitServiceSearchResponse](../../pkg/models/operations/c1apifundsv1subjectapplimitservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetLimit

Cap what one app may take from this subject's fund, creating the row if
 absent. Amount arm only; leaves any pause in place.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.SubjectAppLimitService.SetLimit" method="post" path="/api/v1/funds/subject-app-limits/{user_id}/{app_id}/limit" -->
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

    res, err := s.SubjectAppLimit.SetLimit(ctx, operations.C1APIFundsV1SubjectAppLimitServiceSetLimitRequest{
        AppID: "<id>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubjectAppLimitServiceSetLimitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.C1APIFundsV1SubjectAppLimitServiceSetLimitRequest](../../pkg/models/operations/c1apifundsv1subjectapplimitservicesetlimitrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.C1APIFundsV1SubjectAppLimitServiceSetLimitResponse](../../pkg/models/operations/c1apifundsv1subjectapplimitservicesetlimitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Suspend

Pause one app on this subject's fund. The limit underneath is preserved
 and restored by Unsuspend.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.SubjectAppLimitService.Suspend" method="post" path="/api/v1/funds/subject-app-limits/{user_id}/{app_id}/suspension" -->
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

    res, err := s.SubjectAppLimit.Suspend(ctx, operations.C1APIFundsV1SubjectAppLimitServiceSuspendRequest{
        AppID: "<id>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubjectAppLimitServiceSuspendResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIFundsV1SubjectAppLimitServiceSuspendRequest](../../pkg/models/operations/c1apifundsv1subjectapplimitservicesuspendrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIFundsV1SubjectAppLimitServiceSuspendResponse](../../pkg/models/operations/c1apifundsv1subjectapplimitservicesuspendresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Unsuspend

Lift the pause, restoring the limit it froze. Clearing the last control
 removes the row, so a vacuous result is returned as an absent limit, the
 same way the self-service plane reports it.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.SubjectAppLimitService.Unsuspend" method="delete" path="/api/v1/funds/subject-app-limits/{user_id}/{app_id}/suspension" -->
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

    res, err := s.SubjectAppLimit.Unsuspend(ctx, operations.C1APIFundsV1SubjectAppLimitServiceUnsuspendRequest{
        AppID: "<id>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubjectAppLimitServiceUnsuspendResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIFundsV1SubjectAppLimitServiceUnsuspendRequest](../../pkg/models/operations/c1apifundsv1subjectapplimitserviceunsuspendrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APIFundsV1SubjectAppLimitServiceUnsuspendResponse](../../pkg/models/operations/c1apifundsv1subjectapplimitserviceunsuspendresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |