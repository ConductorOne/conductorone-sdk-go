# Policies

### Available Operations

* [Create](#create) - Invokes the c1.api.policy.v1.Policies.Create method.
* [Delete](#delete) - Invokes the c1.api.policy.v1.Policies.Delete method.
* [Get](#get) - Invokes the c1.api.policy.v1.Policies.Get method.
* [List](#list) - Invokes the c1.api.policy.v1.Policies.List method.

## Create

Invokes the c1.api.policy.v1.Policies.Create method.

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
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Policies.Create(ctx, shared.CreatePolicyRequest{
        Description: conductoroneapi.String("rem"),
        DisplayName: conductoroneapi.String("voluptates"),
        PolicyType: shared.CreatePolicyRequestPolicyTypePolicyTypeUnspecified.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Policy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.CreatePolicyRequest](../../models/shared/createpolicyrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.C1APIPolicyV1PoliciesCreateResponse](../../models/operations/c1apipolicyv1policiescreateresponse.md), error**


## Delete

Invokes the c1.api.policy.v1.Policies.Delete method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Policies.Delete(ctx, operations.C1APIPolicyV1PoliciesDeleteRequest{
        DeletePolicyRequest: &shared.DeletePolicyRequest{},
        ID: "e91e450a-d2ab-4d44-a698-02d502a94bb4",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletePolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.C1APIPolicyV1PoliciesDeleteRequest](../../models/operations/c1apipolicyv1policiesdeleterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.C1APIPolicyV1PoliciesDeleteResponse](../../models/operations/c1apipolicyv1policiesdeleteresponse.md), error**


## Get

Invokes the c1.api.policy.v1.Policies.Get method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Policies.Get(ctx, operations.C1APIPolicyV1PoliciesGetRequest{
        ID: "f63c969e-9a3e-4fa7-bdfb-14cd66ae395e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Policy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.C1APIPolicyV1PoliciesGetRequest](../../models/operations/c1apipolicyv1policiesgetrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.C1APIPolicyV1PoliciesGetResponse](../../models/operations/c1apipolicyv1policiesgetresponse.md), error**


## List

Invokes the c1.api.policy.v1.Policies.List method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Policies.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APIPolicyV1PoliciesListResponse](../../models/operations/c1apipolicyv1policieslistresponse.md), error**

