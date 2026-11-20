# VirtualMCPToolBinding

## Overview

### Available Operations

* [CreateBindings](#createbindings) - Create Bindings
* [DeleteBindings](#deletebindings) - Delete Bindings
* [List](#list) - List

## CreateBindings

CreateBindings assigns up to 100 enabled, visible tools to an Assigned-tools Virtual MCP Server.
 Existing assignments are unchanged and previously deleted assignments are restored.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.VirtualMCPToolBindingService.CreateBindings" method="post" path="/api/v1/virtual_mcp_servers/{virtual_mcp_server_id}/tool_bindings" -->
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

    res, err := s.VirtualMCPToolBinding.CreateBindings(ctx, operations.C1APIAiGovernanceV1VirtualMCPToolBindingServiceCreateBindingsRequest{
        VirtualMcpServerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMCPToolBindingServiceCreateBindingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                              | Type                                                                                                                                                                                   | Required                                                                                                                                                                               | Description                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                     | The context to use for the request.                                                                                                                                                    |
| `request`                                                                                                                                                                              | [operations.C1APIAiGovernanceV1VirtualMCPToolBindingServiceCreateBindingsRequest](../../pkg/models/operations/c1apiaigovernancev1virtualmcptoolbindingservicecreatebindingsrequest.md) | :heavy_check_mark:                                                                                                                                                                     | The request object to use for the request.                                                                                                                                             |
| `opts`                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                     | The options for this request.                                                                                                                                                          |

### Response

**[*operations.C1APIAiGovernanceV1VirtualMCPToolBindingServiceCreateBindingsResponse](../../pkg/models/operations/c1apiaigovernancev1virtualmcptoolbindingservicecreatebindingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteBindings

DeleteBindings removes up to 100 individual-tool assignments from an Assigned-tools Virtual MCP Server.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.VirtualMCPToolBindingService.DeleteBindings" method="post" path="/api/v1/virtual_mcp_servers/{virtual_mcp_server_id}/tool_bindings/delete" -->
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

    res, err := s.VirtualMCPToolBinding.DeleteBindings(ctx, operations.C1APIAiGovernanceV1VirtualMCPToolBindingServiceDeleteBindingsRequest{
        VirtualMcpServerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMCPToolBindingServiceDeleteBindingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                              | Type                                                                                                                                                                                   | Required                                                                                                                                                                               | Description                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                     | The context to use for the request.                                                                                                                                                    |
| `request`                                                                                                                                                                              | [operations.C1APIAiGovernanceV1VirtualMCPToolBindingServiceDeleteBindingsRequest](../../pkg/models/operations/c1apiaigovernancev1virtualmcptoolbindingservicedeletebindingsrequest.md) | :heavy_check_mark:                                                                                                                                                                     | The request object to use for the request.                                                                                                                                             |
| `opts`                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                     | The options for this request.                                                                                                                                                          |

### Response

**[*operations.C1APIAiGovernanceV1VirtualMCPToolBindingServiceDeleteBindingsResponse](../../pkg/models/operations/c1apiaigovernancev1virtualmcptoolbindingservicedeletebindingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List returns one page of individual tools assigned directly to an Assigned-tools Virtual MCP Server.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.VirtualMCPToolBindingService.List" method="get" path="/api/v1/virtual_mcp_servers/{virtual_mcp_server_id}/tool_bindings" -->
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

    res, err := s.VirtualMCPToolBinding.List(ctx, operations.C1APIAiGovernanceV1VirtualMCPToolBindingServiceListRequest{
        VirtualMcpServerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMCPToolBindingServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APIAiGovernanceV1VirtualMCPToolBindingServiceListRequest](../../pkg/models/operations/c1apiaigovernancev1virtualmcptoolbindingservicelistrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APIAiGovernanceV1VirtualMCPToolBindingServiceListResponse](../../pkg/models/operations/c1apiaigovernancev1virtualmcptoolbindingservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |