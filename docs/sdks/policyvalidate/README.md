# PolicyValidate
(*PolicyValidate*)

## Overview

### Available Operations

* [ValidateCEL](#validatecel) - Validate Cel

## ValidateCEL

Validate policies

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.policy.v1.PolicyValidate.ValidateCEL" method="post" path="/api/v1/policies/validate/cel" -->
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

    res, err := s.PolicyValidate.ValidateCEL(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EditorValidateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.EditorValidateRequest](../../pkg/models/shared/editorvalidaterequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.C1APIPolicyV1PolicyValidateValidateCELResponse](../../pkg/models/operations/c1apipolicyv1policyvalidatevalidatecelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |