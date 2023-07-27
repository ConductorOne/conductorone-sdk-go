# Connector

### Available Operations

* [Create](#create) - Create
* [CreateDelegated](#createdelegated) - Create Delegated
* [Delete](#delete) - Delete
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
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.Create(ctx, operations.C1APIAppV1ConnectorServiceCreateRequest{
        ConnectorServiceCreateRequest: &shared.ConnectorServiceCreateRequest{
            ConnectorExpandMask: &shared.ConnectorExpandMask{
                Paths: []string{
                    "saepe",
                    "eius",
                    "aspernatur",
                },
            },
            CatalogID: conductoroneapi.String("perferendis"),
            Config: map[string]interface{}{
                "optio": "accusamus",
            },
            Description: conductoroneapi.String("ad"),
            UserIds: []string{
                "suscipit",
                "deserunt",
                "provident",
                "minima",
            },
        },
        AppID: "repellendus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.C1APIAppV1ConnectorServiceCreateRequest](../../models/operations/c1apiappv1connectorservicecreaterequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.C1APIAppV1ConnectorServiceCreateResponse](../../models/operations/c1apiappv1connectorservicecreateresponse.md), error**


## CreateDelegated

Create a connector that is pending a connector config.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.CreateDelegated(ctx, operations.C1APIAppV1ConnectorServiceCreateDelegatedRequest{
        ConnectorServiceCreateDelegatedRequest: &shared.ConnectorServiceCreateDelegatedRequest{
            ConnectorExpandMask: &shared.ConnectorExpandMask{
                Paths: []string{
                    "similique",
                    "alias",
                    "at",
                },
            },
            CatalogID: conductoroneapi.String("quaerat"),
            Description: conductoroneapi.String("tempora"),
            DisplayName: conductoroneapi.String("vel"),
            UserIds: []string{
                "officiis",
                "qui",
                "dolorum",
                "a",
            },
        },
        AppID: "esse",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.C1APIAppV1ConnectorServiceCreateDelegatedRequest](../../models/operations/c1apiappv1connectorservicecreatedelegatedrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.C1APIAppV1ConnectorServiceCreateDelegatedResponse](../../models/operations/c1apiappv1connectorservicecreatedelegatedresponse.md), error**


## Delete

Delete a connector.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.Delete(ctx, operations.C1APIAppV1ConnectorServiceDeleteRequest{
        ConnectorServiceDeleteRequest: &shared.ConnectorServiceDeleteRequest{},
        AppID: "harum",
        ID: "73cf3be4-53f8-470b-b26b-5a73429cdb1a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.C1APIAppV1ConnectorServiceDeleteRequest](../../models/operations/c1apiappv1connectorservicedeleterequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.C1APIAppV1ConnectorServiceDeleteResponse](../../models/operations/c1apiappv1connectorservicedeleteresponse.md), error**


## Get

Get a connector.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.Get(ctx, operations.C1APIAppV1ConnectorServiceGetRequest{
        AppID: "totam",
        ID: "422bb679-d232-4271-9bf0-cbb1e31b8b90",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.C1APIAppV1ConnectorServiceGetRequest](../../models/operations/c1apiappv1connectorservicegetrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.C1APIAppV1ConnectorServiceGetResponse](../../models/operations/c1apiappv1connectorservicegetresponse.md), error**


## GetCredentials

Get credentials for a connector.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.GetCredentials(ctx, operations.C1APIAppV1ConnectorServiceGetCredentialsRequest{
        AppID: "delectus",
        ConnectorID: "dolorem",
        ID: "443a1108-e0ad-4cf4-b921-879fce953f73",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorServiceGetCredentialsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.C1APIAppV1ConnectorServiceGetCredentialsRequest](../../models/operations/c1apiappv1connectorservicegetcredentialsrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.C1APIAppV1ConnectorServiceGetCredentialsResponse](../../models/operations/c1apiappv1connectorservicegetcredentialsresponse.md), error**


## List

List connectors for an app.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.List(ctx, operations.C1APIAppV1ConnectorServiceListRequest{
        AppID: "vero",
        PageSize: conductoroneapi.Float64(9493.19),
        PageToken: conductoroneapi.String("dignissimos"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.C1APIAppV1ConnectorServiceListRequest](../../models/operations/c1apiappv1connectorservicelistrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.C1APIAppV1ConnectorServiceListResponse](../../models/operations/c1apiappv1connectorservicelistresponse.md), error**


## RevokeCredential

Revoke credentials for a connector.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.RevokeCredential(ctx, operations.C1APIAppV1ConnectorServiceRevokeCredentialRequest{
        ConnectorServiceRevokeCredentialRequest: &shared.ConnectorServiceRevokeCredentialRequest{},
        AppID: "hic",
        ConnectorID: "distinctio",
        ID: "c7abd74d-d39c-40f5-92cf-f7c70a45626d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorServiceRevokeCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIAppV1ConnectorServiceRevokeCredentialRequest](../../models/operations/c1apiappv1connectorservicerevokecredentialrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.C1APIAppV1ConnectorServiceRevokeCredentialResponse](../../models/operations/c1apiappv1connectorservicerevokecredentialresponse.md), error**


## RotateCredential

Rotate credentials for a connector.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.RotateCredential(ctx, operations.C1APIAppV1ConnectorServiceRotateCredentialRequest{
        ConnectorServiceRotateCredentialRequest: &shared.ConnectorServiceRotateCredentialRequest{},
        AppID: "magnam",
        ConnectorID: "ratione",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorServiceRotateCredentialResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIAppV1ConnectorServiceRotateCredentialRequest](../../models/operations/c1apiappv1connectorservicerotatecredentialrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.C1APIAppV1ConnectorServiceRotateCredentialResponse](../../models/operations/c1apiappv1connectorservicerotatecredentialresponse.md), error**


## Update

Update a connector.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/types"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.Update(ctx, operations.C1APIAppV1ConnectorServiceUpdateRequest{
        ConnectorServiceUpdateRequestInput: &shared.ConnectorServiceUpdateRequestInput{
            Connector: &shared.ConnectorInput{
                ConnectorStatus: &shared.ConnectorStatus{
                    CompletedAt: types.MustTimeFromString("2022-06-28T08:50:44.084Z"),
                    LastError: conductoroneapi.String("dicta"),
                    StartedAt: types.MustTimeFromString("2022-01-08T01:04:15.076Z"),
                    Status: shared.ConnectorStatusStatusSyncStatusUnspecified.ToPointer(),
                    UpdatedAt: types.MustTimeFromString("2022-02-20T07:12:08.273Z"),
                },
                OAuth2AuthorizedAs: &shared.OAuth2AuthorizedAs1{},
                AppID: conductoroneapi.String("excepturi"),
                Config: map[string]interface{}{
                    "nostrum": "sapiente",
                    "quisquam": "saepe",
                    "ea": "impedit",
                    "corporis": "veniam",
                },
                Description: conductoroneapi.String("aliquid"),
                DisplayName: conductoroneapi.String("inventore"),
                ID: conductoroneapi.String("46c3e250-fb00-48c4-ae14-1aac366c8dd6"),
                UserIds: []string{
                    "quasi",
                    "tempora",
                    "numquam",
                },
            },
            ConnectorExpandMask: &shared.ConnectorExpandMask{
                Paths: []string{
                    "provident",
                },
            },
            UpdateMask: conductoroneapi.String("ipsa"),
        },
        AppID: "molestiae",
        ID: "474778a7-bd46-46d2-8c10-ab3cdca42519",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.C1APIAppV1ConnectorServiceUpdateRequest](../../models/operations/c1apiappv1connectorserviceupdaterequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.C1APIAppV1ConnectorServiceUpdateResponse](../../models/operations/c1apiappv1connectorserviceupdateresponse.md), error**


## UpdateDelegated

Update a delegated connector.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/types"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.UpdateDelegated(ctx, operations.C1APIAppV1ConnectorServiceUpdateDelegatedRequest{
        ConnectorServiceUpdateDelegatedRequestInput: &shared.ConnectorServiceUpdateDelegatedRequestInput{
            Connector: &shared.ConnectorInput{
                ConnectorStatus: &shared.ConnectorStatus{
                    CompletedAt: types.MustTimeFromString("2022-09-23T10:04:47.931Z"),
                    LastError: conductoroneapi.String("debitis"),
                    StartedAt: types.MustTimeFromString("2022-11-13T06:50:40.250Z"),
                    Status: shared.ConnectorStatusStatusSyncStatusUnspecified.ToPointer(),
                    UpdatedAt: types.MustTimeFromString("2021-08-15T10:59:14.485Z"),
                },
                OAuth2AuthorizedAs: &shared.OAuth2AuthorizedAs1{},
                AppID: conductoroneapi.String("recusandae"),
                Config: map[string]interface{}{
                    "distinctio": "quod",
                },
                Description: conductoroneapi.String("dignissimos"),
                DisplayName: conductoroneapi.String("inventore"),
                ID: conductoroneapi.String("78e4796f-2a70-4c68-8282-aa482562f222"),
                UserIds: []string{
                    "occaecati",
                    "atque",
                    "et",
                    "esse",
                },
            },
            ConnectorExpandMask: &shared.ConnectorExpandMask{
                Paths: []string{
                    "accusamus",
                    "veritatis",
                    "esse",
                    "quod",
                },
            },
            UpdateMask: conductoroneapi.String("nam"),
        },
        ConnectorAppID: "vero",
        ConnectorID: "aliquid",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.C1APIAppV1ConnectorServiceUpdateDelegatedRequest](../../models/operations/c1apiappv1connectorserviceupdatedelegatedrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.C1APIAppV1ConnectorServiceUpdateDelegatedResponse](../../models/operations/c1apiappv1connectorserviceupdatedelegatedresponse.md), error**

