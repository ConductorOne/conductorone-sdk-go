# Task

### Available Operations

* [CreateGrantTask](#creategranttask) - Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.
* [CreateRevokeTask](#createrevoketask) - Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.
* [Get](#get) - Invokes the c1.api.task.v1.TaskService.Get method.

## CreateGrantTask

Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.

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
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Task.CreateGrantTask(ctx, shared.TaskServiceCreateGrantRequest{
        AppEntitlementID: conductoroneapi.String("cum"),
        AppID: conductoroneapi.String("consectetur"),
        AppUserID: conductoroneapi.String("in"),
        Description: conductoroneapi.String("exercitationem"),
        EmergencyAccess: conductoroneapi.Bool(false),
        ExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "facere",
                "numquam",
                "doloribus",
                "suscipit",
            },
        },
        GrantDuration: conductoroneapi.String("reiciendis"),
        IdentityUserID: conductoroneapi.String("quidem"),
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

Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.

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
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Task.CreateRevokeTask(ctx, shared.TaskServiceCreateRevokeRequest{
        AppEntitlementID: conductoroneapi.String("saepe"),
        AppID: conductoroneapi.String("necessitatibus"),
        AppUserID: conductoroneapi.String("dolore"),
        Description: conductoroneapi.String("sunt"),
        ExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "adipisci",
                "non",
                "amet",
                "beatae",
            },
        },
        IdentityUserID: conductoroneapi.String("dignissimos"),
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

Invokes the c1.api.task.v1.TaskService.Get method.

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
    res, err := s.Task.Get(ctx, operations.C1APITaskV1TaskServiceGetRequest{
        ID: "fe35b60e-b1ea-4426-955b-a3c28744ed53",
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

