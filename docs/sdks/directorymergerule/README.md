# DirectoryMergeRule

## Overview

### Available Operations

* [ListDirectoryMergeRules](#listdirectorymergerules) - List Directory Merge Rules
* [SetDirectoryMergeRules](#setdirectorymergerules) - Set Directory Merge Rules

## ListDirectoryMergeRules

ListDirectoryMergeRules returns the directory's rules in evaluation order.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.directory.v1.DirectoryMergeRuleService.ListDirectoryMergeRules" method="get" path="/api/v1/directories/{app_id}/merge_rules" -->
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

    res, err := s.DirectoryMergeRule.ListDirectoryMergeRules(ctx, operations.C1APIDirectoryV1DirectoryMergeRuleServiceListDirectoryMergeRulesRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListDirectoryMergeRulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                    | Type                                                                                                                                                                                         | Required                                                                                                                                                                                     | Description                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                           | The context to use for the request.                                                                                                                                                          |
| `request`                                                                                                                                                                                    | [operations.C1APIDirectoryV1DirectoryMergeRuleServiceListDirectoryMergeRulesRequest](../../pkg/models/operations/c1apidirectoryv1directorymergeruleservicelistdirectorymergerulesrequest.md) | :heavy_check_mark:                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                   |
| `opts`                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                           | The options for this request.                                                                                                                                                                |

### Response

**[*operations.C1APIDirectoryV1DirectoryMergeRuleServiceListDirectoryMergeRulesResponse](../../pkg/models/operations/c1apidirectoryv1directorymergeruleservicelistdirectorymergerulesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetDirectoryMergeRules

SetDirectoryMergeRules replaces the directory's whole rule list in one step. The list order is the
 evaluation order. The list must contain each built-in rule exactly once. A custom rule with an empty
 ID is created; a custom rule with an ID must be one of the directory's current custom rules; current
 custom rules left out of the list are deleted. Read-only fields are ignored. Fails with
 INVALID_ARGUMENT for an invalid list and with RESOURCE_EXHAUSTED when the list exceeds the maximum
 number of rules. When the list matches the saved rules nothing changes and the current rules are
 returned.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.directory.v1.DirectoryMergeRuleService.SetDirectoryMergeRules" method="post" path="/api/v1/directories/{app_id}/merge_rules" -->
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

    res, err := s.DirectoryMergeRule.SetDirectoryMergeRules(ctx, operations.C1APIDirectoryV1DirectoryMergeRuleServiceSetDirectoryMergeRulesRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SetDirectoryMergeRulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                  | Type                                                                                                                                                                                       | Required                                                                                                                                                                                   | Description                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                         | The context to use for the request.                                                                                                                                                        |
| `request`                                                                                                                                                                                  | [operations.C1APIDirectoryV1DirectoryMergeRuleServiceSetDirectoryMergeRulesRequest](../../pkg/models/operations/c1apidirectoryv1directorymergeruleservicesetdirectorymergerulesrequest.md) | :heavy_check_mark:                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                 |
| `opts`                                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                         | The options for this request.                                                                                                                                                              |

### Response

**[*operations.C1APIDirectoryV1DirectoryMergeRuleServiceSetDirectoryMergeRulesResponse](../../pkg/models/operations/c1apidirectoryv1directorymergeruleservicesetdirectorymergerulesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |