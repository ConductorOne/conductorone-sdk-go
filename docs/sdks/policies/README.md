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
        Description: conductoroneapi.String("occaecati"),
        DisplayName: conductoroneapi.String("numquam"),
        PolicySteps: map[string]shared.PolicyStepsInput{
            "explicabo": shared.PolicyStepsInput{
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
                                    AppID: conductoroneapi.String("aut"),
                                    EntitlementID: conductoroneapi.String("dignissimos"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("dicta"),
                                    UserIds: []string{
                                        "natus",
                                        "velit",
                                        "voluptatibus",
                                        "voluptas",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("asperiores"),
                                AppID: conductoroneapi.String("aperiam"),
                                AppUserID: conductoroneapi.String("ea"),
                                GrantDuration: conductoroneapi.String("quaerat"),
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
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
                                    EntitlementID: conductoroneapi.String("repellendus"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("officia"),
                                    UserIds: []string{
                                        "dignissimos",
                                        "officia",
                                        "asperiores",
                                        "nemo",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("quae"),
                                AppID: conductoroneapi.String("quaerat"),
                                AppUserID: conductoroneapi.String("porro"),
                                GrantDuration: conductoroneapi.String("quod"),
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "labore": shared.PolicyStepsInput{
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
                                    AppID: conductoroneapi.String("adipisci"),
                                    EntitlementID: conductoroneapi.String("fuga"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("id"),
                                    UserIds: []string{
                                        "velit",
                                        "culpa",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("est"),
                                AppID: conductoroneapi.String("recusandae"),
                                AppUserID: conductoroneapi.String("totam"),
                                GrantDuration: conductoroneapi.String("fugiat"),
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "vel": shared.PolicyStepsInput{
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
                                    AppID: conductoroneapi.String("quos"),
                                    EntitlementID: conductoroneapi.String("vel"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("labore"),
                                    UserIds: []string{
                                        "facilis",
                                        "cum",
                                        "commodi",
                                        "in",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("corporis"),
                                AppID: conductoroneapi.String("reiciendis"),
                                AppUserID: conductoroneapi.String("assumenda"),
                                GrantDuration: conductoroneapi.String("nemo"),
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
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
                                    AppID: conductoroneapi.String("recusandae"),
                                    EntitlementID: conductoroneapi.String("aliquid"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("aperiam"),
                                    UserIds: []string{
                                        "consectetur",
                                        "in",
                                        "exercitationem",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("earum"),
                                AppID: conductoroneapi.String("facere"),
                                AppUserID: conductoroneapi.String("numquam"),
                                GrantDuration: conductoroneapi.String("doloribus"),
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "suscipit": shared.PolicyStepsInput{
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
                                    AppID: conductoroneapi.String("quidem"),
                                    EntitlementID: conductoroneapi.String("saepe"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("necessitatibus"),
                                    UserIds: []string{
                                        "sunt",
                                        "asperiores",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("adipisci"),
                                AppID: conductoroneapi.String("non"),
                                AppUserID: conductoroneapi.String("amet"),
                                GrantDuration: conductoroneapi.String("beatae"),
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
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
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("harum"),
                                AppID: conductoroneapi.String("laboriosam"),
                                AppUserID: conductoroneapi.String("ipsa"),
                                GrantDuration: conductoroneapi.String("voluptates"),
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
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
                                    AppID: conductoroneapi.String("libero"),
                                    EntitlementID: conductoroneapi.String("vitae"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("accusamus"),
                                    UserIds: []string{
                                        "tempora",
                                        "aspernatur",
                                        "voluptas",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("voluptas"),
                                AppID: conductoroneapi.String("voluptas"),
                                AppUserID: conductoroneapi.String("minima"),
                                GrantDuration: conductoroneapi.String("nobis"),
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
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
                                    AppID: conductoroneapi.String("dolorum"),
                                    EntitlementID: conductoroneapi.String("adipisci"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("minus"),
                                    UserIds: []string{
                                        "blanditiis",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("in"),
                                AppID: conductoroneapi.String("dolore"),
                                AppUserID: conductoroneapi.String("aliquam"),
                                GrantDuration: conductoroneapi.String("officiis"),
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
        ID: "3b88f3a8-d8f5-4c0b-af2f-b7b194a276b2",
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
        ID: "6916fe1f-08f4-4294-a369-8f447f603e8b",
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
        PageSize: conductoroneapi.Float64(3103.81),
        PageToken: conductoroneapi.String("incidunt"),
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
                Description: conductoroneapi.String("ipsam"),
                DisplayName: conductoroneapi.String("debitis"),
                PolicySteps: map[string]shared.PolicyStepsInput{
                    "sit": shared.PolicyStepsInput{
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
                                            AppID: conductoroneapi.String("error"),
                                            EntitlementID: conductoroneapi.String("veniam"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("minima"),
                                            UserIds: []string{
                                                "reiciendis",
                                                "nulla",
                                                "magni",
                                                "aperiam",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("saepe"),
                                        AppID: conductoroneapi.String("numquam"),
                                        AppUserID: conductoroneapi.String("veniam"),
                                        GrantDuration: conductoroneapi.String("in"),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
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
                                            AppID: conductoroneapi.String("officiis"),
                                            EntitlementID: conductoroneapi.String("beatae"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("laudantium"),
                                            UserIds: []string{
                                                "praesentium",
                                                "cum",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("laboriosam"),
                                        AppID: conductoroneapi.String("dolorum"),
                                        AppUserID: conductoroneapi.String("voluptatum"),
                                        GrantDuration: conductoroneapi.String("error"),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
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
                                            AppID: conductoroneapi.String("hic"),
                                            EntitlementID: conductoroneapi.String("expedita"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("debitis"),
                                            UserIds: []string{
                                                "dolorum",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("nostrum"),
                                        AppID: conductoroneapi.String("officia"),
                                        AppUserID: conductoroneapi.String("dolorum"),
                                        GrantDuration: conductoroneapi.String("corrupti"),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
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
                                            AppID: conductoroneapi.String("accusamus"),
                                            EntitlementID: conductoroneapi.String("tempora"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("atque"),
                                            UserIds: []string{
                                                "ut",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("fugiat"),
                                        AppID: conductoroneapi.String("voluptatem"),
                                        AppUserID: conductoroneapi.String("culpa"),
                                        GrantDuration: conductoroneapi.String("expedita"),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                        },
                    },
                    "magnam": shared.PolicyStepsInput{
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
                                            AppID: conductoroneapi.String("esse"),
                                            EntitlementID: conductoroneapi.String("ipsam"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("sit"),
                                            UserIds: []string{
                                                "quas",
                                                "repudiandae",
                                                "corporis",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("et"),
                                        AppID: conductoroneapi.String("blanditiis"),
                                        AppUserID: conductoroneapi.String("ex"),
                                        GrantDuration: conductoroneapi.String("sed"),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                        },
                    },
                    "sit": shared.PolicyStepsInput{
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
                                            AppID: conductoroneapi.String("nostrum"),
                                            EntitlementID: conductoroneapi.String("saepe"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("error"),
                                            UserIds: []string{
                                                "incidunt",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("reiciendis"),
                                        AppID: conductoroneapi.String("dolorem"),
                                        AppUserID: conductoroneapi.String("harum"),
                                        GrantDuration: conductoroneapi.String("dicta"),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
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
                                            AppID: conductoroneapi.String("architecto"),
                                            EntitlementID: conductoroneapi.String("occaecati"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("labore"),
                                            UserIds: []string{
                                                "atque",
                                                "laborum",
                                                "nam",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("tenetur"),
                                        AppID: conductoroneapi.String("laboriosam"),
                                        AppUserID: conductoroneapi.String("alias"),
                                        GrantDuration: conductoroneapi.String("amet"),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                        },
                    },
                },
                PolicyType: shared.PolicyPolicyTypePolicyTypeCertify.ToPointer(),
                PostActions: []shared.PolicyPostActions{
                    shared.PolicyPostActions{
                        CertifyRemediateImmediately: conductoroneapi.Bool(false),
                    },
                    shared.PolicyPostActions{
                        CertifyRemediateImmediately: conductoroneapi.Bool(false),
                    },
                },
                ReassignTasksToDelegates: conductoroneapi.Bool(false),
            },
            UpdateMask: conductoroneapi.String("unde"),
        },
        ID: "f9dfe0ab-7da8-4a50-8e18-7f86bc173d68",
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

