# FundAssignment

## Overview

### Available Operations

* [ClearExtension](#clearextension) - Clear Extension
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [GrantExtension](#grantextension) - Grant Extension
* [ListHistory](#listhistory) - List History
* [Search](#search) - Search
* [SetLimit](#setlimit) - Set Limit
* [Suspend](#suspend) - Suspend
* [Unsuspend](#unsuspend) - Unsuspend

## ClearExtension

Revoke the extension early. The base limit underneath is untouched.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundAssignmentService.ClearExtension" method="delete" path="/api/v1/funds/assignments/{user_id}/extension" -->
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

    res, err := s.FundAssignment.ClearExtension(ctx, operations.C1APIFundsV1FundAssignmentServiceClearExtensionRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FundAssignmentServiceClearExtensionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIFundsV1FundAssignmentServiceClearExtensionRequest](../../pkg/models/operations/c1apifundsv1fundassignmentserviceclearextensionrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIFundsV1FundAssignmentServiceClearExtensionResponse](../../pkg/models/operations/c1apifundsv1fundassignmentserviceclearextensionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the whole assignment. The subject falls back to the rules and the
 tenant default.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundAssignmentService.Delete" method="delete" path="/api/v1/funds/assignments/{user_id}" -->
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

    res, err := s.FundAssignment.Delete(ctx, operations.C1APIFundsV1FundAssignmentServiceDeleteRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FundAssignmentServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.C1APIFundsV1FundAssignmentServiceDeleteRequest](../../pkg/models/operations/c1apifundsv1fundassignmentservicedeleterequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `opts`                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.C1APIFundsV1FundAssignmentServiceDeleteResponse](../../pkg/models/operations/c1apifundsv1fundassignmentservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get returns one subject's exception: their own limit, any extension
 running on top of it, and any suspension. A subject with no exception is
 not found rather than reported at the tenant default, because no row is
 what "this subject is governed by the layers above" looks like.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundAssignmentService.Get" method="get" path="/api/v1/funds/assignments/{user_id}" -->
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

    res, err := s.FundAssignment.Get(ctx, operations.C1APIFundsV1FundAssignmentServiceGetRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FundAssignmentServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.C1APIFundsV1FundAssignmentServiceGetRequest](../../pkg/models/operations/c1apifundsv1fundassignmentservicegetrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.C1APIFundsV1FundAssignmentServiceGetResponse](../../pkg/models/operations/c1apifundsv1fundassignmentservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GrantExtension

Grant a temporary total until expires_at. Never changes the period, and
 never expresses a refusal — a temporary refusal is a suspension.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundAssignmentService.GrantExtension" method="post" path="/api/v1/funds/assignments/{user_id}/extension" -->
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

    res, err := s.FundAssignment.GrantExtension(ctx, operations.C1APIFundsV1FundAssignmentServiceGrantExtensionRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FundAssignmentServiceGrantExtensionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIFundsV1FundAssignmentServiceGrantExtensionRequest](../../pkg/models/operations/c1apifundsv1fundassignmentservicegrantextensionrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIFundsV1FundAssignmentServiceGrantExtensionResponse](../../pkg/models/operations/c1apifundsv1fundassignmentservicegrantextensionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHistory

List the change history for one subject's assignment, newest first.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundAssignmentService.ListHistory" method="get" path="/api/v1/funds/assignments/{user_id}/history" -->
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

    res, err := s.FundAssignment.ListHistory(ctx, operations.C1APIFundsV1FundAssignmentServiceListHistoryRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FundAssignmentServiceListHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.C1APIFundsV1FundAssignmentServiceListHistoryRequest](../../pkg/models/operations/c1apifundsv1fundassignmentservicelisthistoryrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.C1APIFundsV1FundAssignmentServiceListHistoryResponse](../../pkg/models/operations/c1apifundsv1fundassignmentservicelisthistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Search

Search the tenant's assignments. Reads the Postgres mirror: the runtime
 row is keyed on (tenant, user), so there is no cross-subject query on the
 runtime plane at all.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundAssignmentService.Search" method="post" path="/api/v1/funds/assignments/search" -->
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

    res, err := s.FundAssignment.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FundAssignmentServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [shared.FundAssignmentServiceSearchRequest](../../pkg/models/shared/fundassignmentservicesearchrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.C1APIFundsV1FundAssignmentServiceSearchResponse](../../pkg/models/operations/c1apifundsv1fundassignmentservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetLimit

Set the subject's base limit, creating the assignment if absent. Leaves any
 extension and any suspension in place.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundAssignmentService.SetLimit" method="post" path="/api/v1/funds/assignments/{user_id}/limit" -->
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

    res, err := s.FundAssignment.SetLimit(ctx, operations.C1APIFundsV1FundAssignmentServiceSetLimitRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FundAssignmentServiceSetLimitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIFundsV1FundAssignmentServiceSetLimitRequest](../../pkg/models/operations/c1apifundsv1fundassignmentservicesetlimitrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIFundsV1FundAssignmentServiceSetLimitResponse](../../pkg/models/operations/c1apifundsv1fundassignmentservicesetlimitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Suspend

Freeze the subject's fund. The limit and any extension underneath are
 preserved and restored by Unsuspend.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundAssignmentService.Suspend" method="post" path="/api/v1/funds/assignments/{user_id}/suspension" -->
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

    res, err := s.FundAssignment.Suspend(ctx, operations.C1APIFundsV1FundAssignmentServiceSuspendRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FundAssignmentServiceSuspendResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIFundsV1FundAssignmentServiceSuspendRequest](../../pkg/models/operations/c1apifundsv1fundassignmentservicesuspendrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APIFundsV1FundAssignmentServiceSuspendResponse](../../pkg/models/operations/c1apifundsv1fundassignmentservicesuspendresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Unsuspend

Lift the suspension, restoring the numbers it froze.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundAssignmentService.Unsuspend" method="delete" path="/api/v1/funds/assignments/{user_id}/suspension" -->
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

    res, err := s.FundAssignment.Unsuspend(ctx, operations.C1APIFundsV1FundAssignmentServiceUnsuspendRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FundAssignmentServiceUnsuspendResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.C1APIFundsV1FundAssignmentServiceUnsuspendRequest](../../pkg/models/operations/c1apifundsv1fundassignmentserviceunsuspendrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.C1APIFundsV1FundAssignmentServiceUnsuspendResponse](../../pkg/models/operations/c1apifundsv1fundassignmentserviceunsuspendresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |