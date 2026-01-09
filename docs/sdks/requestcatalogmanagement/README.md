# RequestCatalogManagement

## Overview

### Available Operations

* [AddAccessEntitlements](#addaccessentitlements) - Add Access Entitlements
* [AddAppEntitlements](#addappentitlements) - Add App Entitlements
* [Create](#create) - Create
* [CreateBundleAutomation](#createbundleautomation) - Create Bundle Automation
* [CreateRequestableEntry](#createrequestableentry) - Create Requestable Entry
* [Delete](#delete) - Delete
* [DeleteBundleAutomation](#deletebundleautomation) - Delete Bundle Automation
* [DeleteRequestableEntry](#deleterequestableentry) - Delete Requestable Entry
* [ForceRunBundleAutomation](#forcerunbundleautomation) - Force Run Bundle Automation
* [Get](#get) - Get
* [GetBundleAutomation](#getbundleautomation) - Get Bundle Automation
* [GetRequestableEntry](#getrequestableentry) - Get Requestable Entry
* [List](#list) - List
* [ListAllEntitlementIdsPerApp](#listallentitlementidsperapp) - List All Entitlement Ids Per App
* [ListEntitlementsForAccess](#listentitlementsforaccess) - List Entitlements For Access
* [ListEntitlementsPerCatalog](#listentitlementspercatalog) - List Entitlements Per Catalog
* [RemoveAccessEntitlements](#removeaccessentitlements) - Remove Access Entitlements
* [RemoveAppEntitlements](#removeappentitlements) - Remove App Entitlements
* [ResumePausedBundleAutomation](#resumepausedbundleautomation) - Resume Paused Bundle Automation
* [SetBundleAutomation](#setbundleautomation) - Set Bundle Automation
* [Update](#update) - Update
* [UpdateAppEntitlements](#updateappentitlements) - Update App Entitlements

## AddAccessEntitlements

Add visibility bindings (access entitlements) to a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.AddAccessEntitlements" method="post" path="/api/v1/catalogs/{catalog_id}/visibility_bindings" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.AddAccessEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsRequest{
        CatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceAddAccessEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                      | Type                                                                                                                                                                                                           | Required                                                                                                                                                                                                       | Description                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                            |
| `request`                                                                                                                                                                                                      | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddaccessentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                             | The options for this request.                                                                                                                                                                                  |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddaccessentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddAppEntitlements

Add requestable entitlements to a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.AddAppEntitlements" method="post" path="/api/v1/catalogs/{catalog_id}/requestable_entries" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.AddAppEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsRequest{
        CatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceAddAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                | Type                                                                                                                                                                                                     | Required                                                                                                                                                                                                 | Description                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                      |
| `request`                                                                                                                                                                                                | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddappentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                               |
| `opts`                                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                       | The options for this request.                                                                                                                                                                            |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddappentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Create

Creates a new request catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.Create" method="post" path="/api/v1/catalogs" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.Create(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [shared.RequestCatalogManagementServiceCreateRequest](../../pkg/models/shared/requestcatalogmanagementservicecreaterequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateBundleAutomation

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.CreateBundleAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.CreateBundleAutomation" method="post" path="/api/v1/catalogs/{request_catalog_id}/bundle_automation/create" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.CreateBundleAutomation(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceCreateBundleAutomationRequest{
        RequestCatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BundleAutomation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                        | Type                                                                                                                                                                                                             | Required                                                                                                                                                                                                         | Description                                                                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                              |
| `request`                                                                                                                                                                                                        | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceCreateBundleAutomationRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicecreatebundleautomationrequest.md) | :heavy_check_mark:                                                                                                                                                                                               | The request object to use for the request.                                                                                                                                                                       |
| `opts`                                                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                               | The options for this request.                                                                                                                                                                                    |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceCreateBundleAutomationResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicecreatebundleautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateRequestableEntry

Create a single requestable entry

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.CreateRequestableEntry" method="put" path="/api/v1/catalogs/{catalog_id}/requestable_entries/{app_id}/{entitlement_id}" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.CreateRequestableEntry(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceCreateRequestableEntryRequest{
        AppID: "<id>",
        CatalogID: "<id>",
        EntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceCreateRequestableEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                        | Type                                                                                                                                                                                                             | Required                                                                                                                                                                                                         | Description                                                                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                              |
| `request`                                                                                                                                                                                                        | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceCreateRequestableEntryRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicecreaterequestableentryrequest.md) | :heavy_check_mark:                                                                                                                                                                                               | The request object to use for the request.                                                                                                                                                                       |
| `opts`                                                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                               | The options for this request.                                                                                                                                                                                    |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceCreateRequestableEntryResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicecreaterequestableentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.Delete" method="delete" path="/api/v1/catalogs/{id}" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.Delete(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `request`                                                                                                                                                                        | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicedeleterequest.md) | :heavy_check_mark:                                                                                                                                                               | The request object to use for the request.                                                                                                                                       |
| `opts`                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                     | :heavy_minus_sign:                                                                                                                                                               | The options for this request.                                                                                                                                                    |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteBundleAutomation

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.DeleteBundleAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.DeleteBundleAutomation" method="delete" path="/api/v1/catalogs/{request_catalog_id}/bundle_automation" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.DeleteBundleAutomation(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteBundleAutomationRequest{
        RequestCatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteBundleAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                        | Type                                                                                                                                                                                                             | Required                                                                                                                                                                                                         | Description                                                                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                              |
| `request`                                                                                                                                                                                                        | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteBundleAutomationRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicedeletebundleautomationrequest.md) | :heavy_check_mark:                                                                                                                                                                                               | The request object to use for the request.                                                                                                                                                                       |
| `opts`                                                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                               | The options for this request.                                                                                                                                                                                    |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteBundleAutomationResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicedeletebundleautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteRequestableEntry

Delete a single requestable entry

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.DeleteRequestableEntry" method="delete" path="/api/v1/catalogs/{catalog_id}/requestable_entries/{app_id}/{entitlement_id}" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.DeleteRequestableEntry(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequestableEntryRequest{
        AppID: "<id>",
        CatalogID: "<id>",
        EntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceDeleteRequestableEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                        | Type                                                                                                                                                                                                             | Required                                                                                                                                                                                                         | Description                                                                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                              |
| `request`                                                                                                                                                                                                        | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequestableEntryRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicedeleterequestableentryrequest.md) | :heavy_check_mark:                                                                                                                                                                                               | The request object to use for the request.                                                                                                                                                                       |
| `opts`                                                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                               | The options for this request.                                                                                                                                                                                    |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequestableEntryResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicedeleterequestableentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceRunBundleAutomation

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.ForceRunBundleAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.ForceRunBundleAutomation" method="post" path="/api/v1/catalogs/{request_catalog_id}/bundle_automation/run" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.ForceRunBundleAutomation(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceForceRunBundleAutomationRequest{
        RequestCatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ForceRunBundleAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                            | Type                                                                                                                                                                                                                 | Required                                                                                                                                                                                                             | Description                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                  |
| `request`                                                                                                                                                                                                            | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceForceRunBundleAutomationRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceforcerunbundleautomationrequest.md) | :heavy_check_mark:                                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                                           |
| `opts`                                                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                        |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceForceRunBundleAutomationResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceforcerunbundleautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.Get" method="get" path="/api/v1/catalogs/{id}" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.Get(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetrequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |
| `opts`                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                         | The options for this request.                                                                                                                                              |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetBundleAutomation

Get bundle automation

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.GetBundleAutomation" method="get" path="/api/v1/catalogs/{request_catalog_id}/bundle_automation" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.GetBundleAutomation(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetBundleAutomationRequest{
        RequestCatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BundleAutomation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                  | Type                                                                                                                                                                                                       | Required                                                                                                                                                                                                   | Description                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                        |
| `request`                                                                                                                                                                                                  | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetBundleAutomationRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetbundleautomationrequest.md) | :heavy_check_mark:                                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                         | The options for this request.                                                                                                                                                                              |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetBundleAutomationResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetbundleautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetRequestableEntry

Get a single requestable entry

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.GetRequestableEntry" method="get" path="/api/v1/catalogs/{catalog_id}/requestable_entries/{app_id}/{entitlement_id}" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.GetRequestableEntry(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequestableEntryRequest{
        AppID: "<id>",
        CatalogID: "<id>",
        EntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceGetRequestableEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                  | Type                                                                                                                                                                                                       | Required                                                                                                                                                                                                   | Description                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                        |
| `request`                                                                                                                                                                                                  | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequestableEntryRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetrequestableentryrequest.md) | :heavy_check_mark:                                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                         | The options for this request.                                                                                                                                                                              |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequestableEntryResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetrequestableentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

Get a list of request catalogs.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.List" method="get" path="/api/v1/catalogs" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.List(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |
| `opts`                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                           | The options for this request.                                                                                                                                                |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAllEntitlementIdsPerApp

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.ListAllEntitlementIdsPerApp method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.ListAllEntitlementIdsPerApp" method="get" path="/api/v1/catalogs/{catalog_id}/requestable_entitlementIDs" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.ListAllEntitlementIdsPerApp(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListAllEntitlementIdsPerAppRequest{
        CatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceListAllEntitlementIdsPerCatalogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                                        |
| `request`                                                                                                                                                                                                                  | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListAllEntitlementIdsPerAppRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistallentitlementidsperapprequest.md) | :heavy_check_mark:                                                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                         | The options for this request.                                                                                                                                                                                              |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListAllEntitlementIdsPerAppResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistallentitlementidsperappresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEntitlementsForAccess

List visibility bindings (access entitlements) for a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.ListEntitlementsForAccess" method="get" path="/api/v1/catalogs/{catalog_id}/visibility_entitlements" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.ListEntitlementsForAccess(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest{
        CatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceListEntitlementsForAccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                              | Type                                                                                                                                                                                                                   | Required                                                                                                                                                                                                               | Description                                                                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                    |
| `request`                                                                                                                                                                                                              | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementsforaccessrequest.md) | :heavy_check_mark:                                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                          |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementsforaccessresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEntitlementsPerCatalog

List entitlements in a catalog that are requestable.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.ListEntitlementsPerCatalog" method="get" path="/api/v1/catalogs/{catalog_id}/requestable_entitlements" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.ListEntitlementsPerCatalog(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogRequest{
        CatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceListEntitlementsPerCatalogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                | Type                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                 | Description                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                      |
| `request`                                                                                                                                                                                                                | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementspercatalogrequest.md) | :heavy_check_mark:                                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                                               |
| `opts`                                                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                            |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementspercatalogresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccessEntitlements

Remove visibility bindings (access entitlements) to a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.RemoveAccessEntitlements" method="delete" path="/api/v1/catalogs/{catalog_id}/visibility_bindings" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.RemoveAccessEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsRequest{
        CatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceRemoveAccessEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                            | Type                                                                                                                                                                                                                 | Required                                                                                                                                                                                                             | Description                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                  |
| `request`                                                                                                                                                                                                            | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveaccessentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                                           |
| `opts`                                                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                        |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveaccessentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAppEntitlements

Remove requestable entitlements from a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.RemoveAppEntitlements" method="delete" path="/api/v1/catalogs/{catalog_id}/requestable_entries" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.RemoveAppEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsRequest{
        CatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceRemoveAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                      | Type                                                                                                                                                                                                           | Required                                                                                                                                                                                                       | Description                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                            |
| `request`                                                                                                                                                                                                      | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveappentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                             | The options for this request.                                                                                                                                                                                  |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveappentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ResumePausedBundleAutomation

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.ResumePausedBundleAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.ResumePausedBundleAutomation" method="post" path="/api/v1/catalogs/{request_catalog_id}/bundle_automation/resume" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.ResumePausedBundleAutomation(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceResumePausedBundleAutomationRequest{
        RequestCatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResumePausedBundleAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                          |
| `request`                                                                                                                                                                                                                    | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceResumePausedBundleAutomationRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceresumepausedbundleautomationrequest.md) | :heavy_check_mark:                                                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                                                   |
| `opts`                                                                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceResumePausedBundleAutomationResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceresumepausedbundleautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetBundleAutomation

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.SetBundleAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.SetBundleAutomation" method="post" path="/api/v1/catalogs/{request_catalog_id}/bundle_automation" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.SetBundleAutomation(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationRequest{
        RequestCatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BundleAutomation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                  | Type                                                                                                                                                                                                       | Required                                                                                                                                                                                                   | Description                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                        |
| `request`                                                                                                                                                                                                  | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicesetbundleautomationrequest.md) | :heavy_check_mark:                                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                         | The options for this request.                                                                                                                                                                              |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementservicesetbundleautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.Update" method="post" path="/api/v1/catalogs/{id}" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.Update(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `request`                                                                                                                                                                        | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                                                               | The request object to use for the request.                                                                                                                                       |
| `opts`                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                     | :heavy_minus_sign:                                                                                                                                                               | The options for this request.                                                                                                                                                    |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAppEntitlements

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.UpdateAppEntitlements method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.UpdateAppEntitlements" method="post" path="/api/v1/catalogs/{catalog_id}/requestable_entitlements/update" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.RequestCatalogManagement.UpdateAppEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateAppEntitlementsRequest{
        CatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceUpdateAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                      | Type                                                                                                                                                                                                           | Required                                                                                                                                                                                                       | Description                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                            |
| `request`                                                                                                                                                                                                      | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateAppEntitlementsRequest](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceupdateappentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                             | The options for this request.                                                                                                                                                                                  |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateAppEntitlementsResponse](../../pkg/models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceupdateappentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |