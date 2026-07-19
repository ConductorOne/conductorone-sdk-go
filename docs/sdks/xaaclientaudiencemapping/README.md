# XAAClientAudienceMapping

## Overview

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List
* [ListHistory](#listhistory) - List History
* [Search](#search) - Search
* [Update](#update) - Update

## Create

Create a client audience mapping under a resource server.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAClientAudienceMappingService.Create" method="post" path="/api/v1/xaa/resource_servers/{xaa_resource_server_id}/client_audience_mappings" -->
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

    res, err := s.XAAClientAudienceMapping.Create(ctx, operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceCreateRequest{
        XaaResourceServerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAClientAudienceMappingServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `request`                                                                                                                                                                        | [operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceCreateRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingservicecreaterequest.md) | :heavy_check_mark:                                                                                                                                                               | The request object to use for the request.                                                                                                                                       |
| `opts`                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                     | :heavy_minus_sign:                                                                                                                                                               | The options for this request.                                                                                                                                                    |

### Response

**[*operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceCreateResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a mapping (soft delete). Uses a POST .../delete action route
 because the client key travels in the request body, which HTTP DELETE does
 not reliably support.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAClientAudienceMappingService.Delete" method="post" path="/api/v1/xaa/resource_servers/{xaa_resource_server_id}/client_audience_mappings/delete" -->
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

    res, err := s.XAAClientAudienceMapping.Delete(ctx, operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceDeleteRequest{
        XaaResourceServerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAClientAudienceMappingServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `request`                                                                                                                                                                        | [operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceDeleteRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingservicedeleterequest.md) | :heavy_check_mark:                                                                                                                                                               | The request object to use for the request.                                                                                                                                       |
| `opts`                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                     | :heavy_minus_sign:                                                                                                                                                               | The options for this request.                                                                                                                                                    |

### Response

**[*operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceDeleteResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get a client audience mapping by resource server + client key. The client
 key is passed as a query parameter because it may be a URL.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAClientAudienceMappingService.Get" method="get" path="/api/v1/xaa/resource_servers/{xaa_resource_server_id}/client_audience_mappings/by_client_key" -->
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

    res, err := s.XAAClientAudienceMapping.Get(ctx, operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceGetRequest{
        XaaResourceServerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAClientAudienceMappingServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceGetRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingservicegetrequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |
| `opts`                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                         | The options for this request.                                                                                                                                              |

### Response

**[*operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceGetResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List the client audience mappings for a resource server, one page at a
 time.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAClientAudienceMappingService.List" method="get" path="/api/v1/xaa/resource_servers/{xaa_resource_server_id}/client_audience_mappings" -->
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

    res, err := s.XAAClientAudienceMapping.List(ctx, operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceListRequest{
        XaaResourceServerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAClientAudienceMappingServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceListRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingservicelistrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |
| `opts`                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                           | The options for this request.                                                                                                                                                |

### Response

**[*operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceListResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHistory

ListHistory returns the change history (newest first) for a single client
 audience mapping — each entry is a snapshot plus who/when metadata. The
 client key is passed as a query parameter because it may be a URL.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAClientAudienceMappingService.ListHistory" method="get" path="/api/v1/xaa/resource_servers/{xaa_resource_server_id}/client_audience_mappings/history" -->
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

    res, err := s.XAAClientAudienceMapping.ListHistory(ctx, operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceListHistoryRequest{
        XaaResourceServerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAClientAudienceMappingServiceListHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                  | Type                                                                                                                                                                                       | Required                                                                                                                                                                                   | Description                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                         | The context to use for the request.                                                                                                                                                        |
| `request`                                                                                                                                                                                  | [operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceListHistoryRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingservicelisthistoryrequest.md) | :heavy_check_mark:                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                 |
| `opts`                                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                         | The options for this request.                                                                                                                                                              |

### Response

**[*operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceListHistoryResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingservicelisthistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Search

Search client audience mappings across the tenant, filtered by resource
 server, disabled state, or text query, or fetch a specific set by ref.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAClientAudienceMappingService.Search" method="post" path="/api/v1/search/xaa/client_audience_mappings" -->
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

    res, err := s.XAAClientAudienceMapping.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAClientAudienceMappingServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [shared.XAAClientAudienceMappingServiceSearchRequest](../../pkg/models/shared/xaaclientaudiencemappingservicesearchrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceSearchResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update a mapping's editable fields via update_mask. The resource server
 and client key identify the mapping and are immutable. Editable paths:
 audience_client_id, disabled.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.cross_app_access.v1.XAAClientAudienceMappingService.Update" method="post" path="/api/v1/xaa/resource_servers/{client_audience_mapping_xaa_resource_server_id}/client_audience_mappings/update" -->
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

    res, err := s.XAAClientAudienceMapping.Update(ctx, operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceUpdateRequest{
        ClientAudienceMappingXaaResourceServerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.XAAClientAudienceMappingServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `request`                                                                                                                                                                        | [operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceUpdateRequest](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                                                               | The request object to use for the request.                                                                                                                                       |
| `opts`                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                     | :heavy_minus_sign:                                                                                                                                                               | The options for this request.                                                                                                                                                    |

### Response

**[*operations.C1APICrossAppAccessV1XAAClientAudienceMappingServiceUpdateResponse](../../pkg/models/operations/c1apicrossappaccessv1xaaclientaudiencemappingserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |