# SSFReceiverEvent

## Overview

### Available Operations

* [List](#list) - List

## List

Invokes the c1.api.ssf_receiver.v1.SSFReceiverEventService.List method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ssf_receiver.v1.SSFReceiverEventService.List" method="get" path="/api/v1/ssf-receiver-streams/{stream_id}/events" -->
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

    res, err := s.SSFReceiverEvent.List(ctx, operations.C1APISSFReceiverV1SSFReceiverEventServiceListRequest{
        StreamID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SSFReceiverEventServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APISSFReceiverV1SSFReceiverEventServiceListRequest](../../pkg/models/operations/c1apissfreceiverv1ssfreceivereventservicelistrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APISSFReceiverV1SSFReceiverEventServiceListResponse](../../pkg/models/operations/c1apissfreceiverv1ssfreceivereventservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |