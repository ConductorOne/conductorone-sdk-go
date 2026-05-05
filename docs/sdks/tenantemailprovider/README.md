# TenantEmailProvider

## Overview

### Available Operations

* [Get](#get) - Get
* [GetEmailCapabilities](#getemailcapabilities) - Get Email Capabilities
* [SearchAuditEvents](#searchauditevents) - Search Audit Events
* [Test](#test) - Test
* [Update](#update) - Update

## Get

Get retrieves the current tenant email provider configuration.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.settings.v1.TenantEmailProviderService.Get" method="get" path="/api/v1/settings/email-provider" -->
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

    res, err := s.TenantEmailProvider.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetTenantEmailProviderResponse != nil {
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

**[*operations.C1APISettingsV1TenantEmailProviderServiceGetResponse](../../pkg/models/operations/c1apisettingsv1tenantemailproviderservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetEmailCapabilities

GetEmailCapabilities returns a lightweight summary of email capabilities
 for the current tenant. Intended for non-admin users (automation builders,
 secret sharers) to check if external email is available without exposing
 provider configuration details.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.settings.v1.TenantEmailProviderService.GetEmailCapabilities" method="get" path="/api/v1/settings/email-capabilities" -->
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

    res, err := s.TenantEmailProvider.GetEmailCapabilities(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetEmailCapabilitiesResponse != nil {
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

**[*operations.C1APISettingsV1TenantEmailProviderServiceGetEmailCapabilitiesResponse](../../pkg/models/operations/c1apisettingsv1tenantemailproviderservicegetemailcapabilitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchAuditEvents

SearchAuditEvents returns email audit events for the tenant.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.settings.v1.TenantEmailProviderService.SearchAuditEvents" method="post" path="/api/v1/settings/email-provider/audit-events" -->
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

    res, err := s.TenantEmailProvider.SearchAuditEvents(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchEmailAuditEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [shared.SearchEmailAuditEventsRequest](../../pkg/models/shared/searchemailauditeventsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.C1APISettingsV1TenantEmailProviderServiceSearchAuditEventsResponse](../../pkg/models/operations/c1apisettingsv1tenantemailproviderservicesearchauditeventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Test

Test sends a test email to verify the provider configuration works end-to-end.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.settings.v1.TenantEmailProviderService.Test" method="post" path="/api/v1/settings/email-provider/test" -->
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

    res, err := s.TenantEmailProvider.Test(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TestTenantEmailProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [shared.TestTenantEmailProviderRequest](../../pkg/models/shared/testtenantemailproviderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.C1APISettingsV1TenantEmailProviderServiceTestResponse](../../pkg/models/operations/c1apisettingsv1tenantemailproviderservicetestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update creates or updates the tenant email provider configuration.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.settings.v1.TenantEmailProviderService.Update" method="post" path="/api/v1/settings/email-provider" -->
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

    res, err := s.TenantEmailProvider.Update(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateTenantEmailProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [shared.UpdateTenantEmailProviderRequest](../../pkg/models/shared/updatetenantemailproviderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.C1APISettingsV1TenantEmailProviderServiceUpdateResponse](../../pkg/models/operations/c1apisettingsv1tenantemailproviderserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |