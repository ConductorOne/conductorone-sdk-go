# AccessReviewTemplateSetupEntitlement

## Overview

### Available Operations

* [SetScopeAndEntitlements](#setscopeandentitlements) - Set Scope And Entitlements
* [SetScopeByResourceType](#setscopebyresourcetype) - Set Scope By Resource Type

## SetScopeAndEntitlements

Invokes the c1.api.accessreview.v1.AccessReviewTemplateSetupEntitlementService.SetScopeAndEntitlements method.

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

Invokes the c1.api.accessreview.v1.AccessReviewTemplateSetupEntitlementService.SetScopeByResourceType method.

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