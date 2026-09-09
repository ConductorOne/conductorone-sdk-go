# SSOSettings

## Overview

### Available Operations

* [Get](#get) - Get
* [ListHistory](#listhistory) - List History
* [Update](#update) - Update

## Get

Get returns the tenant's SSO provider settings.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.sso.v1.SSOSettingsService.Get" method="get" path="/api/v1/settings/sso" -->
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

    res, err := s.SSOSettings.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.SSOSettingsServiceGetResponse != nil {
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

**[*operations.C1APISSOV1SSOSettingsServiceGetResponse](../../pkg/models/operations/c1apissov1ssosettingsservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHistory

ListHistory returns the SSO settings change history, newest first.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.sso.v1.SSOSettingsService.ListHistory" method="get" path="/api/v1/settings/sso/history" -->
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

    res, err := s.SSOSettings.ListHistory(ctx, operations.C1APISSOV1SSOSettingsServiceListHistoryRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SSOSettingsServiceListHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.C1APISSOV1SSOSettingsServiceListHistoryRequest](../../pkg/models/operations/c1apissov1ssosettingsservicelisthistoryrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `opts`                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.C1APISSOV1SSOSettingsServiceListHistoryResponse](../../pkg/models/operations/c1apissov1ssosettingsservicelisthistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update changes the tenant's SSO provider settings. Supply the settings
 object and an update mask listing the fields to change; only masked fields
 are applied. Editable paths: enabled, default_subject_type,
 default_assertion_lifetime, default_id_token_signed_response_alg.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.sso.v1.SSOSettingsService.Update" method="post" path="/api/v1/settings/sso" -->
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

    res, err := s.SSOSettings.Update(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SSOSettingsServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [shared.SSOSettingsServiceUpdateRequest](../../pkg/models/shared/ssosettingsserviceupdaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.C1APISSOV1SSOSettingsServiceUpdateResponse](../../pkg/models/operations/c1apissov1ssosettingsserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |