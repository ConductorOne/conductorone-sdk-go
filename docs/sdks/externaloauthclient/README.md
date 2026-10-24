# ExternalOAuthClient

## Overview

### Available Operations

* [Create](#create) - Create
* [Revoke](#revoke) - Revoke
* [RotateSecret](#rotatesecret) - Rotate Secret
* [Update](#update) - Update

## Create

Create provisions a durable static OAuth client for a pre-registered
 consumer, eligible for authorization_code + refresh_token only. A
 secret-based client's secret is returned once, on this response.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.ExternalOAuthClientService.Create" method="post" path="/api/v1/iam/external_oauth_clients" -->
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

    res, err := s.ExternalOAuthClient.Create(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ExternalOAuthClientServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [shared.ExternalOAuthClientServiceCreateRequest](../../pkg/models/shared/externaloauthclientservicecreaterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIIamV1ExternalOAuthClientServiceCreateResponse](../../pkg/models/operations/c1apiiamv1externaloauthclientservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Revoke

Revoke disables a static client, then revokes its users' grants. The row is
 kept for audit.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.ExternalOAuthClientService.Revoke" method="post" path="/api/v1/iam/external_oauth_clients/revoke" -->
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

    res, err := s.ExternalOAuthClient.Revoke(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ExternalOAuthClientServiceRevokeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [shared.ExternalOAuthClientServiceRevokeRequest](../../pkg/models/shared/externaloauthclientservicerevokerequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIIamV1ExternalOAuthClientServiceRevokeResponse](../../pkg/models/operations/c1apiiamv1externaloauthclientservicerevokeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RotateSecret

RotateSecret issues a new client secret for a static client and returns it
 once. The client keeps its client_id, redirect URIs, and grants.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.ExternalOAuthClientService.RotateSecret" method="post" path="/api/v1/iam/external_oauth_clients/rotate-secret" -->
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

    res, err := s.ExternalOAuthClient.RotateSecret(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ExternalOAuthClientServiceRotateSecretResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [shared.ExternalOAuthClientServiceRotateSecretRequest](../../pkg/models/shared/externaloauthclientservicerotatesecretrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APIIamV1ExternalOAuthClientServiceRotateSecretResponse](../../pkg/models/operations/c1apiiamv1externaloauthclientservicerotatesecretresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update edits a static client through an explicit field mask. client_id,
 client_id_type, and the auth method are immutable. Update never mints,
 clears, or returns secret material.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.ExternalOAuthClientService.Update" method="post" path="/api/v1/iam/external_oauth_clients/update" -->
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

    res, err := s.ExternalOAuthClient.Update(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ExternalOAuthClientServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [shared.ExternalOAuthClientServiceUpdateRequest](../../pkg/models/shared/externaloauthclientserviceupdaterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIIamV1ExternalOAuthClientServiceUpdateResponse](../../pkg/models/operations/c1apiiamv1externaloauthclientserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |