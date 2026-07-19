# AppEntitlementOwnersV2

## Overview

### Available Operations

* [CreateEntitlementOwner](#createentitlementowner) - Create Entitlement Owner
* [CreateUserOwner](#createuserowner) - Create User Owner
* [DeleteEntitlementOwner](#deleteentitlementowner) - Delete Entitlement Owner
* [DeleteUserOwner](#deleteuserowner) - Delete User Owner
* [GetEntitlementOwner](#getentitlementowner) - Get Entitlement Owner
* [GetUserOwner](#getuserowner) - Get User Owner
* [SearchEntitlementOwners](#searchentitlementowners) - Search Entitlement Owners
* [SearchUserOwners](#searchuserowners) - Search User Owners
* [Set](#set) - Set

## CreateEntitlementOwner

CreateEntitlementOwner creates an entitlement ownership source for an app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppEntitlementOwners.CreateEntitlementOwner" method="post" path="/api/v2/apps/{app_id}/entitlements/{entitlement_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.AppEntitlementOwnersV2.CreateEntitlementOwner(ctx, operations.C1APIAppV2AppEntitlementOwnersCreateEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        EntitlementID: "<id>",
        RoleSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppEntitlementEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.C1APIAppV2AppEntitlementOwnersCreateEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2appentitlementownerscreateentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.C1APIAppV2AppEntitlementOwnersCreateEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2appentitlementownerscreateentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateUserOwner

CreateUserOwner creates a user ownership source for an app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppEntitlementOwners.CreateUserOwner" method="post" path="/api/v2/apps/{app_id}/entitlements/{entitlement_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.AppEntitlementOwnersV2.CreateUserOwner(ctx, operations.C1APIAppV2AppEntitlementOwnersCreateUserOwnerRequest{
        AppID: "<id>",
        EntitlementID: "<id>",
        RoleSlug: "<value>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppEntitlementUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APIAppV2AppEntitlementOwnersCreateUserOwnerRequest](../../pkg/models/operations/c1apiappv2appentitlementownerscreateuserownerrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APIAppV2AppEntitlementOwnersCreateUserOwnerResponse](../../pkg/models/operations/c1apiappv2appentitlementownerscreateuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteEntitlementOwner

DeleteEntitlementOwner deletes an entitlement ownership source for an app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppEntitlementOwners.DeleteEntitlementOwner" method="delete" path="/api/v2/apps/{app_id}/entitlements/{entitlement_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.AppEntitlementOwnersV2.DeleteEntitlementOwner(ctx, operations.C1APIAppV2AppEntitlementOwnersDeleteEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        EntitlementID: "<id>",
        RoleSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppEntitlementEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.C1APIAppV2AppEntitlementOwnersDeleteEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2appentitlementownersdeleteentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.C1APIAppV2AppEntitlementOwnersDeleteEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2appentitlementownersdeleteentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteUserOwner

DeleteUserOwner deletes a user ownership source for an app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppEntitlementOwners.DeleteUserOwner" method="delete" path="/api/v2/apps/{app_id}/entitlements/{entitlement_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.AppEntitlementOwnersV2.DeleteUserOwner(ctx, operations.C1APIAppV2AppEntitlementOwnersDeleteUserOwnerRequest{
        AppID: "<id>",
        EntitlementID: "<id>",
        RoleSlug: "<value>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppEntitlementUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APIAppV2AppEntitlementOwnersDeleteUserOwnerRequest](../../pkg/models/operations/c1apiappv2appentitlementownersdeleteuserownerrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APIAppV2AppEntitlementOwnersDeleteUserOwnerResponse](../../pkg/models/operations/c1apiappv2appentitlementownersdeleteuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetEntitlementOwner

GetEntitlementOwner gets an entitlement ownership source for an app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppEntitlementOwners.GetEntitlementOwner" method="get" path="/api/v2/apps/{app_id}/entitlements/{entitlement_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.AppEntitlementOwnersV2.GetEntitlementOwner(ctx, operations.C1APIAppV2AppEntitlementOwnersGetEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        EntitlementID: "<id>",
        RoleSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAppEntitlementEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.C1APIAppV2AppEntitlementOwnersGetEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2appentitlementownersgetentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |
| `opts`                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                   | :heavy_minus_sign:                                                                                                                                             | The options for this request.                                                                                                                                  |

### Response

**[*operations.C1APIAppV2AppEntitlementOwnersGetEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2appentitlementownersgetentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUserOwner

GetUserOwner gets a user ownership source for an app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppEntitlementOwners.GetUserOwner" method="get" path="/api/v2/apps/{app_id}/entitlements/{entitlement_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.AppEntitlementOwnersV2.GetUserOwner(ctx, operations.C1APIAppV2AppEntitlementOwnersGetUserOwnerRequest{
        AppID: "<id>",
        EntitlementID: "<id>",
        RoleSlug: "<value>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAppEntitlementUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.C1APIAppV2AppEntitlementOwnersGetUserOwnerRequest](../../pkg/models/operations/c1apiappv2appentitlementownersgetuserownerrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.C1APIAppV2AppEntitlementOwnersGetUserOwnerResponse](../../pkg/models/operations/c1apiappv2appentitlementownersgetuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchEntitlementOwners

SearchEntitlementOwners searches for entitlement ownership sources for an app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppEntitlementOwners.SearchEntitlementOwners" method="get" path="/api/v2/apps/{app_id}/entitlements/{entitlement_id}/owners/entitlements" -->
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

    res, err := s.AppEntitlementOwnersV2.SearchEntitlementOwners(ctx, operations.C1APIAppV2AppEntitlementOwnersSearchEntitlementOwnersRequest{
        AppID: "<id>",
        EntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppEntitlementEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.C1APIAppV2AppEntitlementOwnersSearchEntitlementOwnersRequest](../../pkg/models/operations/c1apiappv2appentitlementownerssearchentitlementownersrequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |
| `opts`                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                           | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |

### Response

**[*operations.C1APIAppV2AppEntitlementOwnersSearchEntitlementOwnersResponse](../../pkg/models/operations/c1apiappv2appentitlementownerssearchentitlementownersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchUserOwners

SearchUserOwners searches for user ownership sources of this app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppEntitlementOwners.SearchUserOwners" method="get" path="/api/v2/apps/{app_id}/entitlements/{entitlement_id}/owners/users" -->
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

    res, err := s.AppEntitlementOwnersV2.SearchUserOwners(ctx, operations.C1APIAppV2AppEntitlementOwnersSearchUserOwnersRequest{
        AppID: "<id>",
        EntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppEntitlementUserOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APIAppV2AppEntitlementOwnersSearchUserOwnersRequest](../../pkg/models/operations/c1apiappv2appentitlementownerssearchuserownersrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APIAppV2AppEntitlementOwnersSearchUserOwnersResponse](../../pkg/models/operations/c1apiappv2appentitlementownerssearchuserownersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Set

Set replaces all owners for a given app entitlement and role.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppEntitlementOwners.Set" method="put" path="/api/v2/apps/{app_id}/entitlements/{entitlement_id}/owners" -->
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

    res, err := s.AppEntitlementOwnersV2.Set(ctx, operations.C1APIAppV2AppEntitlementOwnersSetRequest{
        AppID: "<id>",
        EntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SetAppEntitlementOwnersResponseV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APIAppV2AppEntitlementOwnersSetRequest](../../pkg/models/operations/c1apiappv2appentitlementownerssetrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.C1APIAppV2AppEntitlementOwnersSetResponse](../../pkg/models/operations/c1apiappv2appentitlementownerssetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |