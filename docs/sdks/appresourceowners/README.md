# AppResourceOwners
(*AppResourceOwners*)

## Overview

### Available Operations

* [Add](#add) - Add
* [List](#list) - List
* [Remove](#remove) - Remove

## Add

Invokes the c1.api.app.v1.AppResourceOwners.Add method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceOwners.Add" method="post" path="/api/v1/apps/{app_id}/resource_types/{resource_type_id}/resource/{resource_id}/owners" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.AppResourceOwners.Add(ctx, operations.C1APIAppV1AppResourceOwnersAddRequest{
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AddAppResourceOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.C1APIAppV1AppResourceOwnersAddRequest](../../pkg/models/operations/c1apiappv1appresourceownersaddrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.C1APIAppV1AppResourceOwnersAddResponse](../../pkg/models/operations/c1apiappv1appresourceownersaddresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List all owners of an app resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceOwners.List" method="get" path="/api/v1/apps/{app_id}/resource_types/{resource_type_id}/resource/{resource_id}/owners" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.AppResourceOwners.List(ctx, operations.C1APIAppV1AppResourceOwnersListRequest{
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppResourceOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.C1APIAppV1AppResourceOwnersListRequest](../../pkg/models/operations/c1apiappv1appresourceownerslistrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.C1APIAppV1AppResourceOwnersListResponse](../../pkg/models/operations/c1apiappv1appresourceownerslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Remove

Invokes the c1.api.app.v1.AppResourceOwners.Remove method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceOwners.Remove" method="delete" path="/api/v1/apps/{app_id}/resource_types/{resource_type_id}/resource/{resource_id}/owners" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.AppResourceOwners.Remove(ctx, operations.C1APIAppV1AppResourceOwnersRemoveRequest{
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RemoveAppResourceOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APIAppV1AppResourceOwnersRemoveRequest](../../pkg/models/operations/c1apiappv1appresourceownersremoverequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.C1APIAppV1AppResourceOwnersRemoveResponse](../../pkg/models/operations/c1apiappv1appresourceownersremoveresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |