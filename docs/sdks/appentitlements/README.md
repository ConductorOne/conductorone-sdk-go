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
        AppID: "beatae",
        ID: "6742cb73-9205-4929-b96f-ea7596eb10fa",
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
        AppID: "mollitia",
        PageSize: conductoroneapi.Float64(6706.38),
        PageToken: conductoroneapi.String("dolores"),
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
        AppID: "dolorem",
        AppResourceID: "corporis",
        AppResourceTypeID: "explicabo",
        PageSize: conductoroneapi.Float64(7506.86),
        PageToken: conductoroneapi.String("enim"),
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
        AppID: "omnis",
        AppUserID: "nemo",
        PageSize: conductoroneapi.Float64(3250.47),
        PageToken: conductoroneapi.String("excepturi"),
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
        AppEntitlementID: "accusantium",
        AppID: "iure",
        PageSize: conductoroneapi.Float64(6342.74),
        PageToken: conductoroneapi.String("doloribus"),
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
                        AppID: conductoroneapi.String("sapiente"),
                        EntitlementID: conductoroneapi.String("architecto"),
                    },
                    ManualProvision: &shared.ManualProvision{
                        Instructions: conductoroneapi.String("mollitia"),
                        UserIds: []string{
                            "culpa",
                        },
                    },
                },
                AppID: conductoroneapi.String("consequuntur"),
                AppResourceID: conductoroneapi.String("repellat"),
                AppResourceTypeID: conductoroneapi.String("mollitia"),
                CertifyPolicyID: conductoroneapi.String("occaecati"),
                ComplianceFrameworkValueIds: []string{
                    "commodi",
                    "quam",
                },
                Description: conductoroneapi.String("molestiae"),
                DisplayName: conductoroneapi.String("velit"),
                DurationGrant: conductoroneapi.String("error"),
                DurationUnset: &shared.AppEntitlementDurationUnset{},
                EmergencyGrantEnabled: conductoroneapi.Bool(false),
                EmergencyGrantPolicyID: conductoroneapi.String("quia"),
                GrantPolicyID: conductoroneapi.String("quis"),
                RevokePolicyID: conductoroneapi.String("vitae"),
                RiskLevelValueID: conductoroneapi.String("laborum"),
                Slug: conductoroneapi.String("animi"),
            },
            AppEntitlementExpandMask: &shared.AppEntitlementExpandMask{
                Paths: []string{
                    "odit",
                    "quo",
                },
            },
            UpdateMask: conductoroneapi.String("sequi"),
        },
        AppID: "tenetur",
        ID: "5ad019da-1ffe-478f-897b-0074f15471b5",
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

