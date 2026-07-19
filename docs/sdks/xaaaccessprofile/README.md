# XAAAccessProfile

## Overview

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [GetByAppEntitlementID](#getbyappentitlementid) - Get By App Entitlement Id
* [List](#list) - List
* [ListHistory](#listhistory) - List History
* [Search](#search) - Search
* [Update](#update) - Update

## Create

Create an access profile under a resource server. The backend also
 provisions a backing AppEntitlement that users request to be granted the
 profile's scopes.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAAccessProfileService.Create" method="post" path="/api/v1/apps/{app_id}/xaa/access_profiles" -->
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

    res, err := s.XAAAccessProfile.Create(ctx, operations.C1APICrossAppAccessV1XAAAccessProfileServiceCreateRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAAccessProfileServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.C1APICrossAppAccessV1XAAAccessProfileServiceCreateRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicecreaterequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |
| `opts`                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |

### Response

**[*operations.C1APICrossAppAccessV1XAAAccessProfileServiceCreateResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete an access profile (soft delete). Cascades to its scope bindings and
 backing entitlement; outstanding grants end with the entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAAccessProfileService.Delete" method="delete" path="/api/v1/apps/{app_id}/xaa/access_profiles/{id}" -->
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

    res, err := s.XAAAccessProfile.Delete(ctx, operations.C1APICrossAppAccessV1XAAAccessProfileServiceDeleteRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAAccessProfileServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.C1APICrossAppAccessV1XAAAccessProfileServiceDeleteRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicedeleterequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |
| `opts`                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |

### Response

**[*operations.C1APICrossAppAccessV1XAAAccessProfileServiceDeleteResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get an access profile by app_id + id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAAccessProfileService.Get" method="get" path="/api/v1/apps/{app_id}/xaa/access_profiles/{id}" -->
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

    res, err := s.XAAAccessProfile.Get(ctx, operations.C1APICrossAppAccessV1XAAAccessProfileServiceGetRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAAccessProfileServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APICrossAppAccessV1XAAAccessProfileServiceGetRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicegetrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APICrossAppAccessV1XAAAccessProfileServiceGetResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetByAppEntitlementID

GetByAppEntitlementId looks up the access profile linked to an
 entitlement, by app_id + app_entitlement_id. Used by the request catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAAccessProfileService.GetByAppEntitlementId" method="get" path="/api/v1/apps/{app_id}/xaa/access_profiles/by_app_entitlement_id/{app_entitlement_id}" -->
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

    res, err := s.XAAAccessProfile.GetByAppEntitlementID(ctx, operations.C1APICrossAppAccessV1XAAAccessProfileServiceGetByAppEntitlementIDRequest{
        AppEntitlementID: "<id>",
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAAccessProfileServiceGetByAppEntitlementIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                      | Type                                                                                                                                                                                           | Required                                                                                                                                                                                       | Description                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                             | The context to use for the request.                                                                                                                                                            |
| `request`                                                                                                                                                                                      | [operations.C1APICrossAppAccessV1XAAAccessProfileServiceGetByAppEntitlementIDRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicegetbyappentitlementidrequest.md) | :heavy_check_mark:                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                     |
| `opts`                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                             | The options for this request.                                                                                                                                                                  |

### Response

**[*operations.C1APICrossAppAccessV1XAAAccessProfileServiceGetByAppEntitlementIDResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicegetbyappentitlementidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List the access profiles defined for an application, one page at a time.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAAccessProfileService.List" method="get" path="/api/v1/apps/{app_id}/xaa/access_profiles" -->
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

    res, err := s.XAAAccessProfile.List(ctx, operations.C1APICrossAppAccessV1XAAAccessProfileServiceListRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAAccessProfileServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APICrossAppAccessV1XAAAccessProfileServiceListRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicelistrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APICrossAppAccessV1XAAAccessProfileServiceListResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHistory

ListHistory returns the change history (newest first) for a single access
 profile — each entry is a snapshot plus who/when metadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAAccessProfileService.ListHistory" method="get" path="/api/v1/apps/{app_id}/xaa/access_profiles/{id}/history" -->
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

    res, err := s.XAAAccessProfile.ListHistory(ctx, operations.C1APICrossAppAccessV1XAAAccessProfileServiceListHistoryRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAAccessProfileServiceListHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.C1APICrossAppAccessV1XAAAccessProfileServiceListHistoryRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicelisthistoryrequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |
| `opts`                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                         | The options for this request.                                                                                                                                              |

### Response

**[*operations.C1APICrossAppAccessV1XAAAccessProfileServiceListHistoryResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicelisthistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Search

Search access profiles across the tenant, filtered by application,
 resource server, or text query, or fetch a specific set by ref.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAAccessProfileService.Search" method="post" path="/api/v1/search/xaa/access_profiles" -->
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

    res, err := s.XAAAccessProfile.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAAccessProfileServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [shared.XAAAccessProfileServiceSearchRequest](../../pkg/models/shared/xaaaccessprofileservicesearchrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.C1APICrossAppAccessV1XAAAccessProfileServiceSearchResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update an access profile's editable fields via update_mask. Editable
 paths: display_name, description. The profile must include id and app_id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAAccessProfileService.Update" method="post" path="/api/v1/apps/{app_id}/xaa/access_profiles/{id}" -->
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

    res, err := s.XAAAccessProfile.Update(ctx, operations.C1APICrossAppAccessV1XAAAccessProfileServiceUpdateRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAAccessProfileServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.C1APICrossAppAccessV1XAAAccessProfileServiceUpdateRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |
| `opts`                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |

### Response

**[*operations.C1APICrossAppAccessV1XAAAccessProfileServiceUpdateResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofileserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |