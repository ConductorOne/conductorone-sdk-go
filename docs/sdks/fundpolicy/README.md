# FundPolicy

## Overview

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [FreezeTenant](#freezetenant) - Freeze Tenant
* [Get](#get) - Get
* [ListHistory](#listhistory) - List History
* [SetOrgCeiling](#setorgceiling) - Set Org Ceiling
* [UnfreezeTenant](#unfreezetenant) - Unfreeze Tenant
* [Update](#update) - Update

## Create

Create the tenant's fund policy. default_limit is required: a tenant
 states its posture explicitly, and there is no implicit default anywhere
 in the write path.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundPolicyService.Create" method="post" path="/api/v1/funds/policy" -->
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

    res, err := s.FundPolicy.Create(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FundPolicyServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [shared.FundPolicyServiceCreateRequest](../../pkg/models/shared/fundpolicyservicecreaterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.C1APIFundsV1FundPolicyServiceCreateResponse](../../pkg/models/operations/c1apifundsv1fundpolicyservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the tenant's fund policy. Refused while spend governance is enabled
 because enabled inference requires this row. Disable governance and drain
 in-flight leases before calling Delete.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundPolicyService.Delete" method="delete" path="/api/v1/funds/policy" -->
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

    res, err := s.FundPolicy.Delete(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FundPolicyServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [shared.FundPolicyServiceDeleteRequest](../../pkg/models/shared/fundpolicyservicedeleterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.C1APIFundsV1FundPolicyServiceDeleteResponse](../../pkg/models/operations/c1apifundsv1fundpolicyservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## FreezeTenant

Freeze the whole tenant. The ceiling amount underneath is preserved and
 restored by Unfreeze.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundPolicyService.FreezeTenant" method="post" path="/api/v1/funds/policy/ceiling/suspension" -->
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

    res, err := s.FundPolicy.FreezeTenant(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FundPolicyServiceFreezeTenantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [shared.FundPolicyServiceFreezeTenantRequest](../../pkg/models/shared/fundpolicyservicefreezetenantrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.C1APIFundsV1FundPolicyServiceFreezeTenantResponse](../../pkg/models/operations/c1apifundsv1fundpolicyservicefreezetenantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get the tenant's fund policy. An absent policy is a configuration state
 that prevents enabled metered inference from serving traffic.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundPolicyService.Get" method="get" path="/api/v1/funds/policy" -->
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

    res, err := s.FundPolicy.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.FundPolicyServiceGetResponse != nil {
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

**[*operations.C1APIFundsV1FundPolicyServiceGetResponse](../../pkg/models/operations/c1apifundsv1fundpolicyservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHistory

List the change history for the fund policy, newest first. Admin-tier per
 the object-history convention.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundPolicyService.ListHistory" method="get" path="/api/v1/funds/policy/history" -->
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

    res, err := s.FundPolicy.ListHistory(ctx, operations.C1APIFundsV1FundPolicyServiceListHistoryRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.FundPolicyServiceListHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIFundsV1FundPolicyServiceListHistoryRequest](../../pkg/models/operations/c1apifundsv1fundpolicyservicelisthistoryrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APIFundsV1FundPolicyServiceListHistoryResponse](../../pkg/models/operations/c1apifundsv1fundpolicyservicelisthistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetOrgCeiling

Set the org-wide ceiling: the bound on the tenant's total regardless of
 what any principal was granted. Amount arm only. An absent limit clears
 the ceiling only after spend governance is disabled.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundPolicyService.SetOrgCeiling" method="post" path="/api/v1/funds/policy/ceiling/limit" -->
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

    res, err := s.FundPolicy.SetOrgCeiling(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FundPolicyServiceSetOrgCeilingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [shared.FundPolicyServiceSetOrgCeilingRequest](../../pkg/models/shared/fundpolicyservicesetorgceilingrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.C1APIFundsV1FundPolicyServiceSetOrgCeilingResponse](../../pkg/models/operations/c1apifundsv1fundpolicyservicesetorgceilingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UnfreezeTenant

Lift the tenant freeze, restoring the ceiling it froze.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundPolicyService.UnfreezeTenant" method="delete" path="/api/v1/funds/policy/ceiling/suspension" -->
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

    res, err := s.FundPolicy.UnfreezeTenant(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FundPolicyServiceUnfreezeTenantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [shared.FundPolicyServiceUnfreezeTenantRequest](../../pkg/models/shared/fundpolicyserviceunfreezetenantrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.C1APIFundsV1FundPolicyServiceUnfreezeTenantResponse](../../pkg/models/operations/c1apifundsv1fundpolicyserviceunfreezetenantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update the period or the default limit. currency_code is fixed to USD at
 Create and immutable thereafter; update_mask rejects it.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.funds.v1.FundPolicyService.Update" method="post" path="/api/v1/funds/policy/update" -->
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

    res, err := s.FundPolicy.Update(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FundPolicyServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [shared.FundPolicyServiceUpdateRequest](../../pkg/models/shared/fundpolicyserviceupdaterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.C1APIFundsV1FundPolicyServiceUpdateResponse](../../pkg/models/operations/c1apifundsv1fundpolicyserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |