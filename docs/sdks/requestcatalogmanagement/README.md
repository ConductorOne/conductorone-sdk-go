# RequestCatalogManagement

### Available Operations

* [AddAccessEntitlements](#addaccessentitlements) - Add Access Entitlements
* [AddAppEntitlements](#addappentitlements) - Add App Entitlements
* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [ListEntitlementsForAccess](#listentitlementsforaccess) - List Entitlements For Access
* [ListEntitlementsPerCatalog](#listentitlementspercatalog) - List Entitlements Per Catalog
* [RemoveAccessEntitlements](#removeaccessentitlements) - Remove Access Entitlements
* [RemoveAppEntitlements](#removeappentitlements) - Remove App Entitlements
* [Update](#update) - Update

## AddAccessEntitlements

Add visibility bindings (access entitlements) to a catalog.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.AddAccessEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsRequest{
        RequestCatalogManagementServiceAddAccessEntitlementsRequest: &shared.RequestCatalogManagementServiceAddAccessEntitlementsRequest{
            AccessEntitlements: []shared.AppEntitlementRef{
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("eius"),
                    ID: conductoroneapi.String("e3698f44-7f60-43e8-b445-e80ca55efd20"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("saepe"),
                    ID: conductoroneapi.String("457e1858-b6a8-49fb-a3a5-aa8e4824d0ab"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("magnam"),
                    ID: conductoroneapi.String("075088e5-1862-4065-a904-f3b1194b8abf"),
                },
            },
        },
        CatalogID: "laboriosam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceAddAccessEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                  | Type                                                                                                                                                                                                       | Required                                                                                                                                                                                                   | Description                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                        |
| `request`                                                                                                                                                                                                  | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddaccessentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                                 |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddaccessentitlementsresponse.md), error**


## AddAppEntitlements

Add requestable entitlements to a catalog.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.AddAppEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsRequest{
        RequestCatalogManagementServiceAddAppEntitlementsRequest: &shared.RequestCatalogManagementServiceAddAppEntitlementsRequest{
            AppEntitlements: []shared.AppEntitlementRef{
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("amet"),
                    ID: conductoroneapi.String("a79f9dfe-0ab7-4da8-a50c-e187f86bc173"),
                },
            },
        },
        CatalogID: "assumenda",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceAddAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                            | Type                                                                                                                                                                                                 | Required                                                                                                                                                                                             | Description                                                                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                  |
| `request`                                                                                                                                                                                            | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddappentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                           |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddappentitlementsresponse.md), error**


## Create

Creates a new request catalog.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.Create(ctx, shared.RequestCatalogManagementServiceCreateRequest{
        RequestCatalogExpandMask: &shared.RequestCatalogExpandMask{
            Paths: []string{
                "atque",
                "error",
            },
        },
        Description: conductoroneapi.String("officiis"),
        DisplayName: conductoroneapi.String("officiis"),
        Published: conductoroneapi.Bool(false),
        VisibleToEveryone: conductoroneapi.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [shared.RequestCatalogManagementServiceCreateRequest](../../models/shared/requestcatalogmanagementservicecreaterequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicecreateresponse.md), error**


## Delete

Delete a catalog.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.Delete(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest{
        RequestCatalogManagementServiceDeleteRequest: &shared.RequestCatalogManagementServiceDeleteRequest{},
        ID: "e9526f8d-986e-4881-aad4-f0e1012563f9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicedeleterequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicedeleteresponse.md), error**


## Get

Get a catalog.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.Get(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest{
        ID: "4e29e973-e922-4a57-a15b-e3e060807e2b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetrequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetresponse.md), error**


## ListEntitlementsForAccess

List visibility bindings (access entitlements) for a catalog.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.ListEntitlementsForAccess(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest{
        CatalogID: "iure",
        PageSize: conductoroneapi.Float64(8980.63),
        PageToken: conductoroneapi.String("ratione"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceListEntitlementsForAccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                          | Type                                                                                                                                                                                                               | Required                                                                                                                                                                                                           | Description                                                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                |
| `request`                                                                                                                                                                                                          | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementsforaccessrequest.md) | :heavy_check_mark:                                                                                                                                                                                                 | The request object to use for the request.                                                                                                                                                                         |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementsforaccessresponse.md), error**


## ListEntitlementsPerCatalog

List entitlements in a catalog that are requestable.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.ListEntitlementsPerCatalog(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogRequest{
        CatalogID: "laborum",
        PageSize: conductoroneapi.Float64(7152.08),
        PageToken: conductoroneapi.String("voluptatum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceListEntitlementsPerCatalogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                            | Type                                                                                                                                                                                                                 | Required                                                                                                                                                                                                             | Description                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                  |
| `request`                                                                                                                                                                                                            | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementspercatalogrequest.md) | :heavy_check_mark:                                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                                           |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementspercatalogresponse.md), error**


## RemoveAccessEntitlements

Remove visibility bindings (access entitlements) to a catalog.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.RemoveAccessEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsRequest{
        RequestCatalogManagementServiceRemoveAccessEntitlementsRequest: &shared.RequestCatalogManagementServiceRemoveAccessEntitlementsRequest{
            AccessEntitlements: []shared.AppEntitlementRef{
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("aliquam"),
                    ID: conductoroneapi.String("5f0597a6-0ff2-4a54-a31e-94764a3e865e"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("esse"),
                    ID: conductoroneapi.String("956f9251-a5a9-4da6-a0ff-57bfaad4f9ef"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("cumque"),
                    ID: conductoroneapi.String("1b4512c1-0326-448d-82f6-15199ebfd0e9"),
                },
            },
        },
        CatalogID: "maiores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceRemoveAccessEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                        | Type                                                                                                                                                                                                             | Required                                                                                                                                                                                                         | Description                                                                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                              |
| `request`                                                                                                                                                                                                        | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveaccessentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                               | The request object to use for the request.                                                                                                                                                                       |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveaccessentitlementsresponse.md), error**


## RemoveAppEntitlements

Remove requestable entitlements from a catalog.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.RemoveAppEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsRequest{
        RequestCatalogManagementServiceRemoveAppEntitlementsRequest: &shared.RequestCatalogManagementServiceRemoveAppEntitlementsRequest{
            AppEntitlements: []shared.AppEntitlementRef{
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("aliquid"),
                    ID: conductoroneapi.String("c632ca3a-ed01-4179-9631-2fde04771778"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("maiores"),
                    ID: conductoroneapi.String("f61d0174-7636-40a1-9db6-a660659a1ade"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("mollitia"),
                    ID: conductoroneapi.String("ab5851d6-c645-4b08-b618-91baa0fe1ade"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("voluptatem"),
                    ID: conductoroneapi.String("08e6f8c5-f350-4d8c-9b5a-341814301042"),
                },
            },
        },
        CatalogID: "ab",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceRemoveAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                  | Type                                                                                                                                                                                                       | Required                                                                                                                                                                                                   | Description                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                        |
| `request`                                                                                                                                                                                                  | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveappentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                                 |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveappentitlementsresponse.md), error**


## Update

Update a catalog.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.Update(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateRequest{
        RequestCatalogManagementServiceUpdateRequestInput: &shared.RequestCatalogManagementServiceUpdateRequestInput{
            RequestCatalog: &shared.RequestCatalogInput{
                AccessEntitlements: []shared.AppEntitlementInput{
                    shared.AppEntitlementInput{
                        ProvisionPolicy: &shared.ProvisionPolicy{
                            ConnectorProvision: &shared.ConnectorProvision{},
                            DelegatedProvision: &shared.DelegatedProvision{
                                AppID: conductoroneapi.String("quae"),
                                EntitlementID: conductoroneapi.String("dolor"),
                            },
                            ManualProvision: &shared.ManualProvision{
                                Instructions: conductoroneapi.String("fugiat"),
                                UserIds: []string{
                                    "consequuntur",
                                    "ipsa",
                                },
                            },
                        },
                        AppID: conductoroneapi.String("quas"),
                        AppResourceID: conductoroneapi.String("eveniet"),
                        AppResourceTypeID: conductoroneapi.String("impedit"),
                        CertifyPolicyID: conductoroneapi.String("officiis"),
                        ComplianceFrameworkValueIds: []string{
                            "necessitatibus",
                            "sed",
                        },
                        Description: conductoroneapi.String("veniam"),
                        DisplayName: conductoroneapi.String("nesciunt"),
                        DurationGrant: conductoroneapi.String("expedita"),
                        DurationUnset: &shared.AppEntitlementDurationUnset{},
                        EmergencyGrantEnabled: conductoroneapi.Bool(false),
                        EmergencyGrantPolicyID: conductoroneapi.String("eum"),
                        GrantPolicyID: conductoroneapi.String("vel"),
                        RevokePolicyID: conductoroneapi.String("voluptatum"),
                        RiskLevelValueID: conductoroneapi.String("magnam"),
                        Slug: conductoroneapi.String("exercitationem"),
                    },
                    shared.AppEntitlementInput{
                        ProvisionPolicy: &shared.ProvisionPolicy{
                            ConnectorProvision: &shared.ConnectorProvision{},
                            DelegatedProvision: &shared.DelegatedProvision{
                                AppID: conductoroneapi.String("ab"),
                                EntitlementID: conductoroneapi.String("porro"),
                            },
                            ManualProvision: &shared.ManualProvision{
                                Instructions: conductoroneapi.String("autem"),
                                UserIds: []string{
                                    "laboriosam",
                                    "recusandae",
                                    "consequuntur",
                                    "voluptatem",
                                },
                            },
                        },
                        AppID: conductoroneapi.String("exercitationem"),
                        AppResourceID: conductoroneapi.String("necessitatibus"),
                        AppResourceTypeID: conductoroneapi.String("quasi"),
                        CertifyPolicyID: conductoroneapi.String("nisi"),
                        ComplianceFrameworkValueIds: []string{
                            "vero",
                            "est",
                            "harum",
                            "sequi",
                        },
                        Description: conductoroneapi.String("doloribus"),
                        DisplayName: conductoroneapi.String("repudiandae"),
                        DurationGrant: conductoroneapi.String("optio"),
                        DurationUnset: &shared.AppEntitlementDurationUnset{},
                        EmergencyGrantEnabled: conductoroneapi.Bool(false),
                        EmergencyGrantPolicyID: conductoroneapi.String("occaecati"),
                        GrantPolicyID: conductoroneapi.String("nemo"),
                        RevokePolicyID: conductoroneapi.String("voluptate"),
                        RiskLevelValueID: conductoroneapi.String("blanditiis"),
                        Slug: conductoroneapi.String("officia"),
                    },
                    shared.AppEntitlementInput{
                        ProvisionPolicy: &shared.ProvisionPolicy{
                            ConnectorProvision: &shared.ConnectorProvision{},
                            DelegatedProvision: &shared.DelegatedProvision{
                                AppID: conductoroneapi.String("voluptas"),
                                EntitlementID: conductoroneapi.String("numquam"),
                            },
                            ManualProvision: &shared.ManualProvision{
                                Instructions: conductoroneapi.String("nemo"),
                                UserIds: []string{
                                    "eius",
                                    "aspernatur",
                                    "ducimus",
                                },
                            },
                        },
                        AppID: conductoroneapi.String("nesciunt"),
                        AppResourceID: conductoroneapi.String("fuga"),
                        AppResourceTypeID: conductoroneapi.String("laudantium"),
                        CertifyPolicyID: conductoroneapi.String("incidunt"),
                        ComplianceFrameworkValueIds: []string{
                            "rem",
                        },
                        Description: conductoroneapi.String("fugiat"),
                        DisplayName: conductoroneapi.String("dicta"),
                        DurationGrant: conductoroneapi.String("nisi"),
                        DurationUnset: &shared.AppEntitlementDurationUnset{},
                        EmergencyGrantEnabled: conductoroneapi.Bool(false),
                        EmergencyGrantPolicyID: conductoroneapi.String("consequuntur"),
                        GrantPolicyID: conductoroneapi.String("consectetur"),
                        RevokePolicyID: conductoroneapi.String("aperiam"),
                        RiskLevelValueID: conductoroneapi.String("cupiditate"),
                        Slug: conductoroneapi.String("reiciendis"),
                    },
                },
                AppIds: []string{
                    "alias",
                    "omnis",
                    "eos",
                },
                CreatedByUserID: conductoroneapi.String("occaecati"),
                Description: conductoroneapi.String("iste"),
                DisplayName: conductoroneapi.String("magni"),
                ID: conductoroneapi.String("1aefb9f5-8c4d-486e-a8e4-be056013f59d"),
                Published: conductoroneapi.Bool(false),
                VisibleToEveryone: conductoroneapi.Bool(false),
            },
            RequestCatalogExpandMask: &shared.RequestCatalogExpandMask{
                Paths: []string{
                    "esse",
                    "nemo",
                    "reprehenderit",
                },
            },
            UpdateMask: conductoroneapi.String("est"),
        },
        ID: "59ecfef6-6ef1-4caa-b383-c2beb477373c",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceupdateresponse.md), error**

