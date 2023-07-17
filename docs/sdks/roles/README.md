# Roles

### Available Operations

* [Get](#get) - Invokes the c1.api.iam.v1.Roles.Get method.
* [List](#list) - Invokes the c1.api.iam.v1.Roles.List method.
* [Update](#update) - Invokes the c1.api.iam.v1.Roles.Update method.

## Get

Invokes the c1.api.iam.v1.Roles.Get method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Roles.Get(ctx, operations.C1APIIamV1RolesGetRequest{
        RoleID: "rerum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRolesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.C1APIIamV1RolesGetRequest](../../models/operations/c1apiiamv1rolesgetrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.C1APIIamV1RolesGetResponse](../../models/operations/c1apiiamv1rolesgetresponse.md), error**


## List

Invokes the c1.api.iam.v1.Roles.List method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Roles.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListRolesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APIIamV1RolesListResponse](../../models/operations/c1apiiamv1roleslistresponse.md), error**


## Update

Invokes the c1.api.iam.v1.Roles.Update method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/types"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Roles.Update(ctx, operations.C1APIIamV1RolesUpdateRequest{
        UpdateRoleRequest: &shared.UpdateRoleRequest{
            Role: &shared.Role{
                CreatedAt: types.MustTimeFromString("2021-09-24T14:32:36.483Z"),
                DeletedAt: types.MustTimeFromString("2022-01-15T21:54:49.962Z"),
                DisplayName: conductoroneapi.String("officiis"),
                ID: conductoroneapi.String("1ade008e-6f8c-45f3-90d8-cdb5a3418143"),
                Name: conductoroneapi.String("Alice Bailey"),
                Permissions: []string{
                    "laudantium",
                },
                ServiceRoles: []string{
                    "dolor",
                },
                SystemBuiltin: conductoroneapi.Bool(false),
                UpdatedAt: types.MustTimeFromString("2021-11-22T09:48:38.420Z"),
            },
            UpdateMask: conductoroneapi.String("consequuntur"),
        },
        RoleID: "ipsa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateRolesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.C1APIIamV1RolesUpdateRequest](../../models/operations/c1apiiamv1rolesupdaterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.C1APIIamV1RolesUpdateResponse](../../models/operations/c1apiiamv1rolesupdateresponse.md), error**

