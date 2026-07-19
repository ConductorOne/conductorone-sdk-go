# TerraformExport

## Overview

### Available Operations

* [GetSchema](#getschema) - Get Schema

## GetSchema

GetSchema returns the field-by-field Terraform mapping for one C1
 API object type. Cacheable by (object_fqn, block_kind,
 provider_version).

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.terraform_export.v1.TerraformExportService.GetSchema" method="get" path="/api/v1/terraform-export/schema" -->
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

    res, err := s.TerraformExport.GetSchema(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSchemaResponse != nil {
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

**[*operations.C1APITerraformExportV1TerraformExportServiceGetSchemaResponse](../../pkg/models/operations/c1apiterraformexportv1terraformexportservicegetschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |