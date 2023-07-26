# AppEntitlements

### Available Operations

* [Get](#get) - Get
* [List](#list) - List
* [ListForAppResource](#listforappresource) - List For App Resource
* [ListForAppUser](#listforappuser) - List For App User
* [ListGroups](#listgroups) - List Groups
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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlements.List(ctx, operations.C1APIAppV1AppEntitlementsListRequest{
        AppID: "doloribus",
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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlements.ListForAppResource(ctx, operations.C1APIAppV1AppEntitlementsListForAppResourceRequest{
        AppID: "sapiente",
        AppResourceID: "architecto",
        AppResourceTypeID: "mollitia",
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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlements.ListForAppUser(ctx, operations.C1APIAppV1AppEntitlementsListForAppUserRequest{
        AppID: "dolorem",
        AppUserID: "culpa",
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


## ListGroups

 List app groups associated with an app entitlement.


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
    res, err := s.AppEntitlements.ListGroups(ctx, operations.C1APIAppV1AppEntitlementsListGroupsRequest{
        AppEntitlementID: "consequuntur",
        AppID: "repellat",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppEntitlementGroupsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APIAppV1AppEntitlementsListGroupsRequest](../../models/operations/c1apiappv1appentitlementslistgroupsrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.C1APIAppV1AppEntitlementsListGroupsResponse](../../models/operations/c1apiappv1appentitlementslistgroupsresponse.md), error**


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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlements.ListUsers(ctx, operations.C1APIAppV1AppEntitlementsListUsersRequest{
        AppEntitlementID: "mollitia",
        AppID: "occaecati",
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
                        AppID: conductoroneapi.String("numquam"),
                        EntitlementID: conductoroneapi.String("commodi"),
                    },
                    ManualProvision: &shared.ManualProvision{
                        Instructions: conductoroneapi.String("quam"),
                        UserIds: []string{
                            "velit",
                            "error",
                        },
                    },
                },
                AppID: conductoroneapi.String("quia"),
                AppResourceID: conductoroneapi.String("quis"),
                AppResourceTypeID: conductoroneapi.String("vitae"),
                CertifyPolicyID: conductoroneapi.String("laborum"),
                ComplianceFrameworkValueIds: []string{
                    "enim",
                    "odit",
                    "quo",
                },
                Description: conductoroneapi.String("sequi"),
                DisplayName: conductoroneapi.String("tenetur"),
                DurationGrant: conductoroneapi.String("ipsam"),
                DurationUnset: &shared.AppEntitlementDurationUnset{},
                EmergencyGrantEnabled: conductoroneapi.Bool(false),
                EmergencyGrantPolicyID: conductoroneapi.String("id"),
                GrantPolicyID: conductoroneapi.String("possimus"),
                RevokePolicyID: conductoroneapi.String("aut"),
                RiskLevelValueID: conductoroneapi.String("quasi"),
                Slug: conductoroneapi.String("error"),
            },
            AppEntitlementExpandMask: &shared.AppEntitlementExpandMask{
                Paths: []string{
                    "laborum",
                    "quasi",
                    "reiciendis",
                    "voluptatibus",
                },
            },
            UpdateMask: conductoroneapi.String("vero"),
        },
        AppID: "nihil",
        ID: "8f097b00-74f1-4547-9b5e-6e13b99d488e",
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

