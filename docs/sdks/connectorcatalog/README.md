# ConnectorCatalog

## Overview

### Available Operations

* [ConfigurationSchema](#configurationschema) - Configuration Schema

## ConfigurationSchema

Invokes the c1.api.integration.connector.v1.ConnectorCatalogService.ConfigurationSchema method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.integration.connector.v1.ConnectorCatalogService.ConfigurationSchema" method="post" path="/api/v1/connectorcatalog" -->
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

    res, err := s.ConnectorCatalog.ConfigurationSchema(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorCatalogServiceConfigurationSchemaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [shared.ConnectorCatalogServiceConfigurationSchemaRequest](../../pkg/models/shared/connectorcatalogserviceconfigurationschemarequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.C1APIIntegrationConnectorV1ConnectorCatalogServiceConfigurationSchemaResponse](../../pkg/models/operations/c1apiintegrationconnectorv1connectorcatalogserviceconfigurationschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |