# AccessReviewActions

## Overview

### Available Operations

* [GenerateReport](#generatereport) - Generate Report

## GenerateReport

Generate a report of the campaign's reviews and decisions. The format
 defaults to JSON (also available: CSV, XLSX). Works on in-flight (OPEN)
 and closed campaigns. Asynchronous — the report record is created
 immediately; the file is materialized in the background.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewActionsService.GenerateReport" method="post" path="/api/v1/access_review/{access_review_id}/report" -->
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

    res, err := s.AccessReviewActions.GenerateReport(ctx, operations.C1APIAccessreviewV1AccessReviewActionsServiceGenerateReportRequest{
        AccessReviewID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewActionsServiceGenerateReportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                          | Type                                                                                                                                                                               | Required                                                                                                                                                                           | Description                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                              | :heavy_check_mark:                                                                                                                                                                 | The context to use for the request.                                                                                                                                                |
| `request`                                                                                                                                                                          | [operations.C1APIAccessreviewV1AccessReviewActionsServiceGenerateReportRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewactionsservicegeneratereportrequest.md) | :heavy_check_mark:                                                                                                                                                                 | The request object to use for the request.                                                                                                                                         |
| `opts`                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                 | The options for this request.                                                                                                                                                      |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewActionsServiceGenerateReportResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewactionsservicegeneratereportresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |