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
        Description: conductoroneapi.String("temporibus"),
        DisplayName: conductoroneapi.String("accusantium"),
        PolicySteps: map[string]shared.PolicyStepsInput{
            "aut": shared.PolicyStepsInput{
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
                                    AppID: conductoroneapi.String("eum"),
                                    EntitlementID: conductoroneapi.String("mollitia"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("ab"),
                                    UserIds: []string{
                                        "non",
                                        "voluptatem",
                                        "dolor",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("occaecati"),
                                AppID: conductoroneapi.String("numquam"),
                                AppUserID: conductoroneapi.String("impedit"),
                                GrantDuration: conductoroneapi.String("explicabo"),
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
                                    AppID: conductoroneapi.String("voluptas"),
                                    EntitlementID: conductoroneapi.String("aut"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("dignissimos"),
                                    UserIds: []string{
                                        "maiores",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("natus"),
                                AppID: conductoroneapi.String("velit"),
                                AppUserID: conductoroneapi.String("voluptatibus"),
                                GrantDuration: conductoroneapi.String("voluptas"),
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
                                    AppID: conductoroneapi.String("asperiores"),
                                    EntitlementID: conductoroneapi.String("aperiam"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("ea"),
                                    UserIds: []string{
                                        "consequuntur",
                                        "repellendus",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("officia"),
                                AppID: conductoroneapi.String("maxime"),
                                AppUserID: conductoroneapi.String("dignissimos"),
                                GrantDuration: conductoroneapi.String("officia"),
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "asperiores": shared.PolicyStepsInput{
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
                                    AppID: conductoroneapi.String("quae"),
                                    EntitlementID: conductoroneapi.String("quaerat"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("porro"),
                                    UserIds: []string{
                                        "labore",
                                        "ab",
                                        "adipisci",
                                        "fuga",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("id"),
                                AppID: conductoroneapi.String("suscipit"),
                                AppUserID: conductoroneapi.String("velit"),
                                GrantDuration: conductoroneapi.String("culpa"),
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
                                    AppID: conductoroneapi.String("est"),
                                    EntitlementID: conductoroneapi.String("recusandae"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("totam"),
                                    UserIds: []string{
                                        "vel",
                                        "ducimus",
                                        "quos",
                                        "vel",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("labore"),
                                AppID: conductoroneapi.String("possimus"),
                                AppUserID: conductoroneapi.String("facilis"),
                                GrantDuration: conductoroneapi.String("cum"),
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "commodi": shared.PolicyStepsInput{
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
                                    AppID: conductoroneapi.String("corporis"),
                                    EntitlementID: conductoroneapi.String("reiciendis"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("assumenda"),
                                    UserIds: []string{
                                        "recusandae",
                                        "aliquid",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("aperiam"),
                                AppID: conductoroneapi.String("cum"),
                                AppUserID: conductoroneapi.String("consectetur"),
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
                                    AppID: conductoroneapi.String("exercitationem"),
                                    EntitlementID: conductoroneapi.String("earum"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("facere"),
                                    UserIds: []string{
                                        "doloribus",
                                        "suscipit",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("reiciendis"),
                                AppID: conductoroneapi.String("quidem"),
                                AppUserID: conductoroneapi.String("saepe"),
                                GrantDuration: conductoroneapi.String("necessitatibus"),
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
        ID: "f33317fe-35b6-40eb-9ea4-26555ba3c287",
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
        ID: "44ed53b8-8f3a-48d8-b5c0-b2f2fb7b194a",
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
        PageSize: conductoroneapi.Float64(1523.54),
        PageToken: conductoroneapi.String("in"),
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
                Description: conductoroneapi.String("commodi"),
                DisplayName: conductoroneapi.String("quidem"),
                PolicySteps: map[string]shared.PolicyStepsInput{
                    "voluptas": shared.PolicyStepsInput{
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
                                            AppID: conductoroneapi.String("architecto"),
                                            EntitlementID: conductoroneapi.String("suscipit"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("sapiente"),
                                            UserIds: []string{
                                                "illo",
                                                "reiciendis",
                                                "perferendis",
                                                "corrupti",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("maiores"),
                                        AppID: conductoroneapi.String("incidunt"),
                                        AppUserID: conductoroneapi.String("sed"),
                                        GrantDuration: conductoroneapi.String("provident"),
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
                                            AppID: conductoroneapi.String("eius"),
                                            EntitlementID: conductoroneapi.String("necessitatibus"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("ipsum"),
                                            UserIds: []string{
                                                "occaecati",
                                                "quos",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("voluptatibus"),
                                        AppID: conductoroneapi.String("tempora"),
                                        AppUserID: conductoroneapi.String("tempora"),
                                        GrantDuration: conductoroneapi.String("voluptate"),
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
                                            AppID: conductoroneapi.String("reiciendis"),
                                            EntitlementID: conductoroneapi.String("ex"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("sit"),
                                            UserIds: []string{
                                                "officiis",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("praesentium"),
                                        AppID: conductoroneapi.String("facilis"),
                                        AppUserID: conductoroneapi.String("quaerat"),
                                        GrantDuration: conductoroneapi.String("incidunt"),
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
                    shared.PolicyPostActions{
                        CertifyRemediateImmediately: conductoroneapi.Bool(false),
                    },
                },
                ReassignTasksToDelegates: conductoroneapi.Bool(false),
            },
            UpdateMask: conductoroneapi.String("rem"),
        },
        ID: "0ca55efd-20e4-457e-9858-b6a89fbe3a5a",
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

