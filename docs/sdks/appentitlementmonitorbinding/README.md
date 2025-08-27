# AppEntitlementMonitorBinding
(*AppEntitlementMonitorBinding*)

## Overview

### Available Operations

* [CreateAppEntitlementMonitorBinding](#createappentitlementmonitorbinding) - Create App Entitlement Monitor Binding
* [DeleteAppEntitlementMonitorBinding](#deleteappentitlementmonitorbinding) - Delete App Entitlement Monitor Binding
* [GetAppEntitlementMonitorBinding](#getappentitlementmonitorbinding) - Get App Entitlement Monitor Binding

## CreateAppEntitlementMonitorBinding

Invokes the c1.api.accessconflict.v1.AppEntitlementMonitorBindingService.CreateAppEntitlementMonitorBinding method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessconflict.v1.AppEntitlementMonitorBindingService.CreateAppEntitlementMonitorBinding" method="post" path="/api/v1/appentitlementmonitorbinding" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.AppEntitlementMonitorBinding.CreateAppEntitlementMonitorBinding(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementMonitorBinding != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [shared.CreateAppEntitlementMonitorBindingRequest](../../pkg/models/shared/createappentitlementmonitorbindingrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.C1APIAccessconflictV1AppEntitlementMonitorBindingServiceCreateAppEntitlementMonitorBindingResponse](../../pkg/models/operations/c1apiaccessconflictv1appentitlementmonitorbindingservicecreateappentitlementmonitorbindingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteAppEntitlementMonitorBinding

Invokes the c1.api.accessconflict.v1.AppEntitlementMonitorBindingService.DeleteAppEntitlementMonitorBinding method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessconflict.v1.AppEntitlementMonitorBindingService.DeleteAppEntitlementMonitorBinding" method="delete" path="/api/v1/appentitlementmonitorbinding" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.AppEntitlementMonitorBinding.DeleteAppEntitlementMonitorBinding(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppEntitlementMonitorBindingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [shared.DeleteAppEntitlementMonitorBindingRequest](../../pkg/models/shared/deleteappentitlementmonitorbindingrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.C1APIAccessconflictV1AppEntitlementMonitorBindingServiceDeleteAppEntitlementMonitorBindingResponse](../../pkg/models/operations/c1apiaccessconflictv1appentitlementmonitorbindingservicedeleteappentitlementmonitorbindingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAppEntitlementMonitorBinding

Invokes the c1.api.accessconflict.v1.AppEntitlementMonitorBindingService.GetAppEntitlementMonitorBinding method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessconflict.v1.AppEntitlementMonitorBindingService.GetAppEntitlementMonitorBinding" method="post" path="/api/v1/appentitlementmonitorbinding/get" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.AppEntitlementMonitorBinding.GetAppEntitlementMonitorBinding(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementMonitorBinding != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [shared.GetAppEntitlementMonitorBindingRequest](../../pkg/models/shared/getappentitlementmonitorbindingrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.C1APIAccessconflictV1AppEntitlementMonitorBindingServiceGetAppEntitlementMonitorBindingResponse](../../pkg/models/operations/c1apiaccessconflictv1appentitlementmonitorbindingservicegetappentitlementmonitorbindingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |