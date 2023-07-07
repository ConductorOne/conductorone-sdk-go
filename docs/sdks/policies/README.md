# Policies

### Available Operations

* [Create](#create) - Invokes the c1.api.policy.v1.Policies.Create method.
* [Delete](#delete) - Invokes the c1.api.policy.v1.Policies.Delete method.
* [Get](#get) - Invokes the c1.api.policy.v1.Policies.Get method.
* [List](#list) - Invokes the c1.api.policy.v1.Policies.List method.
* [Update](#update) - Invokes the c1.api.policy.v1.Policies.Update method.

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
        Description: conductoroneapi.String("excepturi"),
        DisplayName: conductoroneapi.String("tempora"),
        PolicySteps: map[string]shared.PolicySteps{
            "tempore": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("delectus"),
                                AppID: conductoroneapi.String("eum"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "eligendi",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "aliquid",
                                    "provident",
                                    "necessitatibus",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "officia",
                                    "dolor",
                                    "debitis",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "dolorum",
                                    "in",
                                    "in",
                                    "illum",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "rerum",
                                    "dicta",
                                    "magnam",
                                    "cumque",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "ea",
                                    "aliquid",
                                    "laborum",
                                    "accusamus",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "occaecati",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("enim"),
                                    EntitlementID: conductoroneapi.String("accusamus"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("delectus"),
                                    UserIds: []string{
                                        "provident",
                                        "nam",
                                        "id",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("blanditiis"),
                                AppID: conductoroneapi.String("deleniti"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "amet",
                                    "deserunt",
                                    "nisi",
                                    "vel",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "omnis",
                                    "molestiae",
                                    "perferendis",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "magnam",
                                    "distinctio",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "labore",
                                    "labore",
                                    "suscipit",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "nobis",
                                    "eum",
                                    "vero",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "architecto",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "et",
                                    "excepturi",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("ullam"),
                                    EntitlementID: conductoroneapi.String("provident"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("quos"),
                                    UserIds: []string{
                                        "accusantium",
                                        "mollitia",
                                        "reiciendis",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "mollitia": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("eum"),
                                AppID: conductoroneapi.String("dolor"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "odit",
                                    "nemo",
                                    "quasi",
                                    "iure",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "debitis",
                                    "eius",
                                    "maxime",
                                    "deleniti",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "in",
                                    "architecto",
                                    "architecto",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "ullam",
                                    "expedita",
                                    "nihil",
                                    "repellat",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "sed",
                                    "saepe",
                                    "pariatur",
                                    "accusantium",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "praesentium",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "magni",
                                    "sunt",
                                    "quo",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("illum"),
                                    EntitlementID: conductoroneapi.String("pariatur"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("maxime"),
                                    UserIds: []string{
                                        "excepturi",
                                        "odit",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("ea"),
                                AppID: conductoroneapi.String("accusantium"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "maiores",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "ipsam",
                                    "voluptate",
                                    "autem",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "eaque",
                                    "pariatur",
                                    "nemo",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "perferendis",
                                    "fugiat",
                                    "amet",
                                    "aut",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "corporis",
                                    "hic",
                                    "libero",
                                    "nobis",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "quis",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "dignissimos",
                                    "eaque",
                                    "quis",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("nesciunt"),
                                    EntitlementID: conductoroneapi.String("eos"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("perferendis"),
                                    UserIds: []string{
                                        "minus",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "quam": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("vero"),
                                AppID: conductoroneapi.String("nostrum"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "recusandae",
                                    "omnis",
                                    "facilis",
                                    "perspiciatis",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "porro",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "blanditiis",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "eaque",
                                    "occaecati",
                                    "rerum",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "asperiores",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "modi",
                                    "iste",
                                    "dolorum",
                                    "deleniti",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "provident",
                                    "nobis",
                                    "libero",
                                    "delectus",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("quaerat"),
                                    EntitlementID: conductoroneapi.String("quos"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("aliquid"),
                                    UserIds: []string{
                                        "dolorem",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
        },
        PolicyType: shared.CreatePolicyRequestPolicyTypePolicyTypeGrant.ToPointer(),
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
        ID: "3f9b77f3-a410-4067-8ebf-69280d1ba77a",
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
        ID: "89ebf737-ae42-403c-a5e6-a95d8a0d446c",
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


## Update

Invokes the c1.api.policy.v1.Policies.Update method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/types"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Policies.Update(ctx, operations.C1APIPolicyV1PoliciesUpdateRequest{
        UpdatePolicyRequest: &shared.UpdatePolicyRequest{
            Policy: &shared.Policy{
                CreatedAt: types.MustTimeFromString("2022-06-11T17:29:13.872Z"),
                DeletedAt: types.MustTimeFromString("2021-02-04T11:05:24.484Z"),
                Description: conductoroneapi.String("esse"),
                DisplayName: conductoroneapi.String("harum"),
                ID: conductoroneapi.String("73cf3be4-53f8-470b-b26b-5a73429cdb1a"),
                PolicySteps: map[string]shared.PolicySteps{
                    "incidunt": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("dolores"),
                                        AppID: conductoroneapi.String("distinctio"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "aliquid",
                                            "quam",
                                            "molestias",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "qui",
                                            "neque",
                                            "fugit",
                                            "magni",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "sunt",
                                            "ullam",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "hic",
                                            "voluptatem",
                                            "cumque",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "nobis",
                                            "et",
                                            "saepe",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "veritatis",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "quos",
                                            "tempore",
                                            "cupiditate",
                                        },
                                    },
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    Assigned: conductoroneapi.Bool(false),
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("aperiam"),
                                            EntitlementID: conductoroneapi.String("delectus"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("dolorem"),
                                            UserIds: []string{
                                                "labore",
                                                "adipisci",
                                            },
                                        },
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                        },
                    },
                    "dolorum": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("quae"),
                                        AppID: conductoroneapi.String("aut"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "itaque",
                                            "consequatur",
                                            "est",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "porro",
                                            "doloribus",
                                            "ut",
                                            "facilis",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "qui",
                                            "quae",
                                            "laudantium",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "occaecati",
                                            "voluptatibus",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "vero",
                                            "omnis",
                                            "quis",
                                            "ipsum",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "voluptate",
                                            "consectetur",
                                            "vero",
                                            "tenetur",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "hic",
                                            "distinctio",
                                        },
                                    },
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    Assigned: conductoroneapi.Bool(false),
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("quod"),
                                            EntitlementID: conductoroneapi.String("odio"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("similique"),
                                            UserIds: []string{
                                                "vero",
                                                "ducimus",
                                                "dolore",
                                            },
                                        },
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                        },
                    },
                    "quibusdam": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("sequi"),
                                        AppID: conductoroneapi.String("natus"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "aut",
                                            "voluptatibus",
                                            "exercitationem",
                                            "nulla",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "porro",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "doloribus",
                                            "iusto",
                                            "eligendi",
                                            "ducimus",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "officia",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "ipsam",
                                            "ea",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "vel",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "magnam",
                                            "ratione",
                                            "ex",
                                            "laudantium",
                                        },
                                    },
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    Assigned: conductoroneapi.Bool(false),
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("dicta"),
                                            EntitlementID: conductoroneapi.String("dolor"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("maiores"),
                                            UserIds: []string{
                                                "ex",
                                            },
                                        },
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("nulla"),
                                        AppID: conductoroneapi.String("excepturi"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "nostrum",
                                            "sapiente",
                                            "quisquam",
                                            "saepe",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "impedit",
                                            "corporis",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "aliquid",
                                            "inventore",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "ea",
                                            "quo",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "recusandae",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "minima",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "a",
                                        },
                                    },
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    Assigned: conductoroneapi.Bool(false),
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("libero"),
                                            EntitlementID: conductoroneapi.String("aut"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("aut"),
                                            UserIds: []string{
                                                "impedit",
                                                "aliquam",
                                                "fugit",
                                            },
                                        },
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("accusamus"),
                                        AppID: conductoroneapi.String("inventore"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "et",
                                            "dolorum",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "placeat",
                                            "velit",
                                            "eum",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "nobis",
                                            "quas",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "nulla",
                                            "voluptas",
                                            "libero",
                                            "quasi",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "numquam",
                                            "explicabo",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "ipsa",
                                            "molestiae",
                                            "magnam",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "eius",
                                            "esse",
                                        },
                                    },
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    Assigned: conductoroneapi.Bool(false),
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("esse"),
                                            EntitlementID: conductoroneapi.String("rem"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("fuga"),
                                            UserIds: []string{
                                                "quidem",
                                                "fugiat",
                                            },
                                        },
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("ut"),
                                        AppID: conductoroneapi.String("eum"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "assumenda",
                                            "eos",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "quisquam",
                                            "veritatis",
                                            "ipsa",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "quidem",
                                            "neque",
                                            "quo",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "quo",
                                            "fuga",
                                            "eius",
                                            "eos",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "ab",
                                            "cupiditate",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "tempora",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "ipsam",
                                            "aspernatur",
                                            "sequi",
                                            "quo",
                                        },
                                    },
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    Assigned: conductoroneapi.Bool(false),
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("esse"),
                                            EntitlementID: conductoroneapi.String("recusandae"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("aperiam"),
                                            UserIds: []string{
                                                "quod",
                                                "dignissimos",
                                                "inventore",
                                            },
                                        },
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                        },
                    },
                },
                PolicyType: shared.PolicyPolicyTypePolicyTypeRevoke.ToPointer(),
                PostActions: []shared.PolicyPostActions{
                    shared.PolicyPostActions{
                        CertifyRemediateImmediately: conductoroneapi.Bool(false),
                    },
                    shared.PolicyPostActions{
                        CertifyRemediateImmediately: conductoroneapi.Bool(false),
                    },
                    shared.PolicyPostActions{
                        CertifyRemediateImmediately: conductoroneapi.Bool(false),
                    },
                },
                ReassignTasksToDelegates: conductoroneapi.Bool(false),
                SystemBuiltin: conductoroneapi.Bool(false),
                UpdatedAt: types.MustTimeFromString("2022-01-30T01:01:49.335Z"),
            },
            UpdateMask: conductoroneapi.String("odio"),
        },
        ID: "96f2a70c-6882-482a-a482-562f222e9817",
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

