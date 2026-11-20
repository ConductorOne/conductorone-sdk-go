# OAuthAuthorization

## Overview

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List
* [ListConsents](#listconsents) - List Consents
* [RevokeConsent](#revokeconsent) - Revoke Consent
* [RotateSecret](#rotatesecret) - Rotate Secret
* [SearchConsents](#searchconsents) - Search Consents
* [Update](#update) - Update
* [UpdateClient](#updateclient) - Update Client

## Create

Create creates an authorization and returns its confidential client credential once.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationService.Create" method="post" path="/api/v1/iam/oauth_authorizations" -->
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

    res, err := s.OAuthAuthorization.Create(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [shared.OAuthAuthorizationServiceCreateRequest](../../pkg/models/shared/oauthauthorizationservicecreaterequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.C1APIIamV1OAuthAuthorizationServiceCreateResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete retires an authorization and its client, retaining consent audit history.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationService.Delete" method="delete" path="/api/v1/iam/oauth_authorizations/{id}" -->
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

    res, err := s.OAuthAuthorization.Delete(ctx, operations.C1APIIamV1OAuthAuthorizationServiceDeleteRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIIamV1OAuthAuthorizationServiceDeleteRequest](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicedeleterequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIIamV1OAuthAuthorizationServiceDeleteResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get retrieves an authorization and its redacted client settings.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationService.Get" method="get" path="/api/v1/iam/oauth_authorizations/{id}" -->
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

    res, err := s.OAuthAuthorization.Get(ctx, operations.C1APIIamV1OAuthAuthorizationServiceGetRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.C1APIIamV1OAuthAuthorizationServiceGetRequest](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicegetrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.C1APIIamV1OAuthAuthorizationServiceGetResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List lists authorizations in stable creation order.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationService.List" method="get" path="/api/v1/iam/oauth_authorizations" -->
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

    res, err := s.OAuthAuthorization.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationServiceListResponse != nil {
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

**[*operations.C1APIIamV1OAuthAuthorizationServiceListResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListConsents

ListConsents lists current consents for an authorization.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationService.ListConsents" method="get" path="/api/v1/iam/oauth_authorizations/{id}/consents" -->
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

    res, err := s.OAuthAuthorization.ListConsents(ctx, operations.C1APIIamV1OAuthAuthorizationServiceListConsentsRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationServiceListConsentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIIamV1OAuthAuthorizationServiceListConsentsRequest](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicelistconsentsrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIIamV1OAuthAuthorizationServiceListConsentsResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicelistconsentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RevokeConsent

RevokeConsent revokes an authorization's canonical consent.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationService.RevokeConsent" method="post" path="/api/v1/iam/oauth_authorizations/{id}/revoke-consent" -->
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

    res, err := s.OAuthAuthorization.RevokeConsent(ctx, operations.C1APIIamV1OAuthAuthorizationServiceRevokeConsentRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationServiceRevokeConsentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIIamV1OAuthAuthorizationServiceRevokeConsentRequest](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicerevokeconsentrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIIamV1OAuthAuthorizationServiceRevokeConsentResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicerevokeconsentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RotateSecret

Rotate the shared secret for a client_secret_basic or client_secret_post client and return it once.
 For private_key_jwt clients, replace public_jwks through UpdateClient instead.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationService.RotateSecret" method="post" path="/api/v1/iam/oauth_authorizations/{id}/rotate-secret" -->
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

    res, err := s.OAuthAuthorization.RotateSecret(ctx, operations.C1APIIamV1OAuthAuthorizationServiceRotateSecretRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationServiceRotateSecretResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIIamV1OAuthAuthorizationServiceRotateSecretRequest](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicerotatesecretrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIIamV1OAuthAuthorizationServiceRotateSecretResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicerotatesecretresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchConsents

Search current consents across managed OAuth authorizations, including disabled
 authorizations and consents whose users no longer have authorization access.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationService.SearchConsents" method="get" path="/api/v1/iam/oauth_consents" -->
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

    res, err := s.OAuthAuthorization.SearchConsents(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationServiceSearchConsentsResponse != nil {
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

**[*operations.C1APIIamV1OAuthAuthorizationServiceSearchConsentsResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationservicesearchconsentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update updates the explicitly selected mutable authorization fields.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationService.Update" method="post" path="/api/v1/iam/oauth_authorizations/update" -->
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

    res, err := s.OAuthAuthorization.Update(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [shared.OAuthAuthorizationServiceUpdateRequest](../../pkg/models/shared/oauthauthorizationserviceupdaterequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.C1APIIamV1OAuthAuthorizationServiceUpdateResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateClient

UpdateClient updates registered redirects or public verification keys.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.OAuthAuthorizationService.UpdateClient" method="post" path="/api/v1/iam/oauth_authorizations/{id}/update-client" -->
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

    res, err := s.OAuthAuthorization.UpdateClient(ctx, operations.C1APIIamV1OAuthAuthorizationServiceUpdateClientRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAuthorizationServiceUpdateClientResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIIamV1OAuthAuthorizationServiceUpdateClientRequest](../../pkg/models/operations/c1apiiamv1oauthauthorizationserviceupdateclientrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIIamV1OAuthAuthorizationServiceUpdateClientResponse](../../pkg/models/operations/c1apiiamv1oauthauthorizationserviceupdateclientresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |