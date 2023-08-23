# Task

### Available Operations

* [CreateGrantTask](#creategranttask) - Create Grant Task
* [CreateRevokeTask](#createrevoketask) - Create Revoke Task
* [Get](#get) - Get

## CreateGrantTask

Create a grant task

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
    res, err := s.Task.CreateGrantTask(ctx, shared.TaskServiceCreateGrantRequest{
        TaskExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "vero",
                "est",
                "harum",
                "sequi",
            },
        },
        TaskGrantSource: &shared.TaskGrantSource{
            ExternalURL: conductoroneapi.String("doloribus"),
            IntegrationID: conductoroneapi.String("repudiandae"),
        },
        AppEntitlementID: "optio",
        AppID: "occaecati",
        AppUserID: conductoroneapi.String("nemo"),
        Description: conductoroneapi.String("voluptate"),
        EmergencyAccess: conductoroneapi.Bool(false),
        GrantDuration: conductoroneapi.String("blanditiis"),
        IdentityUserID: conductoroneapi.String("officia"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskServiceCreateGrantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.TaskServiceCreateGrantRequest](../../models/shared/taskservicecreategrantrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.C1APITaskV1TaskServiceCreateGrantTaskResponse](../../models/operations/c1apitaskv1taskservicecreategranttaskresponse.md), error**


## CreateRevokeTask

Create a revoke task

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
    res, err := s.Task.CreateRevokeTask(ctx, shared.TaskServiceCreateRevokeRequest{
        TaskExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "numquam",
                "nemo",
            },
        },
        AppEntitlementID: "quos",
        AppID: "eius",
        AppUserID: conductoroneapi.String("aspernatur"),
        Description: conductoroneapi.String("ducimus"),
        IdentityUserID: conductoroneapi.String("nesciunt"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskServiceCreateRevokeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [shared.TaskServiceCreateRevokeRequest](../../models/shared/taskservicecreaterevokerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.C1APITaskV1TaskServiceCreateRevokeTaskResponse](../../models/operations/c1apitaskv1taskservicecreaterevoketaskresponse.md), error**


## Get

Get a task by ID

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
    res, err := s.Task.Get(ctx, operations.C1APITaskV1TaskServiceGetRequest{
        ID: "a8418d16-2309-4fb0-9299-21aefb9f58c4",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.C1APITaskV1TaskServiceGetRequest](../../models/operations/c1apitaskv1taskservicegetrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.C1APITaskV1TaskServiceGetResponse](../../models/operations/c1apitaskv1taskservicegetresponse.md), error**

