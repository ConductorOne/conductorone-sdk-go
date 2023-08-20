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
    res, err := s.AppEntitlements.Get(ctx, operations.C1APIAppV1AppEntitlementsGetRequest{
        AppID: "iste",
        ID: "396fea75-96eb-410f-aaa2-352c5955907a",
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
    res, err := s.AppEntitlements.List(ctx, operations.C1APIAppV1AppEntitlementsListRequest{
        AppID: "doloribus",
        PageSize: conductoroneapi.Float64(9589.5),
        PageToken: conductoroneapi.String("architecto"),
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
    res, err := s.AppEntitlements.ListForAppResource(ctx, operations.C1APIAppV1AppEntitlementsListForAppResourceRequest{
        AppID: "mollitia",
        AppResourceID: "dolorem",
        AppResourceTypeID: "culpa",
        PageSize: conductoroneapi.Float64(1613.09),
        PageToken: conductoroneapi.String("repellat"),
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
    res, err := s.AppEntitlements.ListForAppUser(ctx, operations.C1APIAppV1AppEntitlementsListForAppUserRequest{
        AppID: "mollitia",
        AppUserID: "occaecati",
        PageSize: conductoroneapi.Float64(2532.91),
        PageToken: conductoroneapi.String("commodi"),
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
    res, err := s.AppEntitlements.ListUsers(ctx, operations.C1APIAppV1AppEntitlementsListUsersRequest{
        AppEntitlementID: "quam",
        AppID: "molestiae",
        PageSize: conductoroneapi.Float64(2444.25),
        PageToken: conductoroneapi.String("error"),
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
    res, err := s.AppEntitlements.Update(ctx, operations.C1APIAppV1AppEntitlementsUpdateRequest{
        UpdateAppEntitlementRequestInput: &shared.UpdateAppEntitlementRequestInput{
            AppEntitlement: &shared.AppEntitlementInput{
                ProvisionPolicy: &shared.ProvisionPolicy{
                    ConnectorProvision: &shared.ConnectorProvision{},
                    DelegatedProvision: &shared.DelegatedProvision{
                        AppID: conductoroneapi.String("quia"),
                        EntitlementID: conductoroneapi.String("quis"),
                    },
                    ManualProvision: &shared.ManualProvision{
                        Instructions: conductoroneapi.String("vitae"),
                        UserIds: []string{
                            "animi",
                            "enim",
                            "odit",
                        },
                    },
                },
                AppID: conductoroneapi.String("quo"),
                AppResourceID: conductoroneapi.String("sequi"),
                AppResourceTypeID: conductoroneapi.String("tenetur"),
                CertifyPolicyID: conductoroneapi.String("ipsam"),
                ComplianceFrameworkValueIds: []string{
                    "possimus",
                    "aut",
                    "quasi",
                },
                Description: conductoroneapi.String("error"),
                DisplayName: conductoroneapi.String("temporibus"),
                DurationGrant: conductoroneapi.String("laborum"),
                DurationUnset: &shared.AppEntitlementDurationUnset{},
                EmergencyGrantEnabled: conductoroneapi.Bool(false),
                EmergencyGrantPolicyID: conductoroneapi.String("quasi"),
                GrantPolicyID: conductoroneapi.String("reiciendis"),
                RevokePolicyID: conductoroneapi.String("voluptatibus"),
                RiskLevelValueID: conductoroneapi.String("vero"),
                Slug: conductoroneapi.String("nihil"),
            },
            AppEntitlementExpandMask: &shared.AppEntitlementExpandMask{
                Paths: []string{
                    "voluptatibus",
                    "ipsa",
                    "omnis",
                },
            },
            UpdateMask: conductoroneapi.String("voluptate"),
        },
        AppID: "cum",
        ID: "0074f154-71b5-4e6e-93b9-9d488e1e91e4",
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

