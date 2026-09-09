# GatewayKey

## Overview

### Available Operations

* [List](#list) - List
* [Mint](#mint) - Mint
* [Revoke](#revoke) - Revoke

## List

List returns the tenant's LLM gateway API keys. Only key metadata and a
 key prefix are returned, never the full key.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.llm_gateway.v1.GatewayKeyService.List" method="get" path="/api/v1/llm-gateway/keys" -->
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

    res, err := s.GatewayKey.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListGatewayKeysResponse != nil {
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

**[*operations.C1APILlmGatewayV1GatewayKeyServiceListResponse](../../pkg/models/operations/c1apillmgatewayv1gatewaykeyservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Mint

Mint creates a new LLM gateway API key. The key value is shown only in
 this response and cannot be retrieved again; store it immediately.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.llm_gateway.v1.GatewayKeyService.Mint" method="post" path="/api/v1/llm-gateway/keys" -->
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

    res, err := s.GatewayKey.Mint(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MintGatewayKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.MintGatewayKeyRequest](../../pkg/models/shared/mintgatewaykeyrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.C1APILlmGatewayV1GatewayKeyServiceMintResponse](../../pkg/models/operations/c1apillmgatewayv1gatewaykeyservicemintresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Revoke

Revoke revokes an LLM gateway API key by ID. The key immediately stops
 authenticating gateway requests.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.llm_gateway.v1.GatewayKeyService.Revoke" method="delete" path="/api/v1/llm-gateway/keys/{id}" -->
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

    res, err := s.GatewayKey.Revoke(ctx, operations.C1APILlmGatewayV1GatewayKeyServiceRevokeRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RevokeGatewayKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APILlmGatewayV1GatewayKeyServiceRevokeRequest](../../pkg/models/operations/c1apillmgatewayv1gatewaykeyservicerevokerequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APILlmGatewayV1GatewayKeyServiceRevokeResponse](../../pkg/models/operations/c1apillmgatewayv1gatewaykeyservicerevokeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |