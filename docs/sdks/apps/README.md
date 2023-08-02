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
        CertifyPolicyID: conductoroneapi.String("ea"),
        Description: conductoroneapi.String("excepturi"),
        DisplayName: conductoroneapi.String("odit"),
        GrantPolicyID: conductoroneapi.String("ea"),
        MonthlyCostUsd: conductoroneapi.Float64(332.22),
        Owners: []string{
            "maiores",
        },
        RevokePolicyID: conductoroneapi.String("quidem"),
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
        ID: "576b0d5f-0d30-4c5f-bb25-87053202c73d",
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
        ID: "5fe9b90c-2890-49b3-be49-a8d9cbf48633",
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
        PageSize: conductoroneapi.Float64(2224.43),
        PageToken: conductoroneapi.String("qui"),
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
                CertifyPolicyID: conductoroneapi.String("ipsum"),
                Description: conductoroneapi.String("hic"),
                DisplayName: conductoroneapi.String("excepturi"),
                GrantPolicyID: conductoroneapi.String("cum"),
                IconURL: conductoroneapi.String("voluptate"),
                MonthlyCostUsd: conductoroneapi.Float64(4904.59),
                RevokePolicyID: conductoroneapi.String("reiciendis"),
            },
            UpdateMask: conductoroneapi.String("amet"),
        },
        ID: "a4100674-ebf6-4928-8d1b-a77a89ebf737",
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

