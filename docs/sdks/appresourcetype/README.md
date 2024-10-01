# AppResourceType
(*AppResourceType*)

## Overview

### Available Operations

* [CreateManuallyManagedResourceType](#createmanuallymanagedresourcetype) - Create Manually Managed Resource Type
* [DeleteManuallyManagedResourceType](#deletemanuallymanagedresourcetype) - Delete Manually Managed Resource Type
* [Get](#get) - Get
* [List](#list) - List
* [UpdateManuallyManagedResourceType](#updatemanuallymanagedresourcetype) - Update Manually Managed Resource Type

## CreateManuallyManagedResourceType

Invokes the c1.api.app.v1.AppResourceTypeService.CreateManuallyManagedResourceType method.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.AppResourceType.CreateManuallyManagedResourceType(ctx, operations.C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateManuallyManagedResourceTypeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                      | Type                                                                                                                                                                                           | Required                                                                                                                                                                                       | Description                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                             | The context to use for the request.                                                                                                                                                            |
| `request`                                                                                                                                                                                      | [operations.C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeRequest](../../pkg/models/operations/c1apiappv1appresourcetypeservicecreatemanuallymanagedresourcetyperequest.md) | :heavy_check_mark:                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                     |
| `opts`                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                             | The options for this request.                                                                                                                                                                  |

### Response

**[*operations.C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeResponse](../../pkg/models/operations/c1apiappv1appresourcetypeservicecreatemanuallymanagedresourcetyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteManuallyManagedResourceType

Invokes the c1.api.app.v1.AppResourceTypeService.DeleteManuallyManagedResourceType method.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.AppResourceType.DeleteManuallyManagedResourceType(ctx, operations.C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteManuallyManagedResourceTypeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                      | Type                                                                                                                                                                                           | Required                                                                                                                                                                                       | Description                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                             | The context to use for the request.                                                                                                                                                            |
| `request`                                                                                                                                                                                      | [operations.C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeRequest](../../pkg/models/operations/c1apiappv1appresourcetypeservicedeletemanuallymanagedresourcetyperequest.md) | :heavy_check_mark:                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                     |
| `opts`                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                             | The options for this request.                                                                                                                                                                  |

### Response

**[*operations.C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeResponse](../../pkg/models/operations/c1apiappv1appresourcetypeservicedeletemanuallymanagedresourcetyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get an app resource type.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.AppResourceType.Get(ctx, operations.C1APIAppV1AppResourceTypeServiceGetRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AppResourceTypeServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.C1APIAppV1AppResourceTypeServiceGetRequest](../../pkg/models/operations/c1apiappv1appresourcetypeservicegetrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.C1APIAppV1AppResourceTypeServiceGetResponse](../../pkg/models/operations/c1apiappv1appresourcetypeservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List app resource types.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.AppResourceType.List(ctx, operations.C1APIAppV1AppResourceTypeServiceListRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AppResourceTypeServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.C1APIAppV1AppResourceTypeServiceListRequest](../../pkg/models/operations/c1apiappv1appresourcetypeservicelistrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.C1APIAppV1AppResourceTypeServiceListResponse](../../pkg/models/operations/c1apiappv1appresourcetypeservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateManuallyManagedResourceType

Invokes the c1.api.app.v1.AppResourceTypeService.UpdateManuallyManagedResourceType method.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.AppResourceType.UpdateManuallyManagedResourceType(ctx, operations.C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateManuallyManagedResourceTypeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                      | Type                                                                                                                                                                                           | Required                                                                                                                                                                                       | Description                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                             | The context to use for the request.                                                                                                                                                            |
| `request`                                                                                                                                                                                      | [operations.C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeRequest](../../pkg/models/operations/c1apiappv1appresourcetypeserviceupdatemanuallymanagedresourcetyperequest.md) | :heavy_check_mark:                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                     |
| `opts`                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                             | The options for this request.                                                                                                                                                                  |

### Response

**[*operations.C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeResponse](../../pkg/models/operations/c1apiappv1appresourcetypeserviceupdatemanuallymanagedresourcetyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |