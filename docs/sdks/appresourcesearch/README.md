# AppResourceSearch

### Available Operations

* [SearchAppResourceTypes](#searchappresourcetypes) - Search App Resource Types

## SearchAppResourceTypes

Search app resources based on filters specified in the request body.

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
    res, err := s.AppResourceSearch.SearchAppResourceTypes(ctx, shared.SearchAppResourceTypesRequest{
        AppIds: []string{
            "officia",
            "dolor",
            "debitis",
        },
        ExcludeResourceTypeIds: []string{
            "dolorum",
            "in",
            "in",
            "illum",
        },
        ExcludeResourceTypeTraitIds: []string{
            "rerum",
            "dicta",
            "magnam",
            "cumque",
        },
        PageSize: conductoroneapi.Float64(8137.98),
        PageToken: conductoroneapi.String("ea"),
        Query: conductoroneapi.String("aliquid"),
        ResourceTypeIds: []string{
            "accusamus",
            "non",
            "occaecati",
        },
        ResourceTypeTraitIds: []string{
            "accusamus",
            "delectus",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchAppResourceTypesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.SearchAppResourceTypesRequest](../../models/shared/searchappresourcetypesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.C1APIAppV1AppResourceSearchSearchAppResourceTypesResponse](../../models/operations/c1apiappv1appresourcesearchsearchappresourcetypesresponse.md), error**

