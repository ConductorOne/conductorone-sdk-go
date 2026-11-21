# AutomationExecutionActions

## Overview

### Available Operations

* [RetryAutomationExecutions](#retryautomationexecutions) - Retry Automation Executions
* [TerminateAutomation](#terminateautomation) - Terminate Automation

## RetryAutomationExecutions

Retry failed automation executions from the step that failed, without re-running steps that already succeeded.
 Executions that are not in a failed state are skipped and reported in the response. Retries run asynchronously.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationExecutionActionsService.RetryAutomationExecutions" method="post" path="/api/v1/automation_executions/actions/retry" -->
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

    res, err := s.AutomationExecutionActions.RetryAutomationExecutions(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RetryAutomationExecutionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [shared.RetryAutomationExecutionsRequest](../../pkg/models/shared/retryautomationexecutionsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.C1APIAutomationsV1AutomationExecutionActionsServiceRetryAutomationExecutionsResponse](../../pkg/models/operations/c1apiautomationsv1automationexecutionactionsserviceretryautomationexecutionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## TerminateAutomation

Terminate a running automation execution asynchronously, stopping it and marking it as terminated.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationExecutionActionsService.TerminateAutomation" method="post" path="/api/v1/automation_executions/{id}/actions/terminate" -->
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

    res, err := s.AutomationExecutionActions.TerminateAutomation(ctx, operations.C1APIAutomationsV1AutomationExecutionActionsServiceTerminateAutomationRequest{
        ID: 839265,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TerminateAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                | Type                                                                                                                                                                                                     | Required                                                                                                                                                                                                 | Description                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                      |
| `request`                                                                                                                                                                                                | [operations.C1APIAutomationsV1AutomationExecutionActionsServiceTerminateAutomationRequest](../../pkg/models/operations/c1apiautomationsv1automationexecutionactionsserviceterminateautomationrequest.md) | :heavy_check_mark:                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                               |
| `opts`                                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                       | The options for this request.                                                                                                                                                                            |

### Response

**[*operations.C1APIAutomationsV1AutomationExecutionActionsServiceTerminateAutomationResponse](../../pkg/models/operations/c1apiautomationsv1automationexecutionactionsserviceterminateautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |