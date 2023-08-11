# AppSearch

### Available Operations

* [Search](#search) - Search

## Search

Search apps based on filters specified in the request body.

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
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppSearch.Search(ctx, shared.SearchAppsRequest{
        AppIds: []string{
            "expedita",
            "nihil",
        },
        DisplayName: conductoroneapi.String("repellat"),
        ExcludeAppIds: []string{
            "sed",
            "saepe",
            "pariatur",
            "accusantium",
        },
        PageSize: conductoroneapi.Float64(1624.93),
        PageToken: conductoroneapi.String("praesentium"),
        Query: conductoroneapi.String("natus"),
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

