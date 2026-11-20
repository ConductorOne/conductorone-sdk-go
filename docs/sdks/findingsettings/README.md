# FindingSettings

## Overview

### Available Operations

* [GetEdgeFindingSettings](#getedgefindingsettings) - Get Edge Finding Settings
* [ListFindingSettings](#listfindingsettings) - List Finding Settings
* [UpdateEdgeFindingSettings](#updateedgefindingsettings) - Update Edge Finding Settings
* [UpdateFindingSettings](#updatefindingsettings) - Update Finding Settings

## GetEdgeFindingSettings

Get whether each Edge finding category is detected. Categories never
 configured read back their shipped default. The Edge finding type's own
 switch (ListFindingSettings) still gates every category.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingSettingsService.GetEdgeFindingSettings" method="get" path="/api/v1/findings/settings/edge" -->
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

    res, err := s.FindingSettings.GetEdgeFindingSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetEdgeFindingSettingsResponse != nil {
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

**[*operations.C1APIFindingV1FindingSettingsServiceGetEdgeFindingSettingsResponse](../../pkg/models/operations/c1apifindingv1findingsettingsservicegetedgefindingsettingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListFindingSettings

List every configurable finding type and whether detection is enabled.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingSettingsService.ListFindingSettings" method="get" path="/api/v1/findings/settings" -->
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

    res, err := s.FindingSettings.ListFindingSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListFindingSettingsResponse != nil {
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

**[*operations.C1APIFindingV1FindingSettingsServiceListFindingSettingsResponse](../../pkg/models/operations/c1apifindingv1findingsettingsservicelistfindingsettingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateEdgeFindingSettings

Enable or disable detection for one or more Edge finding categories in a
 single write.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingSettingsService.UpdateEdgeFindingSettings" method="post" path="/api/v1/findings/settings/edge/update" -->
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

    res, err := s.FindingSettings.UpdateEdgeFindingSettings(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateEdgeFindingSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [shared.UpdateEdgeFindingSettingsRequest](../../pkg/models/shared/updateedgefindingsettingsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.C1APIFindingV1FindingSettingsServiceUpdateEdgeFindingSettingsResponse](../../pkg/models/operations/c1apifindingv1findingsettingsserviceupdateedgefindingsettingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateFindingSettings

Enable or disable detection for one or more finding types in a single
 write. Enabling a type whose detector is a scheduled job also queues an
 immediate run.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.finding.v1.FindingSettingsService.UpdateFindingSettings" method="post" path="/api/v1/findings/settings/update" -->
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

    res, err := s.FindingSettings.UpdateFindingSettings(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateFindingSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [shared.UpdateFindingSettingsRequest](../../pkg/models/shared/updatefindingsettingsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.C1APIFindingV1FindingSettingsServiceUpdateFindingSettingsResponse](../../pkg/models/operations/c1apifindingv1findingsettingsserviceupdatefindingsettingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |