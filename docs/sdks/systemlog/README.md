# SystemLog
(*SystemLog*)

### Available Operations

* [ListEvents](#listevents) - List Events

## ListEvents

ListEvents pulls Events from the ConductorOne system.

 This endpoint should be used to synchorize the
 system log events to external systems.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"os"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: os.Getenv("BEARER_AUTH"),
            Oauth: os.Getenv("OAUTH"),
        }),
    )
    var request *shared.SystemLogServiceListEventsRequest = &shared.SystemLogServiceListEventsRequest{}
    ctx := context.Background()
    res, err := s.SystemLog.ListEvents(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.SystemLogServiceListEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [shared.SystemLogServiceListEventsRequest](../../pkg/models/shared/systemlogservicelisteventsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.C1APISystemlogV1SystemLogServiceListEventsResponse](../../pkg/models/operations/c1apisystemlogv1systemlogservicelisteventsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
