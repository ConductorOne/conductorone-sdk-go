# PaperSecret

## Overview

### Available Operations

* [CreateExternal](#createexternal) - Create External
* [CreateInternal](#createinternal) - Create Internal
* [Get](#get) - Get
* [GetByShareCode](#getbysharecode) - Get By Share Code
* [GetContent](#getcontent) - Get Content
* [Revoke](#revoke) - Revoke
* [SearchAuditEvents](#searchauditevents) - Search Audit Events
* [SearchMySecrets](#searchmysecrets) - Search My Secrets
* [SetTextContent](#settextcontent) - Set Text Content

## CreateExternal

CreateExternal creates a secret vault for external email recipients.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretService.CreateExternal" method="post" path="/api/v1/secrets/external" -->
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

    res, err := s.PaperSecret.CreateExternal(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [shared.PaperSecretServiceCreateExternalRequest](../../pkg/models/shared/papersecretservicecreateexternalrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APISecretsV1PaperSecretServiceCreateExternalResponse](../../pkg/models/operations/c1apisecretsv1papersecretservicecreateexternalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateInternal

CreateInternal creates a secret vault for internal C1 users.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretService.CreateInternal" method="post" path="/api/v1/secrets/internal" -->
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

    res, err := s.PaperSecret.CreateInternal(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [shared.PaperSecretServiceCreateInternalRequest](../../pkg/models/shared/papersecretservicecreateinternalrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APISecretsV1PaperSecretServiceCreateInternalResponse](../../pkg/models/operations/c1apisecretsv1papersecretservicecreateinternalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get retrieves a secret's metadata by vault ID.
 Creator can always get their own secrets. Admins can get any secret.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretService.Get" method="get" path="/api/v1/secrets/{vault_id}" -->
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

    res, err := s.PaperSecret.Get(ctx, operations.C1APISecretsV1PaperSecretServiceGetRequest{
        VaultID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.C1APISecretsV1PaperSecretServiceGetRequest](../../pkg/models/operations/c1apisecretsv1papersecretservicegetrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.C1APISecretsV1PaperSecretServiceGetResponse](../../pkg/models/operations/c1apisecretsv1papersecretservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetByShareCode

GetByShareCode retrieves a secret by its human-friendly share code.
 Share codes are in XXXX-XXXX-XXXX format and are used in share URLs.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretService.GetByShareCode" method="get" path="/api/v1/secrets/code/{share_code}" -->
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

    res, err := s.PaperSecret.GetByShareCode(ctx, operations.C1APISecretsV1PaperSecretServiceGetByShareCodeRequest{
        ShareCode: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APISecretsV1PaperSecretServiceGetByShareCodeRequest](../../pkg/models/operations/c1apisecretsv1papersecretservicegetbysharecoderequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APISecretsV1PaperSecretServiceGetByShareCodeResponse](../../pkg/models/operations/c1apisecretsv1papersecretservicegetbysharecoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetContent

GetContent retrieves the encrypted secret content for an authorized recipient.
 Caller must be in the secret's allowed_user_ids list (for INTERNAL secrets).
 Returns content re-encrypted to caller's ephemeral public key.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretService.GetContent" method="post" path="/api/v1/secrets/{vault_id}/view" -->
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

    res, err := s.PaperSecret.GetContent(ctx, operations.C1APISecretsV1PaperSecretServiceGetContentRequest{
        VaultID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretServiceGetContentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.C1APISecretsV1PaperSecretServiceGetContentRequest](../../pkg/models/operations/c1apisecretsv1papersecretservicegetcontentrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.C1APISecretsV1PaperSecretServiceGetContentResponse](../../pkg/models/operations/c1apisecretsv1papersecretservicegetcontentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Revoke

Revoke soft-deletes a secret (sets Vault.deleted_at, deletes content).

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretService.Revoke" method="delete" path="/api/v1/secrets/{vault_id}" -->
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

    res, err := s.PaperSecret.Revoke(ctx, operations.C1APISecretsV1PaperSecretServiceRevokeRequest{
        VaultID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretServiceRevokeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.C1APISecretsV1PaperSecretServiceRevokeRequest](../../pkg/models/operations/c1apisecretsv1papersecretservicerevokerequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.C1APISecretsV1PaperSecretServiceRevokeResponse](../../pkg/models/operations/c1apisecretsv1papersecretservicerevokeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchAuditEvents

SearchAuditEvents returns audit events for a secret owned by the calling user.
 Returns sanitized OCSF events (IP addresses stripped for non-admin consumption).

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretService.SearchAuditEvents" method="post" path="/api/v1/search/secrets/audit_events" -->
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

    res, err := s.PaperSecret.SearchAuditEvents(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretServiceSearchAuditEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [shared.PaperSecretServiceSearchAuditEventsRequest](../../pkg/models/shared/papersecretservicesearchauditeventsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.C1APISecretsV1PaperSecretServiceSearchAuditEventsResponse](../../pkg/models/operations/c1apisecretsv1papersecretservicesearchauditeventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchMySecrets

SearchMySecrets returns secrets created by the current user.
 Automatically scoped to current user - no user_id filter parameter.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretService.SearchMySecrets" method="post" path="/api/v1/search/secrets/mine" -->
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

    res, err := s.PaperSecret.SearchMySecrets(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [shared.PaperSecretServiceSearchMySecretsRequest](../../pkg/models/shared/papersecretservicesearchmysecretsrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.C1APISecretsV1PaperSecretServiceSearchMySecretsResponse](../../pkg/models/operations/c1apisecretsv1papersecretservicesearchmysecretsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetTextContent

SetTextContent sets the encrypted content for a text secret.
 Client encrypts content using age_recipient from CreateResponse.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretService.SetTextContent" method="post" path="/api/v1/secrets/{vault_id}/content" -->
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

    res, err := s.PaperSecret.SetTextContent(ctx, operations.C1APISecretsV1PaperSecretServiceSetTextContentRequest{
        VaultID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretServiceSetTextContentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APISecretsV1PaperSecretServiceSetTextContentRequest](../../pkg/models/operations/c1apisecretsv1papersecretservicesettextcontentrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APISecretsV1PaperSecretServiceSetTextContentResponse](../../pkg/models/operations/c1apisecretsv1papersecretservicesettextcontentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |