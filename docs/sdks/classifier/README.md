# Classifier

## Overview

### Available Operations

* [CreateClassifier](#createclassifier) - Create Classifier
* [CreateClassifierBinding](#createclassifierbinding) - Create Classifier Binding
* [DeleteClassifier](#deleteclassifier) - Delete Classifier
* [DeleteClassifierBinding](#deleteclassifierbinding) - Delete Classifier Binding
* [GetClassifier](#getclassifier) - Get Classifier
* [ListClassifierBindings](#listclassifierbindings) - List Classifier Bindings
* [ListClassifiers](#listclassifiers) - List Classifiers
* [UpdateClassifier](#updateclassifier) - Update Classifier

## CreateClassifier

Create creates a classifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.ClassifierService.CreateClassifier" method="post" path="/api/v1/classifiers" -->
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

    res, err := s.Classifier.CreateClassifier(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateClassifierResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.CreateClassifierRequest](../../pkg/models/shared/createclassifierrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.C1APIAiGovernanceV1ClassifierServiceCreateClassifierResponse](../../pkg/models/operations/c1apiaigovernancev1classifierservicecreateclassifierresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateClassifierBinding

Create attaches a classifier to an enforcement target.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.ClassifierService.CreateClassifierBinding" method="post" path="/api/v1/classifier_bindings" -->
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

    res, err := s.Classifier.CreateClassifierBinding(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateClassifierBindingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [shared.CreateClassifierBindingRequest](../../pkg/models/shared/createclassifierbindingrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.C1APIAiGovernanceV1ClassifierServiceCreateClassifierBindingResponse](../../pkg/models/operations/c1apiaigovernancev1classifierservicecreateclassifierbindingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteClassifier

Delete removes a classifier by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.ClassifierService.DeleteClassifier" method="post" path="/api/v1/classifiers/{id}/delete" -->
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

    res, err := s.Classifier.DeleteClassifier(ctx, operations.C1APIAiGovernanceV1ClassifierServiceDeleteClassifierRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteClassifierResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.C1APIAiGovernanceV1ClassifierServiceDeleteClassifierRequest](../../pkg/models/operations/c1apiaigovernancev1classifierservicedeleteclassifierrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.C1APIAiGovernanceV1ClassifierServiceDeleteClassifierResponse](../../pkg/models/operations/c1apiaigovernancev1classifierservicedeleteclassifierresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteClassifierBinding

Delete removes a classifier binding by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.ClassifierService.DeleteClassifierBinding" method="post" path="/api/v1/classifier_bindings/{id}/delete" -->
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

    res, err := s.Classifier.DeleteClassifierBinding(ctx, operations.C1APIAiGovernanceV1ClassifierServiceDeleteClassifierBindingRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteClassifierBindingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                          | Type                                                                                                                                                                               | Required                                                                                                                                                                           | Description                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                              | :heavy_check_mark:                                                                                                                                                                 | The context to use for the request.                                                                                                                                                |
| `request`                                                                                                                                                                          | [operations.C1APIAiGovernanceV1ClassifierServiceDeleteClassifierBindingRequest](../../pkg/models/operations/c1apiaigovernancev1classifierservicedeleteclassifierbindingrequest.md) | :heavy_check_mark:                                                                                                                                                                 | The request object to use for the request.                                                                                                                                         |
| `opts`                                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                 | The options for this request.                                                                                                                                                      |

### Response

**[*operations.C1APIAiGovernanceV1ClassifierServiceDeleteClassifierBindingResponse](../../pkg/models/operations/c1apiaigovernancev1classifierservicedeleteclassifierbindingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetClassifier

Get returns a classifier by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.ClassifierService.GetClassifier" method="get" path="/api/v1/classifiers/{id}" -->
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

    res, err := s.Classifier.GetClassifier(ctx, operations.C1APIAiGovernanceV1ClassifierServiceGetClassifierRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetClassifierResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.C1APIAiGovernanceV1ClassifierServiceGetClassifierRequest](../../pkg/models/operations/c1apiaigovernancev1classifierservicegetclassifierrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |
| `opts`                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                   | :heavy_minus_sign:                                                                                                                                             | The options for this request.                                                                                                                                  |

### Response

**[*operations.C1APIAiGovernanceV1ClassifierServiceGetClassifierResponse](../../pkg/models/operations/c1apiaigovernancev1classifierservicegetclassifierresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListClassifierBindings

List returns classifier bindings for the tenant, paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.ClassifierService.ListClassifierBindings" method="get" path="/api/v1/classifier_bindings" -->
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

    res, err := s.Classifier.ListClassifierBindings(ctx, operations.C1APIAiGovernanceV1ClassifierServiceListClassifierBindingsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListClassifierBindingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `request`                                                                                                                                                                        | [operations.C1APIAiGovernanceV1ClassifierServiceListClassifierBindingsRequest](../../pkg/models/operations/c1apiaigovernancev1classifierservicelistclassifierbindingsrequest.md) | :heavy_check_mark:                                                                                                                                                               | The request object to use for the request.                                                                                                                                       |
| `opts`                                                                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                     | :heavy_minus_sign:                                                                                                                                                               | The options for this request.                                                                                                                                                    |

### Response

**[*operations.C1APIAiGovernanceV1ClassifierServiceListClassifierBindingsResponse](../../pkg/models/operations/c1apiaigovernancev1classifierservicelistclassifierbindingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListClassifiers

List returns classifiers for the tenant, paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.ClassifierService.ListClassifiers" method="get" path="/api/v1/classifiers" -->
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

    res, err := s.Classifier.ListClassifiers(ctx, operations.C1APIAiGovernanceV1ClassifierServiceListClassifiersRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListClassifiersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.C1APIAiGovernanceV1ClassifierServiceListClassifiersRequest](../../pkg/models/operations/c1apiaigovernancev1classifierservicelistclassifiersrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.C1APIAiGovernanceV1ClassifierServiceListClassifiersResponse](../../pkg/models/operations/c1apiaigovernancev1classifierservicelistclassifiersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateClassifier

Update modifies a classifier's mutable fields.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.ClassifierService.UpdateClassifier" method="post" path="/api/v1/classifiers/{id}" -->
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

    res, err := s.Classifier.UpdateClassifier(ctx, operations.C1APIAiGovernanceV1ClassifierServiceUpdateClassifierRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateClassifierResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.C1APIAiGovernanceV1ClassifierServiceUpdateClassifierRequest](../../pkg/models/operations/c1apiaigovernancev1classifierserviceupdateclassifierrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.C1APIAiGovernanceV1ClassifierServiceUpdateClassifierResponse](../../pkg/models/operations/c1apiaigovernancev1classifierserviceupdateclassifierresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |