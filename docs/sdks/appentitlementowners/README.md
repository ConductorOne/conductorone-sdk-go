# AppEntitlementOwners
(*AppEntitlementOwners*)

### Available Operations

* [Add](#add) - Add
* [List](#list) - List
* [Remove](#remove) - Remove
* [Set](#set) - Set

## Add

Add an owner to a given app entitlement.

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
    res, err := s.AppEntitlementOwners.Add(ctx, operations.C1APIAppV1AppEntitlementOwnersAddRequest{
        AppID: "<value>",
        EntitlementID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddAppEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APIAppV1AppEntitlementOwnersAddRequest](../../pkg/models/operations/c1apiappv1appentitlementownersaddrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.C1APIAppV1AppEntitlementOwnersAddResponse](../../pkg/models/operations/c1apiappv1appentitlementownersaddresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

List owners for a given app entitlement.

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
    res, err := s.AppEntitlementOwners.List(ctx, operations.C1APIAppV1AppEntitlementOwnersListRequest{
        AppID: "<value>",
        EntitlementID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIAppV1AppEntitlementOwnersListRequest](../../pkg/models/operations/c1apiappv1appentitlementownerslistrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.C1APIAppV1AppEntitlementOwnersListResponse](../../pkg/models/operations/c1apiappv1appentitlementownerslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Remove

Remove an owner from a given app entitlement.

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
    res, err := s.AppEntitlementOwners.Remove(ctx, operations.C1APIAppV1AppEntitlementOwnersRemoveRequest{
        AppID: "<value>",
        EntitlementID: "<value>",
        UserID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RemoveAppEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.C1APIAppV1AppEntitlementOwnersRemoveRequest](../../pkg/models/operations/c1apiappv1appentitlementownersremoverequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.C1APIAppV1AppEntitlementOwnersRemoveResponse](../../pkg/models/operations/c1apiappv1appentitlementownersremoveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Set

Sets the owners for a given app entitlement to the specified list of users.

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
    res, err := s.AppEntitlementOwners.Set(ctx, operations.C1APIAppV1AppEntitlementOwnersSetRequest{
        AppID: "<value>",
        EntitlementID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SetAppEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APIAppV1AppEntitlementOwnersSetRequest](../../pkg/models/operations/c1apiappv1appentitlementownerssetrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.C1APIAppV1AppEntitlementOwnersSetResponse](../../pkg/models/operations/c1apiappv1appentitlementownerssetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
