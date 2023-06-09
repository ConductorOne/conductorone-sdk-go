# Auth

### Available Operations

* [C1APIAuthV1AuthIntrospect](#c1apiauthv1authintrospect) - Invokes the c1.api.auth.v1.Auth.Introspect method.

## C1APIAuthV1AuthIntrospect

Invokes the c1.api.auth.v1.Auth.Introspect method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Auth.C1APIAuthV1AuthIntrospect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAuthV1IntrospectResponse != nil {
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

