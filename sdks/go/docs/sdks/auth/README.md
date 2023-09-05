# Auth

### Available Operations

* [Introspect](#introspect) - Introspect

## Introspect

Introspect returns the current user's principle_id, user_id and a list of roles, permissions, and enabled features.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Auth.Introspect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.IntrospectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APIAuthV1AuthIntrospectResponse](../../models/operations/c1apiauthv1authintrospectresponse.md), error**

