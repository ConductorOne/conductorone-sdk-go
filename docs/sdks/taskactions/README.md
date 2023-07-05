# TaskActions

### Available Operations

* [Approve](#approve) - Invokes the c1.api.task.v1.TaskActionsService.Approve method.
* [Comment](#comment) - Invokes the c1.api.task.v1.TaskActionsService.Comment method.
* [Deny](#deny) - Invokes the c1.api.task.v1.TaskActionsService.Deny method.

## Approve

Invokes the c1.api.task.v1.TaskActionsService.Approve method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.TaskActions.Approve(ctx, operations.C1APITaskV1TaskActionsServiceApproveRequest{
        TaskActionsServiceApproveRequest: &shared.TaskActionsServiceApproveRequest{
            Comment: conductoroneapi.String("cum"),
            ExpandMask: &shared.TaskExpandMask{
                Paths: []string{
                    "quas",
                    "hic",
                    "nesciunt",
                },
            },
            PolicyStepID: conductoroneapi.String("culpa"),
        },
        TaskID: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskActionsServiceApproveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APITaskV1TaskActionsServiceApproveRequest](../../models/operations/c1apitaskv1taskactionsserviceapproverequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.C1APITaskV1TaskActionsServiceApproveResponse](../../models/operations/c1apitaskv1taskactionsserviceapproveresponse.md), error**


## Comment

Invokes the c1.api.task.v1.TaskActionsService.Comment method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.TaskActions.Comment(ctx, operations.C1APITaskV1TaskActionsServiceCommentRequest{
        TaskActionsServiceCommentRequest: &shared.TaskActionsServiceCommentRequest{
            Comment: conductoroneapi.String("pariatur"),
            ExpandMask: &shared.TaskExpandMask{
                Paths: []string{
                    "hic",
                    "exercitationem",
                    "nobis",
                },
            },
        },
        TaskID: "sit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskActionsServiceCommentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APITaskV1TaskActionsServiceCommentRequest](../../models/operations/c1apitaskv1taskactionsservicecommentrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.C1APITaskV1TaskActionsServiceCommentResponse](../../models/operations/c1apitaskv1taskactionsservicecommentresponse.md), error**


## Deny

Invokes the c1.api.task.v1.TaskActionsService.Deny method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.TaskActions.Deny(ctx, operations.C1APITaskV1TaskActionsServiceDenyRequest{
        TaskActionsServiceDenyRequest: &shared.TaskActionsServiceDenyRequest{
            Comment: conductoroneapi.String("rerum"),
            ExpandMask: &shared.TaskExpandMask{
                Paths: []string{
                    "reiciendis",
                },
            },
            PolicyStepID: conductoroneapi.String("explicabo"),
        },
        TaskID: "asperiores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskActionsServiceDenyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.C1APITaskV1TaskActionsServiceDenyRequest](../../models/operations/c1apitaskv1taskactionsservicedenyrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.C1APITaskV1TaskActionsServiceDenyResponse](../../models/operations/c1apitaskv1taskactionsservicedenyresponse.md), error**

