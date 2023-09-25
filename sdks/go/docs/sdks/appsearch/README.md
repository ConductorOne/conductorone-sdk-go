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
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppSearch.Search(ctx, shared.SearchAppsRequest{
        AppIds: []string{
            "aliquid",
        },
        DisplayName: conductoronesdkgo.String("laborum"),
        ExcludeAppIds: []string{
            "accusamus",
        },
        PageSize: conductoronesdkgo.Float64(2497.96),
        PageToken: conductoronesdkgo.String("occaecati"),
        Query: conductoronesdkgo.String("enim"),
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

