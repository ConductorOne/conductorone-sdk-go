# AppUserMergeRules

## Overview

### Available Operations

* [Get](#get) - Get
* [ResetToDefaults](#resettodefaults) - Reset To Defaults
* [Update](#update) - Update

## Get

Get the account merge rules for an application.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppUserMergeRulesService.Get" method="get" path="/api/v1/apps/{app_id}/account_merge_rules" -->
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

    res, err := s.AppUserMergeRules.Get(ctx, operations.C1APIAppV1AppUserMergeRulesServiceGetRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AppUserMergeRulesServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.C1APIAppV1AppUserMergeRulesServiceGetRequest](../../pkg/models/operations/c1apiappv1appusermergerulesservicegetrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.C1APIAppV1AppUserMergeRulesServiceGetResponse](../../pkg/models/operations/c1apiappv1appusermergerulesservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ResetToDefaults

Replace the account merge rules for an application with the default rules: match on
 primary email, then any email, then username.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppUserMergeRulesService.ResetToDefaults" method="post" path="/api/v1/apps/{app_id}/account_merge_rules/reset" -->
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

    res, err := s.AppUserMergeRules.ResetToDefaults(ctx, operations.C1APIAppV1AppUserMergeRulesServiceResetToDefaultsRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AppUserMergeRulesServiceResetToDefaultsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.C1APIAppV1AppUserMergeRulesServiceResetToDefaultsRequest](../../pkg/models/operations/c1apiappv1appusermergerulesserviceresettodefaultsrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |
| `opts`                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                   | :heavy_minus_sign:                                                                                                                                             | The options for this request.                                                                                                                                  |

### Response

**[*operations.C1APIAppV1AppUserMergeRulesServiceResetToDefaultsResponse](../../pkg/models/operations/c1apiappv1appusermergerulesserviceresettodefaultsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Replace the account merge rules for an application. The request must carry the
 rule set version it was based on; if the rules changed since then, the update is rejected.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppUserMergeRulesService.Update" method="post" path="/api/v1/apps/{app_id}/account_merge_rules" -->
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

    res, err := s.AppUserMergeRules.Update(ctx, operations.C1APIAppV1AppUserMergeRulesServiceUpdateRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AppUserMergeRulesServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIAppV1AppUserMergeRulesServiceUpdateRequest](../../pkg/models/operations/c1apiappv1appusermergerulesserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APIAppV1AppUserMergeRulesServiceUpdateResponse](../../pkg/models/operations/c1apiappv1appusermergerulesserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |