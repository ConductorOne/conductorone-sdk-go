# Attributes
(*Attributes*)

## Overview

### Available Operations

* [CreateAttributeValue](#createattributevalue) - Create Attribute Value
* [CreateComplianceFrameworkAttributeValue](#createcomplianceframeworkattributevalue) - Create Compliance Framework Attribute Value
* [CreateRiskLevelAttributeValue](#createrisklevelattributevalue) - Create Risk Level Attribute Value
* [DeleteAttributeValue](#deleteattributevalue) - Delete Attribute Value
* [DeleteComplianceFrameworkAttributeValue](#deletecomplianceframeworkattributevalue) - Delete Compliance Framework Attribute Value
* [DeleteRiskLevelAttributeValue](#deleterisklevelattributevalue) - Delete Risk Level Attribute Value
* [GetAttributeValue](#getattributevalue) - Get Attribute Value
* [GetComplianceFrameworkAttributeValue](#getcomplianceframeworkattributevalue) - Get Compliance Framework Attribute Value
* [GetRiskLevelAttributeValue](#getrisklevelattributevalue) - Get Risk Level Attribute Value
* [ListAttributeTypes](#listattributetypes) - List Attribute Types
* [ListAttributeValues](#listattributevalues) - List Attribute Values
* [ListComplianceFrameworks](#listcomplianceframeworks) - List Compliance Frameworks
* [ListRiskLevels](#listrisklevels) - List Risk Levels

## CreateAttributeValue

Create a new attribute value.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.CreateAttributeValue" method="post" path="/api/v1/attributes" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.CreateAttributeValue(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.CreateAttributeValueRequest](../../pkg/models/shared/createattributevaluerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.C1APIAttributeV1AttributesCreateAttributeValueResponse](../../pkg/models/operations/c1apiattributev1attributescreateattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateComplianceFrameworkAttributeValue

Create a compliance framework value.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.CreateComplianceFrameworkAttributeValue" method="post" path="/api/v1/attributes/compliance_frameworks" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.CreateComplianceFrameworkAttributeValue(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateComplianceFrameworkAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [shared.CreateComplianceFrameworkAttributeValueRequest](../../pkg/models/shared/createcomplianceframeworkattributevaluerequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.C1APIAttributeV1AttributesCreateComplianceFrameworkAttributeValueResponse](../../pkg/models/operations/c1apiattributev1attributescreatecomplianceframeworkattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateRiskLevelAttributeValue

Create a risk level attribute.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.CreateRiskLevelAttributeValue" method="post" path="/api/v1/attributes/risk_levels" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.CreateRiskLevelAttributeValue(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateRiskLevelAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [shared.CreateRiskLevelAttributeValueRequest](../../pkg/models/shared/createrisklevelattributevaluerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.C1APIAttributeV1AttributesCreateRiskLevelAttributeValueResponse](../../pkg/models/operations/c1apiattributev1attributescreaterisklevelattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteAttributeValue

Delete an attribute value by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.DeleteAttributeValue" method="delete" path="/api/v1/attribute/{id}" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.DeleteAttributeValue(ctx, operations.C1APIAttributeV1AttributesDeleteAttributeValueRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.C1APIAttributeV1AttributesDeleteAttributeValueRequest](../../pkg/models/operations/c1apiattributev1attributesdeleteattributevaluerequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.C1APIAttributeV1AttributesDeleteAttributeValueResponse](../../pkg/models/operations/c1apiattributev1attributesdeleteattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteComplianceFrameworkAttributeValue

Delete an attribute value by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.DeleteComplianceFrameworkAttributeValue" method="delete" path="/api/v1/attributes/compliance_frameworks/{id}" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.DeleteComplianceFrameworkAttributeValue(ctx, operations.C1APIAttributeV1AttributesDeleteComplianceFrameworkAttributeValueRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteComplianceFrameworkAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                      | Type                                                                                                                                                                                           | Required                                                                                                                                                                                       | Description                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                             | The context to use for the request.                                                                                                                                                            |
| `request`                                                                                                                                                                                      | [operations.C1APIAttributeV1AttributesDeleteComplianceFrameworkAttributeValueRequest](../../pkg/models/operations/c1apiattributev1attributesdeletecomplianceframeworkattributevaluerequest.md) | :heavy_check_mark:                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                     |
| `opts`                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                             | The options for this request.                                                                                                                                                                  |

### Response

**[*operations.C1APIAttributeV1AttributesDeleteComplianceFrameworkAttributeValueResponse](../../pkg/models/operations/c1apiattributev1attributesdeletecomplianceframeworkattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteRiskLevelAttributeValue

Delete a risk level attribute value by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.DeleteRiskLevelAttributeValue" method="delete" path="/api/v1/attributes/risk_levels/{id}" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.DeleteRiskLevelAttributeValue(ctx, operations.C1APIAttributeV1AttributesDeleteRiskLevelAttributeValueRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteRiskLevelAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.C1APIAttributeV1AttributesDeleteRiskLevelAttributeValueRequest](../../pkg/models/operations/c1apiattributev1attributesdeleterisklevelattributevaluerequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |
| `opts`                                                                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                         | The options for this request.                                                                                                                                              |

### Response

**[*operations.C1APIAttributeV1AttributesDeleteRiskLevelAttributeValueResponse](../../pkg/models/operations/c1apiattributev1attributesdeleterisklevelattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAttributeValue

Get an attribute value by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.GetAttributeValue" method="get" path="/api/v1/attributes/{id}" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.GetAttributeValue(ctx, operations.C1APIAttributeV1AttributesGetAttributeValueRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIAttributeV1AttributesGetAttributeValueRequest](../../pkg/models/operations/c1apiattributev1attributesgetattributevaluerequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.C1APIAttributeV1AttributesGetAttributeValueResponse](../../pkg/models/operations/c1apiattributev1attributesgetattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComplianceFrameworkAttributeValue

Get an attribute value by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.GetComplianceFrameworkAttributeValue" method="get" path="/api/v1/attributes/compliance_frameworks/{id}" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.GetComplianceFrameworkAttributeValue(ctx, operations.C1APIAttributeV1AttributesGetComplianceFrameworkAttributeValueRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetComplianceFrameworkAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                | Type                                                                                                                                                                                     | Required                                                                                                                                                                                 | Description                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                       | The context to use for the request.                                                                                                                                                      |
| `request`                                                                                                                                                                                | [operations.C1APIAttributeV1AttributesGetComplianceFrameworkAttributeValueRequest](../../pkg/models/operations/c1apiattributev1attributesgetcomplianceframeworkattributevaluerequest.md) | :heavy_check_mark:                                                                                                                                                                       | The request object to use for the request.                                                                                                                                               |
| `opts`                                                                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                       | The options for this request.                                                                                                                                                            |

### Response

**[*operations.C1APIAttributeV1AttributesGetComplianceFrameworkAttributeValueResponse](../../pkg/models/operations/c1apiattributev1attributesgetcomplianceframeworkattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetRiskLevelAttributeValue

Get a risk level attribute value by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.GetRiskLevelAttributeValue" method="get" path="/api/v1/attributes/risk_levels/{id}" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.GetRiskLevelAttributeValue(ctx, operations.C1APIAttributeV1AttributesGetRiskLevelAttributeValueRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetRiskLevelAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.C1APIAttributeV1AttributesGetRiskLevelAttributeValueRequest](../../pkg/models/operations/c1apiattributev1attributesgetrisklevelattributevaluerequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.C1APIAttributeV1AttributesGetRiskLevelAttributeValueResponse](../../pkg/models/operations/c1apiattributev1attributesgetrisklevelattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAttributeTypes

List all attribute types.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.ListAttributeTypes" method="get" path="/api/v1/attributes/types" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.ListAttributeTypes(ctx, operations.C1APIAttributeV1AttributesListAttributeTypesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAttributeTypesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.C1APIAttributeV1AttributesListAttributeTypesRequest](../../pkg/models/operations/c1apiattributev1attributeslistattributetypesrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.C1APIAttributeV1AttributesListAttributeTypesResponse](../../pkg/models/operations/c1apiattributev1attributeslistattributetypesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAttributeValues

List all attribute values for a given attribute type.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.ListAttributeValues" method="get" path="/api/v1/attributes/types/{attribute_type_id}/values" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.ListAttributeValues(ctx, operations.C1APIAttributeV1AttributesListAttributeValuesRequest{
        AttributeTypeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAttributeValuesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.C1APIAttributeV1AttributesListAttributeValuesRequest](../../pkg/models/operations/c1apiattributev1attributeslistattributevaluesrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.C1APIAttributeV1AttributesListAttributeValuesResponse](../../pkg/models/operations/c1apiattributev1attributeslistattributevaluesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComplianceFrameworks

Invokes the c1.api.attribute.v1.Attributes.ListComplianceFrameworks method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.ListComplianceFrameworks" method="get" path="/api/v1/attributes/compliance_frameworks" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.ListComplianceFrameworks(ctx, operations.C1APIAttributeV1AttributesListComplianceFrameworksRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListComplianceFrameworksResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.C1APIAttributeV1AttributesListComplianceFrameworksRequest](../../pkg/models/operations/c1apiattributev1attributeslistcomplianceframeworksrequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |
| `opts`                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |

### Response

**[*operations.C1APIAttributeV1AttributesListComplianceFrameworksResponse](../../pkg/models/operations/c1apiattributev1attributeslistcomplianceframeworksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListRiskLevels

Invokes the c1.api.attribute.v1.Attributes.ListRiskLevels method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.ListRiskLevels" method="get" path="/api/v1/attributes/risk_levels" -->
```go
package main

import(
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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

    res, err := s.Attributes.ListRiskLevels(ctx, operations.C1APIAttributeV1AttributesListRiskLevelsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListRiskLevelsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.C1APIAttributeV1AttributesListRiskLevelsRequest](../../pkg/models/operations/c1apiattributev1attributeslistrisklevelsrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.C1APIAttributeV1AttributesListRiskLevelsResponse](../../pkg/models/operations/c1apiattributev1attributeslistrisklevelsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |