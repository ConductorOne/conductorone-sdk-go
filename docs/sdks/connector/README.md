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
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.Create(ctx, operations.C1APIAppV1ConnectorServiceCreateRequest{
        ConnectorServiceCreateRequest: &shared.ConnectorServiceCreateRequest{
            ConnectorExpandMask: &shared.ConnectorExpandMask{
                Paths: []string{
                    "temporibus",
                    "qui",
                    "neque",
                },
            },
            CatalogID: conductoroneapi.String("fugit"),
            Config: map[string]interface{}{
                "odio": "sunt",
            },
            Description: conductoroneapi.String("ullam"),
            UserIds: []string{
                "hic",
                "voluptatem",
                "cumque",
            },
        },
        AppID: "soluta",
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
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.CreateDelegated(ctx, operations.C1APIAppV1ConnectorServiceCreateDelegatedRequest{
        ConnectorServiceCreateDelegatedRequest: &shared.ConnectorServiceCreateDelegatedRequest{
            ConnectorExpandMask: &shared.ConnectorExpandMask{
                Paths: []string{
                    "et",
                    "saepe",
                    "ipsum",
                },
            },
            CatalogID: conductoroneapi.String("veritatis"),
            Description: conductoroneapi.String("nobis"),
            DisplayName: conductoroneapi.String("quos"),
            UserIds: []string{
                "cupiditate",
                "aperiam",
                "delectus",
            },
        },
        AppID: "dolorem",
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
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.Delete(ctx, operations.C1APIAppV1ConnectorServiceDeleteRequest{
        ConnectorServiceDeleteRequest: &shared.ConnectorServiceDeleteRequest{},
        AppID: "dolore",
        ID: "43a1108e-0adc-4f4b-9218-79fce953f73e",
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
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.Get(ctx, operations.C1APIAppV1ConnectorServiceGetRequest{
        AppID: "tenetur",
        ID: "7fbc7abd-74dd-439c-8f5d-2cff7c70a456",
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
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.GetCredentials(ctx, operations.C1APIAppV1ConnectorServiceGetCredentialsRequest{
        AppID: "aspernatur",
        ConnectorID: "vel",
        ID: "d436813f-16d9-4f5f-8e6c-556146c3e250",
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
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.List(ctx, operations.C1APIAppV1ConnectorServiceListRequest{
        AppID: "a",
        PageSize: conductoroneapi.Float64(7255.95),
        PageToken: conductoroneapi.String("aut"),
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
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.RevokeCredential(ctx, operations.C1APIAppV1ConnectorServiceRevokeCredentialRequest{
        ConnectorServiceRevokeCredentialRequest: &shared.ConnectorServiceRevokeCredentialRequest{},
        AppID: "aut",
        ConnectorID: "deleniti",
        ID: "c42e141a-ac36-46c8-9d6b-144290747477",
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
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.RotateCredential(ctx, operations.C1APIAppV1ConnectorServiceRotateCredentialRequest{
        ConnectorServiceRotateCredentialRequest: &shared.ConnectorServiceRotateCredentialRequest{},
        AppID: "rem",
        ConnectorID: "fuga",
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
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.Update(ctx, operations.C1APIAppV1ConnectorServiceUpdateRequest{
        ConnectorServiceUpdateRequestInput: &shared.ConnectorServiceUpdateRequestInput{
            Connector: &shared.ConnectorInput{
                ConnectorStatus: &shared.ConnectorStatus{
                    CompletedAt: types.MustTimeFromString("2022-04-22T02:19:15.622Z"),
                    LastError: conductoroneapi.String("fugiat"),
                    StartedAt: types.MustTimeFromString("2022-07-26T19:04:27.877Z"),
                    Status: shared.ConnectorStatusStatusSyncStatusRunning.ToPointer(),
                    UpdatedAt: types.MustTimeFromString("2022-06-16T15:21:08.890Z"),
                },
                OAuth2AuthorizedAs: &shared.OAuth2AuthorizedAsInput{},
                AppID: conductoroneapi.String("praesentium"),
                Config: map[string]interface{}{
                    "veritatis": "ipsa",
                    "id": "quidem",
                    "neque": "quo",
                    "illum": "quo",
                },
                Description: conductoroneapi.String("fuga"),
                DisplayName: conductoroneapi.String("eius"),
                ID: conductoroneapi.String("251904e5-23c7-4e0b-8717-8e4796f2a70c"),
                UserIds: []string{
                    "quas",
                    "praesentium",
                },
            },
            ConnectorExpandMask: &shared.ConnectorExpandMask{
                Paths: []string{
                    "deleniti",
                },
            },
            UpdateMask: conductoroneapi.String("fugit"),
        },
        AppID: "fuga",
        ID: "a482562f-222e-4981-bee1-7cbe61e6b7b9",
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
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connector.UpdateDelegated(ctx, operations.C1APIAppV1ConnectorServiceUpdateDelegatedRequest{
        ConnectorServiceUpdateDelegatedRequestInput: &shared.ConnectorServiceUpdateDelegatedRequestInput{
            Connector: &shared.ConnectorInput{
                ConnectorStatus: &shared.ConnectorStatus{
                    CompletedAt: types.MustTimeFromString("2022-04-14T13:42:03.513Z"),
                    LastError: conductoroneapi.String("eligendi"),
                    StartedAt: types.MustTimeFromString("2022-05-13T20:06:18.158Z"),
                    Status: shared.ConnectorStatusStatusSyncStatusDone.ToPointer(),
                    UpdatedAt: types.MustTimeFromString("2022-03-27T01:23:58.638Z"),
                },
                OAuth2AuthorizedAs: &shared.OAuth2AuthorizedAsInput{},
                AppID: conductoroneapi.String("consequuntur"),
                Config: map[string]interface{}{
                    "minus": "quaerat",
                },
                Description: conductoroneapi.String("sapiente"),
                DisplayName: conductoroneapi.String("consectetur"),
                ID: conductoroneapi.String("789fd871-f99d-4d2e-bd12-1aa6f1e674bd"),
                UserIds: []string{
                    "accusantium",
                    "aliquam",
                    "sapiente",
                },
            },
            ConnectorExpandMask: &shared.ConnectorExpandMask{
                Paths: []string{
                    "ullam",
                },
            },
            UpdateMask: conductoroneapi.String("reprehenderit"),
        },
        ConnectorAppID: "ullam",
        ConnectorID: "nisi",
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

