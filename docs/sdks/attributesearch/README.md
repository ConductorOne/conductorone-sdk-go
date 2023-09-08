# AttributeSearch

### Available Operations

* [SearchAttributeValues](#searchattributevalues) - Search Attribute Values

## SearchAttributeValues

Search attributes based on filters specified in the request body.

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
    res, err := s.AttributeSearch.SearchAttributeValues(ctx, shared.SearchAttributeValuesRequest{
        AttributeTypeIds: []string{
            "recusandae",
        },
        ExcludeIds: []string{
            "omnis",
        },
        Ids: []string{
            "facilis",
        },
        PageSize: conductoroneapi.Float64(5966.56),
        PageToken: conductoroneapi.String("voluptatem"),
        Query: conductoroneapi.String("porro"),
        Value: conductoroneapi.String("consequuntur"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchAttributeValuesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.SearchAttributeValuesRequest](../../models/shared/searchattributevaluesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse](../../models/operations/c1apiattributev1attributesearchsearchattributevaluesresponse.md), error**

