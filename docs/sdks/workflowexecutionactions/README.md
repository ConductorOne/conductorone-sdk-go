# WorkflowExecutionActions
(*WorkflowExecutionActions*)

## Overview

### Available Operations

* [TerminateWorkflow](#terminateworkflow) - Terminate Workflow

## TerminateWorkflow

Invokes the c1.api.workflows.v1beta.WorkflowExecutionActionsService.TerminateWorkflow method.

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
    res, err := s.WorkflowExecutionActions.TerminateWorkflow(ctx, operations.C1APIWorkflowsV1betaWorkflowExecutionActionsServiceTerminateWorkflowRequest{
        ID: 568176,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TerminateWorkflowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                            | Type                                                                                                                                                                                                 | Required                                                                                                                                                                                             | Description                                                                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                  |
| `request`                                                                                                                                                                                            | [operations.C1APIWorkflowsV1betaWorkflowExecutionActionsServiceTerminateWorkflowRequest](../../pkg/models/operations/c1apiworkflowsv1betaworkflowexecutionactionsserviceterminateworkflowrequest.md) | :heavy_check_mark:                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                           |
| `opts`                                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                   | The options for this request.                                                                                                                                                                        |

### Response

**[*operations.C1APIWorkflowsV1betaWorkflowExecutionActionsServiceTerminateWorkflowResponse](../../pkg/models/operations/c1apiworkflowsv1betaworkflowexecutionactionsserviceterminateworkflowresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |