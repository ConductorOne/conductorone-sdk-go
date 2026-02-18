# WorkloadFederation

## Overview

### Available Operations

* [CreateProvider](#createprovider) - Create Provider
* [CreateTrust](#createtrust) - Create Trust
* [DeleteProvider](#deleteprovider) - Delete Provider
* [DeleteTrust](#deletetrust) - Delete Trust
* [GetProvider](#getprovider) - Get Provider
* [GetTrust](#gettrust) - Get Trust
* [ListProviders](#listproviders) - List Providers
* [ListTrusts](#listtrusts) - List Trusts
* [SearchTrusts](#searchtrusts) - Search Trusts
* [TestCEL](#testcel) - Test Cel
* [TestToken](#testtoken) - Test Token
* [UpdateProvider](#updateprovider) - Update Provider
* [UpdateTrust](#updatetrust) - Update Trust

## CreateProvider

CreateProvider registers a new external OIDC issuer for the tenant.
 Validates the issuer URL via OIDC discovery synchronously.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.CreateProvider" method="post" path="/api/v1/workload_federation/providers" -->
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

    res, err := s.WorkloadFederation.CreateProvider(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceCreateProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [shared.WorkloadFederationServiceCreateProviderRequest](../../pkg/models/shared/workloadfederationservicecreateproviderrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceCreateProviderResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicecreateproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateTrust

CreateTrust creates a trust policy for a service principal.
 Validates the CEL condition_expression at creation time.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.CreateTrust" method="post" path="/api/v1/service_principals/{service_principal_id}/trusts" -->
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

    res, err := s.WorkloadFederation.CreateTrust(ctx, operations.C1APIWorkloadFederationV1WorkloadFederationServiceCreateTrustRequest{
        ServicePrincipalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceCreateTrustResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                              | Type                                                                                                                                                                                   | Required                                                                                                                                                                               | Description                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                     | The context to use for the request.                                                                                                                                                    |
| `request`                                                                                                                                                                              | [operations.C1APIWorkloadFederationV1WorkloadFederationServiceCreateTrustRequest](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicecreatetrustrequest.md) | :heavy_check_mark:                                                                                                                                                                     | The request object to use for the request.                                                                                                                                             |
| `opts`                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                     | The options for this request.                                                                                                                                                          |

### Response

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceCreateTrustResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicecreatetrustresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteProvider

DeleteProvider deletes a provider. Fails if active trusts reference it.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.DeleteProvider" method="delete" path="/api/v1/workload_federation/providers/{id}" -->
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

    res, err := s.WorkloadFederation.DeleteProvider(ctx, operations.C1APIWorkloadFederationV1WorkloadFederationServiceDeleteProviderRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceDeleteProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                    | Type                                                                                                                                                                                         | Required                                                                                                                                                                                     | Description                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                           | The context to use for the request.                                                                                                                                                          |
| `request`                                                                                                                                                                                    | [operations.C1APIWorkloadFederationV1WorkloadFederationServiceDeleteProviderRequest](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicedeleteproviderrequest.md) | :heavy_check_mark:                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                   |
| `opts`                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                           | The options for this request.                                                                                                                                                                |

### Response

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceDeleteProviderResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicedeleteproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteTrust

DeleteTrust deletes a trust for a service principal.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.DeleteTrust" method="delete" path="/api/v1/service_principals/{service_principal_id}/trusts/{client_id}" -->
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

    res, err := s.WorkloadFederation.DeleteTrust(ctx, operations.C1APIWorkloadFederationV1WorkloadFederationServiceDeleteTrustRequest{
        ClientID: "<id>",
        ServicePrincipalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceDeleteTrustResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                              | Type                                                                                                                                                                                   | Required                                                                                                                                                                               | Description                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                     | The context to use for the request.                                                                                                                                                    |
| `request`                                                                                                                                                                              | [operations.C1APIWorkloadFederationV1WorkloadFederationServiceDeleteTrustRequest](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicedeletetrustrequest.md) | :heavy_check_mark:                                                                                                                                                                     | The request object to use for the request.                                                                                                                                             |
| `opts`                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                     | The options for this request.                                                                                                                                                          |

### Response

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceDeleteTrustResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicedeletetrustresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetProvider

GetProvider returns a provider by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.GetProvider" method="get" path="/api/v1/workload_federation/providers/{id}" -->
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

    res, err := s.WorkloadFederation.GetProvider(ctx, operations.C1APIWorkloadFederationV1WorkloadFederationServiceGetProviderRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceGetProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                              | Type                                                                                                                                                                                   | Required                                                                                                                                                                               | Description                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                     | The context to use for the request.                                                                                                                                                    |
| `request`                                                                                                                                                                              | [operations.C1APIWorkloadFederationV1WorkloadFederationServiceGetProviderRequest](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicegetproviderrequest.md) | :heavy_check_mark:                                                                                                                                                                     | The request object to use for the request.                                                                                                                                             |
| `opts`                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                     | The options for this request.                                                                                                                                                          |

### Response

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceGetProviderResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicegetproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTrust

GetTrust returns a trust by ID for a service principal.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.GetTrust" method="get" path="/api/v1/service_principals/{service_principal_id}/trusts/{client_id}" -->
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

    res, err := s.WorkloadFederation.GetTrust(ctx, operations.C1APIWorkloadFederationV1WorkloadFederationServiceGetTrustRequest{
        ClientID: "<id>",
        ServicePrincipalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceGetTrustResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `request`                                                                                                                                                                        | [operations.C1APIWorkloadFederationV1WorkloadFederationServiceGetTrustRequest](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicegettrustrequest.md) | :heavy_check_mark:                                                                                                                                                               | The request object to use for the request.                                                                                                                                       |
| `opts`                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                     | :heavy_minus_sign:                                                                                                                                                               | The options for this request.                                                                                                                                                    |

### Response

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceGetTrustResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicegettrustresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListProviders

ListProviders lists all providers for the tenant.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.ListProviders" method="get" path="/api/v1/workload_federation/providers" -->
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

    res, err := s.WorkloadFederation.ListProviders(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceListProvidersResponse != nil {
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

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceListProvidersResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicelistprovidersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTrusts

ListTrusts lists trusts for a service principal.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.ListTrusts" method="get" path="/api/v1/service_principals/{service_principal_id}/trusts" -->
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

    res, err := s.WorkloadFederation.ListTrusts(ctx, operations.C1APIWorkloadFederationV1WorkloadFederationServiceListTrustsRequest{
        ServicePrincipalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceListTrustsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                            | Type                                                                                                                                                                                 | Required                                                                                                                                                                             | Description                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                | :heavy_check_mark:                                                                                                                                                                   | The context to use for the request.                                                                                                                                                  |
| `request`                                                                                                                                                                            | [operations.C1APIWorkloadFederationV1WorkloadFederationServiceListTrustsRequest](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicelisttrustsrequest.md) | :heavy_check_mark:                                                                                                                                                                   | The request object to use for the request.                                                                                                                                           |
| `opts`                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                   | The options for this request.                                                                                                                                                        |

### Response

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceListTrustsResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicelisttrustsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchTrusts

SearchTrusts searches trusts across all service principals with optional filters.
 Used by the admin providers page to list trusts referencing a provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.SearchTrusts" method="post" path="/api/v1/search/workload_federation_trusts" -->
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

    res, err := s.WorkloadFederation.SearchTrusts(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceSearchTrustsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [shared.WorkloadFederationServiceSearchTrustsRequest](../../pkg/models/shared/workloadfederationservicesearchtrustsrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceSearchTrustsResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicesearchtrustsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## TestCEL

TestCEL evaluates a CEL expression against provided claims without
 requiring a JWT, provider, or trust. Used for expression authoring.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.TestCEL" method="post" path="/api/v1/workload_federation/test-cel" -->
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

    res, err := s.WorkloadFederation.TestCEL(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceTestCELResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [shared.WorkloadFederationServiceTestCELRequest](../../pkg/models/shared/workloadfederationservicetestcelrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceTestCELResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicetestcelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## TestToken

TestToken validates a JWT against a specific trust's configuration without
 issuing an access token. Returns per-step validation results for debugging.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.TestToken" method="post" path="/api/v1/service_principals/{service_principal_id}/trusts/{client_id}/test" -->
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

    res, err := s.WorkloadFederation.TestToken(ctx, operations.C1APIWorkloadFederationV1WorkloadFederationServiceTestTokenRequest{
        ClientID: "<id>",
        ServicePrincipalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceTestTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                          | Type                                                                                                                                                                               | Required                                                                                                                                                                           | Description                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                              | :heavy_check_mark:                                                                                                                                                                 | The context to use for the request.                                                                                                                                                |
| `request`                                                                                                                                                                          | [operations.C1APIWorkloadFederationV1WorkloadFederationServiceTestTokenRequest](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicetesttokenrequest.md) | :heavy_check_mark:                                                                                                                                                                 | The request object to use for the request.                                                                                                                                         |
| `opts`                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                 | The options for this request.                                                                                                                                                      |

### Response

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceTestTokenResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationservicetesttokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateProvider

UpdateProvider updates a provider's mutable fields (display_name, description, disabled).
 The issuer_url is immutable after creation.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.UpdateProvider" method="patch" path="/api/v1/workload_federation/providers/{id}" -->
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

    res, err := s.WorkloadFederation.UpdateProvider(ctx, operations.C1APIWorkloadFederationV1WorkloadFederationServiceUpdateProviderRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceUpdateProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                    | Type                                                                                                                                                                                         | Required                                                                                                                                                                                     | Description                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                           | The context to use for the request.                                                                                                                                                          |
| `request`                                                                                                                                                                                    | [operations.C1APIWorkloadFederationV1WorkloadFederationServiceUpdateProviderRequest](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationserviceupdateproviderrequest.md) | :heavy_check_mark:                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                   |
| `opts`                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                           | The options for this request.                                                                                                                                                                |

### Response

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceUpdateProviderResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationserviceupdateproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateTrust

UpdateTrust updates a trust's mutable fields. The provider_id is immutable.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.workload_federation.v1.WorkloadFederationService.UpdateTrust" method="patch" path="/api/v1/service_principals/{service_principal_id}/trusts/{client_id}" -->
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

    res, err := s.WorkloadFederation.UpdateTrust(ctx, operations.C1APIWorkloadFederationV1WorkloadFederationServiceUpdateTrustRequest{
        ClientID: "<id>",
        ServicePrincipalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkloadFederationServiceUpdateTrustResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                              | Type                                                                                                                                                                                   | Required                                                                                                                                                                               | Description                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                     | The context to use for the request.                                                                                                                                                    |
| `request`                                                                                                                                                                              | [operations.C1APIWorkloadFederationV1WorkloadFederationServiceUpdateTrustRequest](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationserviceupdatetrustrequest.md) | :heavy_check_mark:                                                                                                                                                                     | The request object to use for the request.                                                                                                                                             |
| `opts`                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                     | The options for this request.                                                                                                                                                          |

### Response

**[*operations.C1APIWorkloadFederationV1WorkloadFederationServiceUpdateTrustResponse](../../pkg/models/operations/c1apiworkloadfederationv1workloadfederationserviceupdatetrustresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |