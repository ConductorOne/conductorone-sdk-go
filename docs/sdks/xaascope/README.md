# XAAScope

## Overview

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List
* [ListHistory](#listhistory) - List History
* [Search](#search) - Search
* [Update](#update) - Update

## Create

Declare a scope under a resource server. The scope value is the literal
 OAuth scope string and is immutable after creation.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAScopeService.Create" method="post" path="/api/v1/apps/{app_id}/xaa/scopes" -->
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

    res, err := s.XAAScope.Create(ctx, operations.C1APICrossAppAccessV1XAAScopeServiceCreateRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAScopeServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.C1APICrossAppAccessV1XAAScopeServiceCreateRequest](../../pkg/models/operations/c1apicrossappaccessv1xaascopeservicecreaterequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.C1APICrossAppAccessV1XAAScopeServiceCreateResponse](../../pkg/models/operations/c1apicrossappaccessv1xaascopeservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a scope (soft delete). Cascades to its backing entitlement and to
 any access-profile bindings that reference it.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAScopeService.Delete" method="delete" path="/api/v1/apps/{app_id}/xaa/scopes/{id}" -->
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

    res, err := s.XAAScope.Delete(ctx, operations.C1APICrossAppAccessV1XAAScopeServiceDeleteRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAScopeServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.C1APICrossAppAccessV1XAAScopeServiceDeleteRequest](../../pkg/models/operations/c1apicrossappaccessv1xaascopeservicedeleterequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.C1APICrossAppAccessV1XAAScopeServiceDeleteResponse](../../pkg/models/operations/c1apicrossappaccessv1xaascopeservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get a scope by app_id + id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAScopeService.Get" method="get" path="/api/v1/apps/{app_id}/xaa/scopes/{id}" -->
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

    res, err := s.XAAScope.Get(ctx, operations.C1APICrossAppAccessV1XAAScopeServiceGetRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAScopeServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.C1APICrossAppAccessV1XAAScopeServiceGetRequest](../../pkg/models/operations/c1apicrossappaccessv1xaascopeservicegetrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `opts`                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.C1APICrossAppAccessV1XAAScopeServiceGetResponse](../../pkg/models/operations/c1apicrossappaccessv1xaascopeservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List the scopes defined for an application, one page at a time. To filter
 by resource server, state, or classification, use Search.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAScopeService.List" method="get" path="/api/v1/apps/{app_id}/xaa/scopes" -->
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

    res, err := s.XAAScope.List(ctx, operations.C1APICrossAppAccessV1XAAScopeServiceListRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAScopeServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APICrossAppAccessV1XAAScopeServiceListRequest](../../pkg/models/operations/c1apicrossappaccessv1xaascopeservicelistrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APICrossAppAccessV1XAAScopeServiceListResponse](../../pkg/models/operations/c1apicrossappaccessv1xaascopeservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHistory

ListHistory returns the change history (newest first) for a single scope —
 each entry is a snapshot plus who/when metadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAScopeService.ListHistory" method="get" path="/api/v1/apps/{app_id}/xaa/scopes/{id}/history" -->
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

    res, err := s.XAAScope.ListHistory(ctx, operations.C1APICrossAppAccessV1XAAScopeServiceListHistoryRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAScopeServiceListHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APICrossAppAccessV1XAAScopeServiceListHistoryRequest](../../pkg/models/operations/c1apicrossappaccessv1xaascopeservicelisthistoryrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APICrossAppAccessV1XAAScopeServiceListHistoryResponse](../../pkg/models/operations/c1apicrossappaccessv1xaascopeservicelisthistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Search

Search scopes across the tenant, filtered by resource server, state,
 classification, source, or text query, or fetch a specific set by ref.
 Filter on PENDING_REVIEW to find scopes awaiting approval.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAScopeService.Search" method="post" path="/api/v1/search/xaa/scopes" -->
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

    res, err := s.XAAScope.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAScopeServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [shared.XAAScopeServiceSearchRequest](../../pkg/models/shared/xaascopeservicesearchrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.C1APICrossAppAccessV1XAAScopeServiceSearchResponse](../../pkg/models/operations/c1apicrossappaccessv1xaascopeservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update a scope's editable fields via update_mask. This is how a scope is
 approved: set state to ENABLED to make it mintable, or DISABLED to block
 it. The scope value is immutable. Editable paths: display_name,
 description, state, classification.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAScopeService.Update" method="post" path="/api/v1/apps/{app_id}/xaa/scopes/{id}" -->
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

    res, err := s.XAAScope.Update(ctx, operations.C1APICrossAppAccessV1XAAScopeServiceUpdateRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAScopeServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.C1APICrossAppAccessV1XAAScopeServiceUpdateRequest](../../pkg/models/operations/c1apicrossappaccessv1xaascopeserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.C1APICrossAppAccessV1XAAScopeServiceUpdateResponse](../../pkg/models/operations/c1apicrossappaccessv1xaascopeserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |