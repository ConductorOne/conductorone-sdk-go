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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.AddAccessEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsRequest{
        RequestCatalogManagementServiceAddAccessEntitlementsRequest: &shared.RequestCatalogManagementServiceAddAccessEntitlementsRequest{
            AccessEntitlements: []shared.AppEntitlementRef{
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("velit"),
                    ID: conductoroneapi.String("aae8d678-64db-4b67-9fd5-e60b375ed4f6"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("reiciendis"),
                    ID: conductoroneapi.String("bee41f33-317f-4e35-b60e-b1ea426555ba"),
                },
            },
        },
        CatalogID: "adipisci",
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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.AddAppEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsRequest{
        RequestCatalogManagementServiceAddAppEntitlementsRequest: &shared.RequestCatalogManagementServiceAddAppEntitlementsRequest{
            AppEntitlements: []shared.AppEntitlementRef{
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("dolores"),
                    ID: conductoroneapi.String("8744ed53-b88f-43a8-98f5-c0b2f2fb7b19"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("dolore"),
                    ID: conductoroneapi.String("a276b269-16fe-41f0-8f42-94e3698f447f"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("ex"),
                    ID: conductoroneapi.String("03e8b445-e80c-4a55-afd2-0e457e1858b6"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("dolorum"),
                    ID: conductoroneapi.String("89fbe3a5-aa8e-4482-8d0a-b4075088e518"),
                },
            },
        },
        CatalogID: "ex",
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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.Create(ctx, shared.RequestCatalogManagementServiceCreateRequest{
        RequestCatalogExpandMask: &shared.RequestCatalogExpandMask{
            Paths: []string{
                "sit",
            },
        },
        Description: conductoroneapi.String("vel"),
        DisplayName: conductoroneapi.String("nostrum"),
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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.Delete(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest{
        RequestCatalogManagementServiceDeleteRequest: &shared.RequestCatalogManagementServiceDeleteRequest{},
        ID: "e904f3b1-194b-48ab-b603-a79f9dfe0ab7",
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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.Get(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest{
        ID: "da8a50ce-187f-486b-8173-d689eee9526f",
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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.ListEntitlementsForAccess(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest{
        CatalogID: "corrupti",
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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.ListEntitlementsPerCatalog(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogRequest{
        CatalogID: "at",
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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.RemoveAccessEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsRequest{
        RequestCatalogManagementServiceRemoveAccessEntitlementsRequest: &shared.RequestCatalogManagementServiceRemoveAccessEntitlementsRequest{
            AccessEntitlements: []shared.AppEntitlementRef{
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("blanditiis"),
                    ID: conductoroneapi.String("6e881ead-4f0e-4101-a563-f94e29e973e9"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("consequuntur"),
                    ID: conductoroneapi.String("2a57a15b-e3e0-4608-87e2-b6e3ab8845f0"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("corporis"),
                    ID: conductoroneapi.String("97a60ff2-a54a-431e-9476-4a3e865e7956"),
                },
            },
        },
        CatalogID: "reiciendis",
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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestCatalogManagement.RemoveAppEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsRequest{
        RequestCatalogManagementServiceRemoveAppEntitlementsRequest: &shared.RequestCatalogManagementServiceRemoveAppEntitlementsRequest{
            AppEntitlements: []shared.AppEntitlementRef{
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("aspernatur"),
                    ID: conductoroneapi.String("51a5a9da-660f-4f57-bfaa-d4f9efc1b451"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("fugit"),
                    ID: conductoroneapi.String("c1032648-dc2f-4615-999e-bfd0e9fe6c63"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoroneapi.String("fugit"),
                    ID: conductoroneapi.String("ca3aed01-1799-4631-afde-04771778ff61"),
                },
            },
        },
        CatalogID: "fugiat",
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
                                AppID: conductoroneapi.String("dicta"),
                                EntitlementID: conductoroneapi.String("odio"),
                            },
                            ManualProvision: &shared.ManualProvision{
                                Instructions: conductoroneapi.String("tempora"),
                                UserIds: []string{
                                    "ex",
                                    "consectetur",
                                },
                            },
                        },
                        AppID: conductoroneapi.String("aliquid"),
                        AppResourceID: conductoroneapi.String("ipsa"),
                        AppResourceTypeID: conductoroneapi.String("laborum"),
                        CertifyPolicyID: conductoroneapi.String("sunt"),
                        ComplianceFrameworkValueIds: []string{
                            "fugiat",
                            "expedita",
                        },
                        Description: conductoroneapi.String("aliquid"),
                        DisplayName: conductoroneapi.String("officia"),
                        DurationGrant: conductoroneapi.String("suscipit"),
                        DurationUnset: &shared.AppEntitlementDurationUnset{},
                        EmergencyGrantEnabled: conductoroneapi.Bool(false),
                        EmergencyGrantPolicyID: conductoroneapi.String("aliquid"),
                        GrantPolicyID: conductoroneapi.String("perferendis"),
                        RevokePolicyID: conductoroneapi.String("eum"),
                        RiskLevelValueID: conductoroneapi.String("voluptas"),
                        Slug: conductoroneapi.String("iste"),
                    },
                },
                AppIds: []string{
                    "ab",
                    "error",
                    "possimus",
                },
                CreatedByUserID: conductoroneapi.String("voluptates"),
                Description: conductoroneapi.String("mollitia"),
                DisplayName: conductoroneapi.String("laborum"),
                ID: conductoroneapi.String("b5851d6c-645b-408b-a189-1baa0fe1ade0"),
                Published: conductoroneapi.Bool(false),
                VisibleToEveryone: conductoroneapi.Bool(false),
            },
            RequestCatalogExpandMask: &shared.RequestCatalogExpandMask{
                Paths: []string{
                    "deleniti",
                },
            },
            UpdateMask: conductoroneapi.String("earum"),
        },
        ID: "6f8c5f35-0d8c-4db5-a341-814301042181",
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

