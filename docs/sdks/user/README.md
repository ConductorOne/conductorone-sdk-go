# User

### Available Operations

* [C1APIUserV1UserServiceGet](#c1apiuserv1userserviceget) - Invokes the c1.api.user.v1.UserService.Get method.

## C1APIUserV1UserServiceGet

Invokes the c1.api.user.v1.UserService.Get method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.User.C1APIUserV1UserServiceGet(ctx, operations.C1APIUserV1UserServiceGetRequest{
        ID: "adcf4b92-1879-4fce-953f-73ef7fbc7abd",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIUserV1UserServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.C1APIUserV1UserServiceGetRequest](../../models/operations/c1apiuserv1userservicegetrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.C1APIUserV1UserServiceGetResponse](../../models/operations/c1apiuserv1userservicegetresponse.md), error**

