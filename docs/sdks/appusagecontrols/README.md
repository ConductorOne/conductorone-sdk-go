# AppUsageControls

### Available Operations

* [Get](#get) - Invokes the c1.api.app.v1.AppUsageControlsService.Get method.

## Get

Invokes the c1.api.app.v1.AppUsageControlsService.Get method.

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
    res, err := s.AppUsageControls.Get(ctx, operations.C1APIAppV1AppUsageControlsServiceGetRequest{
        AppID: "aut",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAppUsageControlsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIAppV1AppUsageControlsServiceGetRequest](../../models/operations/c1apiappv1appusagecontrolsservicegetrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.C1APIAppV1AppUsageControlsServiceGetResponse](../../models/operations/c1apiappv1appusagecontrolsservicegetresponse.md), error**

