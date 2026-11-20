# VirtualMCPToolsetBinding

## Overview

### Available Operations

* [CreateBindings](#createbindings) - Create Bindings
* [DeleteBindings](#deletebindings) - Delete Bindings
* [List](#list) - List

## CreateBindings

CreateBindings assigns up to 100 Toolsets to an Assigned-tools Virtual MCP Server.
 Existing assignments are unchanged and previously deleted assignments are restored.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.VirtualMCPToolsetBindingService.CreateBindings" method="post" path="/api/v1/virtual_mcp_servers/{virtual_mcp_server_id}/toolset_bindings" -->
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

    res, err := s.VirtualMCPToolsetBinding.CreateBindings(ctx, operations.C1APIAiGovernanceV1VirtualMCPToolsetBindingServiceCreateBindingsRequest{
        VirtualMcpServerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMCPToolsetBindingServiceCreateBindingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                    | Type                                                                                                                                                                                         | Required                                                                                                                                                                                     | Description                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                           | The context to use for the request.                                                                                                                                                          |
| `request`                                                                                                                                                                                    | [operations.C1APIAiGovernanceV1VirtualMCPToolsetBindingServiceCreateBindingsRequest](../../pkg/models/operations/c1apiaigovernancev1virtualmcptoolsetbindingservicecreatebindingsrequest.md) | :heavy_check_mark:                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                   |
| `opts`                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                           | The options for this request.                                                                                                                                                                |

### Response

**[*operations.C1APIAiGovernanceV1VirtualMCPToolsetBindingServiceCreateBindingsResponse](../../pkg/models/operations/c1apiaigovernancev1virtualmcptoolsetbindingservicecreatebindingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteBindings

DeleteBindings removes up to 100 Toolset assignments from an Assigned-tools Virtual MCP Server.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.VirtualMCPToolsetBindingService.DeleteBindings" method="post" path="/api/v1/virtual_mcp_servers/{virtual_mcp_server_id}/toolset_bindings/delete" -->
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

    res, err := s.VirtualMCPToolsetBinding.DeleteBindings(ctx, operations.C1APIAiGovernanceV1VirtualMCPToolsetBindingServiceDeleteBindingsRequest{
        VirtualMcpServerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMCPToolsetBindingServiceDeleteBindingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                    | Type                                                                                                                                                                                         | Required                                                                                                                                                                                     | Description                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                           | The context to use for the request.                                                                                                                                                          |
| `request`                                                                                                                                                                                    | [operations.C1APIAiGovernanceV1VirtualMCPToolsetBindingServiceDeleteBindingsRequest](../../pkg/models/operations/c1apiaigovernancev1virtualmcptoolsetbindingservicedeletebindingsrequest.md) | :heavy_check_mark:                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                   |
| `opts`                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                           | The options for this request.                                                                                                                                                                |

### Response

**[*operations.C1APIAiGovernanceV1VirtualMCPToolsetBindingServiceDeleteBindingsResponse](../../pkg/models/operations/c1apiaigovernancev1virtualmcptoolsetbindingservicedeletebindingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List returns one page of Toolsets assigned directly to an Assigned-tools Virtual MCP Server.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.VirtualMCPToolsetBindingService.List" method="get" path="/api/v1/virtual_mcp_servers/{virtual_mcp_server_id}/toolset_bindings" -->
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

    res, err := s.VirtualMCPToolsetBinding.List(ctx, operations.C1APIAiGovernanceV1VirtualMCPToolsetBindingServiceListRequest{
        VirtualMcpServerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMCPToolsetBindingServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [operations.C1APIAiGovernanceV1VirtualMCPToolsetBindingServiceListRequest](../../pkg/models/operations/c1apiaigovernancev1virtualmcptoolsetbindingservicelistrequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |
| `opts`                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                                       | The options for this request.                                                                                                                                            |

### Response

**[*operations.C1APIAiGovernanceV1VirtualMCPToolsetBindingServiceListResponse](../../pkg/models/operations/c1apiaigovernancev1virtualmcptoolsetbindingservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |