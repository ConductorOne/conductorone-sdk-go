# AppSecretAdmin

## Overview

### Available Operations

* [Revoke](#revoke) - Revoke

## Revoke

Revoke queues one provider-side revoke for one vended credential.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppSecretAdminService.Revoke" method="post" path="/api/v1/app-secrets-admin/revoke" -->
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

    res, err := s.AppSecretAdmin.Revoke(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppSecretAdminServiceRevokeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [shared.AppSecretAdminServiceRevokeRequest](../../pkg/models/shared/appsecretadminservicerevokerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.C1APIAppV1AppSecretAdminServiceRevokeResponse](../../pkg/models/operations/c1apiappv1appsecretadminservicerevokeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |