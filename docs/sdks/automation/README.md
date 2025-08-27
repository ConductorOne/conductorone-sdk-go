# Automation
(*Automation*)

## Overview

### Available Operations

* [CreateAutomation](#createautomation) - Create Automation
* [DeleteAutomation](#deleteautomation) - Delete Automation
* [ExecuteAutomation](#executeautomation) - Execute Automation
* [GetAutomation](#getautomation) - Get Automation
* [ListAutomations](#listautomations) - List Automations
* [UpdateAutomation](#updateautomation) - Update Automation

## CreateAutomation

Invokes the c1.api.automations.v1.AutomationService.CreateAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.CreateAutomation" method="post" path="/api/v1/automations" -->
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

    res, err := s.Automation.CreateAutomation(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [shared.CreateAutomationRequestInput](../../pkg/models/shared/createautomationrequestinput.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceCreateAutomationResponse](../../pkg/models/operations/c1apiautomationsv1automationservicecreateautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteAutomation

Invokes the c1.api.automations.v1.AutomationService.DeleteAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.DeleteAutomation" method="delete" path="/api/v1/automations/{id}" -->
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

    res, err := s.Automation.DeleteAutomation(ctx, operations.C1APIAutomationsV1AutomationServiceDeleteAutomationRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAutomationResponse != nil {
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

Invokes the c1.api.automations.v1.AutomationService.ExecuteAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.ExecuteAutomation" method="post" path="/api/v1/automations/{id}/execute" -->
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

Invokes the c1.api.automations.v1.AutomationService.GetAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.GetAutomation" method="get" path="/api/v1/automations/{id}" -->
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

Invokes the c1.api.automations.v1.AutomationService.ListAutomations method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.ListAutomations" method="get" path="/api/v1/automations" -->
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

## UpdateAutomation

Invokes the c1.api.automations.v1.AutomationService.UpdateAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.UpdateAutomation" method="post" path="/api/v1/automations/{id}" -->
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