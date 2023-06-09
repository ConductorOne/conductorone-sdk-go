# App

### Available Operations

* [C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant](#c1apiappv1appentitlementuserbindingservicelistappusersforidentitywithgrant) - Invokes the c1.api.app.v1.AppEntitlementUserBindingService.ListAppUsersForIdentityWithGrant method.
* [C1APIAppV1AppEntitlementsGet](#c1apiappv1appentitlementsget) - Invokes the c1.api.app.v1.AppEntitlements.Get method.
* [C1APIAppV1AppResourceServiceGet](#c1apiappv1appresourceserviceget) - Invokes the c1.api.app.v1.AppResourceService.Get method.
* [C1APIAppV1AppResourceTypeServiceGet](#c1apiappv1appresourcetypeserviceget) - Invokes the c1.api.app.v1.AppResourceTypeService.Get method.
* [C1APIAppV1AppsGet](#c1apiappv1appsget) - Invokes the c1.api.app.v1.Apps.Get method.

## C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant

Invokes the c1.api.app.v1.AppEntitlementUserBindingService.ListAppUsersForIdentityWithGrant method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.App.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant(ctx, operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantRequest{
        AppEntitlementID: "quibusdam",
        AppID: "unde",
        IdentityUserID: "nulla",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1ListAppUsersForIdentityWithGrantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                    | Type                                                                                                                                                                                                         | Required                                                                                                                                                                                                     | Description                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                          |
| `request`                                                                                                                                                                                                    | [operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantRequest](../../models/operations/c1apiappv1appentitlementuserbindingservicelistappusersforidentitywithgrantrequest.md) | :heavy_check_mark:                                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                                   |


### Response

**[*operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantResponse](../../models/operations/c1apiappv1appentitlementuserbindingservicelistappusersforidentitywithgrantresponse.md), error**


## C1APIAppV1AppEntitlementsGet

Invokes the c1.api.app.v1.AppEntitlements.Get method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.App.C1APIAppV1AppEntitlementsGet(ctx, operations.C1APIAppV1AppEntitlementsGetRequest{
        AppID: "corrupti",
        ID: "d69a674e-0f46-47cc-8796-ed151a05dfc2",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1GetAppEntitlementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.C1APIAppV1AppEntitlementsGetRequest](../../models/operations/c1apiappv1appentitlementsgetrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.C1APIAppV1AppEntitlementsGetResponse](../../models/operations/c1apiappv1appentitlementsgetresponse.md), error**


## C1APIAppV1AppResourceServiceGet

Invokes the c1.api.app.v1.AppResourceService.Get method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.App.C1APIAppV1AppResourceServiceGet(ctx, operations.C1APIAppV1AppResourceServiceGetRequest{
        AppID: "at",
        AppResourceTypeID: "at",
        ID: "f7cc78ca-1ba9-428f-8816-742cb7392059",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1AppResourceServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.C1APIAppV1AppResourceServiceGetRequest](../../models/operations/c1apiappv1appresourceservicegetrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.C1APIAppV1AppResourceServiceGetResponse](../../models/operations/c1apiappv1appresourceservicegetresponse.md), error**


## C1APIAppV1AppResourceTypeServiceGet

Invokes the c1.api.app.v1.AppResourceTypeService.Get method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.App.C1APIAppV1AppResourceTypeServiceGet(ctx, operations.C1APIAppV1AppResourceTypeServiceGetRequest{
        AppID: "sed",
        ID: "9396fea7-596e-4b10-baaa-2352c5955907",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1AppResourceTypeServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APIAppV1AppResourceTypeServiceGetRequest](../../models/operations/c1apiappv1appresourcetypeservicegetrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.C1APIAppV1AppResourceTypeServiceGetResponse](../../models/operations/c1apiappv1appresourcetypeservicegetresponse.md), error**


## C1APIAppV1AppsGet

Invokes the c1.api.app.v1.Apps.Get method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.App.C1APIAppV1AppsGet(ctx, operations.C1APIAppV1AppsGetRequest{
        ID: "aff1a3a2-fa94-4677-b925-1aa52c3f5ad0",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1GetAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.C1APIAppV1AppsGetRequest](../../models/operations/c1apiappv1appsgetrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.C1APIAppV1AppsGetResponse](../../models/operations/c1apiappv1appsgetresponse.md), error**

