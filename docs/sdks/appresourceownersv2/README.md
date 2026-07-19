# AppResourceOwnersV2

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

CreateEntitlementOwner creates an entitlement ownership source for an app resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppResourceOwnersV2.CreateEntitlementOwner" method="post" path="/api/v2/apps/{app_id}/resource_types/{resource_type_id}/resources/{resource_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.AppResourceOwnersV2.CreateEntitlementOwner(ctx, operations.C1APIAppV2AppResourceOwnersV2CreateEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
        RoleSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppResourceEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APIAppV2AppResourceOwnersV2CreateEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2appresourceownersv2createentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APIAppV2AppResourceOwnersV2CreateEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2appresourceownersv2createentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateUserOwner

CreateUserOwner creates a user ownership source for an app resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppResourceOwnersV2.CreateUserOwner" method="post" path="/api/v2/apps/{app_id}/resource_types/{resource_type_id}/resources/{resource_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.AppResourceOwnersV2.CreateUserOwner(ctx, operations.C1APIAppV2AppResourceOwnersV2CreateUserOwnerRequest{
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
        RoleSlug: "<value>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppResourceUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.C1APIAppV2AppResourceOwnersV2CreateUserOwnerRequest](../../pkg/models/operations/c1apiappv2appresourceownersv2createuserownerrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.C1APIAppV2AppResourceOwnersV2CreateUserOwnerResponse](../../pkg/models/operations/c1apiappv2appresourceownersv2createuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteEntitlementOwner

DeleteEntitlementOwner deletes an entitlement ownership source for an app resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppResourceOwnersV2.DeleteEntitlementOwner" method="delete" path="/api/v2/apps/{app_id}/resource_types/{resource_type_id}/resources/{resource_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.AppResourceOwnersV2.DeleteEntitlementOwner(ctx, operations.C1APIAppV2AppResourceOwnersV2DeleteEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
        RoleSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppResourceEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APIAppV2AppResourceOwnersV2DeleteEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2appresourceownersv2deleteentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APIAppV2AppResourceOwnersV2DeleteEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2appresourceownersv2deleteentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteUserOwner

DeleteUserOwner deletes a user ownership source for an app resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppResourceOwnersV2.DeleteUserOwner" method="delete" path="/api/v2/apps/{app_id}/resource_types/{resource_type_id}/resources/{resource_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.AppResourceOwnersV2.DeleteUserOwner(ctx, operations.C1APIAppV2AppResourceOwnersV2DeleteUserOwnerRequest{
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
        RoleSlug: "<value>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppResourceUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.C1APIAppV2AppResourceOwnersV2DeleteUserOwnerRequest](../../pkg/models/operations/c1apiappv2appresourceownersv2deleteuserownerrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.C1APIAppV2AppResourceOwnersV2DeleteUserOwnerResponse](../../pkg/models/operations/c1apiappv2appresourceownersv2deleteuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetEntitlementOwner

GetEntitlementOwner gets an entitlement ownership source for an app resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppResourceOwnersV2.GetEntitlementOwner" method="get" path="/api/v2/apps/{app_id}/resource_types/{resource_type_id}/resources/{resource_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.AppResourceOwnersV2.GetEntitlementOwner(ctx, operations.C1APIAppV2AppResourceOwnersV2GetEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
        RoleSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAppResourceEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIAppV2AppResourceOwnersV2GetEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2appresourceownersv2getentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIAppV2AppResourceOwnersV2GetEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2appresourceownersv2getentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUserOwner

GetUserOwner gets a user ownership source for an app resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppResourceOwnersV2.GetUserOwner" method="get" path="/api/v2/apps/{app_id}/resource_types/{resource_type_id}/resources/{resource_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.AppResourceOwnersV2.GetUserOwner(ctx, operations.C1APIAppV2AppResourceOwnersV2GetUserOwnerRequest{
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
        RoleSlug: "<value>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAppResourceUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAppV2AppResourceOwnersV2GetUserOwnerRequest](../../pkg/models/operations/c1apiappv2appresourceownersv2getuserownerrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIAppV2AppResourceOwnersV2GetUserOwnerResponse](../../pkg/models/operations/c1apiappv2appresourceownersv2getuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchEntitlementOwners

SearchEntitlementOwners searches for entitlement ownership sources of this app resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppResourceOwnersV2.SearchEntitlementOwners" method="get" path="/api/v2/apps/{app_id}/resource_types/{resource_type_id}/resources/{resource_id}/owners/entitlements" -->
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

    res, err := s.AppResourceOwnersV2.SearchEntitlementOwners(ctx, operations.C1APIAppV2AppResourceOwnersV2SearchEntitlementOwnersRequest{
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppResourceEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.C1APIAppV2AppResourceOwnersV2SearchEntitlementOwnersRequest](../../pkg/models/operations/c1apiappv2appresourceownersv2searchentitlementownersrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.C1APIAppV2AppResourceOwnersV2SearchEntitlementOwnersResponse](../../pkg/models/operations/c1apiappv2appresourceownersv2searchentitlementownersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchUserOwners

SearchUserOwners searches for user ownership sources of this app resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppResourceOwnersV2.SearchUserOwners" method="get" path="/api/v2/apps/{app_id}/resource_types/{resource_type_id}/resources/{resource_id}/owners/users" -->
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

    res, err := s.AppResourceOwnersV2.SearchUserOwners(ctx, operations.C1APIAppV2AppResourceOwnersV2SearchUserOwnersRequest{
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppResourceUserOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APIAppV2AppResourceOwnersV2SearchUserOwnersRequest](../../pkg/models/operations/c1apiappv2appresourceownersv2searchuserownersrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APIAppV2AppResourceOwnersV2SearchUserOwnersResponse](../../pkg/models/operations/c1apiappv2appresourceownersv2searchuserownersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Set

Set replaces all owners for a given app resource and role.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppResourceOwnersV2.Set" method="put" path="/api/v2/apps/{app_id}/resource_types/{resource_type_id}/resources/{resource_id}/owners" -->
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

    res, err := s.AppResourceOwnersV2.Set(ctx, operations.C1APIAppV2AppResourceOwnersV2SetRequest{
        AppID: "<id>",
        ResourceID: "<id>",
        ResourceTypeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SetAppResourceOwnersV2Response != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.C1APIAppV2AppResourceOwnersV2SetRequest](../../pkg/models/operations/c1apiappv2appresourceownersv2setrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.C1APIAppV2AppResourceOwnersV2SetResponse](../../pkg/models/operations/c1apiappv2appresourceownersv2setresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |