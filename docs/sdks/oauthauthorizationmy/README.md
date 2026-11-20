# OAuthAuthorizationMy

## Overview

### Available Operations

* [Disconnect](#disconnect) - Disconnect
* [ListAvailable](#listavailable) - List Available
* [ListConnections](#listconnections) - List Connections

## Disconnect

Disconnect revokes a consent owned by the current human user.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationMyService.Disconnect" method="post" path="/api/v1/iam/my/oauth_authorization_connections/disconnect" -->
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

    res, err := s.OAuthAuthorizationMy.Disconnect(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationMyServiceDisconnectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [shared.OAuthAuthorizationMyServiceDisconnectRequest](../../pkg/models/shared/oauthauthorizationmyservicedisconnectrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.C1APIIamV1OAuthAuthorizationMyServiceDisconnectResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationmyservicedisconnectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAvailable

ListAvailable lists authorizations the current user can use.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationMyService.ListAvailable" method="get" path="/api/v1/iam/my/oauth_authorizations" -->
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

    res, err := s.OAuthAuthorizationMy.ListAvailable(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationMyServiceListAvailableResponse != nil {
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

**[*operations.C1APIIamV1OAuthAuthorizationMyServiceListAvailableResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationmyservicelistavailableresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListConnections

ListConnections lists current consents, including connections with access subsequently removed.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationMyService.ListConnections" method="get" path="/api/v1/iam/my/oauth_authorization_connections" -->
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

    res, err := s.OAuthAuthorizationMy.ListConnections(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationMyServiceListConnectionsResponse != nil {
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

**[*operations.C1APIIamV1OAuthAuthorizationMyServiceListConnectionsResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationmyservicelistconnectionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |