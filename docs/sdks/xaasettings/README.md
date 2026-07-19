# XAASettings

## Overview

### Available Operations

* [Get](#get) - Get
* [ListHistory](#listhistory) - List History
* [Update](#update) - Update

## Get

Get the tenant's cross-app-access settings.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAASettingsService.Get" method="get" path="/api/v1/settings/cross-app-access" -->
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

    res, err := s.XAASettings.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.XAASettingsServiceGetResponse != nil {
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

**[*operations.C1APICrossAppAccessV1XAASettingsServiceGetResponse](../../pkg/models/operations/c1apicrossappaccessv1xaasettingsservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHistory

List the change history for the tenant's cross-app-access settings (newest
 first). Singleton: scoped to the caller's tenant.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAASettingsService.ListHistory" method="get" path="/api/v1/settings/cross-app-access/history" -->
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

    res, err := s.XAASettings.ListHistory(ctx, operations.C1APICrossAppAccessV1XAASettingsServiceListHistoryRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.XAASettingsServiceListHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.C1APICrossAppAccessV1XAASettingsServiceListHistoryRequest](../../pkg/models/operations/c1apicrossappaccessv1xaasettingsservicelisthistoryrequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |
| `opts`                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |

### Response

**[*operations.C1APICrossAppAccessV1XAASettingsServiceListHistoryResponse](../../pkg/models/operations/c1apicrossappaccessv1xaasettingsservicelisthistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update the tenant's cross-app-access settings. Supply the settings object
 and an update mask listing the fields to change; only masked fields are
 applied. Editable paths: enabled, default_grant_lifetime,
 allow_refresh_token_subjects, default_signing_algorithm,
 enabled_signing_algorithms, xaa_id_token_lifetime.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAASettingsService.Update" method="post" path="/api/v1/settings/cross-app-access" -->
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

    res, err := s.XAASettings.Update(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.XAASettingsServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [shared.XAASettingsServiceUpdateRequest](../../pkg/models/shared/xaasettingsserviceupdaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.C1APICrossAppAccessV1XAASettingsServiceUpdateResponse](../../pkg/models/operations/c1apicrossappaccessv1xaasettingsserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |