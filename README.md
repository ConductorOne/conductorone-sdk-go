# conductorone-api

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/ConductorOne/conductorone-sdk-go
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
    res, err := s.AppEntitlementUserBinding.AppEntitlementUserBindingSvcListAppUsersForIdentityWithGrant(ctx, operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantRequest{
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

* [AppEntitlementUserBindingSvcListAppUsersForIdentityWithGrant](docs/sdks/appentitlementuserbinding/README.md#appentitlementuserbindingsvclistappusersforidentitywithgrant) - Invokes the c1.api.app.v1.AppEntitlementUserBindingService.ListAppUsersForIdentityWithGrant method.

### [AppEntitlements](docs/sdks/appentitlements/README.md)

* [AppEntitlementsGet](docs/sdks/appentitlements/README.md#appentitlementsget) - Invokes the c1.api.app.v1.AppEntitlements.Get method.

### [AppResource](docs/sdks/appresource/README.md)

* [AppResourceSvcGet](docs/sdks/appresource/README.md#appresourcesvcget) - Invokes the c1.api.app.v1.AppResourceService.Get method.

### [AppResourceType](docs/sdks/appresourcetype/README.md)

* [AppResourceTypeSvcGet](docs/sdks/appresourcetype/README.md#appresourcetypesvcget) - Invokes the c1.api.app.v1.AppResourceTypeService.Get method.

### [Apps](docs/sdks/apps/README.md)

* [AppsGet](docs/sdks/apps/README.md#appsget) - Invokes the c1.api.app.v1.Apps.Get method.

### [Auth](docs/sdks/auth/README.md)

* [AuthIntrospect](docs/sdks/auth/README.md#authintrospect) - Invokes the c1.api.auth.v1.Auth.Introspect method.

### [RequestCatalogSearch](docs/sdks/requestcatalogsearch/README.md)

* [RequestCatalogSearchSvcSearchEntitlements](docs/sdks/requestcatalogsearch/README.md#requestcatalogsearchsvcsearchentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

### [Task](docs/sdks/task/README.md)

* [TaskSvcCreateGrantTask](docs/sdks/task/README.md#tasksvccreategranttask) - Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.
* [TaskSvcCreateRevokeTask](docs/sdks/task/README.md#tasksvccreaterevoketask) - Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.
* [TaskSvcGet](docs/sdks/task/README.md#tasksvcget) - Invokes the c1.api.task.v1.TaskService.Get method.

### [TaskActions](docs/sdks/taskactions/README.md)

* [TaskActionsSvcApprove](docs/sdks/taskactions/README.md#taskactionssvcapprove) - Invokes the c1.api.task.v1.TaskActionsService.Approve method.
* [TaskActionsSvcComment](docs/sdks/taskactions/README.md#taskactionssvccomment) - Invokes the c1.api.task.v1.TaskActionsService.Comment method.
* [TaskActionsSvcDeny](docs/sdks/taskactions/README.md#taskactionssvcdeny) - Invokes the c1.api.task.v1.TaskActionsService.Deny method.

### [TaskSearch](docs/sdks/tasksearch/README.md)

* [TaskSearchSvcSearch](docs/sdks/tasksearch/README.md#tasksearchsvcsearch) - Invokes the c1.api.task.v1.TaskSearchService.Search method.

### [User](docs/sdks/user/README.md)

* [UserSvcGet](docs/sdks/user/README.md#usersvcget) - Invokes the c1.api.user.v1.UserService.Get method.
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
