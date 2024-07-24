# Attributes
(*Attributes*)

### Available Operations

* [CreateAttributeValue](#createattributevalue) - Create Attribute Value
* [DeleteAttributeValue](#deleteattributevalue) - Delete Attribute Value
* [GetAttributeValue](#getattributevalue) - Get Attribute Value
* [ListAttributeTypes](#listattributetypes) - List Attribute Types
* [ListAttributeValues](#listattributevalues) - List Attribute Values

## CreateAttributeValue

Create a new attribute value.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"os"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: os.Getenv("BEARER_AUTH"),
            Oauth: os.Getenv("OAUTH"),
        }),
    )
    var request *shared.CreateAttributeValueRequest = &shared.CreateAttributeValueRequest{}
    ctx := context.Background()
    res, err := s.Attributes.CreateAttributeValue(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteAttributeValue

Delete an attribute value by id.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"os"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: os.Getenv("BEARER_AUTH"),
            Oauth: os.Getenv("OAUTH"),
        }),
    )
    request := operations.C1APIAttributeV1AttributesDeleteAttributeValueRequest{
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Attributes.DeleteAttributeValue(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAttributeValue

Get an attribute value by id.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"os"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: os.Getenv("BEARER_AUTH"),
            Oauth: os.Getenv("OAUTH"),
        }),
    )
    request := operations.C1APIAttributeV1AttributesGetAttributeValueRequest{
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Attributes.GetAttributeValue(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListAttributeTypes

List all attribute types.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"os"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: os.Getenv("BEARER_AUTH"),
            Oauth: os.Getenv("OAUTH"),
        }),
    )
    request := operations.C1APIAttributeV1AttributesListAttributeTypesRequest{}
    ctx := context.Background()
    res, err := s.Attributes.ListAttributeTypes(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListAttributeValues

List all attribute values for a given attribute type.

### Example Usage

```go
package main

import(
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"os"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: os.Getenv("BEARER_AUTH"),
            Oauth: os.Getenv("OAUTH"),
        }),
    )
    request := operations.C1APIAttributeV1AttributesListAttributeValuesRequest{
        AttributeTypeID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Attributes.ListAttributeValues(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
