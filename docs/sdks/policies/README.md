# Policies

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List
* [Update](#update) - Update

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
    res, err := s.Policies.Create(ctx, shared.CreatePolicyRequestInput{
        Description: conductoroneapi.String("eius"),
        DisplayName: conductoroneapi.String("libero"),
        PolicySteps: map[string]shared.PolicyStepsInput{
            "soluta": shared.PolicyStepsInput{
                Steps: []shared.PolicyStepInput{
                    shared.PolicyStepInput{
                        Approval: &shared.ApprovalInput{
                            AppGroupApproval: &shared.AppGroupApproval1{},
                            AppOwnerApproval: &shared.AppOwnerApproval1{},
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval1{},
                            ManagerApproval: &shared.ManagerApproval1{},
                            SelfApproval: &shared.SelfApproval1{},
                            UserApproval: &shared.UserApproval1{},
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("aliquam"),
                                    EntitlementID: conductoroneapi.String("sapiente"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("dicta"),
                                    UserIds: []string{
                                        "reprehenderit",
                                        "ullam",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "nisi": shared.PolicyStepsInput{
                Steps: []shared.PolicyStepInput{
                    shared.PolicyStepInput{
                        Approval: &shared.ApprovalInput{
                            AppGroupApproval: &shared.AppGroupApproval1{},
                            AppOwnerApproval: &shared.AppOwnerApproval1{},
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval1{},
                            ManagerApproval: &shared.ManagerApproval1{},
                            SelfApproval: &shared.SelfApproval1{},
                            UserApproval: &shared.UserApproval1{},
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("voluptatum"),
                                    EntitlementID: conductoroneapi.String("qui"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("quibusdam"),
                                    UserIds: []string{
                                        "deleniti",
                                        "itaque",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "dolorum": shared.PolicyStepsInput{
                Steps: []shared.PolicyStepInput{
                    shared.PolicyStepInput{
                        Approval: &shared.ApprovalInput{
                            AppGroupApproval: &shared.AppGroupApproval1{},
                            AppOwnerApproval: &shared.AppOwnerApproval1{},
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval1{},
                            ManagerApproval: &shared.ManagerApproval1{},
                            SelfApproval: &shared.SelfApproval1{},
                            UserApproval: &shared.UserApproval1{},
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("omnis"),
                                    EntitlementID: conductoroneapi.String("tenetur"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("quasi"),
                                    UserIds: []string{
                                        "et",
                                        "voluptate",
                                        "ipsa",
                                        "minima",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "veritatis": shared.PolicyStepsInput{
                Steps: []shared.PolicyStepInput{
                    shared.PolicyStepInput{
                        Approval: &shared.ApprovalInput{
                            AppGroupApproval: &shared.AppGroupApproval1{},
                            AppOwnerApproval: &shared.AppOwnerApproval1{},
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval1{},
                            ManagerApproval: &shared.ManagerApproval1{},
                            SelfApproval: &shared.SelfApproval1{},
                            UserApproval: &shared.UserApproval1{},
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("adipisci"),
                                    EntitlementID: conductoroneapi.String("iste"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("temporibus"),
                                    UserIds: []string{
                                        "rem",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
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
        ID: "6a184039-4c26-4071-b93f-5f0642dac7af",
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
        ID: "515cc413-aa63-4aae-8d67-864dbb675fd5",
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
                Description: conductoroneapi.String("recusandae"),
                DisplayName: conductoroneapi.String("aliquid"),
                PolicySteps: map[string]shared.PolicyStepsInput{
                    "cum": shared.PolicyStepsInput{
                        Steps: []shared.PolicyStepInput{
                            shared.PolicyStepInput{
                                Approval: &shared.ApprovalInput{
                                    AppGroupApproval: &shared.AppGroupApproval1{},
                                    AppOwnerApproval: &shared.AppOwnerApproval1{},
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval1{},
                                    ManagerApproval: &shared.ManagerApproval1{},
                                    SelfApproval: &shared.SelfApproval1{},
                                    UserApproval: &shared.UserApproval1{},
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("in"),
                                            EntitlementID: conductoroneapi.String("exercitationem"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("earum"),
                                            UserIds: []string{
                                                "numquam",
                                                "doloribus",
                                                "suscipit",
                                                "reiciendis",
                                            },
                                        },
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
            UpdateMask: conductoroneapi.String("necessitatibus"),
        },
        ID: "41f33317-fe35-4b60-ab1e-a426555ba3c2",
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

