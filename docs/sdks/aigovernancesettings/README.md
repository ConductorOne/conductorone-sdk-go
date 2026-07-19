# AIGovernanceSettings

## Overview

### Available Operations

* [Get](#get) - Get
* [GetTenantDefaults](#gettenantdefaults) - Get Tenant Defaults
* [ListHistory](#listhistory) - List History
* [Update](#update) - Update

## Get

Get the tenant's AI governance settings — the controls behind the admin
 /admin/settings/ai-governance page. Returns the full AIGovernanceSettings:
 allowed MCP client types, default client lifecycle, require_tool_approval,
 default tool classification, audit verbosity, auto-discovery toggle +
 interval, prefer_code_mode_over_direct_tools, and surface_requestable_tools.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.AIGovernanceSettingsService.Get" method="get" path="/api/v1/settings/ai-governance" -->
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

    res, err := s.AIGovernanceSettings.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAIGovernanceSettingsResponse != nil {
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

**[*operations.C1APIAIGovernanceV1AIGovernanceSettingsServiceGetResponse](../../pkg/models/operations/c1apiaigovernancev1aigovernancesettingsservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTenantDefaults

Get the tenant-default subset of AI governance settings. Currently returns
 only require_tool_approval — the default applied to newly registered MCP
 servers/tools. Use Get for the full settings object.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.AIGovernanceSettingsService.GetTenantDefaults" method="get" path="/api/v1/settings/ai-governance/tenant-defaults" -->
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

    res, err := s.AIGovernanceSettings.GetTenantDefaults(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetTenantDefaultsResponse != nil {
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

**[*operations.C1APIAIGovernanceV1AIGovernanceSettingsServiceGetTenantDefaultsResponse](../../pkg/models/operations/c1apiaigovernancev1aigovernancesettingsservicegettenantdefaultsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHistory

List the change history for AI governance settings (newest first).
 Singleton: scoped to the caller's tenant. Admin-tier per object-history convention.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.AIGovernanceSettingsService.ListHistory" method="get" path="/api/v1/settings/ai-governance/history" -->
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

    res, err := s.AIGovernanceSettings.ListHistory(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGovernanceSettingsHistoryResponse != nil {
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

**[*operations.C1APIAIGovernanceV1AIGovernanceSettingsServiceListHistoryResponse](../../pkg/models/operations/c1apiaigovernancev1aigovernancesettingsservicelisthistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update the tenant's AI governance settings. Requires update_mask listing
 which fields to apply (e.g. require_tool_approval,
 default_tool_classification, audit_verbosity, auto_discovery_enabled,
 discovery_interval, prefer_code_mode_over_direct_tools,
 surface_requestable_tools, allowed_client_types, default_client_lifecycle).
 Only masked fields change. Returns the updated settings.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.AIGovernanceSettingsService.Update" method="post" path="/api/v1/settings/ai-governance" -->
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

    res, err := s.AIGovernanceSettings.Update(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateAIGovernanceSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [shared.UpdateAIGovernanceSettingsRequest](../../pkg/models/shared/updateaigovernancesettingsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.C1APIAIGovernanceV1AIGovernanceSettingsServiceUpdateResponse](../../pkg/models/operations/c1apiaigovernancev1aigovernancesettingsserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |