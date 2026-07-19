# IdentityPolicyTenantDefaults

## Overview

### Available Operations

* [Get](#get) - Get
* [Update](#update) - Update

## Get

Get returns the tenant's default identity policies.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.identity_platform.v1.IdentityPolicyTenantDefaultsService.Get" method="get" path="/api/v1/settings/identity-policy-defaults" -->
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

    res, err := s.IdentityPolicyTenantDefaults.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityPolicyTenantDefaultsServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.C1APIIdentityPlatformV1IdentityPolicyTenantDefaultsServiceGetResponse](../../pkg/models/operations/c1apiidentityplatformv1identitypolicytenantdefaultsservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update changes the tenant's default identity policies. Supply the defaults
 object and an update mask listing the fields to change; omitted fields are
 left as-is.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.identity_platform.v1.IdentityPolicyTenantDefaultsService.Update" method="post" path="/api/v1/settings/identity-policy-defaults" -->
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

    res, err := s.IdentityPolicyTenantDefaults.Update(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityPolicyTenantDefaultsServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [shared.IdentityPolicyTenantDefaultsServiceUpdateRequest](../../pkg/models/shared/identitypolicytenantdefaultsserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.C1APIIdentityPlatformV1IdentityPolicyTenantDefaultsServiceUpdateResponse](../../pkg/models/operations/c1apiidentityplatformv1identitypolicytenantdefaultsserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |