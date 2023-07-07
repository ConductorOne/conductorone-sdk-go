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
        RoleID: "vel",
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
                CreatedAt: types.MustTimeFromString("2022-01-20T13:10:25.426Z"),
                DeletedAt: types.MustTimeFromString("2021-11-03T10:56:47.600Z"),
                DisplayName: conductoroneapi.String("minima"),
                ID: conductoroneapi.String("bc0ab3c2-0c4f-4378-9fd8-71f99dd2efd1"),
                Name: conductoroneapi.String("Ann Murphy"),
                Permissions: []string{
                    "quae",
                    "earum",
                    "vel",
                    "in",
                },
                ServiceRoles: []string{
                    "libero",
                    "illum",
                },
                SystemBuiltin: conductoroneapi.Bool(false),
                UpdatedAt: types.MustTimeFromString("2022-12-07T16:30:09.640Z"),
            },
            UpdateMask: conductoroneapi.String("aliquam"),
        },
        RoleID: "sapiente",
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

