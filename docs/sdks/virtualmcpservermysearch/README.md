# VirtualMCPServerMySearch

## Overview

### Available Operations

* [Search](#search) - NOTE: Only shows Virtual MCP Servers available to the current user.

## Search

Search live Virtual MCP Servers for which the caller's own principal has
 an active `use` grant. Results are ordered by display name and server ID
 and may briefly lag writes because search reads the PostgreSQL projection;
 connection authorization always rechecks authoritative state.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.VirtualMCPServerMySearchService.Search" method="post" path="/api/v1/search/my_virtual_mcp_servers" -->
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

    res, err := s.VirtualMCPServerMySearch.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMCPServerMySearchServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [shared.VirtualMCPServerMySearchServiceSearchRequest](../../pkg/models/shared/virtualmcpservermysearchservicesearchrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.C1APIAiGovernanceV1VirtualMCPServerMySearchServiceSearchResponse](../../pkg/models/operations/c1apiaigovernancev1virtualmcpservermysearchservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |