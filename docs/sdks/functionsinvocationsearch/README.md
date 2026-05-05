# FunctionsInvocationSearch

## Overview

### Available Operations

* [Search](#search) - Search

## Search

Search searches for function invocations with filtering and ordering support

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsInvocationSearchService.Search" method="post" path="/api/v1/functions/{function_id}/invocations/search" -->
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

    res, err := s.FunctionsInvocationSearch.Search(ctx, operations.C1APIFunctionsV1FunctionsInvocationSearchServiceSearchRequest{
        FunctionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsInvocationSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [operations.C1APIFunctionsV1FunctionsInvocationSearchServiceSearchRequest](../../pkg/models/operations/c1apifunctionsv1functionsinvocationsearchservicesearchrequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |
| `opts`                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                                       | The options for this request.                                                                                                                                            |

### Response

**[*operations.C1APIFunctionsV1FunctionsInvocationSearchServiceSearchResponse](../../pkg/models/operations/c1apifunctionsv1functionsinvocationsearchservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |