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
	"conductorone-api"
	"conductorone-api/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.App.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant(ctx, operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantRequest{
        AppEntitlementID: "corrupti",
        AppID: "provident",
        IdentityUserID: "distinctio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1ListAppUsersForIdentityWithGrantResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [App](docs/sdks/app/README.md)

* [C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant](docs/sdks/app/README.md#c1apiappv1appentitlementuserbindingservicelistappusersforidentitywithgrant) - Invokes the c1.api.app.v1.AppEntitlementUserBindingService.ListAppUsersForIdentityWithGrant method.
* [C1APIAppV1AppEntitlementsGet](docs/sdks/app/README.md#c1apiappv1appentitlementsget) - Invokes the c1.api.app.v1.AppEntitlements.Get method.
* [C1APIAppV1AppResourceServiceGet](docs/sdks/app/README.md#c1apiappv1appresourceserviceget) - Invokes the c1.api.app.v1.AppResourceService.Get method.
* [C1APIAppV1AppResourceTypeServiceGet](docs/sdks/app/README.md#c1apiappv1appresourcetypeserviceget) - Invokes the c1.api.app.v1.AppResourceTypeService.Get method.
* [C1APIAppV1AppsGet](docs/sdks/app/README.md#c1apiappv1appsget) - Invokes the c1.api.app.v1.Apps.Get method.

### [Auth](docs/sdks/auth/README.md)

* [C1APIAuthV1AuthIntrospect](docs/sdks/auth/README.md#c1apiauthv1authintrospect) - Invokes the c1.api.auth.v1.Auth.Introspect method.

### [Requestcatalog](docs/sdks/requestcatalog/README.md)

* [C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements](docs/sdks/requestcatalog/README.md#c1apirequestcatalogv1requestcatalogsearchservicesearchentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

### [Task](docs/sdks/task/README.md)

* [C1APITaskV1TaskActionsServiceApprove](docs/sdks/task/README.md#c1apitaskv1taskactionsserviceapprove) - Invokes the c1.api.task.v1.TaskActionsService.Approve method.
* [C1APITaskV1TaskActionsServiceComment](docs/sdks/task/README.md#c1apitaskv1taskactionsservicecomment) - Invokes the c1.api.task.v1.TaskActionsService.Comment method.
* [C1APITaskV1TaskActionsServiceDeny](docs/sdks/task/README.md#c1apitaskv1taskactionsservicedeny) - Invokes the c1.api.task.v1.TaskActionsService.Deny method.
* [C1APITaskV1TaskSearchServiceSearch](docs/sdks/task/README.md#c1apitaskv1tasksearchservicesearch) - Invokes the c1.api.task.v1.TaskSearchService.Search method.
* [C1APITaskV1TaskServiceCreateGrantTask](docs/sdks/task/README.md#c1apitaskv1taskservicecreategranttask) - Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.
* [C1APITaskV1TaskServiceCreateRevokeTask](docs/sdks/task/README.md#c1apitaskv1taskservicecreaterevoketask) - Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.
* [C1APITaskV1TaskServiceGet](docs/sdks/task/README.md#c1apitaskv1taskserviceget) - Invokes the c1.api.task.v1.TaskService.Get method.

### [User](docs/sdks/user/README.md)

* [C1APIUserV1UserServiceGet](docs/sdks/user/README.md#c1apiuserv1userserviceget) - Invokes the c1.api.user.v1.UserService.Get method.
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
