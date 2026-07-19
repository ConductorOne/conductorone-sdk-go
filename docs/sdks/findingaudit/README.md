# FindingAudit

## Overview

### Available Operations

* [Search](#search) - Search

## Search

Search returns audit events filtered by finding, actor, type, or
 app. Authorized as VIEWER -- the same role required to read the
 finding itself.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingAuditService.Search" method="post" path="/api/v1/search/finding_audits" -->
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

    res, err := s.FindingAudit.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FindingAuditServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [shared.FindingAuditServiceSearchRequest](../../pkg/models/shared/findingauditservicesearchrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.C1APIFindingV1FindingAuditServiceSearchResponse](../../pkg/models/operations/c1apifindingv1findingauditservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |