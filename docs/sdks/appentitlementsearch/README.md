# AppEntitlementSearch
(*AppEntitlementSearch*)

### Available Operations

* [Search](#search) - Search
* [SearchAppEntitlementsWithExpired](#searchappentitlementswithexpired) - Search App Entitlements With Expired

## Search

Search app entitlements based on filters specified in the request body.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
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

    ctx := context.Background()
    res, err := s.AppEntitlementSearch.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementSearchServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [shared.AppEntitlementSearchServiceSearchRequest](../../pkg/models/shared/appentitlementsearchservicesearchrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchResponse](../../pkg/models/operations/c1apiappv1appentitlementsearchservicesearchresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## SearchAppEntitlementsWithExpired

Search app entitlements, include app users, users, expires, discovered.

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
    request := operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredRequest{
        AppEntitlementID: "<value>",
        AppID: "<value>",
    }
    ctx := context.Background()
    res, err := s.AppEntitlementSearch.SearchAppEntitlementsWithExpired(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppEntitlementsWithExpiredResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                              | Type                                                                                                                                                                                                   | Required                                                                                                                                                                                               | Description                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                    |
| `request`                                                                                                                                                                                              | [operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredRequest](../../pkg/models/operations/c1apiappv1appentitlementsearchservicesearchappentitlementswithexpiredrequest.md) | :heavy_check_mark:                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                             |
| `opts`                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                     | The options for this request.                                                                                                                                                                          |


### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredResponse](../../pkg/models/operations/c1apiappv1appentitlementsearchservicesearchappentitlementswithexpiredresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
