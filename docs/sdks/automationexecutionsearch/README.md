# AutomationExecutionSearch

## Overview

### Available Operations

* [SearchAllAutomationExecutions](#searchallautomationexecutions) - Search All Automation Executions
* [SearchAutomationExecutions](#searchautomationexecutions) - Search Automation Executions

## SearchAllAutomationExecutions

Search across all automation executions in the tenant, with filters for state, template, app, and subject user.
 Each AutomationExecutionView row is large — request a small page_size (≤10) and set expand_mask to only the paths you need; broad expansion inlines whole related objects and bloats the response.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationExecutionSearchService.SearchAllAutomationExecutions" method="post" path="/api/v1/search/all_automation_executions" -->
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

    res, err := s.AutomationExecutionSearch.SearchAllAutomationExecutions(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAllAutomationExecutionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [shared.SearchAllAutomationExecutionsRequest](../../pkg/models/shared/searchallautomationexecutionsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.C1APIAutomationsV1AutomationExecutionSearchServiceSearchAllAutomationExecutionsResponse](../../pkg/models/operations/c1apiautomationsv1automationexecutionsearchservicesearchallautomationexecutionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchAutomationExecutions

Search for automation executions with optional filters for automation_template_id, state, and query.
 Each AutomationExecutionView row is large — request a small page_size (≤10) and set expand_mask to only the paths you need; broad expansion inlines whole related objects and bloats the response.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationExecutionSearchService.SearchAutomationExecutions" method="post" path="/api/v1/search/automation_executions" -->
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