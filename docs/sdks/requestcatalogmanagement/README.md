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
                    AppID: conductoroneapi.String("voluptatem"),
                    ID: conductoroneapi.String("ab407508-8e51-4862-865e-904f3b1194b8"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("laborum"),
                    ID: conductoroneapi.String("bf603a79-f9df-4e0a-b7da-8a50ce187f86"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("quidem"),
                    ID: conductoroneapi.String("c173d689-eee9-4526-b8d9-86e881ead4f0"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("repudiandae"),
                    ID: conductoroneapi.String("1012563f-94e2-49e9-b3e9-22a57a15be3e"),
                },
            },
        },
        CatalogID: "doloremque",
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
                    AppID: conductoroneapi.String("ipsa"),
                    ID: conductoroneapi.String("807e2b6e-3ab8-4845-b059-7a60ff2a54a3"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("quae"),
                    ID: conductoroneapi.String("e94764a3-e865-4e79-96f9-251a5a9da660"),
                },
            },
        },
        CatalogID: "repellat",
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
                "ullam",
                "in",
                "nam",
                "earum",
            },
        },
        Description: conductoroneapi.String("officia"),
        DisplayName: conductoroneapi.String("laborum"),
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
        ID: "d4f9efc1-b451-42c1-8326-48dc2f615199",
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
        ID: "ebfd0e9f-e6c6-432c-a3ae-d0117996312f",
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
        CatalogID: "nulla",
        PageSize: conductoroneapi.Float64(8965.82),
        PageToken: conductoroneapi.String("ipsa"),
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
        CatalogID: "tempora",
        PageSize: conductoroneapi.Float64(4706.21),
        PageToken: conductoroneapi.String("molestiae"),
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
                    AppID: conductoroneapi.String("iusto"),
                    ID: conductoroneapi.String("78ff61d0-1747-4636-8a15-db6a660659a1"),
                },
            },
        },
        CatalogID: "error",
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
                    AppID: conductoroneapi.String("voluptates"),
                    ID: conductoroneapi.String("aab5851d-6c64-45b0-8b61-891baa0fe1ad"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("debitis"),
                    ID: conductoroneapi.String("008e6f8c-5f35-40d8-8db5-a34181430104"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("fugit"),
                    ID: conductoroneapi.String("1813d520-8ece-47e2-93b6-68451c6c6e20"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("exercitationem"),
                    ID: conductoroneapi.String("e16deab3-fec9-4578-a645-84273a8418d1"),
                },
            },
        },
        CatalogID: "nisi",
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
                                AppID: conductoroneapi.String("consectetur"),
                                EntitlementID: conductoroneapi.String("aperiam"),
                            },
                            ManualProvision: &shared.ManualProvision{
                                Instructions: conductoroneapi.String("cupiditate"),
                                UserIds: []string{
                                    "soluta",
                                    "alias",
                                    "omnis",
                                    "eos",
                                },
                            },
                        },
                        AppID: conductoroneapi.String("occaecati"),
                        AppResourceID: conductoroneapi.String("iste"),
                        AppResourceTypeID: conductoroneapi.String("magni"),
                        CertifyPolicyID: conductoroneapi.String("inventore"),
                        ComplianceFrameworkValueIds: []string{
                            "accusamus",
                            "voluptatibus",
                            "distinctio",
                        },
                        Description: conductoroneapi.String("omnis"),
                        DisplayName: conductoroneapi.String("delectus"),
                        DurationGrant: conductoroneapi.String("minima"),
                        DurationUnset: &shared.AppEntitlementDurationUnset{},
                        EmergencyGrantEnabled: conductoroneapi.Bool(false),
                        EmergencyGrantPolicyID: conductoroneapi.String("praesentium"),
                        GrantPolicyID: conductoroneapi.String("maxime"),
                        RevokePolicyID: conductoroneapi.String("magnam"),
                        RiskLevelValueID: conductoroneapi.String("temporibus"),
                        Slug: conductoroneapi.String("quos"),
                    },
                },
                AppIds: []string{
                    "itaque",
                    "commodi",
                },
                CreatedByUserID: conductoroneapi.String("totam"),
                Description: conductoroneapi.String("earum"),
                DisplayName: conductoroneapi.String("modi"),
                ID: conductoroneapi.String("be056013-f59d-4a75-ba59-ecfef66ef1ca"),
                Published: conductoroneapi.Bool(false),
                VisibleToEveryone: conductoroneapi.Bool(false),
            },
            RequestCatalogExpandMask: &shared.RequestCatalogExpandMask{
                Paths: []string{
                    "consectetur",
                    "velit",
                    "atque",
                },
            },
            UpdateMask: conductoroneapi.String("ipsum"),
        },
        ID: "c2beb477-373c-48d7-af64-d1db1f2c4310",
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

