# Roles
(*Roles*)

### Available Operations

* [Get](#get) - Get
* [List](#list) - List
* [Update](#update) - Update

## Get

Get a role by id.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"context"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Roles.Get(ctx, operations.C1APIIamV1RolesGetRequest{
        RoleID: "string",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.C1APIIamV1RolesGetRequest](../../pkg/models/operations/c1apiiamv1rolesgetrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.C1APIIamV1RolesGetResponse](../../pkg/models/operations/c1apiiamv1rolesgetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

List all roles for the current user.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"context"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Roles.List(ctx, operations.C1APIIamV1RolesListRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.ListRolesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.C1APIIamV1RolesListRequest](../../pkg/models/operations/c1apiiamv1roleslistrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.C1APIIamV1RolesListResponse](../../pkg/models/operations/c1apiiamv1roleslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Update a role by passing a Role object.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"context"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Roles.Update(ctx, operations.C1APIIamV1RolesUpdateRequest{
        RoleID: "string",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.C1APIIamV1RolesUpdateRequest](../../pkg/models/operations/c1apiiamv1rolesupdaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.C1APIIamV1RolesUpdateResponse](../../pkg/models/operations/c1apiiamv1rolesupdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
