# Policies

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List
* [Update](#update) - Update

## Create

Create a policy.

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
    res, err := s.Policies.Create(ctx, shared.CreatePolicyRequestInput{
        Description: conductoroneapi.String("nisi"),
        DisplayName: conductoroneapi.String("fugit"),
        PolicySteps: map[string]shared.PolicyStepsInput{
            "sapiente": shared.PolicyStepsInput{
                Steps: []shared.PolicyStepInput{
                    shared.PolicyStepInput{
                        Approval: &shared.ApprovalInput{
                            AppGroupApproval: &shared.AppGroupApprovalInput{},
                            AppOwnerApproval: &shared.AppOwnerApprovalInput{},
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApprovalInput{},
                            ManagerApproval: &shared.ManagerApprovalInput{},
                            SelfApproval: &shared.SelfApprovalInput{},
                            UserApproval: &shared.UserApprovalInput{},
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("consequuntur"),
                                    EntitlementID: conductoroneapi.String("ratione"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("explicabo"),
                                    UserIds: []string{
                                        "saepe",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("occaecati"),
                                AppID: conductoroneapi.String("atque"),
                                AppUserID: conductoroneapi.String("et"),
                                GrantDuration: conductoroneapi.String("esse"),
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
        },
        PolicyType: shared.CreatePolicyRequestPolicyTypePolicyTypeProvision.ToPointer(),
        PostActions: []shared.PolicyPostActions{
            shared.PolicyPostActions{
                CertifyRemediateImmediately: conductoroneapi.Bool(false),
            },
        },
        ReassignTasksToDelegates: conductoroneapi.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.CreatePolicyRequestInput](../../models/shared/createpolicyrequestinput.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.C1APIPolicyV1PoliciesCreateResponse](../../models/operations/c1apipolicyv1policiescreateresponse.md), error**


## Delete

Delete a policy by ID.

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
    res, err := s.Policies.Delete(ctx, operations.C1APIPolicyV1PoliciesDeleteRequest{
        DeletePolicyRequest: &shared.DeletePolicyRequest{},
        ID: "e17cbe61-e6b7-4b95-bc0a-b3c20c4f3789",
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

Get a policy by ID.

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
    res, err := s.Policies.Get(ctx, operations.C1APIPolicyV1PoliciesGetRequest{
        ID: "fd871f99-dd2e-4fd1-a1aa-6f1e674bdb04",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetPolicyResponse != nil {
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

List policies.

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
    res, err := s.Policies.List(ctx, operations.C1APIPolicyV1PoliciesListRequest{
        PageSize: conductoroneapi.Float64(9589.83),
        PageToken: conductoroneapi.String("dicta"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.C1APIPolicyV1PoliciesListRequest](../../models/operations/c1apipolicyv1policieslistrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.C1APIPolicyV1PoliciesListResponse](../../models/operations/c1apipolicyv1policieslistresponse.md), error**


## Update

Update a policy by providing a policy object and an update mask.

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
    res, err := s.Policies.Update(ctx, operations.C1APIPolicyV1PoliciesUpdateRequest{
        UpdatePolicyRequestInput: &shared.UpdatePolicyRequestInput{
            Policy: &shared.PolicyInput{
                Description: conductoroneapi.String("ullam"),
                DisplayName: conductoroneapi.String("reprehenderit"),
                PolicySteps: map[string]shared.PolicyStepsInput{
                    "ullam": shared.PolicyStepsInput{
                        Steps: []shared.PolicyStepInput{
                            shared.PolicyStepInput{
                                Approval: &shared.ApprovalInput{
                                    AppGroupApproval: &shared.AppGroupApprovalInput{},
                                    AppOwnerApproval: &shared.AppOwnerApprovalInput{},
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApprovalInput{},
                                    ManagerApproval: &shared.ManagerApprovalInput{},
                                    SelfApproval: &shared.SelfApprovalInput{},
                                    UserApproval: &shared.UserApprovalInput{},
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("nisi"),
                                            EntitlementID: conductoroneapi.String("aut"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("voluptatum"),
                                            UserIds: []string{
                                                "qui",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("quibusdam"),
                                        AppID: conductoroneapi.String("ex"),
                                        AppUserID: conductoroneapi.String("deleniti"),
                                        GrantDuration: conductoroneapi.String("itaque"),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                        },
                    },
                },
                PolicyType: shared.PolicyPolicyTypePolicyTypeAccessRequest.ToPointer(),
                PostActions: []shared.PolicyPostActions{
                    shared.PolicyPostActions{
                        CertifyRemediateImmediately: conductoroneapi.Bool(false),
                    },
                },
                ReassignTasksToDelegates: conductoroneapi.Bool(false),
            },
            UpdateMask: conductoroneapi.String("architecto"),
        },
        ID: "9f1d1705-1339-4d08-886a-1840394c2607",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdatePolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.C1APIPolicyV1PoliciesUpdateRequest](../../models/operations/c1apipolicyv1policiesupdaterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.C1APIPolicyV1PoliciesUpdateResponse](../../models/operations/c1apipolicyv1policiesupdateresponse.md), error**

