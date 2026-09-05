# Automation

## Overview

### Available Operations

* [ClearAutomationCircuitBreaker](#clearautomationcircuitbreaker) - Clear Automation Circuit Breaker
* [CreateAutomation](#createautomation) - Create Automation
* [DeleteAutomation](#deleteautomation) - Delete Automation
* [ExecuteAutomation](#executeautomation) - Execute Automation
* [GetAutomation](#getautomation) - Get Automation
* [ListAutomations](#listautomations) - List Automations
* [ResolvePausedAutomationExecutions](#resolvepausedautomationexecutions) - Resolve Paused Automation Executions
* [UpdateAutomation](#updateautomation) - Update Automation

## ClearAutomationCircuitBreaker

Clear the circuit breaker on an automation that was auto-disabled by the
 rate cap. Future events flow normally; existing paused executions are not
 affected (use ResolvePausedAutomationExecutions to run or cancel them).

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.ClearAutomationCircuitBreaker" method="post" path="/api/v1/automations/{id}/circuit_breaker/clear" -->
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

    res, err := s.Automation.ClearAutomationCircuitBreaker(ctx, operations.C1APIAutomationsV1AutomationServiceClearAutomationCircuitBreakerRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClearAutomationCircuitBreakerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                    | Type                                                                                                                                                                                         | Required                                                                                                                                                                                     | Description                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                           | The context to use for the request.                                                                                                                                                          |
| `request`                                                                                                                                                                                    | [operations.C1APIAutomationsV1AutomationServiceClearAutomationCircuitBreakerRequest](../../pkg/models/operations/c1apiautomationsv1automationserviceclearautomationcircuitbreakerrequest.md) | :heavy_check_mark:                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                   |
| `opts`                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                           | The options for this request.                                                                                                                                                                |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceClearAutomationCircuitBreakerResponse](../../pkg/models/operations/c1apiautomationsv1automationserviceclearautomationcircuitbreakerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAutomation

Create a new automation with the specified steps, triggers, and
 configuration. See get_authoring_guide for the AutomationStep contract
 (step kinds and their required fields, CEL identifier scope).

 At create time, draft_automation_steps and draft_triggers default to
 their published counterparts when omitted — callers writing a single
 working version don't need to populate both. The draft/publish
 distinction matters only on subsequent edits.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.CreateAutomation" method="post" path="/api/v1/automations" -->
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

    res, err := s.Automation.CreateAutomation(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AutomationsCreateAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [shared.AutomationsCreateAutomationRequest](../../pkg/models/shared/automationscreateautomationrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceCreateAutomationResponse](../../pkg/models/operations/c1apiautomationsv1automationservicecreateautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteAutomation

Delete an automation by its unique identifier, removing it and its associated triggers.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.DeleteAutomation" method="delete" path="/api/v1/automations/{id}" -->
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

    res, err := s.Automation.DeleteAutomation(ctx, operations.C1APIAutomationsV1AutomationServiceDeleteAutomationRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AutomationsDeleteAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APIAutomationsV1AutomationServiceDeleteAutomationRequest](../../pkg/models/operations/c1apiautomationsv1automationservicedeleteautomationrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceDeleteAutomationResponse](../../pkg/models/operations/c1apiautomationsv1automationservicedeleteautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ExecuteAutomation

Trigger an on-demand execution of an automation, returning the new execution's identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.ExecuteAutomation" method="post" path="/api/v1/automations/{id}/execute" -->
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

    res, err := s.Automation.ExecuteAutomation(ctx, operations.C1APIAutomationsV1AutomationServiceExecuteAutomationRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ExecuteAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.C1APIAutomationsV1AutomationServiceExecuteAutomationRequest](../../pkg/models/operations/c1apiautomationsv1automationserviceexecuteautomationrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceExecuteAutomationResponse](../../pkg/models/operations/c1apiautomationsv1automationserviceexecuteautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAutomation

Retrieve a single automation by its unique identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.GetAutomation" method="get" path="/api/v1/automations/{id}" -->
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

    res, err := s.Automation.GetAutomation(ctx, operations.C1APIAutomationsV1AutomationServiceGetAutomationRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIAutomationsV1AutomationServiceGetAutomationRequest](../../pkg/models/operations/c1apiautomationsv1automationservicegetautomationrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceGetAutomationResponse](../../pkg/models/operations/c1apiautomationsv1automationservicegetautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAutomations

List all automations in the tenant with pagination support.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.ListAutomations" method="get" path="/api/v1/automations" -->
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

    res, err := s.Automation.ListAutomations(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAutomationsResponse != nil {
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

**[*operations.C1APIAutomationsV1AutomationServiceListAutomationsResponse](../../pkg/models/operations/c1apiautomationsv1automationservicelistautomationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ResolvePausedAutomationExecutions

Decide what to do with the executions that were paused while the
 automation's circuit breaker was tripped. Idempotent.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.ResolvePausedAutomationExecutions" method="post" path="/api/v1/automations/{id}/circuit_breaker/resolve_paused" -->
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

    res, err := s.Automation.ResolvePausedAutomationExecutions(ctx, operations.C1APIAutomationsV1AutomationServiceResolvePausedAutomationExecutionsRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResolvePausedAutomationExecutionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                            | Type                                                                                                                                                                                                 | Required                                                                                                                                                                                             | Description                                                                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                  |
| `request`                                                                                                                                                                                            | [operations.C1APIAutomationsV1AutomationServiceResolvePausedAutomationExecutionsRequest](../../pkg/models/operations/c1apiautomationsv1automationserviceresolvepausedautomationexecutionsrequest.md) | :heavy_check_mark:                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                           |
| `opts`                                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                   | The options for this request.                                                                                                                                                                        |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceResolvePausedAutomationExecutionsResponse](../../pkg/models/operations/c1apiautomationsv1automationserviceresolvepausedautomationexecutionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAutomation

Update an existing automation's properties, steps, or triggers using a field mask.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.UpdateAutomation" method="post" path="/api/v1/automations/{id}" -->
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

    res, err := s.Automation.UpdateAutomation(ctx, operations.C1APIAutomationsV1AutomationServiceUpdateAutomationRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APIAutomationsV1AutomationServiceUpdateAutomationRequest](../../pkg/models/operations/c1apiautomationsv1automationserviceupdateautomationrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceUpdateAutomationResponse](../../pkg/models/operations/c1apiautomationsv1automationserviceupdateautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |