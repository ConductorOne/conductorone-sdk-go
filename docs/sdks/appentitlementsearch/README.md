# AppEntitlementSearch

## Overview

### Available Operations

* [CountGrantsForUserByApp](#countgrantsforuserbyapp) - Count Grants For User By App
* [Search](#search) - Search
* [SearchAppEntitlementsForAppUser](#searchappentitlementsforappuser) - Search App Entitlements For App User
* [SearchAppEntitlementsWithExpired](#searchappentitlementswithexpired) - Search App Entitlements With Expired
* [SearchGrants](#searchgrants) - Search Grants
* [SearchGraph](#searchgraph) - Search Graph
* [SearchReachableResourcesForUser](#searchreachableresourcesforuser) - Search Reachable Resources For User

## CountGrantsForUserByApp

CountGrantsForUserByApp returns, for a user, the number of grants held per
 application. Use it to size or filter an access graph before calling
 SearchGraph: counts are computed directly from grant bindings and are an
 upper bound on what SearchGraph will render for the same user and app
 filter (SearchGraph applies additional filters that this count does not).

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.CountGrantsForUserByApp" method="post" path="/api/v1/search/graph/counts" -->
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

    res, err := s.AppEntitlementSearch.CountGrantsForUserByApp(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementSearchServiceCountGrantsForUserByAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [shared.AppEntitlementSearchServiceCountGrantsForUserByAppRequest](../../pkg/models/shared/appentitlementsearchservicecountgrantsforuserbyapprequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceCountGrantsForUserByAppResponse](../../pkg/models/operations/c1apiappv1appentitlementsearchservicecountgrantsforuserbyappresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Search

Search app entitlements based on filters specified in the request body.
 Each AppEntitlementView row is large — request a small page_size (≤10) and set expand_mask to only the paths you need; broad expansion inlines whole related objects and bloats the response.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.Search" method="post" path="/api/v1/search/entitlements" -->
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

    res, err := s.AppEntitlementSearch.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementSearchServiceSearchResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [shared.AppEntitlementSearchServiceSearchRequest](../../pkg/models/shared/appentitlementsearchservicesearchrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchResponse](../../pkg/models/operations/c1apiappv1appentitlementsearchservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchAppEntitlementsForAppUser

Search for app entitlements associated with a specific app user, with optional resource type trait filtering.
 Each AppEntitlementView row is large — request a small page_size (≤10) and set expand_mask to only the paths you need; broad expansion inlines whole related objects and bloats the response.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.SearchAppEntitlementsForAppUser" method="get" path="/api/v1/search/apps/{app_id}/entitlements/users/{app_user_id}" -->
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

    res, err := s.AppEntitlementSearch.SearchAppEntitlementsForAppUser(ctx, operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsForAppUserRequest{
        AppID: "<id>",
        AppUserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                            | Type                                                                                                                                                                                                 | Required                                                                                                                                                                                             | Description                                                                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                  |
| `request`                                                                                                                                                                                            | [operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsForAppUserRequest](../../pkg/models/operations/c1apiappv1appentitlementsearchservicesearchappentitlementsforappuserrequest.md) | :heavy_check_mark:                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                           |
| `opts`                                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                   | The options for this request.                                                                                                                                                                        |

### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsForAppUserResponse](../../pkg/models/operations/c1apiappv1appentitlementsearchservicesearchappentitlementsforappuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchAppEntitlementsWithExpired

Search app entitlements, include app users, users, expires, discovered.
 Response rows are large — request a small page_size (≤10) to keep responses small.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.SearchAppEntitlementsWithExpired" method="get" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/grants" -->
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

    res, err := s.AppEntitlementSearch.SearchAppEntitlementsWithExpired(ctx, operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredRequest{
        AppEntitlementID: "<id>",
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppEntitlementsWithExpiredResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                              | Type                                                                                                                                                                                                   | Required                                                                                                                                                                                               | Description                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                    |
| `request`                                                                                                                                                                                              | [operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredRequest](../../pkg/models/operations/c1apiappv1appentitlementsearchservicesearchappentitlementswithexpiredrequest.md) | :heavy_check_mark:                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                             |
| `opts`                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                     | The options for this request.                                                                                                                                                                          |

### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredResponse](../../pkg/models/operations/c1apiappv1appentitlementsearchservicesearchappentitlementswithexpiredresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchGrants

Search grants (user-to-entitlement bindings) across apps, with filters for app, user, resource type, and entitlement.
 Rows are heavy (each binding carries a full entitlement + user view) — request a small page_size (≤10) and set expand_mask to only the paths you need; broad expansion inlines whole related objects and bloats the response.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.SearchGrants" method="post" path="/api/v1/search/grants" -->
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

    res, err := s.AppEntitlementSearch.SearchGrants(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementSearchServiceSearchGrantsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [shared.AppEntitlementSearchServiceSearchGrantsRequest](../../pkg/models/shared/appentitlementsearchservicesearchgrantsrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchGrantsResponse](../../pkg/models/operations/c1apiappv1appentitlementsearchservicesearchgrantsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchGraph

SearchGraph performs a server-side BFS traversal and returns a bounded, filtered subgraph.
 Exactly one of user_id, app_id, or resource_id must be set.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.SearchGraph" method="post" path="/api/v1/search/graph" -->
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

    res, err := s.AppEntitlementSearch.SearchGraph(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementSearchServiceSearchGraphResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [shared.AppEntitlementSearchServiceSearchGraphRequest](../../pkg/models/shared/appentitlementsearchservicesearchgraphrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchGraphResponse](../../pkg/models/operations/c1apiappv1appentitlementsearchservicesearchgraphresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchReachableResourcesForUser

SearchReachableResourcesForUser returns the distinct app resources a user
 can reach through any of their grants, deduplicated across entitlements
 (a resource reachable via more than one grant appears once). Powers the
 Resources lane of the access graph's list view: supports free-text search
 over resource display name and narrowing to specific applications.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.SearchReachableResourcesForUser" method="post" path="/api/v1/search/graph/resources" -->
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

    res, err := s.AppEntitlementSearch.SearchReachableResourcesForUser(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementSearchServiceSearchReachableResourcesForUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [shared.AppEntitlementSearchServiceSearchReachableResourcesForUserRequest](../../pkg/models/shared/appentitlementsearchservicesearchreachableresourcesforuserrequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |
| `opts`                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                                       | The options for this request.                                                                                                                                            |

### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchReachableResourcesForUserResponse](../../pkg/models/operations/c1apiappv1appentitlementsearchservicesearchreachableresourcesforuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |