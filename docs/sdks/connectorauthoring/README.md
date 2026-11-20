# ConnectorAuthoring

## Overview

### Available Operations

* [DeleteCatalogEntry](#deletecatalogentry) - Delete Catalog Entry

## DeleteCatalogEntry

DeleteCatalogEntry deletes a custom connector catalog entry and its
 authoring work. Delete its connector installations first.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.connector_authoring.v1.ConnectorAuthoringService.DeleteCatalogEntry" method="delete" path="/api/v1/connector-authoring/catalog-entries/{catalog_id}" -->
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

    res, err := s.ConnectorAuthoring.DeleteCatalogEntry(ctx, operations.C1APIConnectorAuthoringV1ConnectorAuthoringServiceDeleteCatalogEntryRequest{
        CatalogID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorAuthoringServiceDeleteCatalogEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                            | Type                                                                                                                                                                                                 | Required                                                                                                                                                                                             | Description                                                                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                  |
| `request`                                                                                                                                                                                            | [operations.C1APIConnectorAuthoringV1ConnectorAuthoringServiceDeleteCatalogEntryRequest](../../pkg/models/operations/c1apiconnectorauthoringv1connectorauthoringservicedeletecatalogentryrequest.md) | :heavy_check_mark:                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                           |
| `opts`                                                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                   | The options for this request.                                                                                                                                                                        |

### Response

**[*operations.C1APIConnectorAuthoringV1ConnectorAuthoringServiceDeleteCatalogEntryResponse](../../pkg/models/operations/c1apiconnectorauthoringv1connectorauthoringservicedeletecatalogentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |