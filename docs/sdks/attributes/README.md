# Attributes

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
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Attributes.CreateAttributeValue(ctx, shared.CreateAttributeValueRequest{
        AttributeTypeID: conductoroneapi.String("totam"),
        Value: conductoroneapi.String("similique"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.CreateAttributeValueRequest](../../models/shared/createattributevaluerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.C1APIAttributeV1AttributesCreateAttributeValueResponse](../../models/operations/c1apiattributev1attributescreateattributevalueresponse.md), error**


## DeleteAttributeValue

Delete an attribute value by id.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Attributes.DeleteAttributeValue(ctx, operations.C1APIAttributeV1AttributesDeleteAttributeValueRequest{
        DeleteAttributeValueRequest: &shared.DeleteAttributeValueRequest{},
        ID: "0d446ce2-af7a-473c-b3be-453f870b326b",
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

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.C1APIAttributeV1AttributesDeleteAttributeValueRequest](../../models/operations/c1apiattributev1attributesdeleteattributevaluerequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.C1APIAttributeV1AttributesDeleteAttributeValueResponse](../../models/operations/c1apiattributev1attributesdeleteattributevalueresponse.md), error**


## GetAttributeValue

Get an attribute value by id.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Attributes.GetAttributeValue(ctx, operations.C1APIAttributeV1AttributesGetAttributeValueRequest{
        ID: "5a73429c-db1a-4842-abb6-79d2322715bf",
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

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAttributeV1AttributesGetAttributeValueRequest](../../models/operations/c1apiattributev1attributesgetattributevaluerequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |


### Response

**[*operations.C1APIAttributeV1AttributesGetAttributeValueResponse](../../models/operations/c1apiattributev1attributesgetattributevalueresponse.md), error**


## ListAttributeTypes

List all attribute types.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Attributes.ListAttributeTypes(ctx, operations.C1APIAttributeV1AttributesListAttributeTypesRequest{
        PageSize: conductoroneapi.Float64(304.52),
        PageToken: conductoroneapi.String("cumque"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAttributeTypesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.C1APIAttributeV1AttributesListAttributeTypesRequest](../../models/operations/c1apiattributev1attributeslistattributetypesrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.C1APIAttributeV1AttributesListAttributeTypesResponse](../../models/operations/c1apiattributev1attributeslistattributetypesresponse.md), error**


## ListAttributeValues

List all attribute values for a given attribute type.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Attributes.ListAttributeValues(ctx, operations.C1APIAttributeV1AttributesListAttributeValuesRequest{
        AttributeTypeID: "soluta",
        PageSize: conductoroneapi.Float64(7486.64),
        PageToken: conductoroneapi.String("et"),
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

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.C1APIAttributeV1AttributesListAttributeValuesRequest](../../models/operations/c1apiattributev1attributeslistattributevaluesrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.C1APIAttributeV1AttributesListAttributeValuesResponse](../../models/operations/c1apiattributev1attributeslistattributevaluesresponse.md), error**

