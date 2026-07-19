# MCPAccessProfileToolBinding

## Overview

### Available Operations

* [CreateBindings](#createbindings) - Create Bindings
* [DeleteBindings](#deletebindings) - Delete Bindings
* [GetAccessProfilesForTools](#getaccessprofilesfortools) - Get Access Profiles For Tools
* [List](#list) - List
* [ListProfilesByToolHistory](#listprofilesbytoolhistory) - List Profiles By Tool History
* [ListToolsByProfileHistory](#listtoolsbyprofilehistory) - List Tools By Profile History

## CreateBindings

CreateBindings adds one or more MCP tools (mcp_tool_ids) to a toolset.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileToolBindingService.CreateBindings" method="post" path="/api/v1/apps/{app_id}/connectors/{connector_id}/mcp_toolsets/{access_profile_id}/tool_bindings" -->
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

    res, err := s.MCPAccessProfileToolBinding.CreateBindings(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceCreateBindingsRequest{
        AccessProfileID: "<id>",
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileToolBindingServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                          | Type                                                                                                                                                                                               | Required                                                                                                                                                                                           | Description                                                                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                |
| `request`                                                                                                                                                                                          | [operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceCreateBindingsRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofiletoolbindingservicecreatebindingsrequest.md) | :heavy_check_mark:                                                                                                                                                                                 | The request object to use for the request.                                                                                                                                                         |
| `opts`                                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                 | The options for this request.                                                                                                                                                                      |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceCreateBindingsResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofiletoolbindingservicecreatebindingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteBindings

DeleteBindings removes one or more MCP tools (mcp_tool_ids) from a toolset.
 Uses a POST .../delete action route because the tool IDs travel in the
 request body, which HTTP DELETE does not reliably support.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileToolBindingService.DeleteBindings" method="post" path="/api/v1/apps/{app_id}/connectors/{connector_id}/mcp_toolsets/{access_profile_id}/tool_bindings/delete" -->
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

    res, err := s.MCPAccessProfileToolBinding.DeleteBindings(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceDeleteBindingsRequest{
        AccessProfileID: "<id>",
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileToolBindingServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                          | Type                                                                                                                                                                                               | Required                                                                                                                                                                                           | Description                                                                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                |
| `request`                                                                                                                                                                                          | [operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceDeleteBindingsRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofiletoolbindingservicedeletebindingsrequest.md) | :heavy_check_mark:                                                                                                                                                                                 | The request object to use for the request.                                                                                                                                                         |
| `opts`                                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                 | The options for this request.                                                                                                                                                                      |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceDeleteBindingsResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofiletoolbindingservicedeletebindingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccessProfilesForTools

GetAccessProfilesForTools returns the access profiles bound to each
 of the given MCP tools, hydrated with display_name. Used by the tools
 list to render the "toolset" column for visible rows.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileToolBindingService.GetAccessProfilesForTools" method="post" path="/api/v1/apps/{app_id}/connectors/{connector_id}/tool_bindings/by_tools" -->
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

    res, err := s.MCPAccessProfileToolBinding.GetAccessProfilesForTools(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceGetAccessProfilesForToolsRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileToolBindingServiceGetAccessProfilesForToolsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                | Type                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                 | Description                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                      |
| `request`                                                                                                                                                                                                                | [operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceGetAccessProfilesForToolsRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofiletoolbindingservicegetaccessprofilesfortoolsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                                               |
| `opts`                                                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                            |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceGetAccessProfilesForToolsResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofiletoolbindingservicegetaccessprofilesfortoolsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List returns the tool bindings for a single toolset (access profile) —
 i.e. which MCP tools belong to the toolset — paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileToolBindingService.List" method="get" path="/api/v1/apps/{app_id}/connectors/{connector_id}/mcp_toolsets/{access_profile_id}/tool_bindings" -->
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

    res, err := s.MCPAccessProfileToolBinding.List(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceListRequest{
        AccessProfileID: "<id>",
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileToolBindingServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                          | :heavy_check_mark:                                                                                                                                                             | The context to use for the request.                                                                                                                                            |
| `request`                                                                                                                                                                      | [operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceListRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofiletoolbindingservicelistrequest.md) | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `opts`                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                   | :heavy_minus_sign:                                                                                                                                                             | The options for this request.                                                                                                                                                  |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceListResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofiletoolbindingservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListProfilesByToolHistory

ListProfilesByToolHistory returns the transactional history of
 profiles bound to (app, connector, mcp_tool). Newest first.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileToolBindingService.ListProfilesByToolHistory" method="get" path="/api/v1/apps/{app_id}/connectors/{connector_id}/tool_bindings/by_tool/{mcp_tool_id}/history" -->
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

    res, err := s.MCPAccessProfileToolBinding.ListProfilesByToolHistory(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceListProfilesByToolHistoryRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
        McpToolID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileToolBindingServiceListProfilesByToolHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                | Type                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                 | Description                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                      |
| `request`                                                                                                                                                                                                                | [operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceListProfilesByToolHistoryRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofiletoolbindingservicelistprofilesbytoolhistoryrequest.md) | :heavy_check_mark:                                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                                               |
| `opts`                                                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                            |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceListProfilesByToolHistoryResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofiletoolbindingservicelistprofilesbytoolhistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListToolsByProfileHistory

ListToolsByProfileHistory returns the transactional history of
 tools bound to (app, connector, access_profile). Newest first.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileToolBindingService.ListToolsByProfileHistory" method="get" path="/api/v1/apps/{app_id}/connectors/{connector_id}/mcp_toolsets/{access_profile_id}/tool_bindings/history" -->
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

    res, err := s.MCPAccessProfileToolBinding.ListToolsByProfileHistory(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceListToolsByProfileHistoryRequest{
        AccessProfileID: "<id>",
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileToolBindingServiceListToolsByProfileHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                | Type                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                 | Description                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                      |
| `request`                                                                                                                                                                                                                | [operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceListToolsByProfileHistoryRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofiletoolbindingservicelisttoolsbyprofilehistoryrequest.md) | :heavy_check_mark:                                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                                               |
| `opts`                                                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                            |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileToolBindingServiceListToolsByProfileHistoryResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofiletoolbindingservicelisttoolsbyprofilehistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |