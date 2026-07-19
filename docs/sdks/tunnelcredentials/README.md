# TunnelCredentials

## Overview

### Available Operations

* [CreateBridge](#createbridge) - Create Bridge
* [CreateBridgeCredential](#createbridgecredential) - Create Bridge Credential
* [DeleteBridge](#deletebridge) - Delete Bridge
* [GetBridge](#getbridge) - Get Bridge
* [ListBridgeAnnouncedServices](#listbridgeannouncedservices) - List Bridge Announced Services
* [ListBridgeCredentials](#listbridgecredentials) - List Bridge Credentials
* [ListBridges](#listbridges) - List Bridges
* [RevokeBridgeCredential](#revokebridgecredential) - Revoke Bridge Credential
* [UpdateBridge](#updatebridge) - Update Bridge

## CreateBridge

CreateBridge creates a bridge with no credentials. Use
 CreateBridgeCredential to mint the first credential.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.TunnelCredentialsService.CreateBridge" method="post" path="/api/v1/iam/tunnel/bridges" -->
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

    res, err := s.TunnelCredentials.CreateBridge(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TunnelCredentialsServiceCreateBridgeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [shared.TunnelCredentialsServiceCreateBridgeRequest](../../pkg/models/shared/tunnelcredentialsservicecreatebridgerequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.C1APIIamV1TunnelCredentialsServiceCreateBridgeResponse](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicecreatebridgeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateBridgeCredential

CreateBridgeCredential mints a credential for a bridge. If the bridge
 already has an active credential, it is revoked. The plaintext
 client_secret is returned exactly once on the response.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.TunnelCredentialsService.CreateBridgeCredential" method="post" path="/api/v1/iam/tunnel/bridges/{bridge_id}/credentials" -->
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

    res, err := s.TunnelCredentials.CreateBridgeCredential(ctx, operations.C1APIIamV1TunnelCredentialsServiceCreateBridgeCredentialRequest{
        BridgeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TunnelCredentialsServiceCreateBridgeCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.C1APIIamV1TunnelCredentialsServiceCreateBridgeCredentialRequest](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicecreatebridgecredentialrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |
| `opts`                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                           | The options for this request.                                                                                                                                                |

### Response

**[*operations.C1APIIamV1TunnelCredentialsServiceCreateBridgeCredentialResponse](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicecreatebridgecredentialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteBridge

DeleteBridge hard-deletes a bridge and every credential it owns.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.TunnelCredentialsService.DeleteBridge" method="delete" path="/api/v1/iam/tunnel/bridges/{id}" -->
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

    res, err := s.TunnelCredentials.DeleteBridge(ctx, operations.C1APIIamV1TunnelCredentialsServiceDeleteBridgeRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TunnelCredentialsServiceDeleteBridgeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APIIamV1TunnelCredentialsServiceDeleteBridgeRequest](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicedeletebridgerequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APIIamV1TunnelCredentialsServiceDeleteBridgeResponse](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicedeletebridgeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetBridge

GetBridge returns a bridge by id with live appliance status.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.TunnelCredentialsService.GetBridge" method="get" path="/api/v1/iam/tunnel/bridges/{id}" -->
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

    res, err := s.TunnelCredentials.GetBridge(ctx, operations.C1APIIamV1TunnelCredentialsServiceGetBridgeRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TunnelCredentialsServiceGetBridgeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIIamV1TunnelCredentialsServiceGetBridgeRequest](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicegetbridgerequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APIIamV1TunnelCredentialsServiceGetBridgeResponse](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicegetbridgeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListBridgeAnnouncedServices

ListBridgeAnnouncedServices returns the services the appliance is
 currently announcing for this bridge. Read live from the tunnel store;
 empty when no appliance is connected.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.TunnelCredentialsService.ListBridgeAnnouncedServices" method="get" path="/api/v1/iam/tunnel/bridges/{bridge_id}/announced_services" -->
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

    res, err := s.TunnelCredentials.ListBridgeAnnouncedServices(ctx, operations.C1APIIamV1TunnelCredentialsServiceListBridgeAnnouncedServicesRequest{
        BridgeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TunnelCredentialsServiceListBridgeAnnouncedServicesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                              | Type                                                                                                                                                                                   | Required                                                                                                                                                                               | Description                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                     | The context to use for the request.                                                                                                                                                    |
| `request`                                                                                                                                                                              | [operations.C1APIIamV1TunnelCredentialsServiceListBridgeAnnouncedServicesRequest](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicelistbridgeannouncedservicesrequest.md) | :heavy_check_mark:                                                                                                                                                                     | The request object to use for the request.                                                                                                                                             |
| `opts`                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                     | The options for this request.                                                                                                                                                          |

### Response

**[*operations.C1APIIamV1TunnelCredentialsServiceListBridgeAnnouncedServicesResponse](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicelistbridgeannouncedservicesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListBridgeCredentials

ListBridgeCredentials returns every credential (active + revoked) for
 one bridge.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.TunnelCredentialsService.ListBridgeCredentials" method="get" path="/api/v1/iam/tunnel/bridges/{bridge_id}/credentials" -->
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

    res, err := s.TunnelCredentials.ListBridgeCredentials(ctx, operations.C1APIIamV1TunnelCredentialsServiceListBridgeCredentialsRequest{
        BridgeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TunnelCredentialsServiceListBridgeCredentialsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.C1APIIamV1TunnelCredentialsServiceListBridgeCredentialsRequest](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicelistbridgecredentialsrequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |
| `opts`                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                         | The options for this request.                                                                                                                                              |

### Response

**[*operations.C1APIIamV1TunnelCredentialsServiceListBridgeCredentialsResponse](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicelistbridgecredentialsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListBridges

ListBridges returns the tenant's bridges, paginated, each with live
 appliance status.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.TunnelCredentialsService.ListBridges" method="get" path="/api/v1/iam/tunnel/bridges" -->
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

    res, err := s.TunnelCredentials.ListBridges(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.TunnelCredentialsServiceListBridgesResponse != nil {
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

**[*operations.C1APIIamV1TunnelCredentialsServiceListBridgesResponse](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicelistbridgesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RevokeBridgeCredential

RevokeBridgeCredential soft-revokes one credential by id. The row is
 retained for audit; token mints with this credential are rejected.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.TunnelCredentialsService.RevokeBridgeCredential" method="delete" path="/api/v1/iam/tunnel/credentials/{id}" -->
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

    res, err := s.TunnelCredentials.RevokeBridgeCredential(ctx, operations.C1APIIamV1TunnelCredentialsServiceRevokeBridgeCredentialRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TunnelCredentialsServiceRevokeBridgeCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.C1APIIamV1TunnelCredentialsServiceRevokeBridgeCredentialRequest](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicerevokebridgecredentialrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |
| `opts`                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                           | The options for this request.                                                                                                                                                |

### Response

**[*operations.C1APIIamV1TunnelCredentialsServiceRevokeBridgeCredentialResponse](../../pkg/models/operations/c1apiiamv1tunnelcredentialsservicerevokebridgecredentialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateBridge

UpdateBridge patches a bridge's editable metadata (display_name,
 description). Credentials are not affected.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.iam.v1.TunnelCredentialsService.UpdateBridge" method="patch" path="/api/v1/iam/tunnel/bridges/{id}" -->
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

    res, err := s.TunnelCredentials.UpdateBridge(ctx, operations.C1APIIamV1TunnelCredentialsServiceUpdateBridgeRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TunnelCredentialsServiceUpdateBridgeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APIIamV1TunnelCredentialsServiceUpdateBridgeRequest](../../pkg/models/operations/c1apiiamv1tunnelcredentialsserviceupdatebridgerequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APIIamV1TunnelCredentialsServiceUpdateBridgeResponse](../../pkg/models/operations/c1apiiamv1tunnelcredentialsserviceupdatebridgeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |