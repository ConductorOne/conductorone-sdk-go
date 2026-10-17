# ClassifierTemplate

## Overview

### Available Operations

* [AddClassifierRuleFromTemplate](#addclassifierrulefromtemplate) - Add Classifier Rule From Template
* [GetClassifierTemplate](#getclassifiertemplate) - Get Classifier Template
* [InstantiateClassifierTemplate](#instantiateclassifiertemplate) - Instantiate Classifier Template
* [ListClassifierTemplates](#listclassifiertemplates) - List Classifier Templates

## AddClassifierRuleFromTemplate

Add inserts a template rule into a tenant-owned classifier at insert_index.
 It resolves only that rule's required tool gates and hooks, rewrites their
 references, assigns a new rule ID, and records the source template.
 Grant-policy parameters are validated only for the added rule.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.ClassifierTemplateService.AddClassifierRuleFromTemplate" method="post" path="/api/v1/classifiers/{classifier_id}/rules/from_template" -->
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

    res, err := s.ClassifierTemplate.AddClassifierRuleFromTemplate(ctx, operations.C1APIAiGovernanceV1ClassifierTemplateServiceAddClassifierRuleFromTemplateRequest{
        ClassifierID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AddClassifierRuleFromTemplateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                      | Type                                                                                                                                                                                                           | Required                                                                                                                                                                                                       | Description                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                            |
| `request`                                                                                                                                                                                                      | [operations.C1APIAiGovernanceV1ClassifierTemplateServiceAddClassifierRuleFromTemplateRequest](../../pkg/models/operations/c1apiaigovernancev1classifiertemplateserviceaddclassifierrulefromtemplaterequest.md) | :heavy_check_mark:                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                             | The options for this request.                                                                                                                                                                                  |

### Response

**[*operations.C1APIAiGovernanceV1ClassifierTemplateServiceAddClassifierRuleFromTemplateResponse](../../pkg/models/operations/c1apiaigovernancev1classifiertemplateserviceaddclassifierrulefromtemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetClassifierTemplate

Get returns a classifier template revision. An empty template_version
 selects the latest revision.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.ClassifierTemplateService.GetClassifierTemplate" method="get" path="/api/v1/classifier_templates/{id}" -->
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

    res, err := s.ClassifierTemplate.GetClassifierTemplate(ctx, operations.C1APIAiGovernanceV1ClassifierTemplateServiceGetClassifierTemplateRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetClassifierTemplateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                      | Type                                                                                                                                                                                           | Required                                                                                                                                                                                       | Description                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                             | The context to use for the request.                                                                                                                                                            |
| `request`                                                                                                                                                                                      | [operations.C1APIAiGovernanceV1ClassifierTemplateServiceGetClassifierTemplateRequest](../../pkg/models/operations/c1apiaigovernancev1classifiertemplateservicegetclassifiertemplaterequest.md) | :heavy_check_mark:                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                     |
| `opts`                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                             | The options for this request.                                                                                                                                                                  |

### Response

**[*operations.C1APIAiGovernanceV1ClassifierTemplateServiceGetClassifierTemplateResponse](../../pkg/models/operations/c1apiaigovernancev1classifiertemplateservicegetclassifiertemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## InstantiateClassifierTemplate

Instantiate creates a tenant-owned, editable classifier from a template
 revision. It binds supplied tenant parameters, resolves required tool gates
 and hooks, rewrites their references, and records the source template.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.ClassifierTemplateService.InstantiateClassifierTemplate" method="post" path="/api/v1/classifier_templates/{template_id}/instantiate" -->
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

    res, err := s.ClassifierTemplate.InstantiateClassifierTemplate(ctx, operations.C1APIAiGovernanceV1ClassifierTemplateServiceInstantiateClassifierTemplateRequest{
        TemplateID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.InstantiateClassifierTemplateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                      | Type                                                                                                                                                                                                           | Required                                                                                                                                                                                                       | Description                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                            |
| `request`                                                                                                                                                                                                      | [operations.C1APIAiGovernanceV1ClassifierTemplateServiceInstantiateClassifierTemplateRequest](../../pkg/models/operations/c1apiaigovernancev1classifiertemplateserviceinstantiateclassifiertemplaterequest.md) | :heavy_check_mark:                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                             | The options for this request.                                                                                                                                                                                  |

### Response

**[*operations.C1APIAiGovernanceV1ClassifierTemplateServiceInstantiateClassifierTemplateResponse](../../pkg/models/operations/c1apiaigovernancev1classifiertemplateserviceinstantiateclassifiertemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListClassifierTemplates

List returns classifier templates for the tenant, paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.ClassifierTemplateService.ListClassifierTemplates" method="post" path="/api/v1/classifier_templates/list" -->
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

    res, err := s.ClassifierTemplate.ListClassifierTemplates(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListClassifierTemplatesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [shared.ListClassifierTemplatesRequest](../../pkg/models/shared/listclassifiertemplatesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.C1APIAiGovernanceV1ClassifierTemplateServiceListClassifierTemplatesResponse](../../pkg/models/operations/c1apiaigovernancev1classifiertemplateservicelistclassifiertemplatesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |