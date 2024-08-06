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
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
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

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |


### Response

**[*operations.C1APISettingsV1AWSExternalIDSettingsGetResponse](../../pkg/models/operations/c1apisettingsv1awsexternalidsettingsgetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
