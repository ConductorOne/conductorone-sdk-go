# AppEntitlementOwners

## Overview

### Available Operations

* [Add](#add) - Add
* [Delete](#delete) - Delete
* [List](#list) - List
* [ListOwnerIDs](#listownerids) - List Owner I Ds
* [Remove](#remove) - Remove
* [Set](#set) - Set

## Add

Add an owner to a given app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementOwners.Add" method="post" path="/api/v1/apps/{app_id}/entitlements/{entitlement_id}/owners" -->
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

    res, err := s.AppEntitlementOwners.Add(ctx, operations.C1APIAppV1AppEntitlementOwnersAddRequest{
        AppID: "<id>",
        EntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AddAppEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APIAppV1AppEntitlementOwnersAddRequest](../../pkg/models/operations/c1apiappv1appentitlementownersaddrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.C1APIAppV1AppEntitlementOwnersAddResponse](../../pkg/models/operations/c1apiappv1appentitlementownersaddresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete deletes the owners from a given app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementOwners.Delete" method="delete" path="/api/v1/apps/{app_id}/entitlements/{entitlement_id}/owners" -->
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

    res, err := s.AppEntitlementOwners.Delete(ctx, operations.C1APIAppV1AppEntitlementOwnersDeleteRequest{
        AppID: "<id>",
        EntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.C1APIAppV1AppEntitlementOwnersDeleteRequest](../../pkg/models/operations/c1apiappv1appentitlementownersdeleterequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.C1APIAppV1AppEntitlementOwnersDeleteResponse](../../pkg/models/operations/c1apiappv1appentitlementownersdeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List owners for a given app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementOwners.List" method="get" path="/api/v1/apps/{app_id}/entitlements/{entitlement_id}/owners" -->
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

    res, err := s.AppEntitlementOwners.List(ctx, operations.C1APIAppV1AppEntitlementOwnersListRequest{
        AppID: "<id>",
        EntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIAppV1AppEntitlementOwnersListRequest](../../pkg/models/operations/c1apiappv1appentitlementownerslistrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APIAppV1AppEntitlementOwnersListResponse](../../pkg/models/operations/c1apiappv1appentitlementownerslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListOwnerIDs

ListUserIDs lists owner IDs for a given app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementOwners.ListOwnerIDs" method="get" path="/api/v1/apps/{app_id}/entitlements/{entitlement_id}/ownerids" -->
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

    res, err := s.AppEntitlementOwners.ListOwnerIDs(ctx, operations.C1APIAppV1AppEntitlementOwnersListOwnerIDsRequest{
        AppID: "<id>",
        EntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppEntitlementOwnerIDsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.C1APIAppV1AppEntitlementOwnersListOwnerIDsRequest](../../pkg/models/operations/c1apiappv1appentitlementownerslistowneridsrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.C1APIAppV1AppEntitlementOwnersListOwnerIDsResponse](../../pkg/models/operations/c1apiappv1appentitlementownerslistowneridsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Remove

Remove an owner from a given app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementOwners.Remove" method="delete" path="/api/v1/apps/{app_id}/entitlements/{entitlement_id}/owners/{user_id}" -->
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

    res, err := s.AppEntitlementOwners.Remove(ctx, operations.C1APIAppV1AppEntitlementOwnersRemoveRequest{
        AppID: "<id>",
        EntitlementID: "<id>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RemoveAppEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.C1APIAppV1AppEntitlementOwnersRemoveRequest](../../pkg/models/operations/c1apiappv1appentitlementownersremoverequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.C1APIAppV1AppEntitlementOwnersRemoveResponse](../../pkg/models/operations/c1apiappv1appentitlementownersremoveresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Set

Sets the owners for a given app entitlement to the specified list of users.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementOwners.Set" method="put" path="/api/v1/apps/{app_id}/entitlements/{entitlement_id}/owners" -->
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

    res, err := s.AppEntitlementOwners.Set(ctx, operations.C1APIAppV1AppEntitlementOwnersSetRequest{
        AppID: "<id>",
        EntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SetAppEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APIAppV1AppEntitlementOwnersSetRequest](../../pkg/models/operations/c1apiappv1appentitlementownerssetrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.C1APIAppV1AppEntitlementOwnersSetResponse](../../pkg/models/operations/c1apiappv1appentitlementownerssetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |