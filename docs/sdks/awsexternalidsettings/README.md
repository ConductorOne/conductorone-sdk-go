# AWSExternalIDSettings
(*AWSExternalIDSettings*)

### Available Operations

* [Get](#get) - Get

## Get

Invokes the c1.api.settings.v1.AWSExternalIDSettings.Get method.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"context"
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
    res, err := s.AWSExternalIDSettings.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAWSExternalIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APISettingsV1AWSExternalIDSettingsGetResponse](../../pkg/models/operations/c1apisettingsv1awsexternalidsettingsgetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |