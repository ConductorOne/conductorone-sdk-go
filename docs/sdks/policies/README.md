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
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Policies.Create(ctx, shared.CreatePolicyRequest{
        Description: conductoroneapi.String("dolorem"),
        DisplayName: conductoroneapi.String("dolor"),
        PolicySteps: map[string]shared.PolicySteps{
            "ipsum": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("excepturi"),
                                AppID: conductoroneapi.String("cum"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "dignissimos",
                                    "reiciendis",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "dolorum",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "veritatis",
                                    "ipsa",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "iure",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "quaerat",
                                    "accusamus",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "voluptatibus",
                                    "voluptas",
                                    "natus",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "atque",
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
                                    AppID: conductoroneapi.String("sit"),
                                    EntitlementID: conductoroneapi.String("fugiat"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("ab"),
                                    UserIds: []string{
                                        "dolorum",
                                        "iusto",
                                        "voluptate",
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
                                AppGroupID: conductoroneapi.String("dolorum"),
                                AppID: conductoroneapi.String("deleniti"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "necessitatibus",
                                    "distinctio",
                                    "asperiores",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "ipsum",
                                    "voluptate",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "saepe",
                                    "eius",
                                    "aspernatur",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "amet",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "accusamus",
                                    "ad",
                                    "saepe",
                                    "suscipit",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "provident",
                                    "minima",
                                    "repellendus",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "similique",
                                    "alias",
                                    "at",
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
                                    EntitlementID: conductoroneapi.String("tempora"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("vel"),
                                    UserIds: []string{
                                        "officiis",
                                        "qui",
                                        "dolorum",
                                        "a",
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
                                AppGroupID: conductoroneapi.String("esse"),
                                AppID: conductoroneapi.String("harum"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "ipsum",
                                    "quisquam",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "amet",
                                    "tempore",
                                    "accusamus",
                                    "numquam",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
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
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "sed",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "libero",
                                    "voluptas",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "quam",
                                    "ipsum",
                                    "incidunt",
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
                                    AppID: conductoroneapi.String("qui"),
                                    EntitlementID: conductoroneapi.String("cupiditate"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("maxime"),
                                    UserIds: []string{
                                        "soluta",
                                        "dicta",
                                        "laborum",
                                        "totam",
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
                                AppGroupID: conductoroneapi.String("incidunt"),
                                AppID: conductoroneapi.String("aspernatur"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "distinctio",
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
                                    "quam",
                                    "molestias",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
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
                            SelfApproval: &shared.SelfApproval{
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
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "veritatis",
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
                                    AppID: conductoroneapi.String("nobis"),
                                    EntitlementID: conductoroneapi.String("quos"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("tempore"),
                                    UserIds: []string{
                                        "aperiam",
                                        "delectus",
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
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Policies.Delete(ctx, operations.C1APIPolicyV1PoliciesDeleteRequest{
        DeletePolicyRequest: &shared.DeletePolicyRequest{},
        ID: "3a1108e0-adcf-44b9-a187-9fce953f73ef",
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
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Policies.Get(ctx, operations.C1APIPolicyV1PoliciesGetRequest{
        ID: "7fbc7abd-74dd-439c-8f5d-2cff7c70a456",
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
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

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
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Policies.Update(ctx, operations.C1APIPolicyV1PoliciesUpdateRequest{
        UpdatePolicyRequestInput: &shared.UpdatePolicyRequestInput{
            Policy: &shared.PolicyInput{
                Description: conductoroneapi.String("aspernatur"),
                DisplayName: conductoroneapi.String("vel"),
                ID: conductoroneapi.String("d436813f-16d9-4f5f-8e6c-556146c3e250"),
                PolicySteps: map[string]shared.PolicySteps{
                    "libero": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("aut"),
                                        AppID: conductoroneapi.String("deleniti"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "aliquam",
                                            "fugit",
                                            "accusamus",
                                            "inventore",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "et",
                                            "dolorum",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
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
                                    SelfApproval: &shared.SelfApproval{
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
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "ipsa",
                                            "molestiae",
                                            "magnam",
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
                                            AppID: conductoroneapi.String("odio"),
                                            EntitlementID: conductoroneapi.String("eius"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("esse"),
                                            UserIds: []string{
                                                "rem",
                                                "fuga",
                                            },
                                        },
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                        },
                    },
                    "reprehenderit": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("fugiat"),
                                        AppID: conductoroneapi.String("ut"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "suscipit",
                                            "assumenda",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "praesentium",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "veritatis",
                                            "ipsa",
                                            "id",
                                            "quidem",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "quo",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
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
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "tempora",
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
                                            AppID: conductoroneapi.String("debitis"),
                                            EntitlementID: conductoroneapi.String("ipsam"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("aspernatur"),
                                            UserIds: []string{
                                                "quo",
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
                                        AppGroupID: conductoroneapi.String("esse"),
                                        AppID: conductoroneapi.String("recusandae"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "distinctio",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "dignissimos",
                                            "inventore",
                                            "nihil",
                                            "totam",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "aliquam",
                                            "odio",
                                            "occaecati",
                                            "commodi",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "dolores",
                                            "deserunt",
                                            "molestiae",
                                            "accusantium",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "eum",
                                            "quas",
                                            "praesentium",
                                            "consequuntur",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "fugit",
                                            "fuga",
                                            "mollitia",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "atque",
                                            "explicabo",
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
                                            AppID: conductoroneapi.String("minima"),
                                            EntitlementID: conductoroneapi.String("nisi"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("fugit"),
                                            UserIds: []string{
                                                "consequuntur",
                                                "ratione",
                                                "explicabo",
                                                "saepe",
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
                                        AppGroupID: conductoroneapi.String("occaecati"),
                                        AppID: conductoroneapi.String("atque"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "esse",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "accusamus",
                                            "veritatis",
                                            "esse",
                                            "quod",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
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
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "minima",
                                            "distinctio",
                                            "eligendi",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "culpa",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "adipisci",
                                            "cumque",
                                            "consequuntur",
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
                                            AppID: conductoroneapi.String("consequatur"),
                                            EntitlementID: conductoroneapi.String("minus"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("quaerat"),
                                            UserIds: []string{
                                                "consectetur",
                                                "esse",
                                                "blanditiis",
                                                "provident",
                                            },
                                        },
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                        },
                    },
                    "a": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("quas"),
                                        AppID: conductoroneapi.String("esse"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "a",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "sint",
                                            "pariatur",
                                            "possimus",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "eveniet",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "facere",
                                            "veritatis",
                                            "consequuntur",
                                            "quasi",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "culpa",
                                            "aliquid",
                                            "tenetur",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "earum",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "in",
                                            "eius",
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
                                            EntitlementID: conductoroneapi.String("illum"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("soluta"),
                                            UserIds: []string{
                                                "aliquam",
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
                                        AppGroupID: conductoroneapi.String("sapiente"),
                                        AppID: conductoroneapi.String("dicta"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "reprehenderit",
                                            "ullam",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "aut",
                                            "voluptatum",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "quibusdam",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "deleniti",
                                            "itaque",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "architecto",
                                            "omnis",
                                            "tenetur",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "at",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "voluptate",
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
                                            AppID: conductoroneapi.String("ipsa"),
                                            EntitlementID: conductoroneapi.String("minima"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("veritatis"),
                                            UserIds: []string{
                                                "adipisci",
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
                                        AppGroupID: conductoroneapi.String("iste"),
                                        AppID: conductoroneapi.String("temporibus"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "rem",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "laudantium",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "mollitia",
                                            "ab",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "non",
                                            "voluptatem",
                                            "dolor",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "numquam",
                                            "impedit",
                                            "explicabo",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "aut",
                                            "dignissimos",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "maiores",
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
                                            AppID: conductoroneapi.String("natus"),
                                            EntitlementID: conductoroneapi.String("velit"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("voluptatibus"),
                                            UserIds: []string{
                                                "asperiores",
                                                "aperiam",
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
                                        AppID: conductoroneapi.String("quaerat"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "repellendus",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "maxime",
                                            "dignissimos",
                                            "officia",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "nemo",
                                            "quae",
                                            "quaerat",
                                            "porro",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "labore",
                                            "ab",
                                            "adipisci",
                                            "fuga",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "suscipit",
                                            "velit",
                                            "culpa",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "recusandae",
                                            "totam",
                                            "fugiat",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "ducimus",
                                            "quos",
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
                                            AppID: conductoroneapi.String("vel"),
                                            EntitlementID: conductoroneapi.String("labore"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("possimus"),
                                            UserIds: []string{
                                                "cum",
                                                "commodi",
                                                "in",
                                            },
                                        },
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                        },
                    },
                    "corporis": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("assumenda"),
                                        AppID: conductoroneapi.String("nemo"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "aliquid",
                                            "aperiam",
                                            "cum",
                                            "consectetur",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "exercitationem",
                                            "earum",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "numquam",
                                            "doloribus",
                                            "suscipit",
                                            "reiciendis",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "saepe",
                                            "necessitatibus",
                                            "dolore",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "asperiores",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "non",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "beatae",
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
                                            AppID: conductoroneapi.String("dignissimos"),
                                            EntitlementID: conductoroneapi.String("a"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("debitis"),
                                            UserIds: []string{
                                                "corporis",
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
                                        AppGroupID: conductoroneapi.String("harum"),
                                        AppID: conductoroneapi.String("laboriosam"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "voluptates",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "vitae",
                                            "accusamus",
                                            "similique",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "aspernatur",
                                            "voluptas",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "voluptas",
                                            "minima",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "dolorum",
                                            "adipisci",
                                            "minus",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "blanditiis",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "dolore",
                                            "aliquam",
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
                                            AppID: conductoroneapi.String("officiis"),
                                            EntitlementID: conductoroneapi.String("temporibus"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("ullam"),
                                            UserIds: []string{
                                                "cum",
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
                                        AppID: conductoroneapi.String("quas"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "nesciunt",
                                            "culpa",
                                            "corrupti",
                                            "pariatur",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "hic",
                                            "exercitationem",
                                            "nobis",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "rerum",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "reiciendis",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "asperiores",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "voluptate",
                                            "expedita",
                                            "ab",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "dolore",
                                            "laborum",
                                            "sed",
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
                                            AppID: conductoroneapi.String("in"),
                                            EntitlementID: conductoroneapi.String("commodi"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("quidem"),
                                            UserIds: []string{
                                                "voluptas",
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
                                        AppGroupID: conductoroneapi.String("unde"),
                                        AppID: conductoroneapi.String("architecto"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "sapiente",
                                            "debitis",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "reiciendis",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "corrupti",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "incidunt",
                                            "sed",
                                            "provident",
                                            "eius",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "ipsum",
                                            "ea",
                                            "occaecati",
                                            "quos",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "tempora",
                                            "tempora",
                                            "voluptate",
                                            "reiciendis",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "sit",
                                            "non",
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
                                            AppID: conductoroneapi.String("officiis"),
                                            EntitlementID: conductoroneapi.String("praesentium"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("facilis"),
                                            UserIds: []string{
                                                "incidunt",
                                                "ipsam",
                                            },
                                        },
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                        },
                    },
                },
                PolicyType: shared.PolicyPolicyTypePolicyTypeProvision.ToPointer(),
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
            },
            UpdateMask: conductoroneapi.String("sit"),
        },
        ID: "ca55efd2-0e45-47e1-858b-6a89fbe3a5aa",
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

