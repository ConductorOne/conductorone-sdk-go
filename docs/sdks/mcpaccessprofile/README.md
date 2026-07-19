# MCPAccessProfile

## Overview

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [GetByAppEntitlementID](#getbyappentitlementid) - Get By App Entitlement Id
* [List](#list) - List
* [ListRequestableConnectors](#listrequestableconnectors) - List Requestable Connectors
* [SearchRequestableConnectors](#searchrequestableconnectors) - Search Requestable Connectors
* [Update](#update) - Update

## Create

Create creates a new MCP toolset (access profile) under a connector. The
 backend also provisions a backing AppEntitlement that users request in
 order to be granted the toolset's tools.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileService.Create" method="post" path="/api/v1/apps/{app_id}/connectors/{connector_id}/mcp_toolsets" -->
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

    res, err := s.MCPAccessProfile.Create(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileServiceCreateRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIAiGovernanceV1MCPAccessProfileServiceCreateRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicecreaterequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileServiceCreateResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete soft-deletes a toolset (access profile) and cascades to its tool
 bindings and backing entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileService.Delete" method="delete" path="/api/v1/apps/{app_id}/connectors/{connector_id}/mcp_toolsets/{id}" -->
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

    res, err := s.MCPAccessProfile.Delete(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileServiceDeleteRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIAiGovernanceV1MCPAccessProfileServiceDeleteRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicedeleterequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileServiceDeleteResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get retrieves a single MCP toolset (access profile) by app_id +
 connector_id + id, including its display name, description, linked
 app_entitlement_id, and bound tool count.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileService.Get" method="get" path="/api/v1/apps/{app_id}/connectors/{connector_id}/mcp_toolsets/{id}" -->
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

    res, err := s.MCPAccessProfile.Get(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileServiceGetRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APIAiGovernanceV1MCPAccessProfileServiceGetRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicegetrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileServiceGetResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetByAppEntitlementID

GetByAppEntitlementId looks up the toolset (access profile) linked to a
 synced role entitlement, by app_id + app_entitlement_id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileService.GetByAppEntitlementId" method="get" path="/api/v1/apps/{app_id}/mcp_toolsets/by_app_entitlement_id/{app_entitlement_id}" -->
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

    res, err := s.MCPAccessProfile.GetByAppEntitlementID(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileServiceGetByAppEntitlementIDRequest{
        AppEntitlementID: "<id>",
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileServiceGetByAppEntitlementIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                  | Type                                                                                                                                                                                       | Required                                                                                                                                                                                   | Description                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                         | The context to use for the request.                                                                                                                                                        |
| `request`                                                                                                                                                                                  | [operations.C1APIAiGovernanceV1MCPAccessProfileServiceGetByAppEntitlementIDRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicegetbyappentitlementidrequest.md) | :heavy_check_mark:                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                 |
| `opts`                                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                         | The options for this request.                                                                                                                                                              |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileServiceGetByAppEntitlementIDResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicegetbyappentitlementidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List returns the MCP toolsets (access profiles) defined for a single
 (app_id, connector_id), paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileService.List" method="get" path="/api/v1/apps/{app_id}/connectors/{connector_id}/mcp_toolsets" -->
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

    res, err := s.MCPAccessProfile.List(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileServiceListRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APIAiGovernanceV1MCPAccessProfileServiceListRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicelistrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileServiceListResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListRequestableConnectors

ListRequestableConnectors returns the (app_id, connector_id) pairs that
 have at least one requestable toolset (access profile) for the given user.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileService.ListRequestableConnectors" method="get" path="/api/v1/users/{user_id}/mcp_toolsets/requestable_connectors" -->
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

    res, err := s.MCPAccessProfile.ListRequestableConnectors(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileServiceListRequestableConnectorsRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileServiceListRequestableConnectorsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                          | Type                                                                                                                                                                                               | Required                                                                                                                                                                                           | Description                                                                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                |
| `request`                                                                                                                                                                                          | [operations.C1APIAiGovernanceV1MCPAccessProfileServiceListRequestableConnectorsRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicelistrequestableconnectorsrequest.md) | :heavy_check_mark:                                                                                                                                                                                 | The request object to use for the request.                                                                                                                                                         |
| `opts`                                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                 | The options for this request.                                                                                                                                                                      |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileServiceListRequestableConnectorsResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicelistrequestableconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchRequestableConnectors

SearchRequestableConnectors returns card-ready entries — one per MCP
 connector the user can browse and request tools for — with server-side
 visibility + policy filtering, grant-enrollment-status filtering, text
 search over the connector display name, and pagination. Backs the
 Requests -> AI tools connector cards. Held connectors are surfaced even
 when no longer requestable so "My tools" reflects what the user actually
 holds.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileService.SearchRequestableConnectors" method="get" path="/api/v1/users/{user_id}/mcp_toolsets/requestable_connectors/search" -->
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

    res, err := s.MCPAccessProfile.SearchRequestableConnectors(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileServiceSearchRequestableConnectorsRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileServiceSearchRequestableConnectorsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                              | Type                                                                                                                                                                                                   | Required                                                                                                                                                                                               | Description                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                    |
| `request`                                                                                                                                                                                              | [operations.C1APIAiGovernanceV1MCPAccessProfileServiceSearchRequestableConnectorsRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicesearchrequestableconnectorsrequest.md) | :heavy_check_mark:                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                             |
| `opts`                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                     | The options for this request.                                                                                                                                                                          |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileServiceSearchRequestableConnectorsResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileservicesearchrequestableconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update modifies a toolset's admin-editable fields via update_mask.
 Editable paths: display_name, description. profile must include id,
 app_id, and connector_id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.MCPAccessProfileService.Update" method="post" path="/api/v1/apps/{app_id}/connectors/{connector_id}/mcp_toolsets/{id}" -->
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

    res, err := s.MCPAccessProfile.Update(ctx, operations.C1APIAiGovernanceV1MCPAccessProfileServiceUpdateRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAccessProfileServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIAiGovernanceV1MCPAccessProfileServiceUpdateRequest](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIAiGovernanceV1MCPAccessProfileServiceUpdateResponse](../../pkg/models/operations/c1apiaigovernancev1mcpaccessprofileserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |