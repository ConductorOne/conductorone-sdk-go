# conductorone-api

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/conductorone/conductorone-sdk-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.AppEntitlementUserBinding.ListAppUsersForIdentityWithGrant(ctx, operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantRequest{
        AppEntitlementID: "corrupti",
        AppID: "provident",
        IdentityUserID: "distinctio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppUsersForIdentityWithGrantResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AppEntitlementUserBinding](docs/sdks/appentitlementuserbinding/README.md)

* [ListAppUsersForIdentityWithGrant](docs/sdks/appentitlementuserbinding/README.md#listappusersforidentitywithgrant) - Invokes the c1.api.app.v1.AppEntitlementUserBindingService.ListAppUsersForIdentityWithGrant method.

### [AppEntitlements](docs/sdks/appentitlements/README.md)

* [Get](docs/sdks/appentitlements/README.md#get) - Invokes the c1.api.app.v1.AppEntitlements.Get method.

### [AppResource](docs/sdks/appresource/README.md)

* [Get](docs/sdks/appresource/README.md#get) - Invokes the c1.api.app.v1.AppResourceService.Get method.

### [AppResourceType](docs/sdks/appresourcetype/README.md)

* [Get](docs/sdks/appresourcetype/README.md#get) - Invokes the c1.api.app.v1.AppResourceTypeService.Get method.

### [Apps](docs/sdks/apps/README.md)

* [Get](docs/sdks/apps/README.md#get) - Invokes the c1.api.app.v1.Apps.Get method.

### [Auth](docs/sdks/auth/README.md)

* [Introspect](docs/sdks/auth/README.md#introspect) - Invokes the c1.api.auth.v1.Auth.Introspect method.

### [RequestCatalogSearch](docs/sdks/requestcatalogsearch/README.md)

* [SearchEntitlements](docs/sdks/requestcatalogsearch/README.md#searchentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

### [Task](docs/sdks/task/README.md)

* [CreateGrantTask](docs/sdks/task/README.md#creategranttask) - Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.
* [CreateRevokeTask](docs/sdks/task/README.md#createrevoketask) - Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.
* [Get](docs/sdks/task/README.md#get) - Invokes the c1.api.task.v1.TaskService.Get method.

### [TaskActions](docs/sdks/taskactions/README.md)

* [Approve](docs/sdks/taskactions/README.md#approve) - Invokes the c1.api.task.v1.TaskActionsService.Approve method.
* [Comment](docs/sdks/taskactions/README.md#comment) - Invokes the c1.api.task.v1.TaskActionsService.Comment method.
* [Deny](docs/sdks/taskactions/README.md#deny) - Invokes the c1.api.task.v1.TaskActionsService.Deny method.

### [TaskSearch](docs/sdks/tasksearch/README.md)

* [Search](docs/sdks/tasksearch/README.md#search) - Invokes the c1.api.task.v1.TaskSearchService.Search method.

### [User](docs/sdks/user/README.md)

* [Get](docs/sdks/user/README.md#get) - Invokes the c1.api.user.v1.UserService.Get method.
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
