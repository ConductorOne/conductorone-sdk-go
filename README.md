# conductorone-api

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/ConductorOne/conductorone-sdk-go
```
<!-- End SDK Installation [installation] -->

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
	ctx := context.Background()

	s := NewWithCredentials(ctx, &ClientCredentials{
		ClientID:     "",
		ClientSecret: "",
	} )

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
<!-- No SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Apps](docs/sdks/apps/README.md)

* [List](docs/sdks/apps/README.md#list) - List
* [Create](docs/sdks/apps/README.md#create) - Create
* [Delete](docs/sdks/apps/README.md#delete) - Delete
* [Get](docs/sdks/apps/README.md#get) - Get
* [Update](docs/sdks/apps/README.md#update) - Update

### [AppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md)

* [CreateAppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md#createappaccessrequestsdefaults) - Create App Access Requests Defaults
* [GetAppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md#getappaccessrequestsdefaults) - Get App Access Requests Defaults

### [Connector](docs/sdks/connector/README.md)

* [List](docs/sdks/connector/README.md#list) - List
* [CreateDelegated](docs/sdks/connector/README.md#createdelegated) - Create Delegated
* [RotateCredential](docs/sdks/connector/README.md#rotatecredential) - Rotate Credential
* [GetCredentials](docs/sdks/connector/README.md#getcredentials) - Get Credentials
* [RevokeCredential](docs/sdks/connector/README.md#revokecredential) - Revoke Credential
* [ForceSync](docs/sdks/connector/README.md#forcesync) - Force Sync
* [Delete](docs/sdks/connector/README.md#delete) - Delete
* [Get](docs/sdks/connector/README.md#get) - Get
* [Update](docs/sdks/connector/README.md#update) - Update
* [Create](docs/sdks/connector/README.md#create) - Create
* [UpdateDelegated](docs/sdks/connector/README.md#updatedelegated) - Update Delegated

### [AppEntitlements](docs/sdks/appentitlements/README.md)

* [List](docs/sdks/appentitlements/README.md#list) - List
* [~~ListUsers~~](docs/sdks/appentitlements/README.md#listusers) - List Users :warning: **Deprecated**
* [Get](docs/sdks/appentitlements/README.md#get) - Get
* [Update](docs/sdks/appentitlements/README.md#update) - Update
* [ListForAppResource](docs/sdks/appentitlements/README.md#listforappresource) - List For App Resource
* [ListForAppUser](docs/sdks/appentitlements/README.md#listforappuser) - List For App User

### [AppEntitlementSearch](docs/sdks/appentitlementsearch/README.md)

* [SearchAppEntitlementsWithExpired](docs/sdks/appentitlementsearch/README.md#searchappentitlementswithexpired) - Search App Entitlements With Expired
* [Search](docs/sdks/appentitlementsearch/README.md#search) - Search

### [AppEntitlementUserBinding](docs/sdks/appentitlementuserbinding/README.md)

* [ListAppUsersForIdentityWithGrant](docs/sdks/appentitlementuserbinding/README.md#listappusersforidentitywithgrant) - List App Users For Identity With Grant

### [AppEntitlementOwners](docs/sdks/appentitlementowners/README.md)

* [List](docs/sdks/appentitlementowners/README.md#list) - List
* [Add](docs/sdks/appentitlementowners/README.md#add) - Add
* [Set](docs/sdks/appentitlementowners/README.md#set) - Set
* [Remove](docs/sdks/appentitlementowners/README.md#remove) - Remove

### [AppOwners](docs/sdks/appowners/README.md)

* [List](docs/sdks/appowners/README.md#list) - List
* [Set](docs/sdks/appowners/README.md#set) - Set
* [Remove](docs/sdks/appowners/README.md#remove) - Remove
* [Add](docs/sdks/appowners/README.md#add) - Add

### [AppReport](docs/sdks/appreport/README.md)

* [List](docs/sdks/appreport/README.md#list) - List

### [AppReportAction](docs/sdks/appreportaction/README.md)

* [GenerateReport](docs/sdks/appreportaction/README.md#generatereport) - Generate Report

### [AppResourceType](docs/sdks/appresourcetype/README.md)

* [List](docs/sdks/appresourcetype/README.md#list) - List
* [Get](docs/sdks/appresourcetype/README.md#get) - Get

### [AppResource](docs/sdks/appresource/README.md)

* [List](docs/sdks/appresource/README.md#list) - List
* [Get](docs/sdks/appresource/README.md#get) - Get

### [AppResourceOwners](docs/sdks/appresourceowners/README.md)

* [List](docs/sdks/appresourceowners/README.md#list) - List

### [AppUsageControls](docs/sdks/appusagecontrols/README.md)

* [Get](docs/sdks/appusagecontrols/README.md#get) - Get
* [Update](docs/sdks/appusagecontrols/README.md#update) - Update

### [AppUser](docs/sdks/appuser/README.md)

* [Update](docs/sdks/appuser/README.md#update) - Update

### [Attributes](docs/sdks/attributes/README.md)

* [DeleteAttributeValue](docs/sdks/attributes/README.md#deleteattributevalue) - Delete Attribute Value
* [CreateAttributeValue](docs/sdks/attributes/README.md#createattributevalue) - Create Attribute Value
* [GetAttributeValue](docs/sdks/attributes/README.md#getattributevalue) - Get Attribute Value
* [ListAttributeTypes](docs/sdks/attributes/README.md#listattributetypes) - List Attribute Types
* [ListAttributeValues](docs/sdks/attributes/README.md#listattributevalues) - List Attribute Values

### [Auth](docs/sdks/auth/README.md)

* [Introspect](docs/sdks/auth/README.md#introspect) - Introspect

### [RequestCatalogManagement](docs/sdks/requestcatalogmanagement/README.md)

* [List](docs/sdks/requestcatalogmanagement/README.md#list) - List
* [Create](docs/sdks/requestcatalogmanagement/README.md#create) - Create
* [ListEntitlementsPerCatalog](docs/sdks/requestcatalogmanagement/README.md#listentitlementspercatalog) - List Entitlements Per Catalog
* [RemoveAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeappentitlements) - Remove App Entitlements
* [AddAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#addappentitlements) - Add App Entitlements
* [RemoveAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeaccessentitlements) - Remove Access Entitlements
* [AddAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#addaccessentitlements) - Add Access Entitlements
* [ListEntitlementsForAccess](docs/sdks/requestcatalogmanagement/README.md#listentitlementsforaccess) - List Entitlements For Access
* [Delete](docs/sdks/requestcatalogmanagement/README.md#delete) - Delete
* [Get](docs/sdks/requestcatalogmanagement/README.md#get) - Get
* [Update](docs/sdks/requestcatalogmanagement/README.md#update) - Update

### [Directory](docs/sdks/directory/README.md)

* [List](docs/sdks/directory/README.md#list) - List
* [Create](docs/sdks/directory/README.md#create) - Create
* [Delete](docs/sdks/directory/README.md#delete) - Delete
* [Get](docs/sdks/directory/README.md#get) - Get

### [PersonalClient](docs/sdks/personalclient/README.md)

* [Create](docs/sdks/personalclient/README.md#create) - Create

### [Roles](docs/sdks/roles/README.md)

* [List](docs/sdks/roles/README.md#list) - List
* [Get](docs/sdks/roles/README.md#get) - Get
* [Update](docs/sdks/roles/README.md#update) - Update

### [Policies](docs/sdks/policies/README.md)

* [List](docs/sdks/policies/README.md#list) - List
* [Create](docs/sdks/policies/README.md#create) - Create
* [Delete](docs/sdks/policies/README.md#delete) - Delete
* [Get](docs/sdks/policies/README.md#get) - Get
* [Update](docs/sdks/policies/README.md#update) - Update

### [PolicyValidate](docs/sdks/policyvalidate/README.md)

* [ValidateCEL](docs/sdks/policyvalidate/README.md#validatecel) - Validate Cel

### [AppResourceSearch](docs/sdks/appresourcesearch/README.md)

* [SearchAppResourceTypes](docs/sdks/appresourcesearch/README.md#searchappresourcetypes) - Search App Resource Types

### [AppSearch](docs/sdks/appsearch/README.md)

* [Search](docs/sdks/appsearch/README.md#search) - Search

### [AttributeSearch](docs/sdks/attributesearch/README.md)

* [SearchAttributeValues](docs/sdks/attributesearch/README.md#searchattributevalues) - Search Attribute Values

### [PolicySearch](docs/sdks/policysearch/README.md)

* [Search](docs/sdks/policysearch/README.md#search) - Search

### [RequestCatalogSearch](docs/sdks/requestcatalogsearch/README.md)

* [SearchEntitlements](docs/sdks/requestcatalogsearch/README.md#searchentitlements) - Search Entitlements

### [TaskSearch](docs/sdks/tasksearch/README.md)

* [Search](docs/sdks/tasksearch/README.md#search) - Search

### [UserSearch](docs/sdks/usersearch/README.md)

* [Search](docs/sdks/usersearch/README.md#search) - Search

### [WebhooksSearch](docs/sdks/webhookssearch/README.md)

* [Search](docs/sdks/webhookssearch/README.md#search) - Search

### [AWSExternalIDSettings](docs/sdks/awsexternalidsettings/README.md)

* [Get](docs/sdks/awsexternalidsettings/README.md#get) - Get

### [SessionSettings](docs/sdks/sessionsettings/README.md)

* [Get](docs/sdks/sessionsettings/README.md#get) - Get
* [Update](docs/sdks/sessionsettings/README.md#update) - Update

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

* [List](docs/sdks/user/README.md#list) - List
* [Get](docs/sdks/user/README.md#get) - Get

### [Webhooks](docs/sdks/webhooks/README.md)

* [List](docs/sdks/webhooks/README.md#list) - List
* [Create](docs/sdks/webhooks/README.md#create) - Create
* [Delete](docs/sdks/webhooks/README.md#delete) - Delete
* [Get](docs/sdks/webhooks/README.md#get) - Get
* [Update](docs/sdks/webhooks/README.md#update) - Update
* [Test](docs/sdks/webhooks/README.md#test) - Test
<!-- End Available Resources and Operations [operations] -->


<!-- No Special Types [types] -->


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
	ctx := context.Background()

	s := NewWithCredentials(ctx, &ClientCredentials{
		ClientID:     "",
		ClientSecret: "",
	})
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
<!-- No Error Handling [errors] -->


## SDK Example Usage with Custom Server/Tenant

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
	ctx := context.Background()

	/* Optional Override 
	* Server URL will be extracted from client, optionally, you can
	* provide a server URL or a tenant domain (will create URL https://{tenant_domain}.conductor.one) 
	*/
	opts := []sdk.CustomSDKOption{}
	opt, _ := sdk.WithTenantCustom("Server URL or Tenant Domain")
	opts = append(opts, opt)

	s := NewWithCredentials(ctx, &ClientCredentials{
		ClientID:     "",
		ClientSecret: "",
	} opts...)

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
<!-- No Server Selection [server] -->

<!-- No Custom HTTP Client [http-client] -->

<!-- No Authentication [security] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := conductoronesdkgo.New(
		conductoronesdkgo.WithSecurity(shared.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	ctx := context.Background()
	res, err := s.User.List(ctx, operations.C1APIUserV1UserServiceListRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.UserServiceListResponse != nil {
		for {
			// handle items

			res = res.Next()
			if res == nil {
				break
			}
		}

	}
}

```
<!-- End Pagination [pagination] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
