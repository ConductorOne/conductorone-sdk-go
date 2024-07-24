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
    var request *shared.SearchAttributeValuesRequest = &shared.SearchAttributeValuesRequest{}
    ctx := context.Background()
    res, err := s.AttributeSearch.SearchAttributeValues(ctx, request)
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
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse](../../pkg/models/operations/c1apiattributev1attributesearchsearchattributevaluesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
