# User
(*User*)

### Available Operations

* [Get](#get) - Get
* [List](#list) - List

## Get

Get a user by ID.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"context"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.User.Get(ctx, operations.C1APIUserV1UserServiceGetRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.C1APIUserV1UserServiceGetRequest](../../pkg/models/operations/c1apiuserv1userservicegetrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.C1APIUserV1UserServiceGetResponse](../../pkg/models/operations/c1apiuserv1userservicegetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

List users.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"context"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.User.List(ctx, operations.C1APIUserV1UserServiceListRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.UserServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.C1APIUserV1UserServiceListRequest](../../pkg/models/operations/c1apiuserv1userservicelistrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.C1APIUserV1UserServiceListResponse](../../pkg/models/operations/c1apiuserv1userservicelistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
