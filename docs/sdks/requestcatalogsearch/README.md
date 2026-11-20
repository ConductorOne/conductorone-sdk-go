# RequestCatalogSearch

## Overview

### Available Operations

* [SearchEntitlements](#searchentitlements) - Search Entitlements
* [SearchRequestableCredentialOfferings](#searchrequestablecredentialofferings) - Search Requestable Credential Offerings

## SearchEntitlements

Search entitlements available to the signed-in user through access profiles, ordered by application ID and entitlement ID.
 Results are paginated; continue with next_page_token even when a page is short or empty. Set expand_mask to only the related objects you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements" method="post" path="/api/v1/search/request_catalog/entitlements" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogSearch.SearchEntitlements(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogSearchServiceSearchEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [shared.RequestCatalogSearchServiceSearchEntitlementsRequest](../../pkg/models/shared/requestcatalogsearchservicesearchentitlementsrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogsearchservicesearchentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchRequestableCredentialOfferings

SearchRequestableCredentialOfferings returns the credential offerings the target user can
 request, optionally limited to one app, ordered by display name then ids.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchRequestableCredentialOfferings" method="post" path="/api/v1/search/request_catalog/credential_offerings" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogSearch.SearchRequestableCredentialOfferings(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchRequestableCredentialOfferingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [shared.SearchRequestableCredentialOfferingsRequest](../../pkg/models/shared/searchrequestablecredentialofferingsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchRequestableCredentialOfferingsResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogsearchservicesearchrequestablecredentialofferingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |