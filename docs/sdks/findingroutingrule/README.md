# FindingRoutingRule

## Overview

### Available Operations

* [CreateFindingRoutingRule](#createfindingroutingrule) - Create Finding Routing Rule
* [DeleteFindingRoutingRule](#deletefindingroutingrule) - Delete Finding Routing Rule
* [GetFindingRoutingRule](#getfindingroutingrule) - Get Finding Routing Rule
* [ListFindingRoutingRules](#listfindingroutingrules) - List Finding Routing Rules
* [UpdateFindingRoutingRule](#updatefindingroutingrule) - Update Finding Routing Rule

## CreateFindingRoutingRule

Create a new finding routing rule that defines which policy to use for auto-routing matching findings.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingRoutingRuleService.CreateFindingRoutingRule" method="post" path="/api/v1/findings/routing-rules" -->
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

    res, err := s.FindingRoutingRule.CreateFindingRoutingRule(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateFindingRoutingRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [shared.CreateFindingRoutingRuleRequest](../../pkg/models/shared/createfindingroutingrulerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.C1APIFindingV1FindingRoutingRuleServiceCreateFindingRoutingRuleResponse](../../pkg/models/operations/c1apifindingv1findingroutingruleservicecreatefindingroutingruleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteFindingRoutingRule

Delete a finding routing rule. Findings already routed by this rule are not affected.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingRoutingRuleService.DeleteFindingRoutingRule" method="delete" path="/api/v1/findings/routing-rules/{id}" -->
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

    res, err := s.FindingRoutingRule.DeleteFindingRoutingRule(ctx, operations.C1APIFindingV1FindingRoutingRuleServiceDeleteFindingRoutingRuleRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteFindingRoutingRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                  | Type                                                                                                                                                                                       | Required                                                                                                                                                                                   | Description                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                         | The context to use for the request.                                                                                                                                                        |
| `request`                                                                                                                                                                                  | [operations.C1APIFindingV1FindingRoutingRuleServiceDeleteFindingRoutingRuleRequest](../../pkg/models/operations/c1apifindingv1findingroutingruleservicedeletefindingroutingrulerequest.md) | :heavy_check_mark:                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                 |
| `opts`                                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                         | The options for this request.                                                                                                                                                              |

### Response

**[*operations.C1APIFindingV1FindingRoutingRuleServiceDeleteFindingRoutingRuleResponse](../../pkg/models/operations/c1apifindingv1findingroutingruleservicedeletefindingroutingruleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetFindingRoutingRule

Retrieve a single finding routing rule by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingRoutingRuleService.GetFindingRoutingRule" method="get" path="/api/v1/findings/routing-rules/{id}" -->
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

    res, err := s.FindingRoutingRule.GetFindingRoutingRule(ctx, operations.C1APIFindingV1FindingRoutingRuleServiceGetFindingRoutingRuleRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetFindingRoutingRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                            | Type                                                                                                                                                                                 | Required                                                                                                                                                                             | Description                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                | :heavy_check_mark:                                                                                                                                                                   | The context to use for the request.                                                                                                                                                  |
| `request`                                                                                                                                                                            | [operations.C1APIFindingV1FindingRoutingRuleServiceGetFindingRoutingRuleRequest](../../pkg/models/operations/c1apifindingv1findingroutingruleservicegetfindingroutingrulerequest.md) | :heavy_check_mark:                                                                                                                                                                   | The request object to use for the request.                                                                                                                                           |
| `opts`                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                   | The options for this request.                                                                                                                                                        |

### Response

**[*operations.C1APIFindingV1FindingRoutingRuleServiceGetFindingRoutingRuleResponse](../../pkg/models/operations/c1apifindingv1findingroutingruleservicegetfindingroutingruleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListFindingRoutingRules

List finding routing rules, optionally filtered to a specific app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingRoutingRuleService.ListFindingRoutingRules" method="get" path="/api/v1/findings/routing-rules" -->
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

    res, err := s.FindingRoutingRule.ListFindingRoutingRules(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListFindingRoutingRulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.C1APIFindingV1FindingRoutingRuleServiceListFindingRoutingRulesResponse](../../pkg/models/operations/c1apifindingv1findingroutingruleservicelistfindingroutingrulesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateFindingRoutingRule

Update an existing finding routing rule's match criteria or target policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingRoutingRuleService.UpdateFindingRoutingRule" method="post" path="/api/v1/findings/routing-rules/{routing_rule_id}/update" -->
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

    res, err := s.FindingRoutingRule.UpdateFindingRoutingRule(ctx, operations.C1APIFindingV1FindingRoutingRuleServiceUpdateFindingRoutingRuleRequest{
        RoutingRuleID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateFindingRoutingRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                  | Type                                                                                                                                                                                       | Required                                                                                                                                                                                   | Description                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                         | The context to use for the request.                                                                                                                                                        |
| `request`                                                                                                                                                                                  | [operations.C1APIFindingV1FindingRoutingRuleServiceUpdateFindingRoutingRuleRequest](../../pkg/models/operations/c1apifindingv1findingroutingruleserviceupdatefindingroutingrulerequest.md) | :heavy_check_mark:                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                 |
| `opts`                                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                         | The options for this request.                                                                                                                                                              |

### Response

**[*operations.C1APIFindingV1FindingRoutingRuleServiceUpdateFindingRoutingRuleResponse](../../pkg/models/operations/c1apifindingv1findingroutingruleserviceupdatefindingroutingruleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |