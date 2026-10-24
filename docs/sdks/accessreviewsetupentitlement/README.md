# AccessReviewSetupEntitlement

## Overview

### Available Operations

* [AddCampaignEntitlements](#addcampaignentitlements) - Add Campaign Entitlements
* [AddDeleteExcludedEntitlements](#adddeleteexcludedentitlements) - Add Delete Excluded Entitlements
* [GetCampaignScopeAndEntitlements](#getcampaignscopeandentitlements) - Get Campaign Scope And Entitlements
* [RemoveCampaignEntitlements](#removecampaignentitlements) - Remove Campaign Entitlements
* [SetCampaignScopeAndEntitlements](#setcampaignscopeandentitlements) - Set Campaign Scope And Entitlements
* [SetCampaignScopeAndExcludedEntitlements](#setcampaignscopeandexcludedentitlements) - Set Campaign Scope And Excluded Entitlements
* [~~SetCampaignScopeByExcludedResourceType~~](#setcampaignscopebyexcludedresourcetype) - Set Campaign Scope By Excluded Resource Type :warning: **Deprecated**
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

## AddDeleteExcludedEntitlements

AddDeleteExcludedEntitlements adds and/or removes individual excluded entitlements for a
 user-first campaign without replacing the whole excluded set, so concurrent edits are not
 overwritten. Entitlements that are already excluded are skipped, and entitlements that are not
 excluded are ignored on removal. Exclusions only take effect for user-first campaigns (scope
 type ACCESS_REVIEW_SCOPE_TYPE_BY_USERS). This does not change the campaign's scope_v2, and
 exclusions only take effect while scope_v2 sets excluded_specific_resources; use
 SetCampaignScopeAndExcludedEntitlements to set the scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewSetupEntitlementService.AddDeleteExcludedEntitlements" method="post" path="/api/v1/access_review/{access_review_id}/excluded_entitlements/add_delete" -->
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

    res, err := s.AccessReviewSetupEntitlement.AddDeleteExcludedEntitlements(ctx, operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceAddDeleteExcludedEntitlementsRequest{
        AccessReviewID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewAddDeleteExcludedEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                          | Type                                                                                                                                                                                                                               | Required                                                                                                                                                                                                                           | Description                                                                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                                |
| `request`                                                                                                                                                                                                                          | [operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceAddDeleteExcludedEntitlementsRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementserviceadddeleteexcludedentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                 | The request object to use for the request.                                                                                                                                                                                         |
| `opts`                                                                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                                      |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceAddDeleteExcludedEntitlementsResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementserviceadddeleteexcludedentitlementsresponse.md), error**

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

## SetCampaignScopeAndExcludedEntitlements

SetCampaignScopeAndExcludedEntitlements replaces the scope configuration and the set of
 excluded entitlements for a user-first campaign. Each call replaces the whole excluded set; to
 add or remove individual entitlements, use AddDeleteExcludedEntitlements instead. Exclusions
 only take effect for user-first campaigns (scope type ACCESS_REVIEW_SCOPE_TYPE_BY_USERS). Read
 the current exclusions back with GetCampaignScopeAndEntitlements.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewSetupEntitlementService.SetCampaignScopeAndExcludedEntitlements" method="post" path="/api/v1/access_review/{access_review_id}/scope_and_excluded_entitlements" -->
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

    res, err := s.AccessReviewSetupEntitlement.SetCampaignScopeAndExcludedEntitlements(ctx, operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceSetCampaignScopeAndExcludedEntitlementsRequest{
        AccessReviewID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewSetCampaignScopeAndExcludedEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                    |
| `request`                                                                                                                                                                                                                                              | [operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceSetCampaignScopeAndExcludedEntitlementsRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementservicesetcampaignscopeandexcludedentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                          |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceSetCampaignScopeAndExcludedEntitlementsResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementservicesetcampaignscopeandexcludedentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~SetCampaignScopeByExcludedResourceType~~

SetCampaignScopeByExcludedResourceType replaces the scope configuration and the set of excluded
 resource types for a user-first campaign. Each call replaces the whole excluded set. Exclusions
 only take effect for user-first campaigns (scope type ACCESS_REVIEW_SCOPE_TYPE_BY_USERS). Read
 the current exclusions back with GetCampaignScopeAndEntitlements.

 This route is superseded and scheduled for removal. Requests that set excluded_resource_types
 are rejected with FAILED_PRECONDITION; an empty list, which clears stored resource-type
 exclusions, is still accepted. To exclude entitlements by resource type, call
 SetCampaignScopeAndExcludedEntitlements with scope_v2.excluded_app_selection_criteria and list
 the resource types in resource_type_refs.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewSetupEntitlementService.SetCampaignScopeByExcludedResourceType" method="post" path="/api/v1/access_review/{access_review_id}/scope_by_excluded_resource_type" -->
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

    res, err := s.AccessReviewSetupEntitlement.SetCampaignScopeByExcludedResourceType(ctx, operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceSetCampaignScopeByExcludedResourceTypeRequest{
        AccessReviewID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewSetCampaignScopeByExcludedResourceTypeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                                                  |
| `request`                                                                                                                                                                                                                                            | [operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceSetCampaignScopeByExcludedResourceTypeRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementservicesetcampaignscopebyexcludedresourcetyperequest.md) | :heavy_check_mark:                                                                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                                                                           |
| `opts`                                                                                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                                                        |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewSetupEntitlementServiceSetCampaignScopeByExcludedResourceTypeResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewsetupentitlementservicesetcampaignscopebyexcludedresourcetyperesponse.md), error**

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