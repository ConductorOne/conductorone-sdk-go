# Feedback

## Overview

### Available Operations

* [CreateFeedback](#createfeedback) - Create Feedback

## CreateFeedback

Create feedback with client diagnostics, Datadog correlation context, and
 up to five WebP screenshots. Gated by the feedback feature flag; the
 submitter's identity is taken from the session, not the request.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.feedback.v1.FeedbackService.CreateFeedback" method="post" path="/api/v1/feedback" -->
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

    res, err := s.Feedback.CreateFeedback(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateFeedbackResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.CreateFeedbackRequest](../../pkg/models/shared/createfeedbackrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.C1APIFeedbackV1FeedbackServiceCreateFeedbackResponse](../../pkg/models/operations/c1apifeedbackv1feedbackservicecreatefeedbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |