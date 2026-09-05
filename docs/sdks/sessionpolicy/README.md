# SessionPolicy

## Overview

### Available Operations

* [AssignGroup](#assigngroup) - Assign Group
* [AssignUser](#assignuser) - Assign User
* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [GetEffectiveSessionPolicy](#geteffectivesessionpolicy) - Get Effective Session Policy
* [List](#list) - List
* [ListAssignments](#listassignments) - List Assignments
* [ListUserPolicies](#listuserpolicies) - List User Policies
* [Search](#search) - Search
* [SearchPolicyUsers](#searchpolicyusers) - Search Policy Users
* [UnassignGroup](#unassigngroup) - Unassign Group
* [UnassignUser](#unassignuser) - Unassign User
* [Update](#update) - Update

## AssignGroup

Assign a group to a session policy. Every member of the group becomes
 assigned to the policy; because group membership is expanded
 asynchronously, the per-user effect is eventually consistent (typically
 within a few minutes).

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.AssignGroup" method="post" path="/api/v1/session-policies/{id}/assignments/groups" -->
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

    res, err := s.SessionPolicy.AssignGroup(ctx, operations.C1APISessionPolicyV1SessionPolicyServiceAssignGroupRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceAssignGroupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APISessionPolicyV1SessionPolicyServiceAssignGroupRequest](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyserviceassigngrouprequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceAssignGroupResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyserviceassigngroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AssignUser

Assign a user to a session policy. The assignment takes effect immediately.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.AssignUser" method="post" path="/api/v1/session-policies/{id}/assignments/users" -->
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

    res, err := s.SessionPolicy.AssignUser(ctx, operations.C1APISessionPolicyV1SessionPolicyServiceAssignUserRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceAssignUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.C1APISessionPolicyV1SessionPolicyServiceAssignUserRequest](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyserviceassignuserrequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |
| `opts`                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceAssignUserResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyserviceassignuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Create

Create a session policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.Create" method="post" path="/api/v1/session-policies" -->
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

    res, err := s.SessionPolicy.Create(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [shared.SessionPolicyServiceCreateRequest](../../pkg/models/shared/sessionpolicyservicecreaterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceCreateResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a session policy by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.Delete" method="delete" path="/api/v1/session-policies/{id}" -->
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

    res, err := s.SessionPolicy.Delete(ctx, operations.C1APISessionPolicyV1SessionPolicyServiceDeleteRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APISessionPolicyV1SessionPolicyServiceDeleteRequest](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicedeleterequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceDeleteResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Get a session policy by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.Get" method="get" path="/api/v1/session-policies/{id}" -->
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

    res, err := s.SessionPolicy.Get(ctx, operations.C1APISessionPolicyV1SessionPolicyServiceGetRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APISessionPolicyV1SessionPolicyServiceGetRequest](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicegetrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceGetResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetEffectiveSessionPolicy

Returns the single effective session policy for a user and why it applies:
 the assigned policy with the highest priority, else the tenant default,
 else none. Read-only and diagnostic: it reflects current assignment state
 rather than the resolver's cached resolution.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.GetEffectiveSessionPolicy" method="get" path="/api/v1/users/{user_id}/effective-session-policy" -->
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

    res, err := s.SessionPolicy.GetEffectiveSessionPolicy(ctx, operations.C1APISessionPolicyV1SessionPolicyServiceGetEffectiveSessionPolicyRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceGetEffectiveSessionPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                      | Type                                                                                                                                                                                           | Required                                                                                                                                                                                       | Description                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                             | The context to use for the request.                                                                                                                                                            |
| `request`                                                                                                                                                                                      | [operations.C1APISessionPolicyV1SessionPolicyServiceGetEffectiveSessionPolicyRequest](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicegeteffectivesessionpolicyrequest.md) | :heavy_check_mark:                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                     |
| `opts`                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                             | The options for this request.                                                                                                                                                                  |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceGetEffectiveSessionPolicyResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicegeteffectivesessionpolicyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List all session policies in your tenant, one page at a time.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.List" method="get" path="/api/v1/session-policies" -->
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

    res, err := s.SessionPolicy.List(ctx, operations.C1APISessionPolicyV1SessionPolicyServiceListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.C1APISessionPolicyV1SessionPolicyServiceListRequest](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicelistrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceListResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAssignments

List the principals assigned to a session policy, including both direct
 assignments and those conferred through a group.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.ListAssignments" method="get" path="/api/v1/session-policies/{id}/assignments" -->
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

    res, err := s.SessionPolicy.ListAssignments(ctx, operations.C1APISessionPolicyV1SessionPolicyServiceListAssignmentsRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceListAssignmentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.C1APISessionPolicyV1SessionPolicyServiceListAssignmentsRequest](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicelistassignmentsrequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |
| `opts`                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                         | The options for this request.                                                                                                                                              |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceListAssignmentsResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicelistassignmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUserPolicies

Lists every session policy that applies to a user — direct,
 group-conferred, and the tenant default when set — each with its source.
 Candidates are returned in assignment order; the resolver's priority
 tie-break happens at evaluation, not here.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.ListUserPolicies" method="get" path="/api/v1/users/{user_id}/session-policies" -->
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

    res, err := s.SessionPolicy.ListUserPolicies(ctx, operations.C1APISessionPolicyV1SessionPolicyServiceListUserPoliciesRequest{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceListUserPoliciesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.C1APISessionPolicyV1SessionPolicyServiceListUserPoliciesRequest](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicelistuserpoliciesrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |
| `opts`                                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                           | The options for this request.                                                                                                                                                |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceListUserPoliciesResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicelistuserpoliciesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Search

Search session policies by name, or fetch a specific set by ID. Returns
 one page of matching policies at a time.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.Search" method="post" path="/api/v1/search/session-policies" -->
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

    res, err := s.SessionPolicy.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [shared.SessionPolicyServiceSearchRequest](../../pkg/models/shared/sessionpolicyservicesearchrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceSearchResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchPolicyUsers

Searches the users a policy applies to, expanded: direct assignments plus
 per-user expansion of conferred groups, one page at a time. Served from
 the asynchronously-replicated Postgres mirror of the binding store, so a
 just-made assignment may not appear immediately; the user-scoped reads
 (GetEffectiveSessionPolicy, ListUserPolicies) reflect current state.
 Search, not List: the results are a filtered, server-paginated discovery
 over the binding store (query and source facets), not a bounded
 collection of a named object.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.SearchPolicyUsers" method="post" path="/api/v1/session-policies/{id}/users/search" -->
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

    res, err := s.SessionPolicy.SearchPolicyUsers(ctx, operations.C1APISessionPolicyV1SessionPolicyServiceSearchPolicyUsersRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceSearchPolicyUsersResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                          | :heavy_check_mark:                                                                                                                                                             | The context to use for the request.                                                                                                                                            |
| `request`                                                                                                                                                                      | [operations.C1APISessionPolicyV1SessionPolicyServiceSearchPolicyUsersRequest](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicesearchpolicyusersrequest.md) | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `opts`                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                   | :heavy_minus_sign:                                                                                                                                                             | The options for this request.                                                                                                                                                  |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceSearchPolicyUsersResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyservicesearchpolicyusersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UnassignGroup

Unassign a group from a session policy. The per-user effect is eventually
 consistent, mirroring AssignGroup.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.UnassignGroup" method="delete" path="/api/v1/session-policies/{id}/assignments/groups/{group_app_entitlement_id}" -->
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

    res, err := s.SessionPolicy.UnassignGroup(ctx, operations.C1APISessionPolicyV1SessionPolicyServiceUnassignGroupRequest{
        GroupAppEntitlementID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceUnassignGroupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.C1APISessionPolicyV1SessionPolicyServiceUnassignGroupRequest](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyserviceunassigngrouprequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |
| `opts`                                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                           | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceUnassignGroupResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyserviceunassigngroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UnassignUser

Unassign a user from a session policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.UnassignUser" method="delete" path="/api/v1/session-policies/{id}/assignments/users/{user_id}" -->
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

    res, err := s.SessionPolicy.UnassignUser(ctx, operations.C1APISessionPolicyV1SessionPolicyServiceUnassignUserRequest{
        ID: "<id>",
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceUnassignUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.C1APISessionPolicyV1SessionPolicyServiceUnassignUserRequest](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyserviceunassignuserrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceUnassignUserResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyserviceunassignuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update a session policy. Supply the policy object and an update mask
 listing the fields to change; omitted fields are left as-is.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.session_policy.v1.SessionPolicyService.Update" method="post" path="/api/v1/session-policies/{id}" -->
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

    res, err := s.SessionPolicy.Update(ctx, operations.C1APISessionPolicyV1SessionPolicyServiceUpdateRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionPolicyServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APISessionPolicyV1SessionPolicyServiceUpdateRequest](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APISessionPolicyV1SessionPolicyServiceUpdateResponse](../../pkg/models/operations/c1apisessionpolicyv1sessionpolicyserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |