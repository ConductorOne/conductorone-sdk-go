# AppUserOwnersV2

## Overview

### Available Operations

* [CreateEntitlementOwner](#createentitlementowner) - Create Entitlement Owner
* [CreateUserOwner](#createuserowner) - Create User Owner
* [DeleteEntitlementOwner](#deleteentitlementowner) - Delete Entitlement Owner
* [DeleteUserOwner](#deleteuserowner) - Delete User Owner
* [SearchEntitlementOwners](#searchentitlementowners) - Search Entitlement Owners
* [SearchUserOwners](#searchuserowners) - Search User Owners
* [Set](#set) - Set

## CreateEntitlementOwner

CreateEntitlementOwner creates an entitlement ownership source for an app user.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppUserOwnersV2.CreateEntitlementOwner" method="post" path="/api/v2/apps/{app_id}/users/{user_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.AppUserOwnersV2.CreateEntitlementOwner(ctx, operations.C1APIAppV2AppUserOwnersV2CreateEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        RoleSlug: "<value>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppUserEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIAppV2AppUserOwnersV2CreateEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2appuserownersv2createentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIAppV2AppUserOwnersV2CreateEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2appuserownersv2createentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateUserOwner

CreateUserOwner creates a user ownership source for an app user.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppUserOwnersV2.CreateUserOwner" method="post" path="/api/v2/apps/{app_id}/users/{user_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.AppUserOwnersV2.CreateUserOwner(ctx, operations.C1APIAppV2AppUserOwnersV2CreateUserOwnerRequest{
        AppID: "<id>",
        RoleSlug: "<value>",
        UserID: "<id>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppUserUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIAppV2AppUserOwnersV2CreateUserOwnerRequest](../../pkg/models/operations/c1apiappv2appuserownersv2createuserownerrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APIAppV2AppUserOwnersV2CreateUserOwnerResponse](../../pkg/models/operations/c1apiappv2appuserownersv2createuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteEntitlementOwner

DeleteEntitlementOwner deletes an entitlement ownership source for an app user.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppUserOwnersV2.DeleteEntitlementOwner" method="delete" path="/api/v2/apps/{app_id}/users/{user_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.AppUserOwnersV2.DeleteEntitlementOwner(ctx, operations.C1APIAppV2AppUserOwnersV2DeleteEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        RoleSlug: "<value>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppUserEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIAppV2AppUserOwnersV2DeleteEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2appuserownersv2deleteentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIAppV2AppUserOwnersV2DeleteEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2appuserownersv2deleteentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteUserOwner

DeleteUserOwner deletes a user ownership source for an app user.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppUserOwnersV2.DeleteUserOwner" method="delete" path="/api/v2/apps/{app_id}/users/{user_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.AppUserOwnersV2.DeleteUserOwner(ctx, operations.C1APIAppV2AppUserOwnersV2DeleteUserOwnerRequest{
        AppID: "<id>",
        RoleSlug: "<value>",
        UserID: "<id>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppUserUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIAppV2AppUserOwnersV2DeleteUserOwnerRequest](../../pkg/models/operations/c1apiappv2appuserownersv2deleteuserownerrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APIAppV2AppUserOwnersV2DeleteUserOwnerResponse](../../pkg/models/operations/c1apiappv2appuserownersv2deleteuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchEntitlementOwners

SearchEntitlementOwners searches for entitlement ownership sources of this app user.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppUserOwnersV2.SearchEntitlementOwners" method="get" path="/api/v2/apps/{app_id}/users/{user_id}/owners/entitlements" -->
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

    res, err := s.AppUserOwnersV2.SearchEntitlementOwners(ctx, operations.C1APIAppV2AppUserOwnersV2SearchEntitlementOwnersRequest{
        AppID: "<id>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppUserEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIAppV2AppUserOwnersV2SearchEntitlementOwnersRequest](../../pkg/models/operations/c1apiappv2appuserownersv2searchentitlementownersrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIAppV2AppUserOwnersV2SearchEntitlementOwnersResponse](../../pkg/models/operations/c1apiappv2appuserownersv2searchentitlementownersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchUserOwners

SearchUserOwners searches for user ownership sources of this app user.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppUserOwnersV2.SearchUserOwners" method="get" path="/api/v2/apps/{app_id}/users/{user_id}/owners/users" -->
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

    res, err := s.AppUserOwnersV2.SearchUserOwners(ctx, operations.C1APIAppV2AppUserOwnersV2SearchUserOwnersRequest{
        AppID: "<id>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppUserUserOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAppV2AppUserOwnersV2SearchUserOwnersRequest](../../pkg/models/operations/c1apiappv2appuserownersv2searchuserownersrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIAppV2AppUserOwnersV2SearchUserOwnersResponse](../../pkg/models/operations/c1apiappv2appuserownersv2searchuserownersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Set

Set replaces all owners for a given app user and role.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.AppUserOwnersV2.Set" method="put" path="/api/v2/apps/{app_id}/users/{user_id}/owners" -->
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

    res, err := s.AppUserOwnersV2.Set(ctx, operations.C1APIAppV2AppUserOwnersV2SetRequest{
        AppID: "<id>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SetAppUserOwnersV2Response != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.C1APIAppV2AppUserOwnersV2SetRequest](../../pkg/models/operations/c1apiappv2appuserownersv2setrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIAppV2AppUserOwnersV2SetResponse](../../pkg/models/operations/c1apiappv2appuserownersv2setresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |