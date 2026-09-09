# UIConversations

## Overview

### Available Operations

* [EnsureOnboardingSession](#ensureonboardingsession) - Ensure Onboarding Session

## EnsureOnboardingSession

EnsureOnboardingSession returns the tenant's active onboarding conversation,
 or creates and starts it once. Retries converge on the stored conversation.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.conversations.v1.UIConversationsService.EnsureOnboardingSession" method="post" path="/api/v1/conversations/onboarding:ensure" -->
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

    res, err := s.UIConversations.EnsureOnboardingSession(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EnsureOnboardingSessionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [shared.EnsureOnboardingSessionRequest](../../pkg/models/shared/ensureonboardingsessionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.C1APIConversationsV1UIConversationsServiceEnsureOnboardingSessionResponse](../../pkg/models/operations/c1apiconversationsv1uiconversationsserviceensureonboardingsessionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |