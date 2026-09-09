# AppManagedState

## Overview

### Available Operations

* [Get](#get) - Get
* [List](#list) - List
* [Promote](#promote) - Promote

## Get

Get the managed state of a discovered application.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppManagedStateService.Get" method="get" path="/api/v1/apps/{app_id}/resource_types/{resource_type_id}/managed_state_bindings/{resource_id}" -->
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

    res, err := s.AppManagedState.Get(ctx, operations.C1APIAppV1AppManagedStateServiceGetRequest{
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAppManagedStateBindingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.C1APIAppV1AppManagedStateServiceGetRequest](../../pkg/models/operations/c1apiappv1appmanagedstateservicegetrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.C1APIAppV1AppManagedStateServiceGetResponse](../../pkg/models/operations/c1apiappv1appmanagedstateservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List the managed states of applications discovered by a connector.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppManagedStateService.List" method="get" path="/api/v1/apps/{app_id}/resource_types/{resource_type_id}/managed_state_bindings" -->
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

    res, err := s.AppManagedState.List(ctx, operations.C1APIAppV1AppManagedStateServiceListRequest{
        AppID: "<id>",
        ResourceTypeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppManagedStateBindingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.C1APIAppV1AppManagedStateServiceListRequest](../../pkg/models/operations/c1apiappv1appmanagedstateservicelistrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.C1APIAppV1AppManagedStateServiceListResponse](../../pkg/models/operations/c1apiappv1appmanagedstateservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Promote

Promote an unmanaged application into a managed application.
 Returns AlreadyExists when the application is already managed. The new application inherits source owners when user_ids is omitted.
 Concurrent promotion requests are not supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppManagedStateService.Promote" method="post" path="/api/v1/apps/{app_id}/resource_types/{resource_type_id}/managed_state_bindings/{resource_id}/promote" -->
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

    res, err := s.AppManagedState.Promote(ctx, operations.C1APIAppV1AppManagedStateServicePromoteRequest{
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAppManagedStateBindingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.C1APIAppV1AppManagedStateServicePromoteRequest](../../pkg/models/operations/c1apiappv1appmanagedstateservicepromoterequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `opts`                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.C1APIAppV1AppManagedStateServicePromoteResponse](../../pkg/models/operations/c1apiappv1appmanagedstateservicepromoteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |