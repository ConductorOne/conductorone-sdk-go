# APIKeyUsage

## Overview

### Available Operations

* [BatchGet](#batchget) - Batch Get

## BatchGet

BatchGet returns the live usage of each requested API key the caller may view. Users may view their own personal API keys, device clients and connected apps, and the API keys of service principals they may view. Rate limit administrators may view every key. Keys the caller may not view, and keys that do not exist, are omitted from the response.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ratelimit.v1.APIKeyUsageService.BatchGet" method="post" path="/api/v1/rate_limits/api_keys/usage" -->
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

    res, err := s.APIKeyUsage.BatchGet(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.APIKeyUsageServiceBatchGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [shared.APIKeyUsageServiceBatchGetRequest](../../pkg/models/shared/apikeyusageservicebatchgetrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.C1APIRatelimitV1APIKeyUsageServiceBatchGetResponse](../../pkg/models/operations/c1apiratelimitv1apikeyusageservicebatchgetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |