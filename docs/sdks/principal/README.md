# Principal

## Overview

### Available Operations

* [AddBinding](#addbinding) - Add Binding
* [Create](#create) - Create
* [CreateCredential](#createcredential) - Create Credential
* [Delete](#delete) - Delete
* [DeleteBinding](#deletebinding) - Delete Binding
* [Get](#get) - Get
* [GetCredential](#getcredential) - Get Credential
* [List](#list) - List
* [ListBindings](#listbindings) - List Bindings
* [ListCredentials](#listcredentials) - List Credentials
* [RevokeCredential](#revokecredential) - Revoke Credential
* [Update](#update) - Update
* [UpdateCredential](#updatecredential) - Update Credential

## AddBinding

AddBinding links a tenant-scoped subject (a function today; future kinds
 tomorrow) to a service principal. Outbound c1-api calls made on the
 subject's behalf can then be minted as user:<service_principal_id> via
 an RFC 8693 token-exchange (act-as) flow. Many-aware: a subject may
 hold multiple bindings at the storage layer. Idempotent on
 (subject, service_principal_id) — adds the row if missing,
 resurrects it if soft-deleted, no-op if already active. Consumers
 that need 0-or-1 cardinality (Functions today) enforce it
 client-side via ListBindings + DeleteBinding. Requires the
 SERVICE_PRINCIPALS feature flag.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.AddBinding" method="post" path="/api/v1/service_principals/bindings" -->
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

    res, err := s.Principal.AddBinding(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceAddBindingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [shared.ServicePrincipalServiceAddBindingRequest](../../pkg/models/shared/serviceprincipalserviceaddbindingrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceAddBindingResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalserviceaddbindingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Create

Create creates a new service principal.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.Create" method="post" path="/api/v1/service_principals" -->
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

    res, err := s.Principal.Create(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [shared.ServicePrincipalServiceCreateRequest](../../pkg/models/shared/serviceprincipalservicecreaterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceCreateResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCredential

CreateCredential creates a new client credential for a service principal.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.CreateCredential" method="post" path="/api/v1/service_principals/{service_principal_id}/credentials" -->
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

    res, err := s.Principal.CreateCredential(ctx, operations.C1APIServicePrincipalV1ServicePrincipalServiceCreateCredentialRequest{
        ServicePrincipalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceCreateCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                | Type                                                                                                                                                                                     | Required                                                                                                                                                                                 | Description                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                       | The context to use for the request.                                                                                                                                                      |
| `request`                                                                                                                                                                                | [operations.C1APIServicePrincipalV1ServicePrincipalServiceCreateCredentialRequest](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicecreatecredentialrequest.md) | :heavy_check_mark:                                                                                                                                                                       | The request object to use for the request.                                                                                                                                               |
| `opts`                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                       | The options for this request.                                                                                                                                                            |

### Response

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceCreateCredentialResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicecreatecredentialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete deletes a service principal and all its credentials.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.Delete" method="delete" path="/api/v1/service_principals/{id}" -->
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

    res, err := s.Principal.Delete(ctx, operations.C1APIServicePrincipalV1ServicePrincipalServiceDeleteRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.C1APIServicePrincipalV1ServicePrincipalServiceDeleteRequest](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicedeleterequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceDeleteResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteBinding

DeleteBinding removes a single (subject, service_principal_id) binding
 row. At-most-one delete — does not touch other bindings the subject
 may hold against different service principals. Idempotent — succeeds
 even if no matching row exists.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.DeleteBinding" method="post" path="/api/v1/service_principals/bindings/delete" -->
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

    res, err := s.Principal.DeleteBinding(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceDeleteBindingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [shared.ServicePrincipalServiceDeleteBindingRequest](../../pkg/models/shared/serviceprincipalservicedeletebindingrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceDeleteBindingResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicedeletebindingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get returns a service principal by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.Get" method="get" path="/api/v1/service_principals/{id}" -->
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

    res, err := s.Principal.Get(ctx, operations.C1APIServicePrincipalV1ServicePrincipalServiceGetRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.C1APIServicePrincipalV1ServicePrincipalServiceGetRequest](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicegetrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |
| `opts`                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                   | :heavy_minus_sign:                                                                                                                                             | The options for this request.                                                                                                                                  |

### Response

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceGetResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCredential

GetCredential returns a single client credential for a service principal.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.GetCredential" method="get" path="/api/v1/service_principals/{service_principal_id}/credentials/{id}" -->
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

    res, err := s.Principal.GetCredential(ctx, operations.C1APIServicePrincipalV1ServicePrincipalServiceGetCredentialRequest{
        ID: "<id>",
        ServicePrincipalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceGetCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                          | Type                                                                                                                                                                               | Required                                                                                                                                                                           | Description                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                              | :heavy_check_mark:                                                                                                                                                                 | The context to use for the request.                                                                                                                                                |
| `request`                                                                                                                                                                          | [operations.C1APIServicePrincipalV1ServicePrincipalServiceGetCredentialRequest](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicegetcredentialrequest.md) | :heavy_check_mark:                                                                                                                                                                 | The request object to use for the request.                                                                                                                                         |
| `opts`                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                 | The options for this request.                                                                                                                                                      |

### Response

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceGetCredentialResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicegetcredentialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List lists service principals for the tenant.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.List" method="get" path="/api/v1/service_principals" -->
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

    res, err := s.Principal.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceListResponse != nil {
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

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceListResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListBindings

ListBindings returns every active binding held by a subject. Empty
 list when the subject is unbound. The response is unordered.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.ListBindings" method="post" path="/api/v1/service_principals/bindings/list" -->
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

    res, err := s.Principal.ListBindings(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceListBindingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [shared.ServicePrincipalServiceListBindingsRequest](../../pkg/models/shared/serviceprincipalservicelistbindingsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceListBindingsResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicelistbindingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCredentials

ListCredentials lists client credentials for a service principal.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.ListCredentials" method="get" path="/api/v1/service_principals/{service_principal_id}/credentials" -->
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

    res, err := s.Principal.ListCredentials(ctx, operations.C1APIServicePrincipalV1ServicePrincipalServiceListCredentialsRequest{
        ServicePrincipalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceListCredentialsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                              | Type                                                                                                                                                                                   | Required                                                                                                                                                                               | Description                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                     | The context to use for the request.                                                                                                                                                    |
| `request`                                                                                                                                                                              | [operations.C1APIServicePrincipalV1ServicePrincipalServiceListCredentialsRequest](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicelistcredentialsrequest.md) | :heavy_check_mark:                                                                                                                                                                     | The request object to use for the request.                                                                                                                                             |
| `opts`                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                     | The options for this request.                                                                                                                                                          |

### Response

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceListCredentialsResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicelistcredentialsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RevokeCredential

RevokeCredential revokes (deletes) a client credential for a service principal.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.RevokeCredential" method="delete" path="/api/v1/service_principals/{service_principal_id}/credentials/{id}" -->
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

    res, err := s.Principal.RevokeCredential(ctx, operations.C1APIServicePrincipalV1ServicePrincipalServiceRevokeCredentialRequest{
        ID: "<id>",
        ServicePrincipalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceRevokeCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                | Type                                                                                                                                                                                     | Required                                                                                                                                                                                 | Description                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                       | The context to use for the request.                                                                                                                                                      |
| `request`                                                                                                                                                                                | [operations.C1APIServicePrincipalV1ServicePrincipalServiceRevokeCredentialRequest](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicerevokecredentialrequest.md) | :heavy_check_mark:                                                                                                                                                                       | The request object to use for the request.                                                                                                                                               |
| `opts`                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                       | The options for this request.                                                                                                                                                            |

### Response

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceRevokeCredentialResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalservicerevokecredentialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update updates a service principal's display name.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.Update" method="patch" path="/api/v1/service_principals/{id}" -->
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

    res, err := s.Principal.Update(ctx, operations.C1APIServicePrincipalV1ServicePrincipalServiceUpdateRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.C1APIServicePrincipalV1ServicePrincipalServiceUpdateRequest](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceUpdateResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCredential

UpdateCredential updates a client credential for a service principal.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.service_principal.v1.ServicePrincipalService.UpdateCredential" method="patch" path="/api/v1/service_principals/{service_principal_id}/credentials/{id}" -->
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

    res, err := s.Principal.UpdateCredential(ctx, operations.C1APIServicePrincipalV1ServicePrincipalServiceUpdateCredentialRequest{
        ID: "<id>",
        ServicePrincipalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServicePrincipalServiceUpdateCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                | Type                                                                                                                                                                                     | Required                                                                                                                                                                                 | Description                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                       | The context to use for the request.                                                                                                                                                      |
| `request`                                                                                                                                                                                | [operations.C1APIServicePrincipalV1ServicePrincipalServiceUpdateCredentialRequest](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalserviceupdatecredentialrequest.md) | :heavy_check_mark:                                                                                                                                                                       | The request object to use for the request.                                                                                                                                               |
| `opts`                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                       | The options for this request.                                                                                                                                                            |

### Response

**[*operations.C1APIServicePrincipalV1ServicePrincipalServiceUpdateCredentialResponse](../../pkg/models/operations/c1apiserviceprincipalv1serviceprincipalserviceupdatecredentialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |