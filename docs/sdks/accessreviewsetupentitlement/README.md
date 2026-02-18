# AccessReviewSetupEntitlement

## Overview

### Available Operations

* [SetCampaignScopeAndEntitlements](#setcampaignscopeandentitlements) - Set Campaign Scope And Entitlements
* [SetCampaignScopeByResourceType](#setcampaignscopebyresourcetype) - Set Campaign Scope By Resource Type

## SetCampaignScopeAndEntitlements

Invokes the c1.api.accessreview.v1.AccessReviewSetupEntitlementService.SetCampaignScopeAndEntitlements method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewSetupEntitlementService.SetCampaignScopeAndEntitlements" method="post" path="/api/v1/access_review/{access_review_id}/scope_and_entitlements" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
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

    res, err := s.AccessReviewSetupEntitlement.SetCampaignScopeAndEntitlements(ctx, operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceSetCampaignScopeAndEntitlementsRequest{
        AccessReviewID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewSetupEntitlementAndScopeServiceSetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                    |
| `request`                                                                                                                                                                                                                              | [operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceSetCampaignScopeAndEntitlementsRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementservicesetcampaignscopeandentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                          |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceSetCampaignScopeAndEntitlementsResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementservicesetcampaignscopeandentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetCampaignScopeByResourceType

Invokes the c1.api.accessreview.v1.AccessReviewSetupEntitlementService.SetCampaignScopeByResourceType method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewSetupEntitlementService.SetCampaignScopeByResourceType" method="post" path="/api/v1/access_review/{access_review_id}/scope_by_resource_type" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
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

    res, err := s.AccessReviewSetupEntitlement.SetCampaignScopeByResourceType(ctx, operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceSetCampaignScopeByResourceTypeRequest{
        AccessReviewID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewSetScopeByResourceTypeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                                  |
| `request`                                                                                                                                                                                                                            | [operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceSetCampaignScopeByResourceTypeRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementservicesetcampaignscopebyresourcetyperequest.md) | :heavy_check_mark:                                                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                                                           |
| `opts`                                                                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                                        |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceSetCampaignScopeByResourceTypeResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementservicesetcampaignscopebyresourcetyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |