# PolicyValidate
(*PolicyValidate*)

### Available Operations

* [ValidateCEL](#validatecel) - Validate Cel

## ValidateCEL

Validate policies

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
    res, err := s.PolicyValidate.ValidateCEL(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidatePolicyCELResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.ValidatePolicyCELRequest](../../pkg/models/shared/validatepolicycelrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.C1APIPolicyV1PolicyValidateValidateCELResponse](../../pkg/models/operations/c1apipolicyv1policyvalidatevalidatecelresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
