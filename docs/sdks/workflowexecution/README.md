# WorkflowExecution
(*WorkflowExecution*)

## Overview

### Available Operations

* [GetWorkflowExecution](#getworkflowexecution) - Get Workflow Execution
* [ListWorkflowExecutions](#listworkflowexecutions) - List Workflow Executions

## GetWorkflowExecution

Invokes the c1.api.workflows.v1beta.WorkflowExecutionService.GetWorkflowExecution method.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.WorkflowExecution.GetWorkflowExecution(ctx, operations.C1APIWorkflowsV1betaWorkflowExecutionServiceGetWorkflowExecutionRequest{
        ID: 931086,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetWorkflowExecutionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                    | Type                                                                                                                                                                                         | Required                                                                                                                                                                                     | Description                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                           | The context to use for the request.                                                                                                                                                          |
| `request`                                                                                                                                                                                    | [operations.C1APIWorkflowsV1betaWorkflowExecutionServiceGetWorkflowExecutionRequest](../../pkg/models/operations/c1apiworkflowsv1betaworkflowexecutionservicegetworkflowexecutionrequest.md) | :heavy_check_mark:                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                   |
| `opts`                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                           | The options for this request.                                                                                                                                                                |

### Response

**[*operations.C1APIWorkflowsV1betaWorkflowExecutionServiceGetWorkflowExecutionResponse](../../pkg/models/operations/c1apiworkflowsv1betaworkflowexecutionservicegetworkflowexecutionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListWorkflowExecutions

Invokes the c1.api.workflows.v1beta.WorkflowExecutionService.ListWorkflowExecutions method.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.WorkflowExecution.ListWorkflowExecutions(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListWorkflowExecutionsResponse != nil {
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

**[*operations.C1APIWorkflowsV1betaWorkflowExecutionServiceListWorkflowExecutionsResponse](../../pkg/models/operations/c1apiworkflowsv1betaworkflowexecutionservicelistworkflowexecutionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |