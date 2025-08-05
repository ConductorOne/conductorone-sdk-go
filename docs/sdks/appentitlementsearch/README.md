# AppEntitlementSearch
(*AppEntitlementSearch*)

## Overview

### Available Operations

* [Search](#search) - Search
* [SearchAppEntitlementsForAppUser](#searchappentitlementsforappuser) - Search App Entitlements For App User
* [SearchAppEntitlementsWithExpired](#searchappentitlementswithexpired) - Search App Entitlements With Expired
* [SearchGrants](#searchgrants) - Search Grants

## Search

Search app entitlements based on filters specified in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.Search" method="post" path="/api/v1/search/entitlements" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

Invokes the c1.api.app.v1.AppEntitlementSearchService.SearchAppEntitlementsForAppUser method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.SearchAppEntitlementsForAppUser" method="get" path="/api/v1/search/apps/{app_id}/entitlements/users/{app_user_id}" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.SearchAppEntitlementsWithExpired" method="get" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/grants" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

Invokes the c1.api.app.v1.AppEntitlementSearchService.SearchGrants method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.SearchGrants" method="post" path="/api/v1/search/grants" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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