# AppAccessRequestsDefaults
(*AppAccessRequestsDefaults*)

### Available Operations

* [CreateAppAccessRequestsDefaults](#createappaccessrequestsdefaults) - Create App Access Requests Defaults
* [GetAppAccessRequestsDefaults](#getappaccessrequestsdefaults) - Get App Access Requests Defaults

## CreateAppAccessRequestsDefaults

Invokes the c1.api.app.v1.AppAccessRequestsDefaultsService.CreateAppAccessRequestsDefaults method.

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
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.AppAccessRequestsDefaults.CreateAppAccessRequestsDefaults(ctx, operations.C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsRequest{
        AppID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AppAccessRequestDefaults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                      | Type                                                                                                                                                                                                           | Required                                                                                                                                                                                                       | Description                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                            |
| `request`                                                                                                                                                                                                      | [operations.C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsRequest](../../pkg/models/operations/c1apiappv1appaccessrequestsdefaultsservicecreateappaccessrequestsdefaultsrequest.md) | :heavy_check_mark:                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                     |


### Response

**[*operations.C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsResponse](../../pkg/models/operations/c1apiappv1appaccessrequestsdefaultsservicecreateappaccessrequestsdefaultsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAppAccessRequestsDefaults

Invokes the c1.api.app.v1.AppAccessRequestsDefaultsService.GetAppAccessRequestsDefaults method.

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
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.AppAccessRequestsDefaults.GetAppAccessRequestsDefaults(ctx, operations.C1APIAppV1AppAccessRequestsDefaultsServiceGetAppAccessRequestsDefaultsRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AppAccessRequestDefaults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                | Type                                                                                                                                                                                                     | Required                                                                                                                                                                                                 | Description                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                      |
| `request`                                                                                                                                                                                                | [operations.C1APIAppV1AppAccessRequestsDefaultsServiceGetAppAccessRequestsDefaultsRequest](../../pkg/models/operations/c1apiappv1appaccessrequestsdefaultsservicegetappaccessrequestsdefaultsrequest.md) | :heavy_check_mark:                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                               |


### Response

**[*operations.C1APIAppV1AppAccessRequestsDefaultsServiceGetAppAccessRequestsDefaultsResponse](../../pkg/models/operations/c1apiappv1appaccessrequestsdefaultsservicegetappaccessrequestsdefaultsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
