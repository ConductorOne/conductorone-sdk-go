# AutomationExecution

## Overview

### Available Operations

* [GetAutomationExecution](#getautomationexecution) - Get Automation Execution
* [ListAutomationExecutions](#listautomationexecutions) - List Automation Executions

## GetAutomationExecution

Invokes the c1.api.automations.v1.AutomationExecutionService.GetAutomationExecution method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationExecutionService.GetAutomationExecution" method="get" path="/api/v1/automation_executions/{id}" -->
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

    res, err := s.AutomationExecution.GetAutomationExecution(ctx, operations.C1APIAutomationsV1AutomationExecutionServiceGetAutomationExecutionRequest{
        ID: 728203,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAutomationExecutionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                        | Type                                                                                                                                                                                             | Required                                                                                                                                                                                         | Description                                                                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                               | The context to use for the request.                                                                                                                                                              |
| `request`                                                                                                                                                                                        | [operations.C1APIAutomationsV1AutomationExecutionServiceGetAutomationExecutionRequest](../../pkg/models/operations/c1apiautomationsv1automationexecutionservicegetautomationexecutionrequest.md) | :heavy_check_mark:                                                                                                                                                                               | The request object to use for the request.                                                                                                                                                       |
| `opts`                                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                               | The options for this request.                                                                                                                                                                    |

### Response

**[*operations.C1APIAutomationsV1AutomationExecutionServiceGetAutomationExecutionResponse](../../pkg/models/operations/c1apiautomationsv1automationexecutionservicegetautomationexecutionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAutomationExecutions

Invokes the c1.api.automations.v1.AutomationExecutionService.ListAutomationExecutions method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationExecutionService.ListAutomationExecutions" method="get" path="/api/v1/automation_executions" -->
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

    res, err := s.AutomationExecution.ListAutomationExecutions(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAutomationExecutionsResponse != nil {
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

**[*operations.C1APIAutomationsV1AutomationExecutionServiceListAutomationExecutionsResponse](../../pkg/models/operations/c1apiautomationsv1automationexecutionservicelistautomationexecutionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |