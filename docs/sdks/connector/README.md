# Connector
(*Connector*)

### Available Operations

* [Create](#create) - Create
* [CreateDelegated](#createdelegated) - Create Delegated
* [Delete](#delete) - Delete
* [ForceSync](#forcesync) - Force Sync
* [Get](#get) - Get
* [GetCredentials](#getcredentials) - Get Credentials
* [List](#list) - List
* [RevokeCredential](#revokecredential) - Revoke Credential
* [RotateCredential](#rotatecredential) - Rotate Credential
* [Update](#update) - Update
* [UpdateDelegated](#updatedelegated) - Update Delegated

## Create

Create a configured connector.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: shared.SchemeOauth{
                ClientID: "<YOUR_CLIENT_ID_HERE>",
                ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
                TokenURL: "<YOUR_TOKEN_URL_HERE>",
            },
        }),
    )
    request := operations.C1APIAppV1ConnectorServiceCreateRequest{
        AppID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Connector.Create(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.C1APIAppV1ConnectorServiceCreateRequest](../../pkg/models/operations/c1apiappv1connectorservicecreaterequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |


### Response

**[*operations.C1APIAppV1ConnectorServiceCreateResponse](../../pkg/models/operations/c1apiappv1connectorservicecreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateDelegated

Create a connector that is pending a connector config.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: shared.SchemeOauth{
                ClientID: "<YOUR_CLIENT_ID_HERE>",
                ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
                TokenURL: "<YOUR_TOKEN_URL_HERE>",
            },
        }),
    )
    request := operations.C1APIAppV1ConnectorServiceCreateDelegatedRequest{
        AppID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Connector.CreateDelegated(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAppV1ConnectorServiceCreateDelegatedRequest](../../pkg/models/operations/c1apiappv1connectorservicecreatedelegatedrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |


### Response

**[*operations.C1APIAppV1ConnectorServiceCreateDelegatedResponse](../../pkg/models/operations/c1apiappv1connectorservicecreatedelegatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Delete

Delete a connector.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: shared.SchemeOauth{
                ClientID: "<YOUR_CLIENT_ID_HERE>",
                ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
                TokenURL: "<YOUR_TOKEN_URL_HERE>",
            },
        }),
    )
    request := operations.C1APIAppV1ConnectorServiceDeleteRequest{
        AppID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Connector.Delete(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.C1APIAppV1ConnectorServiceDeleteRequest](../../pkg/models/operations/c1apiappv1connectorservicedeleterequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |


### Response

**[*operations.C1APIAppV1ConnectorServiceDeleteResponse](../../pkg/models/operations/c1apiappv1connectorservicedeleteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ForceSync

Invokes the c1.api.app.v1.ConnectorService.ForceSync method.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: shared.SchemeOauth{
                ClientID: "<YOUR_CLIENT_ID_HERE>",
                ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
                TokenURL: "<YOUR_TOKEN_URL_HERE>",
            },
        }),
    )
    request := operations.C1APIAppV1ConnectorServiceForceSyncRequest{
        AppID: "<value>",
        ConnectorID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Connector.ForceSync(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ForceSyncResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.C1APIAppV1ConnectorServiceForceSyncRequest](../../pkg/models/operations/c1apiappv1connectorserviceforcesyncrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |


### Response

**[*operations.C1APIAppV1ConnectorServiceForceSyncResponse](../../pkg/models/operations/c1apiappv1connectorserviceforcesyncresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Get

Get a connector.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: shared.SchemeOauth{
                ClientID: "<YOUR_CLIENT_ID_HERE>",
                ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
                TokenURL: "<YOUR_TOKEN_URL_HERE>",
            },
        }),
    )
    request := operations.C1APIAppV1ConnectorServiceGetRequest{
        AppID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Connector.Get(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.C1APIAppV1ConnectorServiceGetRequest](../../pkg/models/operations/c1apiappv1connectorservicegetrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.C1APIAppV1ConnectorServiceGetResponse](../../pkg/models/operations/c1apiappv1connectorservicegetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCredentials

Get credentials for a connector.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: shared.SchemeOauth{
                ClientID: "<YOUR_CLIENT_ID_HERE>",
                ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
                TokenURL: "<YOUR_TOKEN_URL_HERE>",
            },
        }),
    )
    request := operations.C1APIAppV1ConnectorServiceGetCredentialsRequest{
        AppID: "<value>",
        ConnectorID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Connector.GetCredentials(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorServiceGetCredentialsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIAppV1ConnectorServiceGetCredentialsRequest](../../pkg/models/operations/c1apiappv1connectorservicegetcredentialsrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |


### Response

**[*operations.C1APIAppV1ConnectorServiceGetCredentialsResponse](../../pkg/models/operations/c1apiappv1connectorservicegetcredentialsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

List connectors for an app.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: shared.SchemeOauth{
                ClientID: "<YOUR_CLIENT_ID_HERE>",
                ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
                TokenURL: "<YOUR_TOKEN_URL_HERE>",
            },
        }),
    )
    request := operations.C1APIAppV1ConnectorServiceListRequest{
        AppID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Connector.List(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.C1APIAppV1ConnectorServiceListRequest](../../pkg/models/operations/c1apiappv1connectorservicelistrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |


### Response

**[*operations.C1APIAppV1ConnectorServiceListResponse](../../pkg/models/operations/c1apiappv1connectorservicelistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RevokeCredential

Revoke credentials for a connector.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: shared.SchemeOauth{
                ClientID: "<YOUR_CLIENT_ID_HERE>",
                ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
                TokenURL: "<YOUR_TOKEN_URL_HERE>",
            },
        }),
    )
    request := operations.C1APIAppV1ConnectorServiceRevokeCredentialRequest{
        AppID: "<value>",
        ConnectorID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Connector.RevokeCredential(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorServiceRevokeCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.C1APIAppV1ConnectorServiceRevokeCredentialRequest](../../pkg/models/operations/c1apiappv1connectorservicerevokecredentialrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |


### Response

**[*operations.C1APIAppV1ConnectorServiceRevokeCredentialResponse](../../pkg/models/operations/c1apiappv1connectorservicerevokecredentialresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RotateCredential

Rotate credentials for a connector.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: shared.SchemeOauth{
                ClientID: "<YOUR_CLIENT_ID_HERE>",
                ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
                TokenURL: "<YOUR_TOKEN_URL_HERE>",
            },
        }),
    )
    request := operations.C1APIAppV1ConnectorServiceRotateCredentialRequest{
        AppID: "<value>",
        ConnectorID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Connector.RotateCredential(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorServiceRotateCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.C1APIAppV1ConnectorServiceRotateCredentialRequest](../../pkg/models/operations/c1apiappv1connectorservicerotatecredentialrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |


### Response

**[*operations.C1APIAppV1ConnectorServiceRotateCredentialResponse](../../pkg/models/operations/c1apiappv1connectorservicerotatecredentialresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Update a connector.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: shared.SchemeOauth{
                ClientID: "<YOUR_CLIENT_ID_HERE>",
                ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
                TokenURL: "<YOUR_TOKEN_URL_HERE>",
            },
        }),
    )
    request := operations.C1APIAppV1ConnectorServiceUpdateRequest{
        AppID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Connector.Update(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.C1APIAppV1ConnectorServiceUpdateRequest](../../pkg/models/operations/c1apiappv1connectorserviceupdaterequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |


### Response

**[*operations.C1APIAppV1ConnectorServiceUpdateResponse](../../pkg/models/operations/c1apiappv1connectorserviceupdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateDelegated

Update a delegated connector.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: shared.SchemeOauth{
                ClientID: "<YOUR_CLIENT_ID_HERE>",
                ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
                TokenURL: "<YOUR_TOKEN_URL_HERE>",
            },
        }),
    )
    request := operations.C1APIAppV1ConnectorServiceUpdateDelegatedRequest{
        ConnectorAppID: "<value>",
        ConnectorID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Connector.UpdateDelegated(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAppV1ConnectorServiceUpdateDelegatedRequest](../../pkg/models/operations/c1apiappv1connectorserviceupdatedelegatedrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |


### Response

**[*operations.C1APIAppV1ConnectorServiceUpdateDelegatedResponse](../../pkg/models/operations/c1apiappv1connectorserviceupdatedelegatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
