# AppResourceSearch
(*AppResourceSearch*)

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
    res, err := s.AppResourceSearch.SearchAppResourceTypes(ctx, shared.SearchAppResourceTypesRequest{
        AppIds: []string{
            "quibusdam",
        },
        ExcludeResourceTypeIds: []string{
            "labore",
        },
        ExcludeResourceTypeTraitIds: []string{
            "modi",
        },
        PageSize: conductoronesdkgo.Float64(1831.91),
        PageToken: conductoronesdkgo.String("aliquid"),
        Query: conductoronesdkgo.String("cupiditate"),
        ResourceTypeIds: []string{
            "quos",
        },
        ResourceTypeTraitIds: []string{
            "perferendis",
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

