# Edge

## Overview

### Available Operations

* [CheckInferenceVaultAccess](#checkinferencevaultaccess) - Check Inference Vault Access
* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [GetAuditInclusionProof](#getauditinclusionproof) - Get Audit Inclusion Proof
* [GetEgressTrafficEvent](#getegresstrafficevent) - Get Egress Traffic Event
* [GetEgressTrafficSummary](#getegresstrafficsummary) - Get Egress Traffic Summary
* [GetEgressTrafficSummaryRollup](#getegresstrafficsummaryrollup) - Get Egress Traffic Summary Rollup
* [GetInferenceUsageAttribution](#getinferenceusageattribution) - Get Inference Usage Attribution
* [GetInferenceUsageAttributionRollup](#getinferenceusageattributionrollup) - Get Inference Usage Attribution Rollup
* [GetInferenceUsageSummary](#getinferenceusagesummary) - Get Inference Usage Summary
* [GetInferenceUsageSummaryRollup](#getinferenceusagesummaryrollup) - Get Inference Usage Summary Rollup
* [List](#list) - List
* [ListAuditAttestationKeys](#listauditattestationkeys) - List Audit Attestation Keys
* [ListAuditAttestationManifests](#listauditattestationmanifests) - List Audit Attestation Manifests
* [ListEdgesRollup](#listedgesrollup) - List Edges Rollup
* [ListEgressDenialSummary](#listegressdenialsummary) - List Egress Denial Summary
* [ListEgressRuleEntitlementIssues](#listegressruleentitlementissues) - List Egress Rule Entitlement Issues
* [ListEgressTopAgents](#listegresstopagents) - List Egress Top Agents
* [ListEgressTopAgentsRollup](#listegresstopagentsrollup) - List Egress Top Agents Rollup
* [ListEgressTopDestinations](#listegresstopdestinations) - List Egress Top Destinations
* [ListEgressTopToolsRollup](#listegresstoptoolsrollup) - List Egress Top Tools Rollup
* [ListEgressTrafficEvents](#listegresstrafficevents) - List Egress Traffic Events
* [ListInferenceCatalog](#listinferencecatalog) - List Inference Catalog
* [ListInferenceRoutes](#listinferenceroutes) - List Inference Routes
* [ListInferenceTopAgents](#listinferencetopagents) - List Inference Top Agents
* [ListInferenceTopTools](#listinferencetoptools) - List Inference Top Tools
* [ListInferenceTrafficEvents](#listinferencetrafficevents) - List Inference Traffic Events
* [ListUnattributedInferenceTrafficEvents](#listunattributedinferencetrafficevents) - List Unattributed Inference Traffic Events
* [ListUsableEdges](#listusableedges) - List Usable Edges
* [PreviewEgressHostDecision](#previewegresshostdecision) - Preview Egress Host Decision
* [PreviewEgressRulesUpdate](#previewegressrulesupdate) - Preview Egress Rules Update
* [Update](#update) - Update

## CheckInferenceVaultAccess

CheckInferenceVaultAccess reports whether the Edge's runtime service
 principal can open secrets in a Vault, using the same check Update runs
 on inference targets, and returns the entitlement that grants it.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.CheckInferenceVaultAccess" method="get" path="/api/v1/apps/{app_id}/edges/{id}/inference/vault-access" -->
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

    res, err := s.Edge.CheckInferenceVaultAccess(ctx, operations.C1APIEdgeV1EdgeServiceCheckInferenceVaultAccessRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceCheckInferenceVaultAccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIEdgeV1EdgeServiceCheckInferenceVaultAccessRequest](../../pkg/models/operations/c1apiedgev1edgeservicecheckinferencevaultaccessrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIEdgeV1EdgeServiceCheckInferenceVaultAccessResponse](../../pkg/models/operations/c1apiedgev1edgeservicecheckinferencevaultaccessresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Create

Create mints an Edge (and one entitlement per enabled_kinds entry) for an
 app. The Edge starts disabled.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.Create" method="post" path="/api/v1/apps/{app_id}/edges" -->
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

    res, err := s.Edge.Create(ctx, operations.C1APIEdgeV1EdgeServiceCreateRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.C1APIEdgeV1EdgeServiceCreateRequest](../../pkg/models/operations/c1apiedgev1edgeservicecreaterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIEdgeV1EdgeServiceCreateResponse](../../pkg/models/operations/c1apiedgev1edgeservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete an Edge and its associated entitlements. New egress requests are
 rejected, and inference tokens can no longer be issued for this Edge.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.Delete" method="delete" path="/api/v1/apps/{app_id}/edges/{id}" -->
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

    res, err := s.Edge.Delete(ctx, operations.C1APIEdgeV1EdgeServiceDeleteRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.C1APIEdgeV1EdgeServiceDeleteRequest](../../pkg/models/operations/c1apiedgev1edgeservicedeleterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIEdgeV1EdgeServiceDeleteResponse](../../pkg/models/operations/c1apiedgev1edgeservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get retrieves one Edge.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.Get" method="get" path="/api/v1/apps/{app_id}/edges/{id}" -->
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

    res, err := s.Edge.Get(ctx, operations.C1APIEdgeV1EdgeServiceGetRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.C1APIEdgeV1EdgeServiceGetRequest](../../pkg/models/operations/c1apiedgev1edgeservicegetrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.C1APIEdgeV1EdgeServiceGetResponse](../../pkg/models/operations/c1apiedgev1edgeservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAuditInclusionProof

Returns the checked inclusion proof of one audit event of the caller's
 own tenant. An event outside the tenant is NOT_FOUND, like a missing one.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.GetAuditInclusionProof" method="get" path="/api/v1/edges/audit_attestation/proofs/{event_id}" -->
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

    res, err := s.Edge.GetAuditInclusionProof(ctx, operations.C1APIEdgeV1EdgeServiceGetAuditInclusionProofRequest{
        EventID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceGetAuditInclusionProofResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.C1APIEdgeV1EdgeServiceGetAuditInclusionProofRequest](../../pkg/models/operations/c1apiedgev1edgeservicegetauditinclusionproofrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.C1APIEdgeV1EdgeServiceGetAuditInclusionProofResponse](../../pkg/models/operations/c1apiedgev1edgeservicegetauditinclusionproofresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetEgressTrafficEvent

Returns one egress traffic event for this Edge by event_id, read from a narrow window around its event_time.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.GetEgressTrafficEvent" method="get" path="/api/v1/apps/{app_id}/edges/{id}/egress/traffic/{event_id}" -->
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

    res, err := s.Edge.GetEgressTrafficEvent(ctx, operations.C1APIEdgeV1EdgeServiceGetEgressTrafficEventRequest{
        AppID: "<id>",
        EventID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceGetEgressTrafficEventResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIEdgeV1EdgeServiceGetEgressTrafficEventRequest](../../pkg/models/operations/c1apiedgev1edgeservicegetegresstrafficeventrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APIEdgeV1EdgeServiceGetEgressTrafficEventResponse](../../pkg/models/operations/c1apiedgev1edgeservicegetegresstrafficeventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetEgressTrafficSummary

Returns egress traffic counts, verdicts and transferred bytes for this Edge in the requested time window.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.GetEgressTrafficSummary" method="get" path="/api/v1/apps/{app_id}/edges/{id}/egress/traffic-summary" -->
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

    res, err := s.Edge.GetEgressTrafficSummary(ctx, operations.C1APIEdgeV1EdgeServiceGetEgressTrafficSummaryRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceGetEgressTrafficSummaryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APIEdgeV1EdgeServiceGetEgressTrafficSummaryRequest](../../pkg/models/operations/c1apiedgev1edgeservicegetegresstrafficsummaryrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APIEdgeV1EdgeServiceGetEgressTrafficSummaryResponse](../../pkg/models/operations/c1apiedgev1edgeservicegetegresstrafficsummaryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetEgressTrafficSummaryRollup

Returns MCP traffic counts, verdicts and transferred bytes across the Edges the caller manages.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.GetEgressTrafficSummaryRollup" method="get" path="/api/v1/edges/egress/traffic-summary" -->
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

    res, err := s.Edge.GetEgressTrafficSummaryRollup(ctx, operations.C1APIEdgeV1EdgeServiceGetEgressTrafficSummaryRollupRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceGetEgressTrafficSummaryRollupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APIEdgeV1EdgeServiceGetEgressTrafficSummaryRollupRequest](../../pkg/models/operations/c1apiedgev1edgeservicegetegresstrafficsummaryrolluprequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APIEdgeV1EdgeServiceGetEgressTrafficSummaryRollupResponse](../../pkg/models/operations/c1apiedgev1edgeservicegetegresstrafficsummaryrollupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetInferenceUsageAttribution

Returns inference request and token usage for this Edge grouped by the selected attribution dimension.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.GetInferenceUsageAttribution" method="get" path="/api/v1/apps/{app_id}/edges/{id}/inference/usage-attribution" -->
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

    res, err := s.Edge.GetInferenceUsageAttribution(ctx, operations.C1APIEdgeV1EdgeServiceGetInferenceUsageAttributionRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceGetInferenceUsageAttributionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.C1APIEdgeV1EdgeServiceGetInferenceUsageAttributionRequest](../../pkg/models/operations/c1apiedgev1edgeservicegetinferenceusageattributionrequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |
| `opts`                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |

### Response

**[*operations.C1APIEdgeV1EdgeServiceGetInferenceUsageAttributionResponse](../../pkg/models/operations/c1apiedgev1edgeservicegetinferenceusageattributionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetInferenceUsageAttributionRollup

Returns inference request and token usage across the Edges the caller manages, grouped by the selected attribution dimension.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.GetInferenceUsageAttributionRollup" method="get" path="/api/v1/edges/inference/usage-attribution" -->
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

    res, err := s.Edge.GetInferenceUsageAttributionRollup(ctx, operations.C1APIEdgeV1EdgeServiceGetInferenceUsageAttributionRollupRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceGetInferenceUsageAttributionRollupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.C1APIEdgeV1EdgeServiceGetInferenceUsageAttributionRollupRequest](../../pkg/models/operations/c1apiedgev1edgeservicegetinferenceusageattributionrolluprequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |
| `opts`                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                           | The options for this request.                                                                                                                                                |

### Response

**[*operations.C1APIEdgeV1EdgeServiceGetInferenceUsageAttributionRollupResponse](../../pkg/models/operations/c1apiedgev1edgeservicegetinferenceusageattributionrollupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetInferenceUsageSummary

Returns inference request and token totals for this Edge, with time series and model breakdowns.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.GetInferenceUsageSummary" method="get" path="/api/v1/apps/{app_id}/edges/{id}/inference/usage" -->
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

    res, err := s.Edge.GetInferenceUsageSummary(ctx, operations.C1APIEdgeV1EdgeServiceGetInferenceUsageSummaryRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceGetInferenceUsageSummaryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APIEdgeV1EdgeServiceGetInferenceUsageSummaryRequest](../../pkg/models/operations/c1apiedgev1edgeservicegetinferenceusagesummaryrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APIEdgeV1EdgeServiceGetInferenceUsageSummaryResponse](../../pkg/models/operations/c1apiedgev1edgeservicegetinferenceusagesummaryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetInferenceUsageSummaryRollup

Returns inference request and token totals across the Edges the caller manages, with time series and model breakdowns.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.GetInferenceUsageSummaryRollup" method="get" path="/api/v1/edges/inference/usage" -->
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

    res, err := s.Edge.GetInferenceUsageSummaryRollup(ctx, operations.C1APIEdgeV1EdgeServiceGetInferenceUsageSummaryRollupRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceGetInferenceUsageSummaryRollupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.C1APIEdgeV1EdgeServiceGetInferenceUsageSummaryRollupRequest](../../pkg/models/operations/c1apiedgev1edgeservicegetinferenceusagesummaryrolluprequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.C1APIEdgeV1EdgeServiceGetInferenceUsageSummaryRollupResponse](../../pkg/models/operations/c1apiedgev1edgeservicegetinferenceusagesummaryrollupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List retrieves Edges for an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.List" method="get" path="/api/v1/apps/{app_id}/edges" -->
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

    res, err := s.Edge.List(ctx, operations.C1APIEdgeV1EdgeServiceListRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.C1APIEdgeV1EdgeServiceListRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAuditAttestationKeys

Lists the public keys that sign audit manifests, for offline verification.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListAuditAttestationKeys" method="get" path="/api/v1/edges/audit_attestation/keys" -->
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

    res, err := s.Edge.ListAuditAttestationKeys(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListAuditAttestationKeysResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListAuditAttestationKeysResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistauditattestationkeysresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAuditAttestationManifests

Lists the signed attestation manifests of the caller's own tenant, which
 chain the audit log into a tamper-evident sequence. Tenant-level: only the
 roles that hold EdgeService:owner (Super Administrator, AI Governance
 Administrator). Fails with FAILED_PRECONDITION when the warehouse has no attestation.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListAuditAttestationManifests" method="get" path="/api/v1/edges/audit_attestation/manifests" -->
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

    res, err := s.Edge.ListAuditAttestationManifests(ctx, operations.C1APIEdgeV1EdgeServiceListAuditAttestationManifestsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListAuditAttestationManifestsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APIEdgeV1EdgeServiceListAuditAttestationManifestsRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistauditattestationmanifestsrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListAuditAttestationManifestsResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistauditattestationmanifestsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEdgesRollup

Lists the Edges the caller manages across every app, for navigating to
 each Edge's own activity. Same authorized set as the other rollups.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListEdgesRollup" method="get" path="/api/v1/edges/inventory" -->
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

    res, err := s.Edge.ListEdgesRollup(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListEdgesRollupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListEdgesRollupResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistedgesrollupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEgressDenialSummary

Groups denied and would-deny egress flows for this Edge by destination, caller, outcome and reason.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListEgressDenialSummary" method="get" path="/api/v1/apps/{app_id}/edges/{id}/egress/denial-summary" -->
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

    res, err := s.Edge.ListEgressDenialSummary(ctx, operations.C1APIEdgeV1EdgeServiceListEgressDenialSummaryRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListEgressDenialSummaryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APIEdgeV1EdgeServiceListEgressDenialSummaryRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistegressdenialsummaryrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListEgressDenialSummaryResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistegressdenialsummaryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEgressRuleEntitlementIssues

ListEgressRuleEntitlementIssues reports each saved egress rule whose
 entitlement does not resolve, and what the Edge does with it instead.
 Nothing is saved.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListEgressRuleEntitlementIssues" method="get" path="/api/v1/apps/{app_id}/edges/{id}/egress-rules/entitlement-issues" -->
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

    res, err := s.Edge.ListEgressRuleEntitlementIssues(ctx, operations.C1APIEdgeV1EdgeServiceListEgressRuleEntitlementIssuesRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListEgressRuleEntitlementIssuesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.C1APIEdgeV1EdgeServiceListEgressRuleEntitlementIssuesRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistegressruleentitlementissuesrequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |
| `opts`                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                           | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListEgressRuleEntitlementIssuesResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistegressruleentitlementissuesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEgressTopAgents

Lists agents ranked by egress request count for this Edge in the requested time window.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListEgressTopAgents" method="get" path="/api/v1/apps/{app_id}/edges/{id}/egress/top-agents" -->
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

    res, err := s.Edge.ListEgressTopAgents(ctx, operations.C1APIEdgeV1EdgeServiceListEgressTopAgentsRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListEgressTopAgentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIEdgeV1EdgeServiceListEgressTopAgentsRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistegresstopagentsrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListEgressTopAgentsResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistegresstopagentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEgressTopAgentsRollup

Lists agents ranked by MCP request count across the Edges the caller manages.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListEgressTopAgentsRollup" method="get" path="/api/v1/edges/egress/top-agents" -->
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

    res, err := s.Edge.ListEgressTopAgentsRollup(ctx, operations.C1APIEdgeV1EdgeServiceListEgressTopAgentsRollupRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListEgressTopAgentsRollupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIEdgeV1EdgeServiceListEgressTopAgentsRollupRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistegresstopagentsrolluprequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListEgressTopAgentsRollupResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistegresstopagentsrollupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEgressTopDestinations

Lists destinations ranked by egress request count for this Edge in the requested time window.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListEgressTopDestinations" method="get" path="/api/v1/apps/{app_id}/edges/{id}/egress/top-destinations" -->
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

    res, err := s.Edge.ListEgressTopDestinations(ctx, operations.C1APIEdgeV1EdgeServiceListEgressTopDestinationsRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListEgressTopDestinationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIEdgeV1EdgeServiceListEgressTopDestinationsRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistegresstopdestinationsrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListEgressTopDestinationsResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistegresstopdestinationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEgressTopToolsRollup

Lists the most frequently used MCP tools across the Edges the caller manages.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListEgressTopToolsRollup" method="get" path="/api/v1/edges/egress/top-tools" -->
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

    res, err := s.Edge.ListEgressTopToolsRollup(ctx, operations.C1APIEdgeV1EdgeServiceListEgressTopToolsRollupRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListEgressTopToolsRollupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APIEdgeV1EdgeServiceListEgressTopToolsRollupRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistegresstoptoolsrolluprequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListEgressTopToolsRollupResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistegresstoptoolsrollupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEgressTrafficEvents

Lists egress traffic events for this Edge in the requested time window.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListEgressTrafficEvents" method="get" path="/api/v1/apps/{app_id}/edges/{id}/egress/traffic" -->
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

    res, err := s.Edge.ListEgressTrafficEvents(ctx, operations.C1APIEdgeV1EdgeServiceListEgressTrafficEventsRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListEgressTrafficEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APIEdgeV1EdgeServiceListEgressTrafficEventsRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistegresstrafficeventsrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListEgressTrafficEventsResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistegresstrafficeventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListInferenceCatalog

ListInferenceCatalog returns the providers, variants and suggested models that inference routes can be built
 from, C1's registered models by family, and the model each C1 tier uses for the caller's tenant. The Edge
 feature is required.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListInferenceCatalog" method="get" path="/api/v1/edges/inference/catalog" -->
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

    res, err := s.Edge.ListInferenceCatalog(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListInferenceCatalogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListInferenceCatalogResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistinferencecatalogresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListInferenceRoutes

ListInferenceRoutes returns routes already configured for this application.
 The Edge feature and application management access are required. New routes
 are authored through Update.inference_routes, not selected from a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListInferenceRoutes" method="get" path="/api/v1/apps/{app_id}/inference-routes" -->
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

    res, err := s.Edge.ListInferenceRoutes(ctx, operations.C1APIEdgeV1EdgeServiceListInferenceRoutesRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListInferenceRoutesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIEdgeV1EdgeServiceListInferenceRoutesRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistinferenceroutesrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListInferenceRoutesResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistinferenceroutesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListInferenceTopAgents

Lists agents ranked by inference request count for this Edge in the requested time window.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListInferenceTopAgents" method="get" path="/api/v1/apps/{app_id}/edges/{id}/inference/top-agents" -->
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

    res, err := s.Edge.ListInferenceTopAgents(ctx, operations.C1APIEdgeV1EdgeServiceListInferenceTopAgentsRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListInferenceTopAgentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.C1APIEdgeV1EdgeServiceListInferenceTopAgentsRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistinferencetopagentsrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListInferenceTopAgentsResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistinferencetopagentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListInferenceTopTools

Lists the most frequently used inference operations for this Edge in the requested time window.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListInferenceTopTools" method="get" path="/api/v1/apps/{app_id}/edges/{id}/inference/top-tools" -->
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

    res, err := s.Edge.ListInferenceTopTools(ctx, operations.C1APIEdgeV1EdgeServiceListInferenceTopToolsRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListInferenceTopToolsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIEdgeV1EdgeServiceListInferenceTopToolsRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistinferencetoptoolsrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListInferenceTopToolsResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistinferencetoptoolsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListInferenceTrafficEvents

Lists inference traffic events for this Edge in the requested time window.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListInferenceTrafficEvents" method="get" path="/api/v1/apps/{app_id}/edges/{id}/inference/traffic" -->
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

    res, err := s.Edge.ListInferenceTrafficEvents(ctx, operations.C1APIEdgeV1EdgeServiceListInferenceTrafficEventsRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListInferenceTrafficEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIEdgeV1EdgeServiceListInferenceTrafficEventsRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistinferencetrafficeventsrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListInferenceTrafficEventsResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistinferencetrafficeventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUnattributedInferenceTrafficEvents

Lists auth-rejected inference requests that Time Bandit could not
 attribute to a single Edge, for the whole tenant. Rejections that resolve
 to exactly one Edge appear in that Edge's own inference traffic instead.
 Tenant-level data: restricted to the roles that hold EdgeService:owner
 (Super Administrator, AI Governance Administrator), not Edge or app owners.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListUnattributedInferenceTrafficEvents" method="get" path="/api/v1/edges/unattributed_inference/traffic" -->
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

    res, err := s.Edge.ListUnattributedInferenceTrafficEvents(ctx, operations.C1APIEdgeV1EdgeServiceListUnattributedInferenceTrafficEventsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListUnattributedInferenceTrafficEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                            | Type                                                                                                                                                                                 | Required                                                                                                                                                                             | Description                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                | :heavy_check_mark:                                                                                                                                                                   | The context to use for the request.                                                                                                                                                  |
| `request`                                                                                                                                                                            | [operations.C1APIEdgeV1EdgeServiceListUnattributedInferenceTrafficEventsRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistunattributedinferencetrafficeventsrequest.md) | :heavy_check_mark:                                                                                                                                                                   | The request object to use for the request.                                                                                                                                           |
| `opts`                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                   | The options for this request.                                                                                                                                                        |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListUnattributedInferenceTrafficEventsResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistunattributedinferencetrafficeventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUsableEdges

List connection targets the calling person can currently use, without
 administrative configuration. Supports the person's own human session
 and device DPoP session; service principals are refused.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.ListUsableEdges" method="get" path="/api/v1/edges/usable" -->
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

    res, err := s.Edge.ListUsableEdges(ctx, operations.C1APIEdgeV1EdgeServiceListUsableEdgesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceListUsableEdgesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.C1APIEdgeV1EdgeServiceListUsableEdgesRequest](../../pkg/models/operations/c1apiedgev1edgeservicelistusableedgesrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.C1APIEdgeV1EdgeServiceListUsableEdgesResponse](../../pkg/models/operations/c1apiedgev1edgeservicelistusableedgesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PreviewEgressHostDecision

PreviewEgressHostDecision reports what blocking or allowing one host for
 every user would do, and the rule list that results. Nothing is saved.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.PreviewEgressHostDecision" method="post" path="/api/v1/apps/{app_id}/edges/{id}/egress-rules/preview-host-decision" -->
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

    res, err := s.Edge.PreviewEgressHostDecision(ctx, operations.C1APIEdgeV1EdgeServicePreviewEgressHostDecisionRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServicePreviewEgressHostDecisionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.C1APIEdgeV1EdgeServicePreviewEgressHostDecisionRequest](../../pkg/models/operations/c1apiedgev1edgeservicepreviewegresshostdecisionrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.C1APIEdgeV1EdgeServicePreviewEgressHostDecisionResponse](../../pkg/models/operations/c1apiedgev1edgeservicepreviewegresshostdecisionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PreviewEgressRulesUpdate

PreviewEgressRulesUpdate checks a proposed egress rule list as Update
 would and compares each host's verdict under the saved rules with its
 verdict under the proposed rules. Nothing is saved.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.PreviewEgressRulesUpdate" method="post" path="/api/v1/apps/{app_id}/edges/{id}/egress-rules/preview-update" -->
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

    res, err := s.Edge.PreviewEgressRulesUpdate(ctx, operations.C1APIEdgeV1EdgeServicePreviewEgressRulesUpdateRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServicePreviewEgressRulesUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APIEdgeV1EdgeServicePreviewEgressRulesUpdateRequest](../../pkg/models/operations/c1apiedgev1edgeservicepreviewegressrulesupdaterequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APIEdgeV1EdgeServicePreviewEgressRulesUpdateResponse](../../pkg/models/operations/c1apiedgev1edgeservicepreviewegressrulesupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update edits an Edge's mode, credential injections, or inference route via update_mask.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.edge.v1.EdgeService.Update" method="post" path="/api/v1/apps/{app_id}/edges/{id}" -->
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

    res, err := s.Edge.Update(ctx, operations.C1APIEdgeV1EdgeServiceUpdateRequest{
        AppID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EdgeServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.C1APIEdgeV1EdgeServiceUpdateRequest](../../pkg/models/operations/c1apiedgev1edgeserviceupdaterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIEdgeV1EdgeServiceUpdateResponse](../../pkg/models/operations/c1apiedgev1edgeserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |