# MCPServer

## Overview

### Available Operations

* [Delete](#delete) - Delete
* [DiscoverOIDCEndpoints](#discoveroidcendpoints) - Discover Oidc Endpoints
* [Get](#get) - Get
* [GetCatalog](#getcatalog) - Get Catalog
* [List](#list) - List
* [ListCatalog](#listcatalog) - List Catalog
* [ListConnections](#listconnections) - List Connections
* [Register](#register) - Register
* [ResyncTools](#resynctools) - Resync Tools
* [SearchWithToolCount](#searchwithtoolcount) - Search With Tool Count
* [TestConnection](#testconnection) - Test Connection
* [Update](#update) - Update
* [UpdateCredentials](#updatecredentials) - Update Credentials

## Delete

Delete an MCP server. Its connector stops and it is soft-deleted from this
 tenant's MCP catalog, and any per-user credentials issued against it are
 revoked.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.Delete" method="delete" path="/api/v1/apps/{app_id}/mcp_servers/{connector_id}" -->
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

    res, err := s.MCPServer.Delete(ctx, operations.C1APIAiGovernanceV1MCPServerServiceDeleteRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAiGovernanceV1MCPServerServiceDeleteRequest](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicedeleterequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceDeleteResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DiscoverOIDCEndpoints

DiscoverOIDCEndpoints fetches the OpenID Connect discovery document for an issuer
 and returns the authorization, token, and supported scopes.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.DiscoverOIDCEndpoints" method="post" path="/api/v1/mcp_servers/discover_oidc" -->
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

    res, err := s.MCPServer.DiscoverOIDCEndpoints(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceDiscoverOIDCEndpointsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [shared.MCPServerServiceDiscoverOIDCEndpointsRequest](../../pkg/models/shared/mcpserverservicediscoveroidcendpointsrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceDiscoverOIDCEndpointsResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicediscoveroidcendpointsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get retrieves a single MCP server.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.Get" method="get" path="/api/v1/apps/{app_id}/mcp_servers/{connector_id}" -->
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

    res, err := s.MCPServer.Get(ctx, operations.C1APIAiGovernanceV1MCPServerServiceGetRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.C1APIAiGovernanceV1MCPServerServiceGetRequest](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicegetrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceGetResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCatalog

GetCatalog retrieves a single MCP server catalog entry by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.GetCatalog" method="get" path="/api/v1/mcp_server_catalog/{catalog_id}" -->
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

    res, err := s.MCPServer.GetCatalog(ctx, operations.C1APIAiGovernanceV1MCPServerServiceGetCatalogRequest{
        CatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceGetCatalogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APIAiGovernanceV1MCPServerServiceGetCatalogRequest](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicegetcatalogrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceGetCatalogResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicegetcatalogresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List retrieves MCP servers, optionally narrowed to an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.List" method="get" path="/api/v1/apps/{app_id}/mcp_servers" -->
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

    res, err := s.MCPServer.List(ctx, operations.C1APIAiGovernanceV1MCPServerServiceListRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.C1APIAiGovernanceV1MCPServerServiceListRequest](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicelistrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `opts`                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceListResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCatalog

ListCatalog returns all available MCP server catalog entries.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.ListCatalog" method="get" path="/api/v1/mcp_server_catalog" -->
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

    res, err := s.MCPServer.ListCatalog(ctx, operations.C1APIAiGovernanceV1MCPServerServiceListCatalogRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceListCatalogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APIAiGovernanceV1MCPServerServiceListCatalogRequest](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicelistcatalogrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceListCatalogResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicelistcatalogresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListConnections

ListConnections returns per-user MCP servers the calling user can
 connect to, filtered to apps where the user has an account. Covers
 OAuth2 authorization-code passthrough as well as bearer-token /
 custom-header per-user methods. Includes per-user connection status.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.ListConnections" method="get" path="/api/v1/mcp_server_connections" -->
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

    res, err := s.MCPServer.ListConnections(ctx, operations.C1APIAiGovernanceV1MCPServerServiceListConnectionsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceListConnectionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.C1APIAiGovernanceV1MCPServerServiceListConnectionsRequest](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicelistconnectionsrequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |
| `opts`                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceListConnectionsResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicelistconnectionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Register

Register a new MCP server under an application. Set server_type to HOSTED
 (C1 runs a catalog integration) or EXTERNAL (a third-party MCP server you
 point C1 at by URL). Auth credentials are validated and stored securely.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.Register" method="post" path="/api/v1/apps/{app_id}/mcp_servers" -->
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

    res, err := s.MCPServer.Register(ctx, operations.C1APIAiGovernanceV1MCPServerServiceRegisterRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceRegisterResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIAiGovernanceV1MCPServerServiceRegisterRequest](../../pkg/models/operations/c1apiaigovernancev1mcpserverserviceregisterrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceRegisterResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverserviceregisterresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ResyncTools

ResyncTools re-runs per-identity tool discovery for the calling user's
 own credential on a per-user MCP server, so a session opened before the
 user connected (or after their visible tools changed) doesn't have to
 wait for the next unrelated MCPTool/AppEntitlementUserBinding change to
 pick it up.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.ResyncTools" method="post" path="/api/v1/apps/{app_id}/mcp_servers/{connector_id}/resync_tools" -->
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

    res, err := s.MCPServer.ResyncTools(ctx, operations.C1APIAiGovernanceV1MCPServerServiceResyncToolsRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceResyncToolsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APIAiGovernanceV1MCPServerServiceResyncToolsRequest](../../pkg/models/operations/c1apiaigovernancev1mcpserverserviceresynctoolsrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceResyncToolsResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverserviceresynctoolsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchWithToolCount

SearchWithToolCount searches MCP servers with filtering and returns per-server tool state counts.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.SearchWithToolCount" method="post" path="/api/v1/apps/{app_id}/mcp_servers/search" -->
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

    res, err := s.MCPServer.SearchWithToolCount(ctx, operations.C1APIAiGovernanceV1MCPServerServiceSearchWithToolCountRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceSearchWithToolCountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [operations.C1APIAiGovernanceV1MCPServerServiceSearchWithToolCountRequest](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicesearchwithtoolcountrequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |
| `opts`                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                                       | The options for this request.                                                                                                                                            |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceSearchWithToolCountResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicesearchwithtoolcountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## TestConnection

TestConnection probes an MCP server with the
 supplied URL + transport + plaintext credentials and reports
 whether a real MCP initialize + tools/list succeeded.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.TestConnection" method="post" path="/api/v1/mcp_servers/test_connection" -->
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

    res, err := s.MCPServer.TestConnection(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceTestConnectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [shared.MCPServerServiceTestConnectionRequest](../../pkg/models/shared/mcpserverservicetestconnectionrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceTestConnectionResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverservicetestconnectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update modifies an existing MCP server's editable fields.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.Update" method="post" path="/api/v1/apps/{app_id}/mcp_servers/{connector_id}" -->
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

    res, err := s.MCPServer.Update(ctx, operations.C1APIAiGovernanceV1MCPServerServiceUpdateRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAiGovernanceV1MCPServerServiceUpdateRequest](../../pkg/models/operations/c1apiaigovernancev1mcpserverserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceUpdateResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCredentials

UpdateCredentials replaces the auth config and/or config fields for an MCP server.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPServerService.UpdateCredentials" method="post" path="/api/v1/apps/{app_id}/mcp_servers/{connector_id}/credentials" -->
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

    res, err := s.MCPServer.UpdateCredentials(ctx, operations.C1APIAiGovernanceV1MCPServerServiceUpdateCredentialsRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerServiceUpdateCredentialsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.C1APIAiGovernanceV1MCPServerServiceUpdateCredentialsRequest](../../pkg/models/operations/c1apiaigovernancev1mcpserverserviceupdatecredentialsrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.C1APIAiGovernanceV1MCPServerServiceUpdateCredentialsResponse](../../pkg/models/operations/c1apiaigovernancev1mcpserverserviceupdatecredentialsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |