# AccessReviewReport

## Overview

### Available Operations

* [List](#list) - List

## List

List the generated reports for an access review campaign, each with a
 time-limited download_url and its output format.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewReportService.List" method="get" path="/api/v1/access_review/{access_review_id}/report" -->
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

    res, err := s.AccessReviewReport.List(ctx, operations.C1APIAccessreviewV1AccessReviewReportServiceListRequest{
        AccessReviewID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewReportServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIAccessreviewV1AccessReviewReportServiceListRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewreportservicelistrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewReportServiceListResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewreportservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |