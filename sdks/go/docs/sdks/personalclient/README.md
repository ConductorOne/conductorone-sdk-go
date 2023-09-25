# PersonalClient

### Available Operations

* [Create](#create) - Create

## Create

Create creates a new PersonalClient object for the current User.

### Example Usage

```go
package main

import(
	"context"
	"log"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PersonalClient.Create(ctx, shared.PersonalClientServiceCreateRequest{
        AllowSourceCidr: []string{
            "sapiente",
        },
        DisplayName: conductoronesdkgo.String("consequuntur"),
        Expires: conductoronesdkgo.String("ratione"),
        ScopedRoles: []string{
            "explicabo",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PersonalClientServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [shared.PersonalClientServiceCreateRequest](../../models/shared/personalclientservicecreaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.C1APIIamV1PersonalClientServiceCreateResponse](../../models/operations/c1apiiamv1personalclientservicecreateresponse.md), error**

