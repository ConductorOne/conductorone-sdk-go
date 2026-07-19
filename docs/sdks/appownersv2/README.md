# AppOwnersV2

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

CreateEntitlementOwner creates an entitlement ownership source for an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppOwners.CreateEntitlementOwner" method="post" path="/api/v2/apps/{app_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.AppOwnersV2.CreateEntitlementOwner(ctx, operations.C1APIAppV2AppOwnersCreateEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        RoleSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAppV2AppOwnersCreateEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2appownerscreateentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIAppV2AppOwnersCreateEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2appownerscreateentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateUserOwner

CreateUserOwner creates a user ownership source for an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppOwners.CreateUserOwner" method="post" path="/api/v2/apps/{app_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.AppOwnersV2.CreateUserOwner(ctx, operations.C1APIAppV2AppOwnersCreateUserOwnerRequest{
        AppID: "<id>",
        RoleSlug: "<value>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIAppV2AppOwnersCreateUserOwnerRequest](../../pkg/models/operations/c1apiappv2appownerscreateuserownerrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APIAppV2AppOwnersCreateUserOwnerResponse](../../pkg/models/operations/c1apiappv2appownerscreateuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteEntitlementOwner

DeleteEntitlementOwner deletes an entitlement ownership source for an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppOwners.DeleteEntitlementOwner" method="delete" path="/api/v2/apps/{app_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.AppOwnersV2.DeleteEntitlementOwner(ctx, operations.C1APIAppV2AppOwnersDeleteEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        RoleSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAppV2AppOwnersDeleteEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2appownersdeleteentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIAppV2AppOwnersDeleteEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2appownersdeleteentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteUserOwner

DeleteUserOwner deletes a user ownership source for an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppOwners.DeleteUserOwner" method="delete" path="/api/v2/apps/{app_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.AppOwnersV2.DeleteUserOwner(ctx, operations.C1APIAppV2AppOwnersDeleteUserOwnerRequest{
        AppID: "<id>",
        RoleSlug: "<value>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIAppV2AppOwnersDeleteUserOwnerRequest](../../pkg/models/operations/c1apiappv2appownersdeleteuserownerrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APIAppV2AppOwnersDeleteUserOwnerResponse](../../pkg/models/operations/c1apiappv2appownersdeleteuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetEntitlementOwner

GetEntitlementOwner gets an entitlement ownership source for an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppOwners.GetEntitlementOwner" method="get" path="/api/v2/apps/{app_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.AppOwnersV2.GetEntitlementOwner(ctx, operations.C1APIAppV2AppOwnersGetEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        RoleSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAppEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.C1APIAppV2AppOwnersGetEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2appownersgetentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.C1APIAppV2AppOwnersGetEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2appownersgetentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUserOwner

GetUserOwner gets a user ownership source for an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppOwners.GetUserOwner" method="get" path="/api/v2/apps/{app_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.AppOwnersV2.GetUserOwner(ctx, operations.C1APIAppV2AppOwnersGetUserOwnerRequest{
        AppID: "<id>",
        RoleSlug: "<value>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAppUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.C1APIAppV2AppOwnersGetUserOwnerRequest](../../pkg/models/operations/c1apiappv2appownersgetuserownerrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.C1APIAppV2AppOwnersGetUserOwnerResponse](../../pkg/models/operations/c1apiappv2appownersgetuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchEntitlementOwners

SearchEntitlementOwners searches for entitlement ownership sources for an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppOwners.SearchEntitlementOwners" method="get" path="/api/v2/apps/{app_id}/owners/entitlements" -->
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

    res, err := s.AppOwnersV2.SearchEntitlementOwners(ctx, operations.C1APIAppV2AppOwnersSearchEntitlementOwnersRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.C1APIAppV2AppOwnersSearchEntitlementOwnersRequest](../../pkg/models/operations/c1apiappv2appownerssearchentitlementownersrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.C1APIAppV2AppOwnersSearchEntitlementOwnersResponse](../../pkg/models/operations/c1apiappv2appownerssearchentitlementownersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchUserOwners

SearchUserOwners searches for user ownership sources for an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppOwners.SearchUserOwners" method="get" path="/api/v2/apps/{app_id}/owners/users" -->
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

    res, err := s.AppOwnersV2.SearchUserOwners(ctx, operations.C1APIAppV2AppOwnersSearchUserOwnersRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppUserOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.C1APIAppV2AppOwnersSearchUserOwnersRequest](../../pkg/models/operations/c1apiappv2appownerssearchuserownersrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.C1APIAppV2AppOwnersSearchUserOwnersResponse](../../pkg/models/operations/c1apiappv2appownerssearchuserownersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Set

Set replaces all user owners for a given app and role.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppOwners.Set" method="put" path="/api/v2/apps/{app_id}/owners" -->
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

    res, err := s.AppOwnersV2.Set(ctx, operations.C1APIAppV2AppOwnersSetRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SetAppOwnersResponseV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.C1APIAppV2AppOwnersSetRequest](../../pkg/models/operations/c1apiappv2appownerssetrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.C1APIAppV2AppOwnersSetResponse](../../pkg/models/operations/c1apiappv2appownerssetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |