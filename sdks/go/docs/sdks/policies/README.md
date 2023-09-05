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
        Description: conductoroneapi.String("dolorum"),
        DisplayName: conductoroneapi.String("architecto"),
        PolicySteps: map[string]shared.PolicyStepsInput{
            "tenetur": shared.PolicyStepsInput{
                Steps: []shared.PolicyStepInput{
                    shared.PolicyStepInput{
                        Approval: &shared.ApprovalInput{
                            AppOwners: &shared.AppOwnerApproval1{},
                            EntitlementOwners: &shared.EntitlementOwnerApproval1{},
                            Group: &shared.AppGroupApproval1{},
                            Manager: &shared.ManagerApproval1{},
                            Self: &shared.SelfApproval1{},
                            Users: &shared.UserApproval1{},
                        },
                        Provision: &shared.Provision{
                            Assigned: conductoroneapi.Bool(false),
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                Connector: &shared.ConnectorProvision{},
                                Delegated: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("at"),
                                    EntitlementID: conductoroneapi.String("et"),
                                },
                                Manual: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("voluptate"),
                                    UserIds: []string{
                                        "minima",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("veritatis"),
                                AppID: conductoroneapi.String("consectetur"),
                                AppUserID: conductoroneapi.String("adipisci"),
                                GrantDuration: conductoroneapi.String("iste"),
                            },
                        },
                    },
                },
            },
            "temporibus": shared.PolicyStepsInput{
                Steps: []shared.PolicyStepInput{
                    shared.PolicyStepInput{
                        Approval: &shared.ApprovalInput{
                            AppOwners: &shared.AppOwnerApproval1{},
                            EntitlementOwners: &shared.EntitlementOwnerApproval1{},
                            Group: &shared.AppGroupApproval1{},
                            Manager: &shared.ManagerApproval1{},
                            Self: &shared.SelfApproval1{},
                            Users: &shared.UserApproval1{},
                        },
                        Provision: &shared.Provision{
                            Assigned: conductoroneapi.Bool(false),
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                Connector: &shared.ConnectorProvision{},
                                Delegated: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("rem"),
                                    EntitlementID: conductoroneapi.String("aut"),
                                },
                                Manual: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("laudantium"),
                                    UserIds: []string{
                                        "mollitia",
                                        "ab",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("corrupti"),
                                AppID: conductoroneapi.String("non"),
                                AppUserID: conductoroneapi.String("voluptatem"),
                                GrantDuration: conductoroneapi.String("dolor"),
                            },
                        },
                    },
                },
            },
            "occaecati": shared.PolicyStepsInput{
                Steps: []shared.PolicyStepInput{
                    shared.PolicyStepInput{
                        Approval: &shared.ApprovalInput{
                            AppOwners: &shared.AppOwnerApproval1{},
                            EntitlementOwners: &shared.EntitlementOwnerApproval1{},
                            Group: &shared.AppGroupApproval1{},
                            Manager: &shared.ManagerApproval1{},
                            Self: &shared.SelfApproval1{},
                            Users: &shared.UserApproval1{},
                        },
                        Provision: &shared.Provision{
                            Assigned: conductoroneapi.Bool(false),
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                Connector: &shared.ConnectorProvision{},
                                Delegated: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("impedit"),
                                    EntitlementID: conductoroneapi.String("explicabo"),
                                },
                                Manual: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("voluptas"),
                                    UserIds: []string{
                                        "dignissimos",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("dicta"),
                                AppID: conductoroneapi.String("maiores"),
                                AppUserID: conductoroneapi.String("natus"),
                                GrantDuration: conductoroneapi.String("velit"),
                            },
                        },
                    },
                    shared.PolicyStepInput{
                        Approval: &shared.ApprovalInput{
                            AppOwners: &shared.AppOwnerApproval1{},
                            EntitlementOwners: &shared.EntitlementOwnerApproval1{},
                            Group: &shared.AppGroupApproval1{},
                            Manager: &shared.ManagerApproval1{},
                            Self: &shared.SelfApproval1{},
                            Users: &shared.UserApproval1{},
                        },
                        Provision: &shared.Provision{
                            Assigned: conductoroneapi.Bool(false),
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                Connector: &shared.ConnectorProvision{},
                                Delegated: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("voluptatibus"),
                                    EntitlementID: conductoroneapi.String("voluptas"),
                                },
                                Manual: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("asperiores"),
                                    UserIds: []string{
                                        "ea",
                                    },
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoroneapi.String("quaerat"),
                                AppID: conductoroneapi.String("consequuntur"),
                                AppUserID: conductoroneapi.String("repellendus"),
                                GrantDuration: conductoroneapi.String("officia"),
                            },
                        },
                    },
                },
            },
        },
        PolicyType: shared.CreatePolicyRequestPolicyTypePolicyTypeAccessRequest.ToPointer(),
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
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
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
    res, err := s.Policies.Delete(ctx, operations.C1APIPolicyV1PoliciesDeleteRequest{
        DeletePolicyRequest: &shared.DeletePolicyRequest{},
        ID: "af515cc4-13aa-463a-ae8d-67864dbb675f",
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
        ID: "d5e60b37-5ed4-4f6f-bee4-1f33317fe35b",
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
        PageSize: conductoroneapi.Float64(3852.37),
        PageToken: conductoroneapi.String("ipsa"),
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
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
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
    res, err := s.Policies.Update(ctx, operations.C1APIPolicyV1PoliciesUpdateRequest{
        UpdatePolicyRequestInput: &shared.UpdatePolicyRequestInput{
            Policy: &shared.PolicyInput{
                Description: conductoroneapi.String("voluptates"),
                DisplayName: conductoroneapi.String("libero"),
                PolicySteps: map[string]shared.PolicyStepsInput{
                    "accusamus": shared.PolicyStepsInput{
                        Steps: []shared.PolicyStepInput{
                            shared.PolicyStepInput{
                                Approval: &shared.ApprovalInput{
                                    AppOwners: &shared.AppOwnerApproval1{},
                                    EntitlementOwners: &shared.EntitlementOwnerApproval1{},
                                    Group: &shared.AppGroupApproval1{},
                                    Manager: &shared.ManagerApproval1{},
                                    Self: &shared.SelfApproval1{},
                                    Users: &shared.UserApproval1{},
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("tempora"),
                                            EntitlementID: conductoroneapi.String("aspernatur"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("voluptas"),
                                            UserIds: []string{
                                                "voluptas",
                                                "minima",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("nobis"),
                                        AppID: conductoroneapi.String("dolorum"),
                                        AppUserID: conductoroneapi.String("adipisci"),
                                        GrantDuration: conductoroneapi.String("minus"),
                                    },
                                },
                            },
                            shared.PolicyStepInput{
                                Approval: &shared.ApprovalInput{
                                    AppOwners: &shared.AppOwnerApproval1{},
                                    EntitlementOwners: &shared.EntitlementOwnerApproval1{},
                                    Group: &shared.AppGroupApproval1{},
                                    Manager: &shared.ManagerApproval1{},
                                    Self: &shared.SelfApproval1{},
                                    Users: &shared.UserApproval1{},
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("dolores"),
                                            EntitlementID: conductoroneapi.String("blanditiis"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("in"),
                                            UserIds: []string{
                                                "aliquam",
                                                "officiis",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("temporibus"),
                                        AppID: conductoroneapi.String("ullam"),
                                        AppUserID: conductoroneapi.String("adipisci"),
                                        GrantDuration: conductoroneapi.String("cum"),
                                    },
                                },
                            },
                            shared.PolicyStepInput{
                                Approval: &shared.ApprovalInput{
                                    AppOwners: &shared.AppOwnerApproval1{},
                                    EntitlementOwners: &shared.EntitlementOwnerApproval1{},
                                    Group: &shared.AppGroupApproval1{},
                                    Manager: &shared.ManagerApproval1{},
                                    Self: &shared.SelfApproval1{},
                                    Users: &shared.UserApproval1{},
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("blanditiis"),
                                            EntitlementID: conductoroneapi.String("quas"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("hic"),
                                            UserIds: []string{
                                                "culpa",
                                            },
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoroneapi.String("corrupti"),
                                        AppID: conductoroneapi.String("pariatur"),
                                        AppUserID: conductoroneapi.String("totam"),
                                        GrantDuration: conductoroneapi.String("hic"),
                                    },
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
            UpdateMask: conductoroneapi.String("sit"),
        },
        ID: "b2f2fb7b-194a-4276-b269-16fe1f08f429",
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

