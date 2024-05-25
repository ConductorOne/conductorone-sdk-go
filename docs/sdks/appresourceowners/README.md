# AppResourceOwners
(*AppResourceOwners*)

### Available Operations

* [List](#list) - List

## List

List all owners of an app resource.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )
    request := operations.C1APIAppV1AppResourceOwnersListRequest{
        AppID: "<value>",
        ResourceID: "<value>",
        ResourceTypeID: "<value>",
    }
    ctx := context.Background()
    res, err := s.AppResourceOwners.List(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppResourceOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.C1APIAppV1AppResourceOwnersListRequest](../../pkg/models/operations/c1apiappv1appresourceownerslistrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.C1APIAppV1AppResourceOwnersListResponse](../../pkg/models/operations/c1apiappv1appresourceownerslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
