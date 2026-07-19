# TaskActions

## Overview

### Available Operations

* [Approve](#approve) - Approve
* [ApproveWithStepUp](#approvewithstepup) - Approve With Step Up
* [Close](#close) - Close
* [Comment](#comment) - Comment
* [Deny](#deny) - Deny
* [EscalateToEmergencyAccess](#escalatetoemergencyaccess) - Escalate To Emergency Access
* [HardReset](#hardreset) - Hard Reset
* [ProcessNow](#processnow) - Process Now
* [Reassign](#reassign) - Reassign
* [Restart](#restart) - Restart
* [SkipStep](#skipstep) - Skip Step
* [UpdateGrantDuration](#updategrantduration) - Update Grant Duration
* [UpdateRequestData](#updaterequestdata) - Update Request Data

## Approve

Approve the specified policy step on a task.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.Approve" method="post" path="/api/v1/tasks/{task_id}/action/approve" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.Approve(ctx, operations.C1APITaskV1TaskActionsServiceApproveRequest{
        TaskID: "<id>",
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

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.C1APITaskV1TaskActionsServiceApproveRequest](../../pkg/models/operations/c1apitaskv1taskactionsserviceapproverequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.C1APITaskV1TaskActionsServiceApproveResponse](../../pkg/models/operations/c1apitaskv1taskactionsserviceapproveresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ApproveWithStepUp

Approve a task that requires step-up authentication. If a verified step-up transaction ID is provided, the approval is processed immediately. Otherwise, a redirect URL is returned for the caller to complete authentication first.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.ApproveWithStepUp" method="post" path="/api/v1/tasks/{task_id}/action/approve-with-step-up" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.ApproveWithStepUp(ctx, operations.C1APITaskV1TaskActionsServiceApproveWithStepUpRequest{
        TaskID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceApproveWithStepUpResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APITaskV1TaskActionsServiceApproveWithStepUpRequest](../../pkg/models/operations/c1apitaskv1taskactionsserviceapprovewithstepuprequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APITaskV1TaskActionsServiceApproveWithStepUpResponse](../../pkg/models/operations/c1apitaskv1taskactionsserviceapprovewithstepupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Close

Close a task, ending its workflow. Async — returns a ticketActionId on
 accept and queues the close; the task's state field is not updated
 synchronously. Poll task_service_get to observe the transition to
 TASK_STATE_CLOSED. When the task has an active provision step, close
 may no-op — finish or cancel the provision step first
 (mark_provision_complete / _errored / _cancelled) before closing.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.Close" method="post" path="/api/v1/tasks/{task_id}/action/close" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.Close(ctx, operations.C1APITaskV1TaskActionsServiceCloseRequest{
        TaskID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceCloseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APITaskV1TaskActionsServiceCloseRequest](../../pkg/models/operations/c1apitaskv1taskactionsservicecloserequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APITaskV1TaskActionsServiceCloseResponse](../../pkg/models/operations/c1apitaskv1taskactionsservicecloseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Comment

Post a comment on a task without changing its state.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.Comment" method="post" path="/api/v1/tasks/{task_id}/action/comment" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.Comment(ctx, operations.C1APITaskV1TaskActionsServiceCommentRequest{
        TaskID: "<id>",
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

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.C1APITaskV1TaskActionsServiceCommentRequest](../../pkg/models/operations/c1apitaskv1taskactionsservicecommentrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.C1APITaskV1TaskActionsServiceCommentResponse](../../pkg/models/operations/c1apitaskv1taskactionsservicecommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Deny

Deny the specified policy step on a task. In multi-step policies, this may route to fallback steps rather than finalizing the task outcome.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.Deny" method="post" path="/api/v1/tasks/{task_id}/action/deny" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.Deny(ctx, operations.C1APITaskV1TaskActionsServiceDenyRequest{
        TaskID: "<id>",
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APITaskV1TaskActionsServiceDenyRequest](../../pkg/models/operations/c1apitaskv1taskactionsservicedenyrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.C1APITaskV1TaskActionsServiceDenyResponse](../../pkg/models/operations/c1apitaskv1taskactionsservicedenyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## EscalateToEmergencyAccess

Escalate a grant task to use the emergency access policy, bypassing the normal approval flow. Only valid for grant tasks.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.EscalateToEmergencyAccess" method="post" path="/api/v1/tasks/{task_id}/action/escalate" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.EscalateToEmergencyAccess(ctx, operations.C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessRequest{
        TaskID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskServiceActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [operations.C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessRequest](../../pkg/models/operations/c1apitaskv1taskactionsserviceescalatetoemergencyaccessrequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |
| `opts`                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                                       | The options for this request.                                                                                                                                            |

### Response

**[*operations.C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessResponse](../../pkg/models/operations/c1apitaskv1taskactionsserviceescalatetoemergencyaccessresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## HardReset

Reset a task and recalculate its policy from scratch. Unlike Restart, this re-evaluates which policy applies to the task.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.HardReset" method="post" path="/api/v1/tasks/{task_id}/action/reset" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.HardReset(ctx, operations.C1APITaskV1TaskActionsServiceHardResetRequest{
        TaskID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceHardResetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.C1APITaskV1TaskActionsServiceHardResetRequest](../../pkg/models/operations/c1apitaskv1taskactionsservicehardresetrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.C1APITaskV1TaskActionsServiceHardResetResponse](../../pkg/models/operations/c1apitaskv1taskactionsservicehardresetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ProcessNow

Trigger immediate processing of a task, bypassing any scheduled wait. For tasks linked to an external system, this also attempts to sync the external state.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.ProcessNow" method="post" path="/api/v1/tasks/{task_id}/action/process" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.ProcessNow(ctx, operations.C1APITaskV1TaskActionsServiceProcessNowRequest{
        TaskID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceProcessNowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.C1APITaskV1TaskActionsServiceProcessNowRequest](../../pkg/models/operations/c1apitaskv1taskactionsserviceprocessnowrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `opts`                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.C1APITaskV1TaskActionsServiceProcessNowResponse](../../pkg/models/operations/c1apitaskv1taskactionsserviceprocessnowresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Reassign

Reassign a task's current policy step to a different set of users. The target step must be an approval, provision, or form step.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.Reassign" method="post" path="/api/v1/tasks/{task_id}/action/reassign" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.Reassign(ctx, operations.C1APITaskV1TaskActionsServiceReassignRequest{
        TaskID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceReassignResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.C1APITaskV1TaskActionsServiceReassignRequest](../../pkg/models/operations/c1apitaskv1taskactionsservicereassignrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.C1APITaskV1TaskActionsServiceReassignResponse](../../pkg/models/operations/c1apitaskv1taskactionsservicereassignresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Restart

Restart a task, returning it to the beginning of its current policy workflow.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.Restart" method="post" path="/api/v1/tasks/{task_id}/action/restart" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.Restart(ctx, operations.C1APITaskV1TaskActionsServiceRestartRequest{
        TaskID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceRestartResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.C1APITaskV1TaskActionsServiceRestartRequest](../../pkg/models/operations/c1apitaskv1taskactionsservicerestartrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.C1APITaskV1TaskActionsServiceRestartResponse](../../pkg/models/operations/c1apitaskv1taskactionsservicerestartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SkipStep

Skip a specific policy step in a task, advancing the task to the next step in the workflow.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.SkipStep" method="post" path="/api/v1/tasks/{task_id}/action/skip-step" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.SkipStep(ctx, operations.C1APITaskV1TaskActionsServiceSkipStepRequest{
        TaskID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskServiceActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.C1APITaskV1TaskActionsServiceSkipStepRequest](../../pkg/models/operations/c1apitaskv1taskactionsserviceskipsteprequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.C1APITaskV1TaskActionsServiceSkipStepResponse](../../pkg/models/operations/c1apitaskv1taskactionsserviceskipstepresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateGrantDuration

Update the grant duration for a task. Only applies to grant tasks with a single entitlement that are not in a provision step. The new duration must not exceed the entitlement's maximum allowed provision time.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.UpdateGrantDuration" method="post" path="/api/v1/tasks/{task_id}/action/update-grant-duration" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.UpdateGrantDuration(ctx, operations.C1APITaskV1TaskActionsServiceUpdateGrantDurationRequest{
        TaskID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskServiceActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APITaskV1TaskActionsServiceUpdateGrantDurationRequest](../../pkg/models/operations/c1apitaskv1taskactionsserviceupdategrantdurationrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APITaskV1TaskActionsServiceUpdateGrantDurationResponse](../../pkg/models/operations/c1apitaskv1taskactionsserviceupdategrantdurationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateRequestData

Update the request data on a task that is currently in a form step. The submitted data is validated against the form schema before being applied.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.UpdateRequestData" method="post" path="/api/v1/tasks/{task_id}/action/update-request-data" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.TaskActions.UpdateRequestData(ctx, operations.C1APITaskV1TaskActionsServiceUpdateRequestDataRequest{
        TaskID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskServiceActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APITaskV1TaskActionsServiceUpdateRequestDataRequest](../../pkg/models/operations/c1apitaskv1taskactionsserviceupdaterequestdatarequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APITaskV1TaskActionsServiceUpdateRequestDataResponse](../../pkg/models/operations/c1apitaskv1taskactionsserviceupdaterequestdataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |