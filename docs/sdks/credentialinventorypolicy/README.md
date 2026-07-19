# CredentialInventoryPolicy

## Overview

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List
* [Search](#search) - Search
* [Update](#update) - Update

## Create

Create a credential inventory policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.credential_inventory.v1.CredentialInventoryPolicyService.Create" method="post" path="/api/v1/credential-inventory-policies" -->
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

    res, err := s.CredentialInventoryPolicy.Create(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CredentialInventoryPolicyServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [shared.CredentialInventoryPolicyServiceCreateRequest](../../pkg/models/shared/credentialinventorypolicyservicecreaterequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceCreateResponse](../../pkg/models/operations/c1apicredentialinventoryv1credentialinventorypolicyservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a credential inventory policy by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.credential_inventory.v1.CredentialInventoryPolicyService.Delete" method="delete" path="/api/v1/credential-inventory-policies/{id}" -->
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

    res, err := s.CredentialInventoryPolicy.Delete(ctx, operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceDeleteRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CredentialInventoryPolicyServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                    | Type                                                                                                                                                                                         | Required                                                                                                                                                                                     | Description                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                           | The context to use for the request.                                                                                                                                                          |
| `request`                                                                                                                                                                                    | [operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceDeleteRequest](../../pkg/models/operations/c1apicredentialinventoryv1credentialinventorypolicyservicedeleterequest.md) | :heavy_check_mark:                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                   |
| `opts`                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                           | The options for this request.                                                                                                                                                                |

### Response

**[*operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceDeleteResponse](../../pkg/models/operations/c1apicredentialinventoryv1credentialinventorypolicyservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get a credential inventory policy by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.credential_inventory.v1.CredentialInventoryPolicyService.Get" method="get" path="/api/v1/credential-inventory-policies/{id}" -->
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

    res, err := s.CredentialInventoryPolicy.Get(ctx, operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceGetRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CredentialInventoryPolicyServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                              | Type                                                                                                                                                                                   | Required                                                                                                                                                                               | Description                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                     | The context to use for the request.                                                                                                                                                    |
| `request`                                                                                                                                                                              | [operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceGetRequest](../../pkg/models/operations/c1apicredentialinventoryv1credentialinventorypolicyservicegetrequest.md) | :heavy_check_mark:                                                                                                                                                                     | The request object to use for the request.                                                                                                                                             |
| `opts`                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                     | The options for this request.                                                                                                                                                          |

### Response

**[*operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceGetResponse](../../pkg/models/operations/c1apicredentialinventoryv1credentialinventorypolicyservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List all credential inventory policies in your tenant, one page at a time.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.credential_inventory.v1.CredentialInventoryPolicyService.List" method="get" path="/api/v1/credential-inventory-policies" -->
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

    res, err := s.CredentialInventoryPolicy.List(ctx, operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.CredentialInventoryPolicyServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                | Type                                                                                                                                                                                     | Required                                                                                                                                                                                 | Description                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                       | The context to use for the request.                                                                                                                                                      |
| `request`                                                                                                                                                                                | [operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceListRequest](../../pkg/models/operations/c1apicredentialinventoryv1credentialinventorypolicyservicelistrequest.md) | :heavy_check_mark:                                                                                                                                                                       | The request object to use for the request.                                                                                                                                               |
| `opts`                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                       | The options for this request.                                                                                                                                                            |

### Response

**[*operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceListResponse](../../pkg/models/operations/c1apicredentialinventoryv1credentialinventorypolicyservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Search

Search credential inventory policies by name, or fetch a specific set by
 ID. Returns one page of matching policies at a time.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.credential_inventory.v1.CredentialInventoryPolicyService.Search" method="post" path="/api/v1/search/credential-inventory-policies" -->
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

    res, err := s.CredentialInventoryPolicy.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CredentialInventoryPolicyServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [shared.CredentialInventoryPolicyServiceSearchRequest](../../pkg/models/shared/credentialinventorypolicyservicesearchrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceSearchResponse](../../pkg/models/operations/c1apicredentialinventoryv1credentialinventorypolicyservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update a credential inventory policy. Supply the policy object and an
 update mask listing the fields to change; omitted fields are left as-is.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.credential_inventory.v1.CredentialInventoryPolicyService.Update" method="post" path="/api/v1/credential-inventory-policies/{id}" -->
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

    res, err := s.CredentialInventoryPolicy.Update(ctx, operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceUpdateRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CredentialInventoryPolicyServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                    | Type                                                                                                                                                                                         | Required                                                                                                                                                                                     | Description                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                           | The context to use for the request.                                                                                                                                                          |
| `request`                                                                                                                                                                                    | [operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceUpdateRequest](../../pkg/models/operations/c1apicredentialinventoryv1credentialinventorypolicyserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                   |
| `opts`                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                           | The options for this request.                                                                                                                                                                |

### Response

**[*operations.C1APICredentialInventoryV1CredentialInventoryPolicyServiceUpdateResponse](../../pkg/models/operations/c1apicredentialinventoryv1credentialinventorypolicyserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |