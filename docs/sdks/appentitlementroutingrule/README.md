# AppEntitlementRoutingRule

## Overview

### Available Operations

* [CreateAppEntitlementRoutingRule](#createappentitlementroutingrule) - Create App Entitlement Routing Rule
* [DeleteAppEntitlementRoutingRule](#deleteappentitlementroutingrule) - Delete App Entitlement Routing Rule
* [GetAppEntitlementRoutingRule](#getappentitlementroutingrule) - Get App Entitlement Routing Rule
* [ListAppEntitlementRoutingRules](#listappentitlementroutingrules) - List App Entitlement Routing Rules
* [ReorderAppEntitlementRoutingRules](#reorderappentitlementroutingrules) - Reorder App Entitlement Routing Rules
* [UpdateAppEntitlementRoutingRule](#updateappentitlementroutingrule) - Update App Entitlement Routing Rule

## CreateAppEntitlementRoutingRule

Invokes the c1.api.app.v1.AppEntitlementRoutingRuleService.CreateAppEntitlementRoutingRule method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementRoutingRuleService.CreateAppEntitlementRoutingRule" method="post" path="/api/v1/apps/{app_id}/entitlement_configuration_rules" -->
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

    res, err := s.AppEntitlementRoutingRule.CreateAppEntitlementRoutingRule(ctx, operations.C1APIAppV1AppEntitlementRoutingRuleServiceCreateAppEntitlementRoutingRuleRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppEntitlementRoutingRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                      | Type                                                                                                                                                                                                           | Required                                                                                                                                                                                                       | Description                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                            |
| `request`                                                                                                                                                                                                      | [operations.C1APIAppV1AppEntitlementRoutingRuleServiceCreateAppEntitlementRoutingRuleRequest](../../pkg/models/operations/c1apiappv1appentitlementroutingruleservicecreateappentitlementroutingrulerequest.md) | :heavy_check_mark:                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                             | The options for this request.                                                                                                                                                                                  |

### Response

**[*operations.C1APIAppV1AppEntitlementRoutingRuleServiceCreateAppEntitlementRoutingRuleResponse](../../pkg/models/operations/c1apiappv1appentitlementroutingruleservicecreateappentitlementroutingruleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteAppEntitlementRoutingRule

Invokes the c1.api.app.v1.AppEntitlementRoutingRuleService.DeleteAppEntitlementRoutingRule method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementRoutingRuleService.DeleteAppEntitlementRoutingRule" method="delete" path="/api/v1/apps/{app_id}/entitlement_configuration_rules/{id}" -->
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

    res, err := s.AppEntitlementRoutingRule.DeleteAppEntitlementRoutingRule(ctx, operations.C1APIAppV1AppEntitlementRoutingRuleServiceDeleteAppEntitlementRoutingRuleRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppEntitlementRoutingRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                      | Type                                                                                                                                                                                                           | Required                                                                                                                                                                                                       | Description                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                            |
| `request`                                                                                                                                                                                                      | [operations.C1APIAppV1AppEntitlementRoutingRuleServiceDeleteAppEntitlementRoutingRuleRequest](../../pkg/models/operations/c1apiappv1appentitlementroutingruleservicedeleteappentitlementroutingrulerequest.md) | :heavy_check_mark:                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                             | The options for this request.                                                                                                                                                                                  |

### Response

**[*operations.C1APIAppV1AppEntitlementRoutingRuleServiceDeleteAppEntitlementRoutingRuleResponse](../../pkg/models/operations/c1apiappv1appentitlementroutingruleservicedeleteappentitlementroutingruleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAppEntitlementRoutingRule

Invokes the c1.api.app.v1.AppEntitlementRoutingRuleService.GetAppEntitlementRoutingRule method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementRoutingRuleService.GetAppEntitlementRoutingRule" method="get" path="/api/v1/apps/{app_id}/entitlement_configuration_rules/{id}" -->
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

    res, err := s.AppEntitlementRoutingRule.GetAppEntitlementRoutingRule(ctx, operations.C1APIAppV1AppEntitlementRoutingRuleServiceGetAppEntitlementRoutingRuleRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAppEntitlementRoutingRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                | Type                                                                                                                                                                                                     | Required                                                                                                                                                                                                 | Description                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                      |
| `request`                                                                                                                                                                                                | [operations.C1APIAppV1AppEntitlementRoutingRuleServiceGetAppEntitlementRoutingRuleRequest](../../pkg/models/operations/c1apiappv1appentitlementroutingruleservicegetappentitlementroutingrulerequest.md) | :heavy_check_mark:                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                               |
| `opts`                                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                       | The options for this request.                                                                                                                                                                            |

### Response

**[*operations.C1APIAppV1AppEntitlementRoutingRuleServiceGetAppEntitlementRoutingRuleResponse](../../pkg/models/operations/c1apiappv1appentitlementroutingruleservicegetappentitlementroutingruleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAppEntitlementRoutingRules

Invokes the c1.api.app.v1.AppEntitlementRoutingRuleService.ListAppEntitlementRoutingRules method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementRoutingRuleService.ListAppEntitlementRoutingRules" method="get" path="/api/v1/apps/{app_id}/entitlement_configuration_rules" -->
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

    res, err := s.AppEntitlementRoutingRule.ListAppEntitlementRoutingRules(ctx, operations.C1APIAppV1AppEntitlementRoutingRuleServiceListAppEntitlementRoutingRulesRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppEntitlementRoutingRulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                    | Type                                                                                                                                                                                                         | Required                                                                                                                                                                                                     | Description                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                          |
| `request`                                                                                                                                                                                                    | [operations.C1APIAppV1AppEntitlementRoutingRuleServiceListAppEntitlementRoutingRulesRequest](../../pkg/models/operations/c1apiappv1appentitlementroutingruleservicelistappentitlementroutingrulesrequest.md) | :heavy_check_mark:                                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                                   |
| `opts`                                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                |

### Response

**[*operations.C1APIAppV1AppEntitlementRoutingRuleServiceListAppEntitlementRoutingRulesResponse](../../pkg/models/operations/c1apiappv1appentitlementroutingruleservicelistappentitlementroutingrulesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ReorderAppEntitlementRoutingRules

Reorder all configuration rules for an app in a single call. The caller
 supplies the full ordered list of rule IDs; the server assigns dense
 priorities (1..N) in that order. The request must be a permutation of every
 active rule in the app — missing or extra IDs are rejected.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementRoutingRuleService.ReorderAppEntitlementRoutingRules" method="post" path="/api/v1/apps/{app_id}/entitlement_configuration_rules/reorder" -->
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

    res, err := s.AppEntitlementRoutingRule.ReorderAppEntitlementRoutingRules(ctx, operations.C1APIAppV1AppEntitlementRoutingRuleServiceReorderAppEntitlementRoutingRulesRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ReorderAppEntitlementRoutingRulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                          | Type                                                                                                                                                                                                               | Required                                                                                                                                                                                                           | Description                                                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                |
| `request`                                                                                                                                                                                                          | [operations.C1APIAppV1AppEntitlementRoutingRuleServiceReorderAppEntitlementRoutingRulesRequest](../../pkg/models/operations/c1apiappv1appentitlementroutingruleservicereorderappentitlementroutingrulesrequest.md) | :heavy_check_mark:                                                                                                                                                                                                 | The request object to use for the request.                                                                                                                                                                         |
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |

### Response

**[*operations.C1APIAppV1AppEntitlementRoutingRuleServiceReorderAppEntitlementRoutingRulesResponse](../../pkg/models/operations/c1apiappv1appentitlementroutingruleservicereorderappentitlementroutingrulesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAppEntitlementRoutingRule

Update an existing app entitlement configuration rule. The app_id field is
 immutable; moving a rule between apps is modeled as delete + create.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementRoutingRuleService.UpdateAppEntitlementRoutingRule" method="post" path="/api/v1/apps/{app_id}/entitlement_configuration_rules/{id}" -->
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

    res, err := s.AppEntitlementRoutingRule.UpdateAppEntitlementRoutingRule(ctx, operations.C1APIAppV1AppEntitlementRoutingRuleServiceUpdateAppEntitlementRoutingRuleRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateAppEntitlementRoutingRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                      | Type                                                                                                                                                                                                           | Required                                                                                                                                                                                                       | Description                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                            |
| `request`                                                                                                                                                                                                      | [operations.C1APIAppV1AppEntitlementRoutingRuleServiceUpdateAppEntitlementRoutingRuleRequest](../../pkg/models/operations/c1apiappv1appentitlementroutingruleserviceupdateappentitlementroutingrulerequest.md) | :heavy_check_mark:                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                             | The options for this request.                                                                                                                                                                                  |

### Response

**[*operations.C1APIAppV1AppEntitlementRoutingRuleServiceUpdateAppEntitlementRoutingRuleResponse](../../pkg/models/operations/c1apiappv1appentitlementroutingruleserviceupdateappentitlementroutingruleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |