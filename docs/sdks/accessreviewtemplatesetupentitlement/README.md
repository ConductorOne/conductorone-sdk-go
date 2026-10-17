# AccessReviewTemplateSetupEntitlement

## Overview

### Available Operations

* [AddTemplateEntitlements](#addtemplateentitlements) - Add Template Entitlements
* [GetScopeAndEntitlements](#getscopeandentitlements) - Get Scope And Entitlements
* [RemoveTemplateEntitlements](#removetemplateentitlements) - Remove Template Entitlements
* [SetScopeAndEntitlements](#setscopeandentitlements) - Set Scope And Entitlements
* [SetScopeByResourceType](#setscopebyresourcetype) - Set Scope By Resource Type

## AddTemplateEntitlements

AddTemplateEntitlements adds entitlements to an access review template's selection without
 replacing existing selections or their policies. For large selections, send chunks of at most
 200 entitlements and wait for each call to finish before sending the next. Entitlements that
 are already selected are skipped, so retrying a chunk is safe. If the template has no apps and
 resources scope yet, it is set to the selected entitlements. The per-campaign entitlement quota
 applies to the total selection. To change user, account or grant scope afterwards, send the
 template from AccessReviewTemplateService.Get to AccessReviewTemplateService.Update with
 update_mask scope rather than using SetScopeAndEntitlements, which would replace the whole selection.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewTemplateSetupEntitlementService.AddTemplateEntitlements" method="post" path="/api/v1/access_review_template/{access_review_template_id}/entitlements/add" -->
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

    res, err := s.AccessReviewTemplateSetupEntitlement.AddTemplateEntitlements(ctx, operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceAddTemplateEntitlementsRequest{
        AccessReviewTemplateID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewAddTemplateEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                    |
| `request`                                                                                                                                                                                                                              | [operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceAddTemplateEntitlementsRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewtemplatesetupentitlementserviceaddtemplateentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                          |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceAddTemplateEntitlementsResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewtemplatesetupentitlementserviceaddtemplateentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetScopeAndEntitlements

GetScopeAndEntitlements retrieves the current scope configuration and selected entitlements for an access review template.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewTemplateSetupEntitlementService.GetScopeAndEntitlements" method="get" path="/api/v1/access_review_template/{access_review_template_id}/scope_and_entitlements" -->
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

    res, err := s.AccessReviewTemplateSetupEntitlement.GetScopeAndEntitlements(ctx, operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceGetScopeAndEntitlementsRequest{
        AccessReviewTemplateID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewTemplateSetupEntitlementServiceSetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                    |
| `request`                                                                                                                                                                                                                              | [operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceGetScopeAndEntitlementsRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewtemplatesetupentitlementservicegetscopeandentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                          |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceGetScopeAndEntitlementsResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewtemplatesetupentitlementservicegetscopeandentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveTemplateEntitlements

RemoveTemplateEntitlements removes entitlements from an access review template's selection
 without changing other selections or their policies. Entitlements that are not selected are
 ignored, so retrying a chunk is safe. Send at most 200 entitlements per call.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewTemplateSetupEntitlementService.RemoveTemplateEntitlements" method="post" path="/api/v1/access_review_template/{access_review_template_id}/entitlements/remove" -->
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

    res, err := s.AccessReviewTemplateSetupEntitlement.RemoveTemplateEntitlements(ctx, operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceRemoveTemplateEntitlementsRequest{
        AccessReviewTemplateID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewRemoveTemplateEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                          |
| `request`                                                                                                                                                                                                                                    | [operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceRemoveTemplateEntitlementsRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewtemplatesetupentitlementserviceremovetemplateentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                                                                   |
| `opts`                                                                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                                |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceRemoveTemplateEntitlementsResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewtemplatesetupentitlementserviceremovetemplateentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetScopeAndEntitlements

SetScopeAndEntitlements replaces the scope configuration and selected entitlements for an access review template.
 Each call replaces the whole selection. To build a large selection, use AddTemplateEntitlements in chunks instead.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewTemplateSetupEntitlementService.SetScopeAndEntitlements" method="post" path="/api/v1/access_review_template/{access_review_template_id}/scope_and_entitlements" -->
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

    res, err := s.AccessReviewTemplateSetupEntitlement.SetScopeAndEntitlements(ctx, operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceSetScopeAndEntitlementsRequest{
        AccessReviewTemplateID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewTemplateSetupEntitlementServiceSetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                    |
| `request`                                                                                                                                                                                                                              | [operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceSetScopeAndEntitlementsRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewtemplatesetupentitlementservicesetscopeandentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                          |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceSetScopeAndEntitlementsResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewtemplatesetupentitlementservicesetscopeandentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetScopeByResourceType

SetScopeByResourceType sets the template scope by selecting specific resource types to include in campaigns created from this template.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessreview.v1.AccessReviewTemplateSetupEntitlementService.SetScopeByResourceType" method="post" path="/api/v1/access_review_template/{access_review_template_id}/scope_by_resource_type" -->
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

    res, err := s.AccessReviewTemplateSetupEntitlement.SetScopeByResourceType(ctx, operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceSetScopeByResourceTypeRequest{
        AccessReviewTemplateID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccessReviewTemplateSetScopeByResourceTypeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                                  |
| `request`                                                                                                                                                                                                                            | [operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceSetScopeByResourceTypeRequest](../../pkg/models/operations/c1apiaccessreviewv1accessreviewtemplatesetupentitlementservicesetscopebyresourcetyperequest.md) | :heavy_check_mark:                                                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                                                           |
| `opts`                                                                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                                        |

### Response

**[*operations.C1APIAccessreviewV1AccessReviewTemplateSetupEntitlementServiceSetScopeByResourceTypeResponse](../../pkg/models/operations/c1apiaccessreviewv1accessreviewtemplatesetupentitlementservicesetscopebyresourcetyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |