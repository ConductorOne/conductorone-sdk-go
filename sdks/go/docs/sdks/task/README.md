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
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Task.CreateGrantTask(ctx, shared.TaskServiceCreateGrantRequest{
        TaskExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "ratione",
            },
        },
        TaskGrantSource: &shared.TaskGrantSource{
            ExternalURL: conductoronesdkgo.String("laborum"),
            IntegrationID: conductoronesdkgo.String("distinctio"),
        },
        AppEntitlementID: "voluptatum",
        AppID: "rem",
        AppUserID: conductoronesdkgo.String("aliquam"),
        Description: conductoronesdkgo.String("ad"),
        EmergencyAccess: conductoronesdkgo.Bool(false),
        GrantDuration: conductoronesdkgo.String("repellat"),
        IdentityUserID: conductoronesdkgo.String("alias"),
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
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Task.CreateRevokeTask(ctx, shared.TaskServiceCreateRevokeRequest{
        TaskExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "corporis",
            },
        },
        AppEntitlementID: "perspiciatis",
        AppID: "nihil",
        AppUserID: conductoronesdkgo.String("mollitia"),
        Description: conductoronesdkgo.String("voluptas"),
        IdentityUserID: conductoronesdkgo.String("alias"),
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
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Task.Get(ctx, operations.C1APITaskV1TaskServiceGetRequest{
        ID: "ff2a54a3-1e94-4764-a3e8-65e7956f9251",
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

