# ConnectorSyncDiagnostics

## Overview

### Available Operations

* [CompareSyncStats](#comparesyncstats) - Compare Sync Stats
* [GetSyncChangeTimeline](#getsyncchangetimeline) - Get Sync Change Timeline
* [GetSyncErrorDigest](#getsyncerrordigest) - Get Sync Error Digest
* [GetSyncHistory](#getsynchistory) - Get Sync History
* [GetSyncStatus](#getsyncstatus) - Get Sync Status
* [ListDebuggableConnectors](#listdebuggableconnectors) - List Debuggable Connectors

## CompareSyncStats

Compare recorded aggregate counts between two runs or a production run and a draft test for the same connector.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.ConnectorSyncDiagnosticsService.CompareSyncStats" method="post" path="/api/v1/apps/{app_id}/connectors/{connector_id}/sync_diagnostics/compare_stats" -->
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

    res, err := s.ConnectorSyncDiagnostics.CompareSyncStats(ctx, operations.C1APIAppV1ConnectorSyncDiagnosticsServiceCompareSyncStatsRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CompareSyncStatsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                          | :heavy_check_mark:                                                                                                                                                             | The context to use for the request.                                                                                                                                            |
| `request`                                                                                                                                                                      | [operations.C1APIAppV1ConnectorSyncDiagnosticsServiceCompareSyncStatsRequest](../../pkg/models/operations/c1apiappv1connectorsyncdiagnosticsservicecomparesyncstatsrequest.md) | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `opts`                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                   | :heavy_minus_sign:                                                                                                                                                             | The options for this request.                                                                                                                                                  |

### Response

**[*operations.C1APIAppV1ConnectorSyncDiagnosticsServiceCompareSyncStatsResponse](../../pkg/models/operations/c1apiappv1connectorsyncdiagnosticsservicecomparesyncstatsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSyncChangeTimeline

Get a bounded metadata-only timeline of configuration, release and deployment changes around sync failures.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.ConnectorSyncDiagnosticsService.GetSyncChangeTimeline" method="post" path="/api/v1/apps/{app_id}/connectors/{connector_id}/sync_diagnostics/change_timeline" -->
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

    res, err := s.ConnectorSyncDiagnostics.GetSyncChangeTimeline(ctx, operations.C1APIAppV1ConnectorSyncDiagnosticsServiceGetSyncChangeTimelineRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSyncChangeTimelineResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                | Type                                                                                                                                                                                     | Required                                                                                                                                                                                 | Description                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                       | The context to use for the request.                                                                                                                                                      |
| `request`                                                                                                                                                                                | [operations.C1APIAppV1ConnectorSyncDiagnosticsServiceGetSyncChangeTimelineRequest](../../pkg/models/operations/c1apiappv1connectorsyncdiagnosticsservicegetsyncchangetimelinerequest.md) | :heavy_check_mark:                                                                                                                                                                       | The request object to use for the request.                                                                                                                                               |
| `opts`                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                       | The options for this request.                                                                                                                                                            |

### Response

**[*operations.C1APIAppV1ConnectorSyncDiagnosticsServiceGetSyncChangeTimelineResponse](../../pkg/models/operations/c1apiappv1connectorsyncdiagnosticsservicegetsyncchangetimelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSyncErrorDigest

Get a bounded digest of classified, redacted recent errors rather than raw log pages.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.ConnectorSyncDiagnosticsService.GetSyncErrorDigest" method="get" path="/api/v1/apps/{app_id}/connectors/{connector_id}/sync_diagnostics/error_digest" -->
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

    res, err := s.ConnectorSyncDiagnostics.GetSyncErrorDigest(ctx, operations.C1APIAppV1ConnectorSyncDiagnosticsServiceGetSyncErrorDigestRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSyncErrorDigestResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                          | Type                                                                                                                                                                               | Required                                                                                                                                                                           | Description                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                              | :heavy_check_mark:                                                                                                                                                                 | The context to use for the request.                                                                                                                                                |
| `request`                                                                                                                                                                          | [operations.C1APIAppV1ConnectorSyncDiagnosticsServiceGetSyncErrorDigestRequest](../../pkg/models/operations/c1apiappv1connectorsyncdiagnosticsservicegetsyncerrordigestrequest.md) | :heavy_check_mark:                                                                                                                                                                 | The request object to use for the request.                                                                                                                                         |
| `opts`                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                 | The options for this request.                                                                                                                                                      |

### Response

**[*operations.C1APIAppV1ConnectorSyncDiagnosticsServiceGetSyncErrorDigestResponse](../../pkg/models/operations/c1apiappv1connectorsyncdiagnosticsservicegetsyncerrordigestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSyncHistory

Get recent sync runs and their failure trend without configuration or synced records.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.ConnectorSyncDiagnosticsService.GetSyncHistory" method="get" path="/api/v1/apps/{app_id}/connectors/{connector_id}/sync_diagnostics/history" -->
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

    res, err := s.ConnectorSyncDiagnostics.GetSyncHistory(ctx, operations.C1APIAppV1ConnectorSyncDiagnosticsServiceGetSyncHistoryRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSyncHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.C1APIAppV1ConnectorSyncDiagnosticsServiceGetSyncHistoryRequest](../../pkg/models/operations/c1apiappv1connectorsyncdiagnosticsservicegetsynchistoryrequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |
| `opts`                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                         | The options for this request.                                                                                                                                              |

### Response

**[*operations.C1APIAppV1ConnectorSyncDiagnosticsServiceGetSyncHistoryResponse](../../pkg/models/operations/c1apiappv1connectorsyncdiagnosticsservicegetsynchistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSyncStatus

Get current sync progress and classified, redacted errors for a connector the caller manages.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.ConnectorSyncDiagnosticsService.GetSyncStatus" method="get" path="/api/v1/apps/{app_id}/connectors/{connector_id}/sync_diagnostics/status" -->
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

    res, err := s.ConnectorSyncDiagnostics.GetSyncStatus(ctx, operations.C1APIAppV1ConnectorSyncDiagnosticsServiceGetSyncStatusRequest{
        AppID: "<id>",
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSyncStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [operations.C1APIAppV1ConnectorSyncDiagnosticsServiceGetSyncStatusRequest](../../pkg/models/operations/c1apiappv1connectorsyncdiagnosticsservicegetsyncstatusrequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |
| `opts`                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                                       | The options for this request.                                                                                                                                            |

### Response

**[*operations.C1APIAppV1ConnectorSyncDiagnosticsServiceGetSyncStatusResponse](../../pkg/models/operations/c1apiappv1connectorsyncdiagnosticsservicegetsyncstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDebuggableConnectors

List connectors the caller may debug; each result contains only identification and sync metadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.ConnectorSyncDiagnosticsService.ListDebuggableConnectors" method="get" path="/api/v1/sync_diagnostics/connectors" -->
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

    res, err := s.ConnectorSyncDiagnostics.ListDebuggableConnectors(ctx, operations.C1APIAppV1ConnectorSyncDiagnosticsServiceListDebuggableConnectorsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListDebuggableConnectorsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                      | Type                                                                                                                                                                                           | Required                                                                                                                                                                                       | Description                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                             | The context to use for the request.                                                                                                                                                            |
| `request`                                                                                                                                                                                      | [operations.C1APIAppV1ConnectorSyncDiagnosticsServiceListDebuggableConnectorsRequest](../../pkg/models/operations/c1apiappv1connectorsyncdiagnosticsservicelistdebuggableconnectorsrequest.md) | :heavy_check_mark:                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                     |
| `opts`                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                             | The options for this request.                                                                                                                                                                  |

### Response

**[*operations.C1APIAppV1ConnectorSyncDiagnosticsServiceListDebuggableConnectorsResponse](../../pkg/models/operations/c1apiappv1connectorsyncdiagnosticsservicelistdebuggableconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |