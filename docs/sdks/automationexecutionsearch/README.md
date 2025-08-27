# AutomationExecutionSearch
(*AutomationExecutionSearch*)

## Overview

### Available Operations

* [SearchAutomationExecutions](#searchautomationexecutions) - Search Automation Executions

## SearchAutomationExecutions

Invokes the c1.api.automations.v1.AutomationExecutionSearchService.SearchAutomationExecutions method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationExecutionSearchService.SearchAutomationExecutions" method="post" path="/api/v1/automation_executions/search" -->
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

    res, err := s.AutomationExecutionSearch.SearchAutomationExecutions(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAutomationExecutionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [shared.SearchAutomationExecutionsRequest](../../pkg/models/shared/searchautomationexecutionsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.C1APIAutomationsV1AutomationExecutionSearchServiceSearchAutomationExecutionsResponse](../../pkg/models/operations/c1apiautomationsv1automationexecutionsearchservicesearchautomationexecutionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |