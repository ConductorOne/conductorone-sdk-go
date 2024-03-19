# Webhooks
(*Webhooks*)

### Available Operations

* [Test](#test) - Test

## Test

Invokes the c1.api.webhooks.v1.WebhooksService.Test method.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
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
    res, err := s.Webhooks.Test(ctx, operations.C1APIWebhooksV1WebhooksServiceTestRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhooksServiceTestResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIWebhooksV1WebhooksServiceTestRequest](../../pkg/models/operations/c1apiwebhooksv1webhooksservicetestrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.C1APIWebhooksV1WebhooksServiceTestResponse](../../pkg/models/operations/c1apiwebhooksv1webhooksservicetestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
