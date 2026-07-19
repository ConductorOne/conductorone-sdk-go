# PersonalDevice

## Overview

### Available Operations

* [GetDevice](#getdevice) - Get Device
* [ListDeviceClients](#listdeviceclients) - List Device Clients
* [RevokeDevice](#revokedevice) - Revoke Device
* [RevokeDeviceClient](#revokedeviceclient) - Revoke Device Client
* [Search](#search) - NOTE: Only shows devices for the current user.
* [UpdateDevice](#updatedevice) - NOTE: Only updates devices owned by the current user.

## GetDevice

GetDevice retrieves a single device (by device_id) with its child clients.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.PersonalDeviceService.GetDevice" method="get" path="/api/v1/iam/personal_devices/{device_id}" -->
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

    res, err := s.PersonalDevice.GetDevice(ctx, operations.C1APIIamV1PersonalDeviceServiceGetDeviceRequest{
        DeviceID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PersonalDeviceServiceGetDeviceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIIamV1PersonalDeviceServiceGetDeviceRequest](../../pkg/models/operations/c1apiiamv1personaldeviceservicegetdevicerequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APIIamV1PersonalDeviceServiceGetDeviceResponse](../../pkg/models/operations/c1apiiamv1personaldeviceservicegetdeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDeviceClients

ListDeviceClients returns the app clients registered on a device, one page at
 a time. A device can accrue many app clients over time, so the clients are
 served from this dedicated paginated endpoint rather than inlined on the
 device.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.PersonalDeviceService.ListDeviceClients" method="get" path="/api/v1/iam/personal_devices/{device_id}/clients" -->
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

    res, err := s.PersonalDevice.ListDeviceClients(ctx, operations.C1APIIamV1PersonalDeviceServiceListDeviceClientsRequest{
        DeviceID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PersonalDeviceServiceListDeviceClientsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIIamV1PersonalDeviceServiceListDeviceClientsRequest](../../pkg/models/operations/c1apiiamv1personaldeviceservicelistdeviceclientsrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIIamV1PersonalDeviceServiceListDeviceClientsResponse](../../pkg/models/operations/c1apiiamv1personaldeviceservicelistdeviceclientsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RevokeDevice

RevokeDevice revokes a whole device: it revokes the device and removes every
 app client on it, so no app on that machine can mint further tokens.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.PersonalDeviceService.RevokeDevice" method="delete" path="/api/v1/iam/personal_devices/{device_id}" -->
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

    res, err := s.PersonalDevice.RevokeDevice(ctx, operations.C1APIIamV1PersonalDeviceServiceRevokeDeviceRequest{
        DeviceID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PersonalDeviceServiceRevokeDeviceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIIamV1PersonalDeviceServiceRevokeDeviceRequest](../../pkg/models/operations/c1apiiamv1personaldeviceservicerevokedevicerequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APIIamV1PersonalDeviceServiceRevokeDeviceResponse](../../pkg/models/operations/c1apiiamv1personaldeviceservicerevokedeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RevokeDeviceClient

RevokeDeviceClient revokes a single app client on a device: it removes that
 client, leaving the device and its other clients intact.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.PersonalDeviceService.RevokeDeviceClient" method="delete" path="/api/v1/iam/personal_devices/{device_id}/clients/{id}" -->
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

    res, err := s.PersonalDevice.RevokeDeviceClient(ctx, operations.C1APIIamV1PersonalDeviceServiceRevokeDeviceClientRequest{
        DeviceID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PersonalDeviceServiceRevokeDeviceClientResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.C1APIIamV1PersonalDeviceServiceRevokeDeviceClientRequest](../../pkg/models/operations/c1apiiamv1personaldeviceservicerevokedeviceclientrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |
| `opts`                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                   | :heavy_minus_sign:                                                                                                                                             | The options for this request.                                                                                                                                  |

### Response

**[*operations.C1APIIamV1PersonalDeviceServiceRevokeDeviceClientResponse](../../pkg/models/operations/c1apiiamv1personaldeviceservicerevokedeviceclientresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Search

Search returns the calling user's registered devices, ordered by display name.
 By default only active devices are returned; use the status filter to include
 revoked devices. Optionally filter by a display-name query.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.PersonalDeviceService.Search" method="post" path="/api/v1/search/iam/personal_devices" -->
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

    res, err := s.PersonalDevice.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PersonalDeviceServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [shared.PersonalDeviceServiceSearchRequest](../../pkg/models/shared/personaldeviceservicesearchrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.C1APIIamV1PersonalDeviceServiceSearchResponse](../../pkg/models/operations/c1apiiamv1personaldeviceservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDevice

UpdateDevice renames one of the calling user's devices. Use the update mask
 to specify which fields to change; only the display name is mutable, so
 device identity and keys never change.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.PersonalDeviceService.UpdateDevice" method="put" path="/api/v1/iam/personal_devices/{device_id}" -->
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

    res, err := s.PersonalDevice.UpdateDevice(ctx, operations.C1APIIamV1PersonalDeviceServiceUpdateDeviceRequest{
        DeviceID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PersonalDeviceServiceUpdateDeviceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIIamV1PersonalDeviceServiceUpdateDeviceRequest](../../pkg/models/operations/c1apiiamv1personaldeviceserviceupdatedevicerequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APIIamV1PersonalDeviceServiceUpdateDeviceResponse](../../pkg/models/operations/c1apiiamv1personaldeviceserviceupdatedeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |