# AttributeSearch
(*AttributeSearch*)

### Available Operations

* [SearchAttributeValues](#searchattributevalues) - Search Attribute Values

## SearchAttributeValues

Search attributes based on filters specified in the request body.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AttributeSearch.SearchAttributeValues(ctx, &shared.SearchAttributeValuesRequest{
        AttributeTypeIds: []string{
            "string",
        },
        ExcludeIds: []string{
            "string",
        },
        Ids: []string{
            "string",
        },
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [shared.SearchAttributeValuesRequest](../../pkg/models/shared/searchattributevaluesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse](../../pkg/models/operations/c1apiattributev1attributesearchsearchattributevaluesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
