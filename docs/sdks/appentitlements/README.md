# AppEntitlements

### Available Operations

* [Get](#get) - Get
* [List](#list) - List
* [ListForAppResource](#listforappresource) - List For App Resource
* [ListForAppUser](#listforappuser) - List For App User
* [ListUsers](#listusers) - List Users
* [Update](#update) - Update

## Get

Get an app entitlement by ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlements.Get(ctx, operations.C1APIAppV1AppEntitlementsGetRequest{
        AppID: "sapiente",
        ID: "c2ddf7cc-78ca-41ba-928f-c816742cb739",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAppEntitlementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.C1APIAppV1AppEntitlementsGetRequest](../../models/operations/c1apiappv1appentitlementsgetrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.C1APIAppV1AppEntitlementsGetResponse](../../models/operations/c1apiappv1appentitlementsgetresponse.md), error**


## List

List app entitlements associated with an app.

### Example Usage

```go
package main

import(
	"context"
	"log"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlements.List(ctx, operations.C1APIAppV1AppEntitlementsListRequest{
        AppID: "aspernatur",
        PageSize: conductoronesdkgo.Float64(187.89),
        PageToken: conductoronesdkgo.String("ad"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.C1APIAppV1AppEntitlementsListRequest](../../models/operations/c1apiappv1appentitlementslistrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.C1APIAppV1AppEntitlementsListResponse](../../models/operations/c1apiappv1appentitlementslistresponse.md), error**


## ListForAppResource

List app entitlements associated with an app resource.

### Example Usage

```go
package main

import(
	"context"
	"log"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlements.ListForAppResource(ctx, operations.C1APIAppV1AppEntitlementsListForAppResourceRequest{
        AppID: "natus",
        AppResourceID: "sed",
        AppResourceTypeID: "iste",
        PageSize: conductoronesdkgo.Float64(2223.21),
        PageToken: conductoronesdkgo.String("natus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAppV1AppEntitlementsListForAppResourceRequest](../../models/operations/c1apiappv1appentitlementslistforappresourcerequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |


### Response

**[*operations.C1APIAppV1AppEntitlementsListForAppResourceResponse](../../models/operations/c1apiappv1appentitlementslistforappresourceresponse.md), error**


## ListForAppUser

List app entitlements associated with an app user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlements.ListForAppUser(ctx, operations.C1APIAppV1AppEntitlementsListForAppUserRequest{
        AppID: "laboriosam",
        AppUserID: "hic",
        PageSize: conductoronesdkgo.Float64(9025.99),
        PageToken: conductoronesdkgo.String("fuga"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.C1APIAppV1AppEntitlementsListForAppUserRequest](../../models/operations/c1apiappv1appentitlementslistforappuserrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.C1APIAppV1AppEntitlementsListForAppUserResponse](../../models/operations/c1apiappv1appentitlementslistforappuserresponse.md), error**


## ListUsers

List the users, as AppEntitlementUsers objects, of an app entitlement.

### Example Usage

```go
package main

import(
	"context"
	"log"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlements.ListUsers(ctx, operations.C1APIAppV1AppEntitlementsListUsersRequest{
        AppEntitlementID: "in",
        AppID: "corporis",
        PageSize: conductoronesdkgo.Float64(6130.64),
        PageToken: conductoronesdkgo.String("iure"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppEntitlementUsersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.C1APIAppV1AppEntitlementsListUsersRequest](../../models/operations/c1apiappv1appentitlementslistusersrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.C1APIAppV1AppEntitlementsListUsersResponse](../../models/operations/c1apiappv1appentitlementslistusersresponse.md), error**


## Update

Update an app entitlement by ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlements.Update(ctx, operations.C1APIAppV1AppEntitlementsUpdateRequest{
        UpdateAppEntitlementRequestInput: &shared.UpdateAppEntitlementRequestInput{
            AppEntitlement: &shared.AppEntitlementInput{
                ProvisionPolicy: &shared.ProvisionPolicy{
                    ConnectorProvision: &shared.ConnectorProvision{},
                    DelegatedProvision: &shared.DelegatedProvision{
                        AppID: conductoronesdkgo.String("saepe"),
                        EntitlementID: conductoronesdkgo.String("quidem"),
                    },
                    ManualProvision: &shared.ManualProvision{
                        Instructions: conductoronesdkgo.String("architecto"),
                        UserIds: []string{
                            "ipsa",
                        },
                    },
                },
                AppID: conductoronesdkgo.String("reiciendis"),
                AppResourceID: conductoronesdkgo.String("est"),
                AppResourceTypeID: conductoronesdkgo.String("mollitia"),
                CertifyPolicyID: conductoronesdkgo.String("laborum"),
                ComplianceFrameworkValueIds: []string{
                    "dolores",
                },
                Description: conductoronesdkgo.String("dolorem"),
                DisplayName: conductoronesdkgo.String("corporis"),
                DurationGrant: conductoronesdkgo.String("explicabo"),
                DurationUnset: &shared.AppEntitlementDurationUnset{},
                EmergencyGrantEnabled: conductoronesdkgo.Bool(false),
                EmergencyGrantPolicyID: conductoronesdkgo.String("nobis"),
                GrantPolicyID: conductoronesdkgo.String("enim"),
                RevokePolicyID: conductoronesdkgo.String("omnis"),
                RiskLevelValueID: conductoronesdkgo.String("nemo"),
                Slug: conductoronesdkgo.String("minima"),
            },
            AppEntitlementExpandMask: &shared.AppEntitlementExpandMask{
                Paths: []string{
                    "excepturi",
                },
            },
            UpdateMask: conductoronesdkgo.String("accusantium"),
        },
        AppID: "iure",
        ID: "aff1a3a2-fa94-4677-b925-1aa52c3f5ad0",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateAppEntitlementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.C1APIAppV1AppEntitlementsUpdateRequest](../../models/operations/c1apiappv1appentitlementsupdaterequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.C1APIAppV1AppEntitlementsUpdateResponse](../../models/operations/c1apiappv1appentitlementsupdateresponse.md), error**

