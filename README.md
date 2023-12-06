# conductorone-api

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/ConductorOne/conductorone-sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"log"
)

func main() {
	s := conductoronesdkgo.New(
		conductoronesdkgo.WithSecurity(shared.Security{
			BearerAuth: "",
			Oauth:      "",
		}),
	)

	ctx := context.Background()
	res, err := s.Apps.Create(ctx, &shared.CreateAppRequest{
		Owners: []string{
			"string",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateAppResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Apps](docs/sdks/apps/README.md)

* [Create](docs/sdks/apps/README.md#create) - Create
* [Delete](docs/sdks/apps/README.md#delete) - Delete
* [Get](docs/sdks/apps/README.md#get) - Get
* [List](docs/sdks/apps/README.md#list) - List
* [Update](docs/sdks/apps/README.md#update) - Update

### [Connector](docs/sdks/connector/README.md)

* [Create](docs/sdks/connector/README.md#create) - Create
* [CreateDelegated](docs/sdks/connector/README.md#createdelegated) - Create Delegated
* [Delete](docs/sdks/connector/README.md#delete) - Delete
* [Get](docs/sdks/connector/README.md#get) - Get
* [GetCredentials](docs/sdks/connector/README.md#getcredentials) - Get Credentials
* [List](docs/sdks/connector/README.md#list) - List
* [RevokeCredential](docs/sdks/connector/README.md#revokecredential) - Revoke Credential
* [RotateCredential](docs/sdks/connector/README.md#rotatecredential) - Rotate Credential
* [Update](docs/sdks/connector/README.md#update) - Update
* [UpdateDelegated](docs/sdks/connector/README.md#updatedelegated) - Update Delegated

### [AppEntitlements](docs/sdks/appentitlements/README.md)

* [Get](docs/sdks/appentitlements/README.md#get) - Get
* [List](docs/sdks/appentitlements/README.md#list) - List
* [ListForAppResource](docs/sdks/appentitlements/README.md#listforappresource) - List For App Resource
* [ListForAppUser](docs/sdks/appentitlements/README.md#listforappuser) - List For App User
* [ListUsers](docs/sdks/appentitlements/README.md#listusers) - List Users
* [Update](docs/sdks/appentitlements/README.md#update) - Update

### [AppEntitlementUserBinding](docs/sdks/appentitlementuserbinding/README.md)

* [ListAppUsersForIdentityWithGrant](docs/sdks/appentitlementuserbinding/README.md#listappusersforidentitywithgrant) - List App Users For Identity With Grant

### [AppEntitlementOwners](docs/sdks/appentitlementowners/README.md)

* [Add](docs/sdks/appentitlementowners/README.md#add) - Add
* [List](docs/sdks/appentitlementowners/README.md#list) - List
* [Remove](docs/sdks/appentitlementowners/README.md#remove) - Remove
* [Set](docs/sdks/appentitlementowners/README.md#set) - Set

### [AppOwners](docs/sdks/appowners/README.md)

* [Add](docs/sdks/appowners/README.md#add) - Add
* [List](docs/sdks/appowners/README.md#list) - List
* [Remove](docs/sdks/appowners/README.md#remove) - Remove
* [Set](docs/sdks/appowners/README.md#set) - Set

### [AppReport](docs/sdks/appreport/README.md)

* [List](docs/sdks/appreport/README.md#list) - List

### [AppReportAction](docs/sdks/appreportaction/README.md)

* [GenerateReport](docs/sdks/appreportaction/README.md#generatereport) - Generate Report

### [AppResourceType](docs/sdks/appresourcetype/README.md)

* [Get](docs/sdks/appresourcetype/README.md#get) - Get
* [List](docs/sdks/appresourcetype/README.md#list) - List

### [AppResource](docs/sdks/appresource/README.md)

* [Get](docs/sdks/appresource/README.md#get) - Get
* [List](docs/sdks/appresource/README.md#list) - List

### [AppResourceOwners](docs/sdks/appresourceowners/README.md)

* [List](docs/sdks/appresourceowners/README.md#list) - List

### [AppUsageControls](docs/sdks/appusagecontrols/README.md)

* [Get](docs/sdks/appusagecontrols/README.md#get) - Get
* [Update](docs/sdks/appusagecontrols/README.md#update) - Update

### [AppUser](docs/sdks/appuser/README.md)

* [Update](docs/sdks/appuser/README.md#update) - Update

### [Attributes](docs/sdks/attributes/README.md)

* [CreateAttributeValue](docs/sdks/attributes/README.md#createattributevalue) - Create Attribute Value
* [DeleteAttributeValue](docs/sdks/attributes/README.md#deleteattributevalue) - Delete Attribute Value
* [GetAttributeValue](docs/sdks/attributes/README.md#getattributevalue) - Get Attribute Value
* [ListAttributeTypes](docs/sdks/attributes/README.md#listattributetypes) - List Attribute Types
* [ListAttributeValues](docs/sdks/attributes/README.md#listattributevalues) - List Attribute Values

### [Auth](docs/sdks/auth/README.md)

* [Introspect](docs/sdks/auth/README.md#introspect) - Introspect

### [RequestCatalogManagement](docs/sdks/requestcatalogmanagement/README.md)

* [AddAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#addaccessentitlements) - Add Access Entitlements
* [AddAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#addappentitlements) - Add App Entitlements
* [Create](docs/sdks/requestcatalogmanagement/README.md#create) - Create
* [Delete](docs/sdks/requestcatalogmanagement/README.md#delete) - Delete
* [Get](docs/sdks/requestcatalogmanagement/README.md#get) - Get
* [List](docs/sdks/requestcatalogmanagement/README.md#list) - List
* [ListEntitlementsForAccess](docs/sdks/requestcatalogmanagement/README.md#listentitlementsforaccess) - List Entitlements For Access
* [ListEntitlementsPerCatalog](docs/sdks/requestcatalogmanagement/README.md#listentitlementspercatalog) - List Entitlements Per Catalog
* [RemoveAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeaccessentitlements) - Remove Access Entitlements
* [RemoveAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeappentitlements) - Remove App Entitlements
* [Update](docs/sdks/requestcatalogmanagement/README.md#update) - Update

### [Directory](docs/sdks/directory/README.md)

* [Create](docs/sdks/directory/README.md#create) - Create
* [Delete](docs/sdks/directory/README.md#delete) - Delete
* [Get](docs/sdks/directory/README.md#get) - Get
* [List](docs/sdks/directory/README.md#list) - List

### [PersonalClient](docs/sdks/personalclient/README.md)

* [Create](docs/sdks/personalclient/README.md#create) - Create

### [Roles](docs/sdks/roles/README.md)

* [Get](docs/sdks/roles/README.md#get) - Get
* [List](docs/sdks/roles/README.md#list) - List
* [Update](docs/sdks/roles/README.md#update) - Update

### [Policies](docs/sdks/policies/README.md)

* [Create](docs/sdks/policies/README.md#create) - Create
* [Delete](docs/sdks/policies/README.md#delete) - Delete
* [Get](docs/sdks/policies/README.md#get) - Get
* [List](docs/sdks/policies/README.md#list) - List
* [Update](docs/sdks/policies/README.md#update) - Update

### [PolicyValidate](docs/sdks/policyvalidate/README.md)

* [ValidateCEL](docs/sdks/policyvalidate/README.md#validatecel) - Validate Cel

### [AppResourceSearch](docs/sdks/appresourcesearch/README.md)

* [SearchAppResourceTypes](docs/sdks/appresourcesearch/README.md#searchappresourcetypes) - Search App Resource Types

### [AppSearch](docs/sdks/appsearch/README.md)

* [Search](docs/sdks/appsearch/README.md#search) - Search

### [AttributeSearch](docs/sdks/attributesearch/README.md)

* [SearchAttributeValues](docs/sdks/attributesearch/README.md#searchattributevalues) - Search Attribute Values

### [AppEntitlementSearch](docs/sdks/appentitlementsearch/README.md)

* [Search](docs/sdks/appentitlementsearch/README.md#search) - Search

### [PolicySearch](docs/sdks/policysearch/README.md)

* [Search](docs/sdks/policysearch/README.md#search) - Search

### [RequestCatalogSearch](docs/sdks/requestcatalogsearch/README.md)

* [SearchEntitlements](docs/sdks/requestcatalogsearch/README.md#searchentitlements) - Search Entitlements

### [TaskSearch](docs/sdks/tasksearch/README.md)

* [Search](docs/sdks/tasksearch/README.md#search) - Search

### [UserSearch](docs/sdks/usersearch/README.md)

* [Search](docs/sdks/usersearch/README.md#search) - Search

### [Task](docs/sdks/task/README.md)

* [CreateGrantTask](docs/sdks/task/README.md#creategranttask) - Create Grant Task
* [CreateRevokeTask](docs/sdks/task/README.md#createrevoketask) - Create Revoke Task
* [Get](docs/sdks/task/README.md#get) - Get

### [TaskActions](docs/sdks/taskactions/README.md)

* [Approve](docs/sdks/taskactions/README.md#approve) - Approve
* [Comment](docs/sdks/taskactions/README.md#comment) - Comment
* [Deny](docs/sdks/taskactions/README.md#deny) - Deny
* [EscalateToEmergencyAccess](docs/sdks/taskactions/README.md#escalatetoemergencyaccess) - Escalate To Emergency Access
* [Restart](docs/sdks/taskactions/README.md#restart) - Restart

### [User](docs/sdks/user/README.md)

* [Get](docs/sdks/user/README.md#get) - Get
* [List](docs/sdks/user/README.md#list) - List
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/sdkerrors"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"log"
)

func main() {
	s := conductoronesdkgo.New(
		conductoronesdkgo.WithSecurity(shared.Security{
			BearerAuth: "",
			Oauth:      "",
		}),
	)

	ctx := context.Background()
	res, err := s.Apps.Create(ctx, &shared.CreateAppRequest{
		Owners: []string{
			"string",
		},
	})
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://{tenantDomain}.conductor.one` | `tenantDomain` (default is `example`) |

#### Example

```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"log"
)

func main() {
	s := conductoronesdkgo.New(
		conductoronesdkgo.WithServerIndex(0),
		conductoronesdkgo.WithSecurity(shared.Security{
			BearerAuth: "",
			Oauth:      "",
		}),
	)

	ctx := context.Background()
	res, err := s.Apps.Create(ctx, &shared.CreateAppRequest{
		Owners: []string{
			"string",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateAppResponse != nil {
		// handle response
	}
}

```

#### Variables

Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithTenantDomain string`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"log"
)

func main() {
	s := conductoronesdkgo.New(
		conductoronesdkgo.WithServerURL("https://{tenantDomain}.conductor.one"),
		conductoronesdkgo.WithSecurity(shared.Security{
			BearerAuth: "",
			Oauth:      "",
		}),
	)

	ctx := context.Background()
	res, err := s.Apps.Create(ctx, &shared.CreateAppRequest{
		Owners: []string{
			"string",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateAppResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `BearerAuth` | http         | HTTP Bearer  |
| `Oauth`      | oauth2       | OAuth2 token |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"log"
)

func main() {
	s := conductoronesdkgo.New(
		conductoronesdkgo.WithSecurity(shared.Security{
			BearerAuth: "",
			Oauth:      "",
		}),
	)

	ctx := context.Background()
	res, err := s.Apps.Create(ctx, &shared.CreateAppRequest{
		Owners: []string{
			"string",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateAppResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
