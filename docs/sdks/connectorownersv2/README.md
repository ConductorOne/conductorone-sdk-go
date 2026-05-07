# ConnectorOwnersV2

## Overview

### Available Operations

* [SearchEntitlementOwners](#searchentitlementowners) - Search Entitlement Owners
* [SearchUserOwners](#searchuserowners) - Search User Owners
* [Set](#set) - Set

## SearchEntitlementOwners

SearchEntitlementOwners searches for the entitlement ownership for a connector.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.ConnectorOwners.SearchEntitlementOwners" method="get" path="/api/v2/apps/{app_id}/connectors/{connector_id}/owners/entitlements" -->
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

    res, err := s.ConnectorOwnersV2.SearchEntitlementOwners(ctx, operations.C1APIAppV2ConnectorOwnersSearchEntitlementOwnersRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchConnectorEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIAppV2ConnectorOwnersSearchEntitlementOwnersRequest](../../pkg/models/operations/c1apiappv2connectorownerssearchentitlementownersrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIAppV2ConnectorOwnersSearchEntitlementOwnersResponse](../../pkg/models/operations/c1apiappv2connectorownerssearchentitlementownersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchUserOwners

SearchUserOwners searches for users who are owners of this connector.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.ConnectorOwners.SearchUserOwners" method="get" path="/api/v2/apps/{app_id}/connectors/{connector_id}/owners/users" -->
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

    res, err := s.ConnectorOwnersV2.SearchUserOwners(ctx, operations.C1APIAppV2ConnectorOwnersSearchUserOwnersRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchConnectorUserOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAppV2ConnectorOwnersSearchUserOwnersRequest](../../pkg/models/operations/c1apiappv2connectorownerssearchuserownersrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIAppV2ConnectorOwnersSearchUserOwnersResponse](../../pkg/models/operations/c1apiappv2connectorownerssearchuserownersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Set

Set replaces all owners for a given connector and role.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v2.ConnectorOwners.Set" method="put" path="/api/v2/apps/{app_id}/connectors/{connector_id}/owners" -->
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

    res, err := s.ConnectorOwnersV2.Set(ctx, operations.C1APIAppV2ConnectorOwnersSetRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SetConnectorOwnersV2Response != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.C1APIAppV2ConnectorOwnersSetRequest](../../pkg/models/operations/c1apiappv2connectorownerssetrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIAppV2ConnectorOwnersSetResponse](../../pkg/models/operations/c1apiappv2connectorownerssetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |