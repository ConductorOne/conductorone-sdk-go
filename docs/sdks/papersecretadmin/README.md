# PaperSecretAdmin

## Overview

### Available Operations

* [Get](#get) - Get
* [Revoke](#revoke) - Revoke
* [Search](#search) - Search
* [SearchAuditEvents](#searchauditevents) - Search Audit Events

## Get

Get retrieves any secret's metadata by vault ID (admin override).

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretAdminService.Get" method="get" path="/api/v1/secrets_admin/{vault_id}" -->
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

    res, err := s.PaperSecretAdmin.Get(ctx, operations.C1APISecretsV1PaperSecretAdminServiceGetRequest{
        VaultID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretAdminServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APISecretsV1PaperSecretAdminServiceGetRequest](../../pkg/models/operations/c1apisecretsv1papersecretadminservicegetrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APISecretsV1PaperSecretAdminServiceGetResponse](../../pkg/models/operations/c1apisecretsv1papersecretadminservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Revoke

Revoke allows admin to revoke any secret (not just their own).

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretAdminService.Revoke" method="delete" path="/api/v1/secrets_admin/{vault_id}" -->
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

    res, err := s.PaperSecretAdmin.Revoke(ctx, operations.C1APISecretsV1PaperSecretAdminServiceRevokeRequest{
        VaultID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretAdminServiceRevokeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APISecretsV1PaperSecretAdminServiceRevokeRequest](../../pkg/models/operations/c1apisecretsv1papersecretadminservicerevokerequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APISecretsV1PaperSecretAdminServiceRevokeResponse](../../pkg/models/operations/c1apisecretsv1papersecretadminservicerevokeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Search

Search returns secrets across the tenant.
 Can filter by creator, sharing mode, status, time range, etc.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretAdminService.Search" method="post" path="/api/v1/search/secrets_admin" -->
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

    res, err := s.PaperSecretAdmin.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretAdminServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [shared.PaperSecretAdminServiceSearchRequest](../../pkg/models/shared/papersecretadminservicesearchrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.C1APISecretsV1PaperSecretAdminServiceSearchResponse](../../pkg/models/operations/c1apisecretsv1papersecretadminservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchAuditEvents

SearchAuditEvents returns audit events for paper secrets.
 Can filter by vault_id, actor (user ID or email), client IP.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.secrets.v1.PaperSecretAdminService.SearchAuditEvents" method="post" path="/api/v1/search/secrets_admin/audit_events" -->
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

    res, err := s.PaperSecretAdmin.SearchAuditEvents(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaperSecretAdminServiceSearchAuditEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [shared.PaperSecretAdminServiceSearchAuditEventsRequest](../../pkg/models/shared/papersecretadminservicesearchauditeventsrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.C1APISecretsV1PaperSecretAdminServiceSearchAuditEventsResponse](../../pkg/models/operations/c1apisecretsv1papersecretadminservicesearchauditeventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |