# FindingSettings

## Overview

### Available Operations

* [ListFindingSettings](#listfindingsettings) - List Finding Settings
* [UpdateFindingSettings](#updatefindingsettings) - Update Finding Settings

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