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
        Description: conductoroneapi.String("labore"),
        DisplayName: conductoroneapi.String("modi"),
        PolicySteps: map[string]shared.PolicySteps{
            "aliquid": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Step: &shared.PolicyStepStep{
                            Approval: &shared.Approval{
                                AllowReassignment: conductoroneapi.Bool(false),
                                Assigned: conductoroneapi.Bool(false),
                                RequireApprovalReason: conductoroneapi.Bool(false),
                                RequireReassignmentReason: conductoroneapi.Bool(false),
                                Typ: &shared.ApprovalTyp{
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "perferendis",
                                            "magni",
                                            "assumenda",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("ipsam"),
                                        AppID: conductoroneapi.String("alias"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "dolorum",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "tempora",
                                            "facilis",
                                            "tempore",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "delectus",
                                            "eum",
                                        },
                                    },
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "eligendi",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "aliquid",
                                            "provident",
                                            "necessitatibus",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "officia",
                                            "dolor",
                                            "debitis",
                                        },
                                    },
                                },
                            },
                            Provision: &shared.Provision{
                                Assigned: conductoroneapi.Bool(false),
                                ProvisionPolicy: &shared.ProvisionPolicy{
                                    Typ: &shared.ProvisionPolicyTyp{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("a"),
                                            EntitlementID: conductoroneapi.String("dolorum"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("in"),
                                            UserIds: []string{
                                                "illum",
                                                "maiores",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    shared.PolicyStep{
                        Step: &shared.PolicyStepStep{
                            Approval: &shared.Approval{
                                AllowReassignment: conductoroneapi.Bool(false),
                                Assigned: conductoroneapi.Bool(false),
                                RequireApprovalReason: conductoroneapi.Bool(false),
                                RequireReassignmentReason: conductoroneapi.Bool(false),
                                Typ: &shared.ApprovalTyp{
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "dicta",
                                            "magnam",
                                            "cumque",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("facere"),
                                        AppID: conductoroneapi.String("ea"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "laborum",
                                            "accusamus",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "occaecati",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "accusamus",
                                            "delectus",
                                        },
                                    },
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "provident",
                                            "nam",
                                            "id",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "deleniti",
                                            "sapiente",
                                            "amet",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "nisi",
                                            "vel",
                                            "natus",
                                        },
                                    },
                                },
                            },
                            Provision: &shared.Provision{
                                Assigned: conductoroneapi.Bool(false),
                                ProvisionPolicy: &shared.ProvisionPolicy{
                                    Typ: &shared.ProvisionPolicyTyp{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("omnis"),
                                            EntitlementID: conductoroneapi.String("molestiae"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("perferendis"),
                                            UserIds: []string{
                                                "magnam",
                                                "distinctio",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    shared.PolicyStep{
                        Step: &shared.PolicyStepStep{
                            Approval: &shared.Approval{
                                AllowReassignment: conductoroneapi.Bool(false),
                                Assigned: conductoroneapi.Bool(false),
                                RequireApprovalReason: conductoroneapi.Bool(false),
                                RequireReassignmentReason: conductoroneapi.Bool(false),
                                Typ: &shared.ApprovalTyp{
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "labore",
                                            "labore",
                                            "suscipit",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("natus"),
                                        AppID: conductoroneapi.String("nobis"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "vero",
                                            "aspernatur",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "magnam",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "excepturi",
                                        },
                                    },
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "provident",
                                            "quos",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "accusantium",
                                            "mollitia",
                                            "reiciendis",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "ad",
                                            "eum",
                                            "dolor",
                                        },
                                    },
                                },
                            },
                            Provision: &shared.Provision{
                                Assigned: conductoroneapi.Bool(false),
                                ProvisionPolicy: &shared.ProvisionPolicy{
                                    Typ: &shared.ProvisionPolicyTyp{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("necessitatibus"),
                                            EntitlementID: conductoroneapi.String("odit"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("nemo"),
                                            UserIds: []string{
                                                "iure",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
        PolicyType: shared.CreatePolicyRequestPolicyTypePolicyTypeProvision.ToPointer(),
        PostActions: []shared.PolicyPostActions{
            shared.PolicyPostActions{
                Action: &shared.PolicyPostActionsAction{
                    CertifyRemediateImmediately: conductoroneapi.Bool(false),
                },
            },
            shared.PolicyPostActions{
                Action: &shared.PolicyPostActionsAction{
                    CertifyRemediateImmediately: conductoroneapi.Bool(false),
                },
            },
            shared.PolicyPostActions{
                Action: &shared.PolicyPostActionsAction{
                    CertifyRemediateImmediately: conductoroneapi.Bool(false),
                },
            },
            shared.PolicyPostActions{
                Action: &shared.PolicyPostActionsAction{
                    CertifyRemediateImmediately: conductoroneapi.Bool(false),
                },
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
        ID: "4c8b711e-5b7f-4d2e-9028-921cddc69260",
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
        ID: "1fb576b0-d5f0-4d30-85fb-b2587053202c",
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

