# AppSearch

### Available Operations

* [Search](#search) - Invokes the c1.api.app.v1.AppSearch.Search method.

## Search

Invokes the c1.api.app.v1.AppSearch.Search method.

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
    res, err := s.AppSearch.Search(ctx, shared.SearchAppsRequest{
        AppIds: []string{
            "quo",
        },
        ExcludeAppIds: []string{
            "tenetur",
        },
        PageSize: conductoroneapi.Float64(3687.25),
        PageToken: conductoroneapi.String("id"),
        Query: conductoroneapi.String("possimus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchAppsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.SearchAppsRequest](../../models/shared/searchappsrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.C1APIAppV1AppSearchSearchResponse](../../models/operations/c1apiappv1appsearchsearchresponse.md), error**
