# AppOwners
(*.AppOwners*)

### Available Operations

* [Add](#add) - Add
* [List](#list) - List
* [Remove](#remove) - Remove

## Add

Adds an owner to an app.

### Example Usage

```go
package main

import(
	"context"
	"log"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/operations"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppOwners.Add(ctx, operations.C1APIAppV1AppOwnersAddRequest{
        AddAppOwnerRequest: &shared.AddAppOwnerRequest{},
        AppID: "string",
        UserID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddAppOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.C1APIAppV1AppOwnersAddRequest](../../models/operations/c1apiappv1appownersaddrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.C1APIAppV1AppOwnersAddResponse](../../models/operations/c1apiappv1appownersaddresponse.md), error**


## List

List owners of an app.

### Example Usage

```go
package main

import(
	"context"
	"log"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/operations"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppOwners.List(ctx, operations.C1APIAppV1AppOwnersListRequest{
        AppID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.C1APIAppV1AppOwnersListRequest](../../models/operations/c1apiappv1appownerslistrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.C1APIAppV1AppOwnersListResponse](../../models/operations/c1apiappv1appownerslistresponse.md), error**


## Remove

Removes an owner from an app.

### Example Usage

```go
package main

import(
	"context"
	"log"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/operations"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppOwners.Remove(ctx, operations.C1APIAppV1AppOwnersRemoveRequest{
        RemoveAppOwnerRequest: &shared.RemoveAppOwnerRequest{},
        AppID: "string",
        UserID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RemoveAppOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.C1APIAppV1AppOwnersRemoveRequest](../../models/operations/c1apiappv1appownersremoverequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.C1APIAppV1AppOwnersRemoveResponse](../../models/operations/c1apiappv1appownersremoveresponse.md), error**

