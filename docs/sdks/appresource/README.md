# AppResource

## Overview

### Available Operations

* [CreateManuallyManagedAppResource](#createmanuallymanagedappresource) - Create Manually Managed App Resource
* [DeleteManuallyManagedAppResource](#deletemanuallymanagedappresource) - Delete Manually Managed App Resource
* [Get](#get) - Get
* [List](#list) - List
* [Update](#update) - Update

## CreateManuallyManagedAppResource

Create a manually managed app resource tracked directly by ConductorOne under an existing resource type.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceService.CreateManuallyManagedAppResource" method="post" path="/api/v1/apps/{app_id}/resource_types/{app_resource_type_id}/resources" -->
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

    res, err := s.AppResource.CreateManuallyManagedAppResource(ctx, operations.C1APIAppV1AppResourceServiceCreateManuallyManagedAppResourceRequest{
        AppID: "<id>",
        AppResourceTypeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateManuallyManagedAppResourceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                            | Type                                                                                                                                                                                 | Required                                                                                                                                                                             | Description                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                | :heavy_check_mark:                                                                                                                                                                   | The context to use for the request.                                                                                                                                                  |
| `request`                                                                                                                                                                            | [operations.C1APIAppV1AppResourceServiceCreateManuallyManagedAppResourceRequest](../../pkg/models/operations/c1apiappv1appresourceservicecreatemanuallymanagedappresourcerequest.md) | :heavy_check_mark:                                                                                                                                                                   | The request object to use for the request.                                                                                                                                           |
| `opts`                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                   | The options for this request.                                                                                                                                                        |

### Response

**[*operations.C1APIAppV1AppResourceServiceCreateManuallyManagedAppResourceResponse](../../pkg/models/operations/c1apiappv1appresourceservicecreatemanuallymanagedappresourceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteManuallyManagedAppResource

Delete a manually managed app resource and its associated entitlements from an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceService.DeleteManuallyManagedAppResource" method="delete" path="/api/v1/apps/{app_id}/resource_types/{app_resource_type_id}/resources/{id}" -->
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

    res, err := s.AppResource.DeleteManuallyManagedAppResource(ctx, operations.C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceRequest{
        AppID: "<id>",
        AppResourceTypeID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteManuallyManagedAppResourceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                            | Type                                                                                                                                                                                 | Required                                                                                                                                                                             | Description                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                | :heavy_check_mark:                                                                                                                                                                   | The context to use for the request.                                                                                                                                                  |
| `request`                                                                                                                                                                            | [operations.C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceRequest](../../pkg/models/operations/c1apiappv1appresourceservicedeletemanuallymanagedappresourcerequest.md) | :heavy_check_mark:                                                                                                                                                                   | The request object to use for the request.                                                                                                                                           |
| `opts`                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                   | The options for this request.                                                                                                                                                        |

### Response

**[*operations.C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceResponse](../../pkg/models/operations/c1apiappv1appresourceservicedeletemanuallymanagedappresourceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a single app resource by its app, resource type, and resource ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceService.Get" method="get" path="/api/v1/apps/{app_id}/resource_types/{app_resource_type_id}/resources/{id}" -->
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

    res, err := s.AppResource.Get(ctx, operations.C1APIAppV1AppResourceServiceGetRequest{
        AppID: "<id>",
        AppResourceTypeID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AppResourceServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.C1APIAppV1AppResourceServiceGetRequest](../../pkg/models/operations/c1apiappv1appresourceservicegetrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.C1APIAppV1AppResourceServiceGetResponse](../../pkg/models/operations/c1apiappv1appresourceservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List app resources for a given app and optionally filter by resource type.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceService.List" method="get" path="/api/v1/apps/{app_id}/resource_types/{app_resource_type_id}/resources" -->
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

    res, err := s.AppResource.List(ctx, operations.C1APIAppV1AppResourceServiceListRequest{
        AppID: "<id>",
        AppResourceTypeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AppResourceServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.C1APIAppV1AppResourceServiceListRequest](../../pkg/models/operations/c1apiappv1appresourceservicelistrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.C1APIAppV1AppResourceServiceListResponse](../../pkg/models/operations/c1apiappv1appresourceservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update an app resource's fields. Only the fields specified in the update mask are modified.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceService.Update" method="post" path="/api/v1/apps/{app_id}/resource_types/{app_resource_type_id}/resources/{id}" -->
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

    res, err := s.AppResource.Update(ctx, operations.C1APIAppV1AppResourceServiceUpdateRequest{
        AppID: "<id>",
        AppResourceTypeID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AppResourceServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIAppV1AppResourceServiceUpdateRequest](../../pkg/models/operations/c1apiappv1appresourceserviceupdaterequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APIAppV1AppResourceServiceUpdateResponse](../../pkg/models/operations/c1apiappv1appresourceserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |