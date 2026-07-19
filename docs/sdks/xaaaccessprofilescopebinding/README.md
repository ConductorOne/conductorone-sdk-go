# XAAAccessProfileScopeBinding

## Overview

### Available Operations

* [CreateBindings](#createbindings) - Create Bindings
* [DeleteBindings](#deletebindings) - Delete Bindings
* [List](#list) - List
* [Search](#search) - Search

## CreateBindings

CreateBindings binds one or more scopes (xaa_scope_ids) to an access
 profile. Every scope must belong to the profile's resource server.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAAccessProfileScopeBindingService.CreateBindings" method="post" path="/api/v1/apps/{app_id}/xaa/access_profiles/{access_profile_id}/scope_bindings" -->
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

    res, err := s.XAAAccessProfileScopeBinding.CreateBindings(ctx, operations.C1APICrossAppAccessV1XAAAccessProfileScopeBindingServiceCreateBindingsRequest{
        AccessProfileID: "<id>",
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAAccessProfileScopeBindingServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                | Type                                                                                                                                                                                                     | Required                                                                                                                                                                                                 | Description                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                      |
| `request`                                                                                                                                                                                                | [operations.C1APICrossAppAccessV1XAAAccessProfileScopeBindingServiceCreateBindingsRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofilescopebindingservicecreatebindingsrequest.md) | :heavy_check_mark:                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                               |
| `opts`                                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                       | The options for this request.                                                                                                                                                                            |

### Response

**[*operations.C1APICrossAppAccessV1XAAAccessProfileScopeBindingServiceCreateBindingsResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofilescopebindingservicecreatebindingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteBindings

DeleteBindings unbinds one or more scopes (xaa_scope_ids) from an access
 profile. Uses a POST .../delete action route because the scope IDs travel
 in the request body, which HTTP DELETE does not reliably support.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAAccessProfileScopeBindingService.DeleteBindings" method="post" path="/api/v1/apps/{app_id}/xaa/access_profiles/{access_profile_id}/scope_bindings/delete" -->
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

    res, err := s.XAAAccessProfileScopeBinding.DeleteBindings(ctx, operations.C1APICrossAppAccessV1XAAAccessProfileScopeBindingServiceDeleteBindingsRequest{
        AccessProfileID: "<id>",
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAAccessProfileScopeBindingServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                | Type                                                                                                                                                                                                     | Required                                                                                                                                                                                                 | Description                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                      |
| `request`                                                                                                                                                                                                | [operations.C1APICrossAppAccessV1XAAAccessProfileScopeBindingServiceDeleteBindingsRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofilescopebindingservicedeletebindingsrequest.md) | :heavy_check_mark:                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                               |
| `opts`                                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                       | The options for this request.                                                                                                                                                                            |

### Response

**[*operations.C1APICrossAppAccessV1XAAAccessProfileScopeBindingServiceDeleteBindingsResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofilescopebindingservicedeletebindingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List the scopes bound to an access profile, one page at a time.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAAccessProfileScopeBindingService.List" method="get" path="/api/v1/apps/{app_id}/xaa/access_profiles/{access_profile_id}/scope_bindings" -->
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

    res, err := s.XAAAccessProfileScopeBinding.List(ctx, operations.C1APICrossAppAccessV1XAAAccessProfileScopeBindingServiceListRequest{
        AccessProfileID: "<id>",
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAAccessProfileScopeBindingServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                            | Type                                                                                                                                                                                 | Required                                                                                                                                                                             | Description                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                | :heavy_check_mark:                                                                                                                                                                   | The context to use for the request.                                                                                                                                                  |
| `request`                                                                                                                                                                            | [operations.C1APICrossAppAccessV1XAAAccessProfileScopeBindingServiceListRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofilescopebindingservicelistrequest.md) | :heavy_check_mark:                                                                                                                                                                   | The request object to use for the request.                                                                                                                                           |
| `opts`                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                   | The options for this request.                                                                                                                                                        |

### Response

**[*operations.C1APICrossAppAccessV1XAAAccessProfileScopeBindingServiceListResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofilescopebindingservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Search

Search scope bindings, filtered by access profile or by scope, or fetch a
 specific set by ref. The by-scope direction answers "which profiles
 contain this scope" for impact analysis.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAAccessProfileScopeBindingService.Search" method="post" path="/api/v1/search/xaa/access_profile_scope_bindings" -->
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

    res, err := s.XAAAccessProfileScopeBinding.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAAccessProfileScopeBindingServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [shared.XAAAccessProfileScopeBindingServiceSearchRequest](../../pkg/models/shared/xaaaccessprofilescopebindingservicesearchrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.C1APICrossAppAccessV1XAAAccessProfileScopeBindingServiceSearchResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaaccessprofilescopebindingservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |