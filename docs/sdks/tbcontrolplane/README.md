# TBControlPlane

## Overview

### Available Operations

* [GetDiscoverySnapshot](#getdiscoverysnapshot) - Get Discovery Snapshot
* [GetEgressPolicy](#getegresspolicy) - Get Egress Policy
* [~~GetTrafficSummary~~](#gettrafficsummary) - Get Traffic Summary :warning: **Deprecated**
* [~~GetUsageAttribution~~](#getusageattribution) - Get Usage Attribution :warning: **Deprecated**
* [~~GetUsageSummary~~](#getusagesummary) - Get Usage Summary :warning: **Deprecated**
* [~~ListTopAgents~~](#listtopagents) - List Top Agents :warning: **Deprecated**
* [~~ListTopTools~~](#listtoptools) - List Top Tools :warning: **Deprecated**
* [~~ListTrafficEvents~~](#listtrafficevents) - List Traffic Events :warning: **Deprecated**
* [PushDiscovery](#pushdiscovery) - Push Discovery
* [SaveEgressPolicy](#saveegresspolicy) - Save Egress Policy

## GetDiscoverySnapshot

GetDiscoverySnapshot returns the instance's latest self-reported
 vocabulary -- used by the authoring UI to build discovery-driven
 pickers (principals/scopes/destinations/postures/credentials/routes)
 independent of any policy read or write.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.tbcontrolplane.v1.TBControlPlaneService.GetDiscoverySnapshot" method="get" path="/api/v1/tb-control-plane/discovery/{tb_instance_id}" -->
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

    res, err := s.TBControlPlane.GetDiscoverySnapshot(ctx, operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetDiscoverySnapshotRequest{
        TbInstanceID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TBControlPlaneServiceGetDiscoverySnapshotResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                | Type                                                                                                                                                                                     | Required                                                                                                                                                                                 | Description                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                       | The context to use for the request.                                                                                                                                                      |
| `request`                                                                                                                                                                                | [operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetDiscoverySnapshotRequest](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicegetdiscoverysnapshotrequest.md) | :heavy_check_mark:                                                                                                                                                                       | The request object to use for the request.                                                                                                                                               |
| `opts`                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                       | The options for this request.                                                                                                                                                            |

### Response

**[*operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetDiscoverySnapshotResponse](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicegetdiscoverysnapshotresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetEgressPolicy

GetEgressPolicy returns both the typed rules (for the authoring UI) and
 the compiled TB policy YAML + generation (for timebanditd's poller) in
 one call -- one instance's policy is small enough that splitting the
 two reads into separate RPCs would be pure surface area, not a real
 cost saving for either caller.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.tbcontrolplane.v1.TBControlPlaneService.GetEgressPolicy" method="get" path="/api/v1/tb-control-plane/egress-policy/{tb_instance_id}" -->
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

    res, err := s.TBControlPlane.GetEgressPolicy(ctx, operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetEgressPolicyRequest{
        TbInstanceID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TBControlPlaneServiceGetEgressPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                          | :heavy_check_mark:                                                                                                                                                             | The context to use for the request.                                                                                                                                            |
| `request`                                                                                                                                                                      | [operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetEgressPolicyRequest](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicegetegresspolicyrequest.md) | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `opts`                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                   | :heavy_minus_sign:                                                                                                                                                             | The options for this request.                                                                                                                                                  |

### Response

**[*operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetEgressPolicyResponse](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicegetegresspolicyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~GetTrafficSummary~~

Deprecated: analytics require an Edge service resource. Use the EdgeService
 egress or inference endpoint; this operation returns UNIMPLEMENTED.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.tbcontrolplane.v1.TBControlPlaneService.GetTrafficSummary" method="get" path="/api/v1/apps/{app_id}/traffic-summary" -->
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

    res, err := s.TBControlPlane.GetTrafficSummary(ctx, operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetTrafficSummaryRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TBControlPlaneServiceGetTrafficSummaryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                          | Type                                                                                                                                                                               | Required                                                                                                                                                                           | Description                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                              | :heavy_check_mark:                                                                                                                                                                 | The context to use for the request.                                                                                                                                                |
| `request`                                                                                                                                                                          | [operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetTrafficSummaryRequest](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicegettrafficsummaryrequest.md) | :heavy_check_mark:                                                                                                                                                                 | The request object to use for the request.                                                                                                                                         |
| `opts`                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                 | The options for this request.                                                                                                                                                      |

### Response

**[*operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetTrafficSummaryResponse](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicegettrafficsummaryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~GetUsageAttribution~~

Deprecated: analytics require an Edge service resource. Use the EdgeService
 egress or inference endpoint; this operation returns UNIMPLEMENTED.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.tbcontrolplane.v1.TBControlPlaneService.GetUsageAttribution" method="get" path="/api/v1/tb-control-plane/usage-attribution" -->
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

    res, err := s.TBControlPlane.GetUsageAttribution(ctx, operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetUsageAttributionRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.TBControlPlaneServiceGetUsageAttributionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                              | Type                                                                                                                                                                                   | Required                                                                                                                                                                               | Description                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                     | The context to use for the request.                                                                                                                                                    |
| `request`                                                                                                                                                                              | [operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetUsageAttributionRequest](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicegetusageattributionrequest.md) | :heavy_check_mark:                                                                                                                                                                     | The request object to use for the request.                                                                                                                                             |
| `opts`                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                     | The options for this request.                                                                                                                                                          |

### Response

**[*operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetUsageAttributionResponse](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicegetusageattributionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~GetUsageSummary~~

Deprecated: analytics require an Edge service resource. Use the EdgeService
 egress or inference endpoint; this operation returns UNIMPLEMENTED.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.tbcontrolplane.v1.TBControlPlaneService.GetUsageSummary" method="get" path="/api/v1/tb-control-plane/usage-summary" -->
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

    res, err := s.TBControlPlane.GetUsageSummary(ctx, operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetUsageSummaryRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.TBControlPlaneServiceGetUsageSummaryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                          | :heavy_check_mark:                                                                                                                                                             | The context to use for the request.                                                                                                                                            |
| `request`                                                                                                                                                                      | [operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetUsageSummaryRequest](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicegetusagesummaryrequest.md) | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `opts`                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                   | :heavy_minus_sign:                                                                                                                                                             | The options for this request.                                                                                                                                                  |

### Response

**[*operations.C1APITbcontrolplaneV1TBControlPlaneServiceGetUsageSummaryResponse](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicegetusagesummaryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~ListTopAgents~~

Deprecated: analytics require an Edge service resource. Use the EdgeService
 egress or inference endpoint; this operation returns UNIMPLEMENTED.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.tbcontrolplane.v1.TBControlPlaneService.ListTopAgents" method="get" path="/api/v1/apps/{app_id}/traffic-top-agents" -->
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

    res, err := s.TBControlPlane.ListTopAgents(ctx, operations.C1APITbcontrolplaneV1TBControlPlaneServiceListTopAgentsRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TBControlPlaneServiceListTopAgentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.C1APITbcontrolplaneV1TBControlPlaneServiceListTopAgentsRequest](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicelisttopagentsrequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |
| `opts`                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                         | The options for this request.                                                                                                                                              |

### Response

**[*operations.C1APITbcontrolplaneV1TBControlPlaneServiceListTopAgentsResponse](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicelisttopagentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~ListTopTools~~

Deprecated: analytics require an Edge service resource. Use the EdgeService
 egress or inference endpoint; this operation returns UNIMPLEMENTED.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.tbcontrolplane.v1.TBControlPlaneService.ListTopTools" method="get" path="/api/v1/apps/{app_id}/traffic-top-tools" -->
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

    res, err := s.TBControlPlane.ListTopTools(ctx, operations.C1APITbcontrolplaneV1TBControlPlaneServiceListTopToolsRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TBControlPlaneServiceListTopToolsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [operations.C1APITbcontrolplaneV1TBControlPlaneServiceListTopToolsRequest](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicelisttoptoolsrequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |
| `opts`                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                                       | The options for this request.                                                                                                                                            |

### Response

**[*operations.C1APITbcontrolplaneV1TBControlPlaneServiceListTopToolsResponse](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicelisttoptoolsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~ListTrafficEvents~~

Deprecated: analytics require an Edge service resource. Use the EdgeService
 egress or inference endpoint; this operation returns UNIMPLEMENTED.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.tbcontrolplane.v1.TBControlPlaneService.ListTrafficEvents" method="get" path="/api/v1/apps/{app_id}/traffic-events" -->
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

    res, err := s.TBControlPlane.ListTrafficEvents(ctx, operations.C1APITbcontrolplaneV1TBControlPlaneServiceListTrafficEventsRequest{
        AppID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TBControlPlaneServiceListTrafficEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                          | Type                                                                                                                                                                               | Required                                                                                                                                                                           | Description                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                              | :heavy_check_mark:                                                                                                                                                                 | The context to use for the request.                                                                                                                                                |
| `request`                                                                                                                                                                          | [operations.C1APITbcontrolplaneV1TBControlPlaneServiceListTrafficEventsRequest](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicelisttrafficeventsrequest.md) | :heavy_check_mark:                                                                                                                                                                 | The request object to use for the request.                                                                                                                                         |
| `opts`                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                 | The options for this request.                                                                                                                                                      |

### Response

**[*operations.C1APITbcontrolplaneV1TBControlPlaneServiceListTrafficEventsResponse](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicelisttrafficeventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PushDiscovery

PushDiscovery records the vocabulary a Time Bandit instance reports for
 itself -- principal, scope, destination, credential recipe, posture and
 route names -- replacing any prior snapshot. Called by timebanditd when
 its local configuration changes.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.tbcontrolplane.v1.TBControlPlaneService.PushDiscovery" method="post" path="/api/v1/tb-control-plane/discovery" -->
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

    res, err := s.TBControlPlane.PushDiscovery(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TBControlPlaneServicePushDiscoveryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [shared.TBControlPlaneServicePushDiscoveryRequest](../../pkg/models/shared/tbcontrolplaneservicepushdiscoveryrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.C1APITbcontrolplaneV1TBControlPlaneServicePushDiscoveryResponse](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicepushdiscoveryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SaveEgressPolicy

SaveEgressPolicy validates every rule against the instance's latest
 discovery snapshot (principal/scope/destination/posture/credential
 must all be names discovery actually reported), bumps generation, and
 stores the result. Used by both the authoring UI (stage 3) and any
 direct API caller.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.tbcontrolplane.v1.TBControlPlaneService.SaveEgressPolicy" method="put" path="/api/v1/tb-control-plane/egress-policy" -->
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

    res, err := s.TBControlPlane.SaveEgressPolicy(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TBControlPlaneServiceSaveEgressPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [shared.TBControlPlaneServiceSaveEgressPolicyRequest](../../pkg/models/shared/tbcontrolplaneservicesaveegresspolicyrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.C1APITbcontrolplaneV1TBControlPlaneServiceSaveEgressPolicyResponse](../../pkg/models/operations/c1apitbcontrolplanev1tbcontrolplaneservicesaveegresspolicyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |