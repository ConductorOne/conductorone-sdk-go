# ConnectorOwnersV2

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

CreateEntitlementOwner creates an entitlement ownership source for a connector.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.ConnectorOwners.CreateEntitlementOwner" method="post" path="/api/v2/apps/{app_id}/connectors/{connector_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.ConnectorOwnersV2.CreateEntitlementOwner(ctx, operations.C1APIAppV2ConnectorOwnersCreateEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        ConnectorID: "<id>",
        RoleSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateConnectorEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIAppV2ConnectorOwnersCreateEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2connectorownerscreateentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIAppV2ConnectorOwnersCreateEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2connectorownerscreateentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateUserOwner

CreateUserOwner creates a user ownership source for a connector.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.ConnectorOwners.CreateUserOwner" method="post" path="/api/v2/apps/{app_id}/connectors/{connector_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.ConnectorOwnersV2.CreateUserOwner(ctx, operations.C1APIAppV2ConnectorOwnersCreateUserOwnerRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
        RoleSlug: "<value>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateConnectorUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIAppV2ConnectorOwnersCreateUserOwnerRequest](../../pkg/models/operations/c1apiappv2connectorownerscreateuserownerrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APIAppV2ConnectorOwnersCreateUserOwnerResponse](../../pkg/models/operations/c1apiappv2connectorownerscreateuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteEntitlementOwner

DeleteEntitlementOwner deletes an entitlement ownership source for a connector.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.ConnectorOwners.DeleteEntitlementOwner" method="delete" path="/api/v2/apps/{app_id}/connectors/{connector_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.ConnectorOwnersV2.DeleteEntitlementOwner(ctx, operations.C1APIAppV2ConnectorOwnersDeleteEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        ConnectorID: "<id>",
        RoleSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteConnectorEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIAppV2ConnectorOwnersDeleteEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2connectorownersdeleteentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIAppV2ConnectorOwnersDeleteEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2connectorownersdeleteentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteUserOwner

DeleteUserOwner deletes a user ownership source for a connector.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.ConnectorOwners.DeleteUserOwner" method="delete" path="/api/v2/apps/{app_id}/connectors/{connector_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.ConnectorOwnersV2.DeleteUserOwner(ctx, operations.C1APIAppV2ConnectorOwnersDeleteUserOwnerRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
        RoleSlug: "<value>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteConnectorUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIAppV2ConnectorOwnersDeleteUserOwnerRequest](../../pkg/models/operations/c1apiappv2connectorownersdeleteuserownerrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APIAppV2ConnectorOwnersDeleteUserOwnerResponse](../../pkg/models/operations/c1apiappv2connectorownersdeleteuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetEntitlementOwner

GetEntitlementOwner gets an entitlement ownership source for a connector.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.ConnectorOwners.GetEntitlementOwner" method="get" path="/api/v2/apps/{app_id}/connectors/{connector_id}/owners/entitlements/{role_slug}/{app_entitlement_ref_app_id}/{app_entitlement_ref_id}" -->
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

    res, err := s.ConnectorOwnersV2.GetEntitlementOwner(ctx, operations.C1APIAppV2ConnectorOwnersGetEntitlementOwnerRequest{
        AppEntitlementRefAppID: "<id>",
        AppEntitlementRefID: "<id>",
        AppID: "<id>",
        ConnectorID: "<id>",
        RoleSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetConnectorEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.C1APIAppV2ConnectorOwnersGetEntitlementOwnerRequest](../../pkg/models/operations/c1apiappv2connectorownersgetentitlementownerrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.C1APIAppV2ConnectorOwnersGetEntitlementOwnerResponse](../../pkg/models/operations/c1apiappv2connectorownersgetentitlementownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUserOwner

GetUserOwner gets a user ownership source for a connector.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.ConnectorOwners.GetUserOwner" method="get" path="/api/v2/apps/{app_id}/connectors/{connector_id}/owners/users/{role_slug}/{user_ref_id}" -->
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

    res, err := s.ConnectorOwnersV2.GetUserOwner(ctx, operations.C1APIAppV2ConnectorOwnersGetUserOwnerRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
        RoleSlug: "<value>",
        UserRefID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetConnectorUserOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.C1APIAppV2ConnectorOwnersGetUserOwnerRequest](../../pkg/models/operations/c1apiappv2connectorownersgetuserownerrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.C1APIAppV2ConnectorOwnersGetUserOwnerResponse](../../pkg/models/operations/c1apiappv2connectorownersgetuserownerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchEntitlementOwners

SearchEntitlementOwners searches for entitlement ownership sources for a connector.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.ConnectorOwners.SearchEntitlementOwners" method="get" path="/api/v2/apps/{app_id}/connectors/{connector_id}/owners/entitlements" -->
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

    res, err := s.ConnectorOwnersV2.SearchEntitlementOwners(ctx, operations.C1APIAppV2ConnectorOwnersSearchEntitlementOwnersRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchConnectorEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIAppV2ConnectorOwnersSearchEntitlementOwnersRequest](../../pkg/models/operations/c1apiappv2connectorownerssearchentitlementownersrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIAppV2ConnectorOwnersSearchEntitlementOwnersResponse](../../pkg/models/operations/c1apiappv2connectorownerssearchentitlementownersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchUserOwners

SearchUserOwners searches for user ownership sources for a connector.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.ConnectorOwners.SearchUserOwners" method="get" path="/api/v2/apps/{app_id}/connectors/{connector_id}/owners/users" -->
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

    res, err := s.ConnectorOwnersV2.SearchUserOwners(ctx, operations.C1APIAppV2ConnectorOwnersSearchUserOwnersRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchConnectorUserOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAppV2ConnectorOwnersSearchUserOwnersRequest](../../pkg/models/operations/c1apiappv2connectorownerssearchuserownersrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIAppV2ConnectorOwnersSearchUserOwnersResponse](../../pkg/models/operations/c1apiappv2connectorownerssearchuserownersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Set

Set replaces all owners for a given connector and role.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.ConnectorOwners.Set" method="put" path="/api/v2/apps/{app_id}/connectors/{connector_id}/owners" -->
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

    res, err := s.ConnectorOwnersV2.Set(ctx, operations.C1APIAppV2ConnectorOwnersSetRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SetConnectorOwnersV2Response != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.C1APIAppV2ConnectorOwnersSetRequest](../../pkg/models/operations/c1apiappv2connectorownerssetrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIAppV2ConnectorOwnersSetResponse](../../pkg/models/operations/c1apiappv2connectorownerssetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |