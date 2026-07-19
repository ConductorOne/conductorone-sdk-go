# XAAResourceServer

## Overview

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List
* [ListHistory](#listhistory) - List History
* [Search](#search) - Search
* [Update](#update) - Update

## Create

Register a resource server (a third-party authorization server) as a
 permitted cross-app-access audience for an application. The audience must
 be unique within the application and must not equal your own tenant's
 issuer — C1 cannot be both the granting IdP and the resource server in the
 same flow.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAResourceServerService.Create" method="post" path="/api/v1/apps/{app_id}/xaa/resource_servers" -->
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

    res, err := s.XAAResourceServer.Create(ctx, operations.C1APICrossAppAccessV1XAAResourceServerServiceCreateRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAResourceServerServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APICrossAppAccessV1XAAResourceServerServiceCreateRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverservicecreaterequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APICrossAppAccessV1XAAResourceServerServiceCreateResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a resource server (soft delete). Cascades to its scopes, access
 profiles, and client audience mappings.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAResourceServerService.Delete" method="delete" path="/api/v1/apps/{app_id}/xaa/resource_servers/{id}" -->
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

    res, err := s.XAAResourceServer.Delete(ctx, operations.C1APICrossAppAccessV1XAAResourceServerServiceDeleteRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAResourceServerServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APICrossAppAccessV1XAAResourceServerServiceDeleteRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverservicedeleterequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APICrossAppAccessV1XAAResourceServerServiceDeleteResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get a registered resource server by app_id + id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAResourceServerService.Get" method="get" path="/api/v1/apps/{app_id}/xaa/resource_servers/{id}" -->
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

    res, err := s.XAAResourceServer.Get(ctx, operations.C1APICrossAppAccessV1XAAResourceServerServiceGetRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAResourceServerServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APICrossAppAccessV1XAAResourceServerServiceGetRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverservicegetrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APICrossAppAccessV1XAAResourceServerServiceGetResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List the resource servers registered for an application, one page at a time.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAResourceServerService.List" method="get" path="/api/v1/apps/{app_id}/xaa/resource_servers" -->
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

    res, err := s.XAAResourceServer.List(ctx, operations.C1APICrossAppAccessV1XAAResourceServerServiceListRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAResourceServerServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.C1APICrossAppAccessV1XAAResourceServerServiceListRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverservicelistrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |
| `opts`                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                   | :heavy_minus_sign:                                                                                                                                             | The options for this request.                                                                                                                                  |

### Response

**[*operations.C1APICrossAppAccessV1XAAResourceServerServiceListResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHistory

ListHistory returns the change history (newest first) for a single
 resource server — each entry is a snapshot plus who/when metadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAResourceServerService.ListHistory" method="get" path="/api/v1/apps/{app_id}/xaa/resource_servers/{id}/history" -->
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

    res, err := s.XAAResourceServer.ListHistory(ctx, operations.C1APICrossAppAccessV1XAAResourceServerServiceListHistoryRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAResourceServerServiceListHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.C1APICrossAppAccessV1XAAResourceServerServiceListHistoryRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverservicelisthistoryrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |
| `opts`                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                           | The options for this request.                                                                                                                                                |

### Response

**[*operations.C1APICrossAppAccessV1XAAResourceServerServiceListHistoryResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverservicelisthistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Search

Search resource servers across the tenant, filtered by application,
 audience substring, signing algorithm, or disabled state, or fetch a
 specific set by ref. Returns one page at a time.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAResourceServerService.Search" method="post" path="/api/v1/search/xaa/resource_servers" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
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

    res, err := s.XAAResourceServer.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAResourceServerServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [shared.XAAResourceServerServiceSearchRequest](../../pkg/models/shared/xaaresourceserverservicesearchrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.C1APICrossAppAccessV1XAAResourceServerServiceSearchResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update a resource server's editable fields via update_mask. The audience
 is immutable (delete and recreate to change it); supplying a different
 audience is rejected. Editable paths: display_name, description,
 resource_uris, max_grant_lifetime, signing_algorithm,
 require_proof_of_possession, modify_claims_hook, disabled. sector_id is
 immutable once set (delete and recreate to change it).

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAResourceServerService.Update" method="post" path="/api/v1/apps/{app_id}/xaa/resource_servers/{id}" -->
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

    res, err := s.XAAResourceServer.Update(ctx, operations.C1APICrossAppAccessV1XAAResourceServerServiceUpdateRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAResourceServerServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APICrossAppAccessV1XAAResourceServerServiceUpdateRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APICrossAppAccessV1XAAResourceServerServiceUpdateResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaresourceserverserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |