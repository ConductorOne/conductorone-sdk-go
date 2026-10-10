# AccessReviewSetupEntitlement

## Overview

### Available Operations

* [AddCampaignEntitlements](#addcampaignentitlements) - Add Campaign Entitlements
* [GetCampaignScopeAndEntitlements](#getcampaignscopeandentitlements) - Get Campaign Scope And Entitlements
* [RemoveCampaignEntitlements](#removecampaignentitlements) - Remove Campaign Entitlements
* [SetCampaignScopeAndEntitlements](#setcampaignscopeandentitlements) - Set Campaign Scope And Entitlements
* [SetCampaignScopeByResourceType](#setcampaignscopebyresourcetype) - Set Campaign Scope By Resource Type

## AddCampaignEntitlements

AddCampaignEntitlements adds entitlements to an access review campaign's selection without
 replacing existing selections or their policies. For large selections, send chunks of at most
 200 entitlements and wait for each call to finish before sending the next. Entitlements that
 are already selected are skipped, so retrying a chunk is safe. If the campaign has no apps and
 resources scope yet, it is set to the selected entitlements. The per-campaign entitlement quota
 applies to the total selection. To change user, account or grant scope afterwards, use
 AccessReviewService.Update with update_mask scope_v2 rather than SetCampaignScopeAndEntitlements,
 which would replace the whole selection.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewSetupEntitlementService.AddCampaignEntitlements" method="post" path="/api/v1/access_review/{access_review_id}/entitlements/add" -->
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

    res, err := s.AccessReviewSetupEntitlement.AddCampaignEntitlements(ctx, operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceAddCampaignEntitlementsRequest{
        AccessReviewID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewAddCampaignEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                              | Type                                                                                                                                                                                                                   | Required                                                                                                                                                                                                               | Description                                                                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                    |
| `request`                                                                                                                                                                                                              | [operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceAddCampaignEntitlementsRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementserviceaddcampaignentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                          |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceAddCampaignEntitlementsResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementserviceaddcampaignentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCampaignScopeAndEntitlements

GetCampaignScopeAndEntitlements retrieves the current scope configuration and selected entitlements for an access review campaign.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewSetupEntitlementService.GetCampaignScopeAndEntitlements" method="get" path="/api/v1/access_review/{access_review_id}/scope_and_entitlements" -->
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

    res, err := s.AccessReviewSetupEntitlement.GetCampaignScopeAndEntitlements(ctx, operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceGetCampaignScopeAndEntitlementsRequest{
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
| `request`                                                                                                                                                                                                                              | [operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceGetCampaignScopeAndEntitlementsRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementservicegetcampaignscopeandentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                          |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceGetCampaignScopeAndEntitlementsResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementservicegetcampaignscopeandentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCampaignEntitlements

RemoveCampaignEntitlements removes entitlements from an access review campaign's selection
 without changing other selections or their policies. Entitlements that are not selected are
 ignored, so retrying a chunk is safe. Send at most 200 entitlements per call.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewSetupEntitlementService.RemoveCampaignEntitlements" method="post" path="/api/v1/access_review/{access_review_id}/entitlements/remove" -->
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

    res, err := s.AccessReviewSetupEntitlement.RemoveCampaignEntitlements(ctx, operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceRemoveCampaignEntitlementsRequest{
        AccessReviewID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewRemoveCampaignEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                          |
| `request`                                                                                                                                                                                                                    | [operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceRemoveCampaignEntitlementsRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementserviceremovecampaignentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                                                   |
| `opts`                                                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceRemoveCampaignEntitlementsResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementserviceremovecampaignentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetCampaignScopeAndEntitlements

SetCampaignScopeAndEntitlements replaces the scope configuration and selected entitlements for an access review campaign.
 Each call replaces the whole selection. To build a large selection, use AddCampaignEntitlements in chunks instead.

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

SetCampaignScopeByResourceType sets the campaign scope by selecting specific resource types to include in the review.

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