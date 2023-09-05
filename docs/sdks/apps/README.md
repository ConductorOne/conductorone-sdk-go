# Apps

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List
* [Update](#update) - Update

## Create

Create a new app.

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
    res, err := s.Apps.Create(ctx, shared.CreateAppRequest{
        CertifyPolicyID: conductoroneapi.String("nihil"),
        Description: conductoroneapi.String("repellat"),
        DisplayName: conductoroneapi.String("quibusdam"),
        GrantPolicyID: conductoroneapi.String("sed"),
        MonthlyCostUsd: conductoroneapi.Float64(9046.48),
        Owners: []string{
            "accusantium",
            "consequuntur",
            "praesentium",
            "natus",
        },
        RevokePolicyID: conductoroneapi.String("magni"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.CreateAppRequest](../../models/shared/createapprequest.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.C1APIAppV1AppsCreateResponse](../../models/operations/c1apiappv1appscreateresponse.md), error**


## Delete

Delete an app.

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
    res, err := s.Apps.Delete(ctx, operations.C1APIAppV1AppsDeleteRequest{
        DeleteAppRequest: &shared.DeleteAppRequest{},
        ID: "1cddc692-601f-4b57-ab0d-5f0d30c5fbb2",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.C1APIAppV1AppsDeleteRequest](../../models/operations/c1apiappv1appsdeleterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.C1APIAppV1AppsDeleteResponse](../../models/operations/c1apiappv1appsdeleteresponse.md), error**


## Get

Get an app by ID.

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
    res, err := s.Apps.Get(ctx, operations.C1APIAppV1AppsGetRequest{
        ID: "58705320-2c73-4d5f-a9b9-0c28909b3fe4",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.C1APIAppV1AppsGetRequest](../../models/operations/c1apiappv1appsgetrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.C1APIAppV1AppsGetResponse](../../models/operations/c1apiappv1appsgetresponse.md), error**


## List

List all apps.

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
    res, err := s.Apps.List(ctx, operations.C1APIAppV1AppsListRequest{
        PageSize: conductoroneapi.Float64(6139.66),
        PageToken: conductoroneapi.String("dolorum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.C1APIAppV1AppsListRequest](../../models/operations/c1apiappv1appslistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.C1APIAppV1AppsListResponse](../../models/operations/c1apiappv1appslistresponse.md), error**


## Update

Update an existing app.

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
    res, err := s.Apps.Update(ctx, operations.C1APIAppV1AppsUpdateRequest{
        UpdateAppRequestInput: &shared.UpdateAppRequestInput{
            App: &shared.AppInput{
                CertifyPolicyID: conductoroneapi.String("deleniti"),
                Description: conductoroneapi.String("pariatur"),
                DisplayName: conductoroneapi.String("provident"),
                GrantPolicyID: conductoroneapi.String("nobis"),
                IconURL: conductoroneapi.String("libero"),
                MonthlyCostUsd: conductoroneapi.Float64(9644.9),
                RevokePolicyID: conductoroneapi.String("quaerat"),
            },
            UpdateMask: conductoroneapi.String("quos"),
        },
        ID: "633323f9-b77f-43a4-9006-74ebf69280d1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.C1APIAppV1AppsUpdateRequest](../../models/operations/c1apiappv1appsupdaterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.C1APIAppV1AppsUpdateResponse](../../models/operations/c1apiappv1appsupdateresponse.md), error**

