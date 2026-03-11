# A2Ui

## Overview

### Available Operations

* [CreateSurfaceFeedback](#createsurfacefeedback) - Create Surface Feedback
* [ListSurfaceFeedback](#listsurfacefeedback) - List Surface Feedback
* [ListSurfaces](#listsurfaces) - List Surfaces
* [SubmitAction](#submitaction) - Submit Action

## CreateSurfaceFeedback

CreateSurfaceFeedback submits feedback for a surface with a snapshot.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.a2ui.v1.A2UIService.CreateSurfaceFeedback" method="post" path="/api/v1/a2ui/surfaces/{surface_id}/feedback" -->
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

    res, err := s.A2UI.CreateSurfaceFeedback(ctx, operations.C1APIA2uiV1A2UIServiceCreateSurfaceFeedbackRequest{
        SurfaceID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.A2UIServiceCreateSurfaceFeedbackResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIA2uiV1A2UIServiceCreateSurfaceFeedbackRequest](../../pkg/models/operations/c1apia2uiv1a2uiservicecreatesurfacefeedbackrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APIA2uiV1A2UIServiceCreateSurfaceFeedbackResponse](../../pkg/models/operations/c1apia2uiv1a2uiservicecreatesurfacefeedbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSurfaceFeedback

ListSurfaceFeedback lists feedback for a surface.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.a2ui.v1.A2UIService.ListSurfaceFeedback" method="get" path="/api/v1/a2ui/surfaces/{surface_id}/feedback" -->
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

    res, err := s.A2UI.ListSurfaceFeedback(ctx, operations.C1APIA2uiV1A2UIServiceListSurfaceFeedbackRequest{
        SurfaceID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.A2UIServiceListSurfaceFeedbackResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIA2uiV1A2UIServiceListSurfaceFeedbackRequest](../../pkg/models/operations/c1apia2uiv1a2uiservicelistsurfacefeedbackrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIA2uiV1A2UIServiceListSurfaceFeedbackResponse](../../pkg/models/operations/c1apia2uiv1a2uiservicelistsurfacefeedbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSurfaces

ListSurfaces returns active surfaces for a conversation.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.a2ui.v1.A2UIService.ListSurfaces" method="get" path="/api/v1/a2ui/conversations/{conversation_id}/surfaces" -->
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

    res, err := s.A2UI.ListSurfaces(ctx, operations.C1APIA2uiV1A2UIServiceListSurfacesRequest{
        ConversationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.A2UIServiceListSurfacesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIA2uiV1A2UIServiceListSurfacesRequest](../../pkg/models/operations/c1apia2uiv1a2uiservicelistsurfacesrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APIA2uiV1A2UIServiceListSurfacesResponse](../../pkg/models/operations/c1apia2uiv1a2uiservicelistsurfacesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SubmitAction

SubmitAction handles user actions on A2UI surfaces.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.a2ui.v1.A2UIService.SubmitAction" method="post" path="/api/v1/a2ui/surfaces/{surface_id}/actions" -->
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

    res, err := s.A2UI.SubmitAction(ctx, operations.C1APIA2uiV1A2UIServiceSubmitActionRequest{
        SurfaceID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.A2UIServiceSubmitActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIA2uiV1A2UIServiceSubmitActionRequest](../../pkg/models/operations/c1apia2uiv1a2uiservicesubmitactionrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APIA2uiV1A2UIServiceSubmitActionResponse](../../pkg/models/operations/c1apia2uiv1a2uiservicesubmitactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |