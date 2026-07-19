# Functions

## Overview

### Available Operations

* [CreateFinalCommit](#createfinalcommit) - Create Final Commit
* [CreateFunction](#createfunction) - Create Function
* [CreateInitialCommit](#createinitialcommit) - Create Initial Commit
* [CreateTag](#createtag) - Create Tag
* [DeleteFunction](#deletefunction) - Delete Function
* [GetCommitContent](#getcommitcontent) - Get Commit Content
* [GetFunction](#getfunction) - Get Function
* [GetLockFile](#getlockfile) - Get Lock File
* [Invoke](#invoke) - Invoke
* [ListCommits](#listcommits) - List Commits
* [ListFunctions](#listfunctions) - List Functions
* [ListTags](#listtags) - List Tags
* [Test](#test) - Test
* [UpdateFunction](#updatefunction) - Update Function

## CreateFinalCommit

CreateFinalCommit completes a commit after files are uploaded

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.CreateFinalCommit" method="post" path="/api/v1/functions/{function_id}/commits/{commit_id}/finalize" -->
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

    res, err := s.Functions.CreateFinalCommit(ctx, operations.C1APIFunctionsV1FunctionsServiceCreateFinalCommitRequest{
        CommitID: "<id>",
        FunctionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceCreateFinalCommitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.C1APIFunctionsV1FunctionsServiceCreateFinalCommitRequest](../../pkg/models/operations/c1apifunctionsv1functionsservicecreatefinalcommitrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |
| `opts`                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                   | :heavy_minus_sign:                                                                                                                                             | The options for this request.                                                                                                                                  |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceCreateFinalCommitResponse](../../pkg/models/operations/c1apifunctionsv1functionsservicecreatefinalcommitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateFunction

CreateFunction registers a new serverless function and creates its
 initial code commit. Functions run as TypeScript modules in a sandboxed
 runtime; see initial_content for the entry-file signature and SDK import.

 The new function is unpublished. To make the commit the default
 runnable version (and have the function appear as runnable in the
 Functions UI), call UpdateFunction with function.published_commit_id
 set and update_mask=["published_commit_id"].

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.CreateFunction" method="post" path="/api/v1/functions" -->
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

    res, err := s.Functions.CreateFunction(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceCreateFunctionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [shared.FunctionsServiceCreateFunctionRequest](../../pkg/models/shared/functionsservicecreatefunctionrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceCreateFunctionResponse](../../pkg/models/operations/c1apifunctionsv1functionsservicecreatefunctionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateInitialCommit

CreateInitialCommit starts a new commit and returns upload URLs for files

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.CreateInitialCommit" method="post" path="/api/v1/functions/{function_id}/commits" -->
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

    res, err := s.Functions.CreateInitialCommit(ctx, operations.C1APIFunctionsV1FunctionsServiceCreateInitialCommitRequest{
        FunctionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceCreateInitialCommitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APIFunctionsV1FunctionsServiceCreateInitialCommitRequest](../../pkg/models/operations/c1apifunctionsv1functionsservicecreateinitialcommitrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceCreateInitialCommitResponse](../../pkg/models/operations/c1apifunctionsv1functionsservicecreateinitialcommitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateTag

CreateTag creates a named reference to a specific commit

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.CreateTag" method="post" path="/api/v1/functions/{function_id}/tags" -->
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

    res, err := s.Functions.CreateTag(ctx, operations.C1APIFunctionsV1FunctionsServiceCreateTagRequest{
        FunctionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceCreateTagResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIFunctionsV1FunctionsServiceCreateTagRequest](../../pkg/models/operations/c1apifunctionsv1functionsservicecreatetagrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceCreateTagResponse](../../pkg/models/operations/c1apifunctionsv1functionsservicecreatetagresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteFunction

Delete removes a function

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.DeleteFunction" method="delete" path="/api/v1/functions/{id}" -->
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

    res, err := s.Functions.DeleteFunction(ctx, operations.C1APIFunctionsV1FunctionsServiceDeleteFunctionRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceDeleteFunctionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APIFunctionsV1FunctionsServiceDeleteFunctionRequest](../../pkg/models/operations/c1apifunctionsv1functionsservicedeletefunctionrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceDeleteFunctionResponse](../../pkg/models/operations/c1apifunctionsv1functionsservicedeletefunctionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCommitContent

GetCommitContent retrieves a commit and all its file contents in a single unary response.
 This is a non-streaming alternative to GetCommit for REST API consumers.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.GetCommitContent" method="get" path="/api/v1/functions/{function_id}/commits/{id}" -->
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

    res, err := s.Functions.GetCommitContent(ctx, operations.C1APIFunctionsV1FunctionsServiceGetCommitContentRequest{
        FunctionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceGetCommitContentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.C1APIFunctionsV1FunctionsServiceGetCommitContentRequest](../../pkg/models/operations/c1apifunctionsv1functionsservicegetcommitcontentrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceGetCommitContentResponse](../../pkg/models/operations/c1apifunctionsv1functionsservicegetcommitcontentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetFunction

Get retrieves a specific function by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.GetFunction" method="get" path="/api/v1/functions/{id}" -->
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

    res, err := s.Functions.GetFunction(ctx, operations.C1APIFunctionsV1FunctionsServiceGetFunctionRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceGetFunctionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIFunctionsV1FunctionsServiceGetFunctionRequest](../../pkg/models/operations/c1apifunctionsv1functionsservicegetfunctionrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceGetFunctionResponse](../../pkg/models/operations/c1apifunctionsv1functionsservicegetfunctionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLockFile

GetLockFile retrieves the deno lock file for a specific commit, if it exists.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.GetLockFile" method="get" path="/api/v1/functions/{function_id}/commits/{commit_id}/lockfile" -->
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

    res, err := s.Functions.GetLockFile(ctx, operations.C1APIFunctionsV1FunctionsServiceGetLockFileRequest{
        CommitID: "<id>",
        FunctionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceGetLockFileResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIFunctionsV1FunctionsServiceGetLockFileRequest](../../pkg/models/operations/c1apifunctionsv1functionsservicegetlockfilerequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceGetLockFileResponse](../../pkg/models/operations/c1apifunctionsv1functionsservicegetlockfileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Invoke

Invoke executes a function at a specific commit with the provided input data.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.Invoke" method="post" path="/api/v1/functions/{function_id}/invoke" -->
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

    res, err := s.Functions.Invoke(ctx, operations.C1APIFunctionsV1FunctionsServiceInvokeRequest{
        FunctionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceInvokeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.C1APIFunctionsV1FunctionsServiceInvokeRequest](../../pkg/models/operations/c1apifunctionsv1functionsserviceinvokerequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceInvokeResponse](../../pkg/models/operations/c1apifunctionsv1functionsserviceinvokeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommits

ListCommits retrieves the commit history

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.ListCommits" method="get" path="/api/v1/functions/{function_id}/commits" -->
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

    res, err := s.Functions.ListCommits(ctx, operations.C1APIFunctionsV1FunctionsServiceListCommitsRequest{
        FunctionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceListCommitsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIFunctionsV1FunctionsServiceListCommitsRequest](../../pkg/models/operations/c1apifunctionsv1functionsservicelistcommitsrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceListCommitsResponse](../../pkg/models/operations/c1apifunctionsv1functionsservicelistcommitsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListFunctions

List retrieves all functions with pagination

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.ListFunctions" method="get" path="/api/v1/functions" -->
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

    res, err := s.Functions.ListFunctions(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceListFunctionsResponse != nil {
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

**[*operations.C1APIFunctionsV1FunctionsServiceListFunctionsResponse](../../pkg/models/operations/c1apifunctionsv1functionsservicelistfunctionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTags

ListTags lists all tags for a function

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.ListTags" method="get" path="/api/v1/functions/{function_id}/tags" -->
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

    res, err := s.Functions.ListTags(ctx, operations.C1APIFunctionsV1FunctionsServiceListTagsRequest{
        FunctionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceListTagsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIFunctionsV1FunctionsServiceListTagsRequest](../../pkg/models/operations/c1apifunctionsv1functionsservicelisttagsrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceListTagsResponse](../../pkg/models/operations/c1apifunctionsv1functionsservicelisttagsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Test

Test runs a function's test suite in a sandboxed environment and returns the results.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.Test" method="post" path="/api/v1/functions/{function_id}/test" -->
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

    res, err := s.Functions.Test(ctx, operations.C1APIFunctionsV1FunctionsServiceTestRequest{
        FunctionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceTestResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.C1APIFunctionsV1FunctionsServiceTestRequest](../../pkg/models/operations/c1apifunctionsv1functionsservicetestrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceTestResponse](../../pkg/models/operations/c1apifunctionsv1functionsservicetestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateFunction

Update an existing function's metadata. Also the publish path: set
 function.published_commit_id and include "published_commit_id" in
 update_mask to make a commit the default runnable version.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.UpdateFunction" method="post" path="/api/v1/functions/update" -->
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

    res, err := s.Functions.UpdateFunction(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceUpdateFunctionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [shared.FunctionsServiceUpdateFunctionRequest](../../pkg/models/shared/functionsserviceupdatefunctionrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceUpdateFunctionResponse](../../pkg/models/operations/c1apifunctionsv1functionsserviceupdatefunctionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |