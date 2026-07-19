# FindingTransformationRule

## Overview

### Available Operations

* [CreateFindingTransformationRule](#createfindingtransformationrule) - Create Finding Transformation Rule
* [DeleteFindingTransformationRule](#deletefindingtransformationrule) - Delete Finding Transformation Rule
* [GetFindingTransformationRule](#getfindingtransformationrule) - Get Finding Transformation Rule
* [ListFindingTransformationRules](#listfindingtransformationrules) - List Finding Transformation Rules
* [UpdateFindingTransformationRule](#updatefindingtransformationrule) - Update Finding Transformation Rule

## CreateFindingTransformationRule

Create a new finding transformation rule.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingTransformationRuleService.CreateFindingTransformationRule" method="post" path="/api/v1/findings/transformation-rules" -->
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

    res, err := s.FindingTransformationRule.CreateFindingTransformationRule(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateFindingTransformationRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [shared.CreateFindingTransformationRuleRequest](../../pkg/models/shared/createfindingtransformationrulerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.C1APIFindingV1FindingTransformationRuleServiceCreateFindingTransformationRuleResponse](../../pkg/models/operations/c1apifindingv1findingtransformationruleservicecreatefindingtransformationruleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteFindingTransformationRule

Delete a finding transformation rule. Findings already transformed by this rule are not affected.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingTransformationRuleService.DeleteFindingTransformationRule" method="delete" path="/api/v1/findings/transformation-rules/{id}" -->
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

    res, err := s.FindingTransformationRule.DeleteFindingTransformationRule(ctx, operations.C1APIFindingV1FindingTransformationRuleServiceDeleteFindingTransformationRuleRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteFindingTransformationRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                              | Type                                                                                                                                                                                                                   | Required                                                                                                                                                                                                               | Description                                                                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                    |
| `request`                                                                                                                                                                                                              | [operations.C1APIFindingV1FindingTransformationRuleServiceDeleteFindingTransformationRuleRequest](../../pkg/models/operations/c1apifindingv1findingtransformationruleservicedeletefindingtransformationrulerequest.md) | :heavy_check_mark:                                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                          |

### Response

**[*operations.C1APIFindingV1FindingTransformationRuleServiceDeleteFindingTransformationRuleResponse](../../pkg/models/operations/c1apifindingv1findingtransformationruleservicedeletefindingtransformationruleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetFindingTransformationRule

Retrieve a single finding transformation rule by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingTransformationRuleService.GetFindingTransformationRule" method="get" path="/api/v1/findings/transformation-rules/{id}" -->
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

    res, err := s.FindingTransformationRule.GetFindingTransformationRule(ctx, operations.C1APIFindingV1FindingTransformationRuleServiceGetFindingTransformationRuleRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetFindingTransformationRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                        | Type                                                                                                                                                                                                             | Required                                                                                                                                                                                                         | Description                                                                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                              |
| `request`                                                                                                                                                                                                        | [operations.C1APIFindingV1FindingTransformationRuleServiceGetFindingTransformationRuleRequest](../../pkg/models/operations/c1apifindingv1findingtransformationruleservicegetfindingtransformationrulerequest.md) | :heavy_check_mark:                                                                                                                                                                                               | The request object to use for the request.                                                                                                                                                                       |
| `opts`                                                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                               | The options for this request.                                                                                                                                                                                    |

### Response

**[*operations.C1APIFindingV1FindingTransformationRuleServiceGetFindingTransformationRuleResponse](../../pkg/models/operations/c1apifindingv1findingtransformationruleservicegetfindingtransformationruleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListFindingTransformationRules

List finding transformation rules, optionally filtered to a specific app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingTransformationRuleService.ListFindingTransformationRules" method="get" path="/api/v1/findings/transformation-rules" -->
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

    res, err := s.FindingTransformationRule.ListFindingTransformationRules(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListFindingTransformationRulesResponse != nil {
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

**[*operations.C1APIFindingV1FindingTransformationRuleServiceListFindingTransformationRulesResponse](../../pkg/models/operations/c1apifindingv1findingtransformationruleservicelistfindingtransformationrulesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateFindingTransformationRule

Update an existing finding transformation rule's match criteria or transforms.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingTransformationRuleService.UpdateFindingTransformationRule" method="post" path="/api/v1/findings/transformation-rules/{transformation_rule_id}/update" -->
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

    res, err := s.FindingTransformationRule.UpdateFindingTransformationRule(ctx, operations.C1APIFindingV1FindingTransformationRuleServiceUpdateFindingTransformationRuleRequest{
        TransformationRuleID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateFindingTransformationRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                              | Type                                                                                                                                                                                                                   | Required                                                                                                                                                                                                               | Description                                                                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                    |
| `request`                                                                                                                                                                                                              | [operations.C1APIFindingV1FindingTransformationRuleServiceUpdateFindingTransformationRuleRequest](../../pkg/models/operations/c1apifindingv1findingtransformationruleserviceupdatefindingtransformationrulerequest.md) | :heavy_check_mark:                                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                          |

### Response

**[*operations.C1APIFindingV1FindingTransformationRuleServiceUpdateFindingTransformationRuleResponse](../../pkg/models/operations/c1apifindingv1findingtransformationruleserviceupdatefindingtransformationruleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |