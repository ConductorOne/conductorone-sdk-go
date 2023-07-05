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
        Description: conductoroneapi.String("dolorum"),
        DisplayName: conductoroneapi.String("excepturi"),
        PolicySteps: map[string]shared.PolicySteps{
            "facilis": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AllowReassignment: conductoroneapi.Bool(false),
                            AppOwners: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            Assigned: conductoroneapi.Bool(false),
                            EntitlementOwners: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "delectus",
                                    "eum",
                                },
                            },
                            Group: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("non"),
                                AppID: conductoroneapi.String("eligendi"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "aliquid",
                                    "provident",
                                    "necessitatibus",
                                },
                            },
                            Manager: &shared.ManagerApproval{
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
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                            Self: &shared.SelfApproval{
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
                            Users: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "occaecati",
                                },
                            },
                        },
                        Provision: &shared.Provision{
                            Assigned: conductoroneapi.Bool(false),
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                Connector: &shared.ConnectorProvision{},
                                Delegated: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("enim"),
                                    EntitlementID: conductoroneapi.String("accusamus"),
                                },
                                Manual: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("delectus"),
                                    UserIds: []string{
                                        "provident",
                                        "nam",
                                        "id",
                                    },
                                },
                            },
                        },
                    },
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AllowReassignment: conductoroneapi.Bool(false),
                            AppOwners: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            Assigned: conductoroneapi.Bool(false),
                            EntitlementOwners: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "deleniti",
                                    "sapiente",
                                    "amet",
                                },
                            },
                            Group: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("deserunt"),
                                AppID: conductoroneapi.String("nisi"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "natus",
                                    "omnis",
                                },
                            },
                            Manager: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "perferendis",
                                    "nihil",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "distinctio",
                                    "id",
                                },
                            },
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                            Self: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "labore",
                                    "suscipit",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "nobis",
                                    "eum",
                                    "vero",
                                },
                            },
                            Users: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "architecto",
                                },
                            },
                        },
                        Provision: &shared.Provision{
                            Assigned: conductoroneapi.Bool(false),
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                Connector: &shared.ConnectorProvision{},
                                Delegated: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("magnam"),
                                    EntitlementID: conductoroneapi.String("et"),
                                },
                                Manual: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("excepturi"),
                                    UserIds: []string{
                                        "provident",
                                        "quos",
                                    },
                                },
                            },
                        },
                    },
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AllowReassignment: conductoroneapi.Bool(false),
                            AppOwners: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            Assigned: conductoroneapi.Bool(false),
                            EntitlementOwners: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "accusantium",
                                    "mollitia",
                                    "reiciendis",
                                },
                            },
                            Group: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("mollitia"),
                                AppID: conductoroneapi.String("ad"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "dolor",
                                    "necessitatibus",
                                },
                            },
                            Manager: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "nemo",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "iure",
                                },
                            },
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                            Self: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "debitis",
                                    "eius",
                                    "maxime",
                                    "deleniti",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "in",
                                    "architecto",
                                    "architecto",
                                },
                            },
                            Users: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "ullam",
                                    "expedita",
                                    "nihil",
                                    "repellat",
                                },
                            },
                        },
                        Provision: &shared.Provision{
                            Assigned: conductoroneapi.Bool(false),
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                Connector: &shared.ConnectorProvision{},
                                Delegated: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("quibusdam"),
                                    EntitlementID: conductoroneapi.String("sed"),
                                },
                                Manual: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("saepe"),
                                    UserIds: []string{
                                        "accusantium",
                                        "consequuntur",
                                        "praesentium",
                                        "natus",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "magni": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AllowReassignment: conductoroneapi.Bool(false),
                            AppOwners: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            Assigned: conductoroneapi.Bool(false),
                            EntitlementOwners: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "illum",
                                    "pariatur",
                                    "maxime",
                                    "ea",
                                },
                            },
                            Group: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("excepturi"),
                                AppID: conductoroneapi.String("odit"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "accusantium",
                                    "ab",
                                },
                            },
                            Manager: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "quidem",
                                    "ipsam",
                                    "voluptate",
                                    "autem",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "eaque",
                                    "pariatur",
                                    "nemo",
                                },
                            },
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                            Self: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "perferendis",
                                    "fugiat",
                                    "amet",
                                    "aut",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "corporis",
                                    "hic",
                                    "libero",
                                    "nobis",
                                },
                            },
                            Users: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "quis",
                                },
                            },
                        },
                        Provision: &shared.Provision{
                            Assigned: conductoroneapi.Bool(false),
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                Connector: &shared.ConnectorProvision{},
                                Delegated: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("totam"),
                                    EntitlementID: conductoroneapi.String("dignissimos"),
                                },
                                Manual: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("eaque"),
                                    UserIds: []string{
                                        "nesciunt",
                                        "eos",
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
        PolicyType: shared.CreatePolicyRequestPolicyTypePolicyTypeUnspecified.ToPointer(),
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
        ID: "c73d5fe9-b90c-4289-89b3-fe49a8d9cbf4",
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
        ID: "8633323f-9b77-4f3a-8100-674ebf69280d",
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
                CreatedAt: types.MustTimeFromString("2022-04-04T12:00:33.616Z"),
                DeletedAt: types.MustTimeFromString("2022-01-16T14:59:31.978Z"),
                Description: conductoroneapi.String("voluptate"),
                DisplayName: conductoroneapi.String("dolorum"),
                ID: conductoroneapi.String("89ebf737-ae42-403c-a5e6-a95d8a0d446c"),
                PolicySteps: map[string]shared.PolicySteps{
                    "qui": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "esse",
                                            "harum",
                                            "iusto",
                                            "ipsum",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("quisquam"),
                                        AppID: conductoroneapi.String("tenetur"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "tempore",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "numquam",
                                            "enim",
                                            "dolorem",
                                            "sapiente",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "nihil",
                                            "sit",
                                            "expedita",
                                        },
                                    },
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "sed",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "libero",
                                            "voluptas",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "quam",
                                            "ipsum",
                                            "incidunt",
                                        },
                                    },
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("qui"),
                                            EntitlementID: conductoroneapi.String("cupiditate"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("maxime"),
                                            UserIds: []string{
                                                "soluta",
                                                "dicta",
                                                "laborum",
                                                "totam",
                                            },
                                        },
                                    },
                                },
                            },
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "aspernatur",
                                            "dolores",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("distinctio"),
                                        AppID: conductoroneapi.String("facilis"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "quam",
                                            "molestias",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "qui",
                                            "neque",
                                            "fugit",
                                            "magni",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "sunt",
                                            "ullam",
                                        },
                                    },
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "hic",
                                            "voluptatem",
                                            "cumque",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "nobis",
                                            "et",
                                            "saepe",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "veritatis",
                                        },
                                    },
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("nobis"),
                                            EntitlementID: conductoroneapi.String("quos"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("tempore"),
                                            UserIds: []string{
                                                "aperiam",
                                                "delectus",
                                                "dolorem",
                                            },
                                        },
                                    },
                                },
                            },
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "labore",
                                            "adipisci",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("dolorum"),
                                        AppID: conductoroneapi.String("architecto"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "aut",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "itaque",
                                            "consequatur",
                                            "est",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "porro",
                                            "doloribus",
                                            "ut",
                                            "facilis",
                                        },
                                    },
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                    Self: &shared.SelfApproval{
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
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "vero",
                                            "omnis",
                                            "quis",
                                            "ipsum",
                                        },
                                    },
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("delectus"),
                                            EntitlementID: conductoroneapi.String("voluptate"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("consectetur"),
                                            UserIds: []string{
                                                "tenetur",
                                                "dignissimos",
                                                "hic",
                                                "distinctio",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    "quod": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "facilis",
                                            "vero",
                                            "ducimus",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("dolore"),
                                        AppID: conductoroneapi.String("quibusdam"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "sequi",
                                            "natus",
                                            "impedit",
                                            "aut",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "exercitationem",
                                            "nulla",
                                            "fugit",
                                            "porro",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "doloribus",
                                            "iusto",
                                            "eligendi",
                                            "ducimus",
                                        },
                                    },
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "officia",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "ipsam",
                                            "ea",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "vel",
                                        },
                                    },
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("possimus"),
                                            EntitlementID: conductoroneapi.String("magnam"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("ratione"),
                                            UserIds: []string{
                                                "laudantium",
                                                "dicta",
                                            },
                                        },
                                    },
                                },
                            },
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "maiores",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("quasi"),
                                        AppID: conductoroneapi.String("ex"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "excepturi",
                                            "voluptatibus",
                                            "nostrum",
                                            "sapiente",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "saepe",
                                            "ea",
                                            "impedit",
                                            "corporis",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "aliquid",
                                            "inventore",
                                        },
                                    },
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "ea",
                                            "quo",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "recusandae",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "minima",
                                        },
                                    },
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("eaque"),
                                            EntitlementID: conductoroneapi.String("a"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("libero"),
                                            UserIds: []string{
                                                "aut",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    "deleniti": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "fugit",
                                            "accusamus",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("inventore"),
                                        AppID: conductoroneapi.String("non"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "dolorum",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "placeat",
                                            "velit",
                                            "eum",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "nobis",
                                            "quas",
                                        },
                                    },
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "nulla",
                                            "voluptas",
                                            "libero",
                                            "quasi",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "numquam",
                                            "explicabo",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "ipsa",
                                            "molestiae",
                                            "magnam",
                                        },
                                    },
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("odio"),
                                            EntitlementID: conductoroneapi.String("eius"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("esse"),
                                            UserIds: []string{
                                                "rem",
                                                "fuga",
                                            },
                                        },
                                    },
                                },
                            },
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "quidem",
                                            "fugiat",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("ut"),
                                        AppID: conductoroneapi.String("eum"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "assumenda",
                                            "eos",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "quisquam",
                                            "veritatis",
                                            "ipsa",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "quidem",
                                            "neque",
                                            "quo",
                                        },
                                    },
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "quo",
                                            "fuga",
                                            "eius",
                                            "eos",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "ab",
                                            "cupiditate",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "tempora",
                                        },
                                    },
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("debitis"),
                                            EntitlementID: conductoroneapi.String("ipsam"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("aspernatur"),
                                            UserIds: []string{
                                                "quo",
                                            },
                                        },
                                    },
                                },
                            },
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "recusandae",
                                            "aperiam",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("distinctio"),
                                        AppID: conductoroneapi.String("quod"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "inventore",
                                            "nihil",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "accusamus",
                                            "aliquam",
                                            "odio",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "commodi",
                                            "sapiente",
                                            "dolores",
                                        },
                                    },
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "molestiae",
                                            "accusantium",
                                            "porro",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "quas",
                                            "praesentium",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "deleniti",
                                        },
                                    },
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("fugit"),
                                            EntitlementID: conductoroneapi.String("fuga"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("mollitia"),
                                            UserIds: []string{
                                                "atque",
                                                "explicabo",
                                            },
                                        },
                                    },
                                },
                            },
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "nisi",
                                            "fugit",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("sapiente"),
                                        AppID: conductoroneapi.String("consequuntur"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "explicabo",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "occaecati",
                                            "atque",
                                            "et",
                                            "esse",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "accusamus",
                                            "veritatis",
                                            "esse",
                                            "quod",
                                        },
                                    },
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "vero",
                                            "aliquid",
                                            "quasi",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "vel",
                                            "harum",
                                            "molestiae",
                                            "rerum",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "minima",
                                            "distinctio",
                                            "eligendi",
                                        },
                                    },
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("sit"),
                                            EntitlementID: conductoroneapi.String("culpa"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("tempore"),
                                            UserIds: []string{
                                                "cumque",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    "consequuntur": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "quaerat",
                                            "sapiente",
                                            "consectetur",
                                            "esse",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("blanditiis"),
                                        AppID: conductoroneapi.String("provident"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "nulla",
                                            "quas",
                                            "esse",
                                            "quasi",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "error",
                                            "sint",
                                            "pariatur",
                                            "possimus",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "eveniet",
                                        },
                                    },
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "facere",
                                            "veritatis",
                                            "consequuntur",
                                            "quasi",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "culpa",
                                            "aliquid",
                                            "tenetur",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "earum",
                                        },
                                    },
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("vel"),
                                            EntitlementID: conductoroneapi.String("in"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("eius"),
                                            UserIds: []string{
                                                "illum",
                                                "soluta",
                                                "accusantium",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
                PolicyType: shared.PolicyPolicyTypePolicyTypeGrant.ToPointer(),
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
                    shared.PolicyPostActions{
                        CertifyRemediateImmediately: conductoroneapi.Bool(false),
                    },
                },
                ReassignTasksToDelegates: conductoroneapi.Bool(false),
                SystemBuiltin: conductoroneapi.Bool(false),
                UpdatedAt: types.MustTimeFromString("2022-08-24T06:58:07.315Z"),
            },
            UpdateMask: conductoroneapi.String("reprehenderit"),
        },
        ID: "56082d68-ea19-4f1d-9705-1339d08086a1",
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

