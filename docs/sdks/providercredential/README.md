# ProviderCredential

## Overview

### Available Operations

* [Clear](#clear) - Clear
* [Get](#get) - Get
* [Set](#set) - Set

## Clear

Clear deletes the provider credential stored in the given slot.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.llm_gateway.v1.ProviderCredentialService.Clear" method="delete" path="/api/v1/llm-gateway/provider-credentials/{slot_id}" -->
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

    res, err := s.ProviderCredential.Clear(ctx, operations.C1APILlmGatewayV1ProviderCredentialServiceClearRequest{
        SlotID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClearProviderCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APILlmGatewayV1ProviderCredentialServiceClearRequest](../../pkg/models/operations/c1apillmgatewayv1providercredentialserviceclearrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APILlmGatewayV1ProviderCredentialServiceClearResponse](../../pkg/models/operations/c1apillmgatewayv1providercredentialserviceclearresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get returns metadata for the provider credential in the given slot. The
 stored API key is never returned. Returns an empty response if no
 credential has ever been set for the slot; a cleared slot returns
 FAILED_PRECONDITION.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.llm_gateway.v1.ProviderCredentialService.Get" method="get" path="/api/v1/llm-gateway/provider-credentials/{slot_id}" -->
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

    res, err := s.ProviderCredential.Get(ctx, operations.C1APILlmGatewayV1ProviderCredentialServiceGetRequest{
        SlotID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetProviderCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APILlmGatewayV1ProviderCredentialServiceGetRequest](../../pkg/models/operations/c1apillmgatewayv1providercredentialservicegetrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APILlmGatewayV1ProviderCredentialServiceGetResponse](../../pkg/models/operations/c1apillmgatewayv1providercredentialservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Set

Set stores or replaces the provider API key used by the LLM gateway for
 the given slot. The API key is stored encrypted and is never returned by
 the API.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.llm_gateway.v1.ProviderCredentialService.Set" method="put" path="/api/v1/llm-gateway/provider-credentials/{slot_id}" -->
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

    res, err := s.ProviderCredential.Set(ctx, operations.C1APILlmGatewayV1ProviderCredentialServiceSetRequest{
        SlotID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SetProviderCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APILlmGatewayV1ProviderCredentialServiceSetRequest](../../pkg/models/operations/c1apillmgatewayv1providercredentialservicesetrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APILlmGatewayV1ProviderCredentialServiceSetResponse](../../pkg/models/operations/c1apillmgatewayv1providercredentialservicesetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |