# RoleMiningManagement

## Overview

### Available Operations

* [CreateAccessProfileFromCohort](#createaccessprofilefromcohort) - Create Access Profile From Cohort
* [GetLatestRun](#getlatestrun) - Get Latest Run
* [GetRoleMiningConfig](#getroleminingconfig) - Get Role Mining Config
* [GetSuggestion](#getsuggestion) - Get Suggestion
* [ListRuns](#listruns) - List Runs
* [ListSuggestions](#listsuggestions) - List Suggestions
* [SearchCohortUsers](#searchcohortusers) - Search Cohort Users
* [TriggerAnalysis](#triggeranalysis) - Trigger Analysis
* [UpdateRoleMiningConfig](#updateroleminingconfig) - Update Role Mining Config
* [UpdateSuggestionState](#updatesuggestionstate) - Update Suggestion State

## CreateAccessProfileFromCohort

CreateAccessProfileFromCohort creates an access profile from a cohort definition,
 adds the specified entitlements, and sets up dynamic membership automation using
 a CEL expression derived from the profile filters.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.role_mining_management.v1.RoleMiningManagementService.CreateAccessProfileFromCohort" method="post" path="/api/v1/role-mining/access-profiles" -->
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

    res, err := s.RoleMiningManagement.CreateAccessProfileFromCohort(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAccessProfileFromCohortResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [shared.CreateAccessProfileFromCohortRequest](../../pkg/models/shared/createaccessprofilefromcohortrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceCreateAccessProfileFromCohortResponse](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementservicecreateaccessprofilefromcohortresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLatestRun

Invokes the c1.api.role_mining_management.v1.RoleMiningManagementService.GetLatestRun method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.role_mining_management.v1.RoleMiningManagementService.GetLatestRun" method="get" path="/api/v1/role-mining/runs/latest" -->
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

    res, err := s.RoleMiningManagement.GetLatestRun(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetLatestRunResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceGetLatestRunResponse](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementservicegetlatestrunresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetRoleMiningConfig

Invokes the c1.api.role_mining_management.v1.RoleMiningManagementService.GetRoleMiningConfig method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.role_mining_management.v1.RoleMiningManagementService.GetRoleMiningConfig" method="get" path="/api/v1/role-mining/config" -->
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

    res, err := s.RoleMiningManagement.GetRoleMiningConfig(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetRoleMiningConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceGetRoleMiningConfigResponse](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementservicegetroleminingconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSuggestion

Invokes the c1.api.role_mining_management.v1.RoleMiningManagementService.GetSuggestion method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.role_mining_management.v1.RoleMiningManagementService.GetSuggestion" method="get" path="/api/v1/role-mining/suggestions/{id}" -->
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

    res, err := s.RoleMiningManagement.GetSuggestion(ctx, operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceGetSuggestionRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSuggestionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                          | Type                                                                                                                                                                                               | Required                                                                                                                                                                                           | Description                                                                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                |
| `request`                                                                                                                                                                                          | [operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceGetSuggestionRequest](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementservicegetsuggestionrequest.md) | :heavy_check_mark:                                                                                                                                                                                 | The request object to use for the request.                                                                                                                                                         |
| `opts`                                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                 | The options for this request.                                                                                                                                                                      |

### Response

**[*operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceGetSuggestionResponse](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementservicegetsuggestionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListRuns

Invokes the c1.api.role_mining_management.v1.RoleMiningManagementService.ListRuns method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.role_mining_management.v1.RoleMiningManagementService.ListRuns" method="get" path="/api/v1/role-mining/runs" -->
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

    res, err := s.RoleMiningManagement.ListRuns(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListRunsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceListRunsResponse](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementservicelistrunsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSuggestions

Invokes the c1.api.role_mining_management.v1.RoleMiningManagementService.ListSuggestions method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.role_mining_management.v1.RoleMiningManagementService.ListSuggestions" method="get" path="/api/v1/role-mining/suggestions" -->
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

    res, err := s.RoleMiningManagement.ListSuggestions(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSuggestionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceListSuggestionsResponse](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementservicelistsuggestionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchCohortUsers

Invokes the c1.api.role_mining_management.v1.RoleMiningManagementService.SearchCohortUsers method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.role_mining_management.v1.RoleMiningManagementService.SearchCohortUsers" method="post" path="/api/v1/role-mining/suggestions/{suggestion_id}/users" -->
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

    res, err := s.RoleMiningManagement.SearchCohortUsers(ctx, operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceSearchCohortUsersRequest{
        SuggestionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchCohortUsersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                  | Type                                                                                                                                                                                                       | Required                                                                                                                                                                                                   | Description                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                        |
| `request`                                                                                                                                                                                                  | [operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceSearchCohortUsersRequest](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementservicesearchcohortusersrequest.md) | :heavy_check_mark:                                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                         | The options for this request.                                                                                                                                                                              |

### Response

**[*operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceSearchCohortUsersResponse](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementservicesearchcohortusersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## TriggerAnalysis

Invokes the c1.api.role_mining_management.v1.RoleMiningManagementService.TriggerAnalysis method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.role_mining_management.v1.RoleMiningManagementService.TriggerAnalysis" method="post" path="/api/v1/role-mining/trigger" -->
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

    res, err := s.RoleMiningManagement.TriggerAnalysis(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TriggerAnalysisResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.TriggerAnalysisRequest](../../pkg/models/shared/triggeranalysisrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceTriggerAnalysisResponse](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementservicetriggeranalysisresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateRoleMiningConfig

Invokes the c1.api.role_mining_management.v1.RoleMiningManagementService.UpdateRoleMiningConfig method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.role_mining_management.v1.RoleMiningManagementService.UpdateRoleMiningConfig" method="post" path="/api/v1/role-mining/config" -->
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

    res, err := s.RoleMiningManagement.UpdateRoleMiningConfig(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateRoleMiningConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [shared.UpdateRoleMiningConfigRequest](../../pkg/models/shared/updateroleminingconfigrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceUpdateRoleMiningConfigResponse](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementserviceupdateroleminingconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateSuggestionState

Invokes the c1.api.role_mining_management.v1.RoleMiningManagementService.UpdateSuggestionState method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.role_mining_management.v1.RoleMiningManagementService.UpdateSuggestionState" method="post" path="/api/v1/role-mining/suggestions/{id}/state" -->
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

    res, err := s.RoleMiningManagement.UpdateSuggestionState(ctx, operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceUpdateSuggestionStateRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateSuggestionStateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                          | Type                                                                                                                                                                                                               | Required                                                                                                                                                                                                           | Description                                                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                |
| `request`                                                                                                                                                                                                          | [operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceUpdateSuggestionStateRequest](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementserviceupdatesuggestionstaterequest.md) | :heavy_check_mark:                                                                                                                                                                                                 | The request object to use for the request.                                                                                                                                                                         |
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |

### Response

**[*operations.C1APIRoleMiningManagementV1RoleMiningManagementServiceUpdateSuggestionStateResponse](../../pkg/models/operations/c1apiroleminingmanagementv1roleminingmanagementserviceupdatesuggestionstateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |