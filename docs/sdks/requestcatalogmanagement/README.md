# RequestCatalogManagement

### Available Operations

* [AddAccessEntitlements](#addaccessentitlements) - Add Access Entitlements
* [AddAppEntitlements](#addappentitlements) - Add App Entitlements
* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List
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
                    AppID: conductoroneapi.String("labore"),
                    ID: conductoroneapi.String("dbb675fd-5e60-4b37-9ed4-f6fbee41f333"),
                },
            },
        },
        CatalogID: "beatae",
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
                    AppID: conductoroneapi.String("dignissimos"),
                    ID: conductoroneapi.String("fe35b60e-b1ea-4426-955b-a3c28744ed53"),
                },
            },
        },
        CatalogID: "cum",
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
                "blanditiis",
            },
        },
        Description: conductoroneapi.String("quas"),
        DisplayName: conductoroneapi.String("hic"),
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
        ID: "3a8d8f5c-0b2f-42fb-bb19-4a276b26916f",
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
        ID: "e1f08f42-94e3-4698-b447-f603e8b445e8",
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


## List

Get a list of request catalogs.

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
    res, err := s.RequestCatalogManagement.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistresponse.md), error**


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
        CatalogID: "sit",
        PageSize: conductoroneapi.Float64(7505.95),
        PageToken: conductoroneapi.String("error"),
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
        CatalogID: "veniam",
        PageSize: conductoroneapi.Float64(3295.43),
        PageToken: conductoroneapi.String("recusandae"),
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
                    AppID: conductoroneapi.String("reiciendis"),
                    ID: conductoroneapi.String("d20e457e-1858-4b6a-89fb-e3a5aa8e4824"),
                },
            },
        },
        CatalogID: "fugiat",
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
                    AppID: conductoroneapi.String("voluptatem"),
                    ID: conductoroneapi.String("ab407508-8e51-4862-865e-904f3b1194b8"),
                },
            },
        },
        CatalogID: "laborum",
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
                                AppID: conductoroneapi.String("nam"),
                                EntitlementID: conductoroneapi.String("tenetur"),
                            },
                            ManualProvision: &shared.ManualProvision{
                                Instructions: conductoroneapi.String("laboriosam"),
                                UserIds: []string{
                                    "alias",
                                },
                            },
                        },
                        AppID: conductoroneapi.String("amet"),
                        AppResourceID: conductoroneapi.String("deserunt"),
                        AppResourceTypeID: conductoroneapi.String("voluptate"),
                        CertifyPolicyID: conductoroneapi.String("unde"),
                        ComplianceFrameworkValueIds: []string{
                            "reiciendis",
                        },
                        Description: conductoroneapi.String("provident"),
                        DisplayName: conductoroneapi.String("repellendus"),
                        DurationGrant: conductoroneapi.String("delectus"),
                        DurationUnset: &shared.AppEntitlementDurationUnset{},
                        EmergencyGrantEnabled: conductoroneapi.Bool(false),
                        EmergencyGrantPolicyID: conductoroneapi.String("voluptates"),
                        GrantPolicyID: conductoroneapi.String("perferendis"),
                        RevokePolicyID: conductoroneapi.String("est"),
                        RiskLevelValueID: conductoroneapi.String("quidem"),
                        Slug: conductoroneapi.String("reprehenderit"),
                    },
                },
                AppIds: []string{
                    "facere",
                },
                CreatedByUserID: conductoroneapi.String("fuga"),
                Description: conductoroneapi.String("praesentium"),
                DisplayName: conductoroneapi.String("mollitia"),
                ID: conductoroneapi.String("50ce187f-86bc-4173-9689-eee9526f8d98"),
                Published: conductoroneapi.Bool(false),
                VisibleToEveryone: conductoroneapi.Bool(false),
            },
            RequestCatalogExpandMask: &shared.RequestCatalogExpandMask{
                Paths: []string{
                    "suscipit",
                },
            },
            UpdateMask: conductoroneapi.String("repudiandae"),
        },
        ID: "881ead4f-0e10-4125-a3f9-4e29e973e922",
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

