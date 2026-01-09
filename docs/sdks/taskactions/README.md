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

Invokes the c1.api.task.v1.TaskActionsService.Approve method.

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

Invokes the c1.api.task.v1.TaskActionsService.ApproveWithStepUp method.

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

Invokes the c1.api.task.v1.TaskActionsService.Close method.

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

Invokes the c1.api.task.v1.TaskActionsService.Comment method.

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

Invokes the c1.api.task.v1.TaskActionsService.Deny method.

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

Invokes the c1.api.task.v1.TaskActionsService.EscalateToEmergencyAccess method.

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

Invokes the c1.api.task.v1.TaskActionsService.HardReset method.

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

Invokes the c1.api.task.v1.TaskActionsService.ProcessNow method.

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

Invokes the c1.api.task.v1.TaskActionsService.Reassign method.

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

Invokes the c1.api.task.v1.TaskActionsService.Restart method.

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

Invokes the c1.api.task.v1.TaskActionsService.SkipStep method.

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

Invokes the c1.api.task.v1.TaskActionsService.UpdateGrantDuration method.

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

Invokes the c1.api.task.v1.TaskActionsService.UpdateRequestData method.

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