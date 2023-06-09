# Task

### Available Operations

* [C1APITaskV1TaskActionsServiceApprove](#c1apitaskv1taskactionsserviceapprove) - Invokes the c1.api.task.v1.TaskActionsService.Approve method.
* [C1APITaskV1TaskActionsServiceComment](#c1apitaskv1taskactionsservicecomment) - Invokes the c1.api.task.v1.TaskActionsService.Comment method.
* [C1APITaskV1TaskActionsServiceDeny](#c1apitaskv1taskactionsservicedeny) - Invokes the c1.api.task.v1.TaskActionsService.Deny method.
* [C1APITaskV1TaskSearchServiceSearch](#c1apitaskv1tasksearchservicesearch) - Invokes the c1.api.task.v1.TaskSearchService.Search method.
* [C1APITaskV1TaskServiceCreateGrantTask](#c1apitaskv1taskservicecreategranttask) - Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.
* [C1APITaskV1TaskServiceCreateRevokeTask](#c1apitaskv1taskservicecreaterevoketask) - Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.
* [C1APITaskV1TaskServiceGet](#c1apitaskv1taskserviceget) - Invokes the c1.api.task.v1.TaskService.Get method.

## C1APITaskV1TaskActionsServiceApprove

Invokes the c1.api.task.v1.TaskActionsService.Approve method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/operations"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Task.C1APITaskV1TaskActionsServiceApprove(ctx, operations.C1APITaskV1TaskActionsServiceApproveRequest{
        C1APITaskV1TaskActionsServiceApproveRequestInput: &shared.C1APITaskV1TaskActionsServiceApproveRequestInput{
            Comment: conductoroneapi.String("nihil"),
            ExpandMask: &shared.C1APITaskV1TaskExpandMask{
                Paths: []string{
                    "voluptatibus",
                    "ipsa",
                    "omnis",
                },
            },
            PolicyStepID: conductoroneapi.String("voluptate"),
        },
        TaskID: "cum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskActionsServiceApproveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APITaskV1TaskActionsServiceApproveRequest](../../models/operations/c1apitaskv1taskactionsserviceapproverequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.C1APITaskV1TaskActionsServiceApproveResponse](../../models/operations/c1apitaskv1taskactionsserviceapproveresponse.md), error**


## C1APITaskV1TaskActionsServiceComment

Invokes the c1.api.task.v1.TaskActionsService.Comment method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/operations"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Task.C1APITaskV1TaskActionsServiceComment(ctx, operations.C1APITaskV1TaskActionsServiceCommentRequest{
        C1APITaskV1TaskActionsServiceCommentRequestInput: &shared.C1APITaskV1TaskActionsServiceCommentRequestInput{
            Comment: conductoroneapi.String("perferendis"),
            ExpandMask: &shared.C1APITaskV1TaskExpandMask{
                Paths: []string{
                    "reprehenderit",
                },
            },
        },
        TaskID: "ut",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskActionsServiceCommentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APITaskV1TaskActionsServiceCommentRequest](../../models/operations/c1apitaskv1taskactionsservicecommentrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.C1APITaskV1TaskActionsServiceCommentResponse](../../models/operations/c1apitaskv1taskactionsservicecommentresponse.md), error**


## C1APITaskV1TaskActionsServiceDeny

Invokes the c1.api.task.v1.TaskActionsService.Deny method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/operations"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Task.C1APITaskV1TaskActionsServiceDeny(ctx, operations.C1APITaskV1TaskActionsServiceDenyRequest{
        C1APITaskV1TaskActionsServiceDenyRequestInput: &shared.C1APITaskV1TaskActionsServiceDenyRequestInput{
            Comment: conductoroneapi.String("maiores"),
            ExpandMask: &shared.C1APITaskV1TaskExpandMask{
                Paths: []string{
                    "corporis",
                },
            },
            PolicyStepID: conductoroneapi.String("dolore"),
        },
        TaskID: "iusto",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskActionsServiceDenyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.C1APITaskV1TaskActionsServiceDenyRequest](../../models/operations/c1apitaskv1taskactionsservicedenyrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.C1APITaskV1TaskActionsServiceDenyResponse](../../models/operations/c1apitaskv1taskactionsservicedenyresponse.md), error**


## C1APITaskV1TaskSearchServiceSearch

Invokes the c1.api.task.v1.TaskSearchService.Search method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/shared"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/types"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Task.C1APITaskV1TaskSearchServiceSearch(ctx, shared.C1APITaskV1TaskSearchRequest{
        AccessReviewIds: []string{
            "harum",
        },
        AccountOwnerIds: []string{
            "accusamus",
            "commodi",
        },
        ActorID: conductoroneapi.String("repudiandae"),
        AppEntitlementIds: []string{
            "ipsum",
        },
        AppResourceIds: []string{
            "molestias",
            "excepturi",
            "pariatur",
        },
        AppResourceTypeIds: []string{
            "praesentium",
            "rem",
        },
        AppUserSubjectIds: []string{
            "quasi",
            "repudiandae",
            "sint",
            "veritatis",
        },
        ApplicationIds: []string{
            "incidunt",
            "enim",
            "consequatur",
            "est",
        },
        AssigneesInIds: []string{
            "explicabo",
            "deserunt",
            "distinctio",
            "quibusdam",
        },
        CreatedAfter: types.MustTimeFromString("2022-09-26T08:57:48.803Z"),
        CreatedBefore: types.MustTimeFromString("2022-08-08T19:05:24.174Z"),
        CurrentStep: shared.C1APITaskV1TaskSearchRequestCurrentStepTaskSearchCurrentStepApproval.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "perferendis",
            "magni",
            "assumenda",
        },
        ExcludeIds: []string{
            "alias",
            "fugit",
        },
        ExpandMask: &shared.C1APITaskV1TaskExpandMask{
            Paths: []string{
                "excepturi",
                "tempora",
                "facilis",
            },
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "labore",
            "delectus",
            "eum",
        },
        OpenerIds: []string{
            "eligendi",
        },
        PageSize: conductoroneapi.Float64(5761.57),
        PageToken: conductoroneapi.String("aliquid"),
        PreviouslyActedOnIds: []string{
            "necessitatibus",
            "sint",
            "officia",
        },
        Query: conductoroneapi.String("dolor"),
        Refs: []shared.C1APITaskV1TaskRef{
            shared.C1APITaskV1TaskRef{
                ID: conductoroneapi.String("fa77dfb1-4cd6-46ae-b95e-fb9ba88f3a66"),
            },
            shared.C1APITaskV1TaskRef{
                ID: conductoroneapi.String("997074ba-4469-4b6e-a141-959890afa563"),
            },
            shared.C1APITaskV1TaskRef{
                ID: conductoroneapi.String("e2516fe4-c8b7-411e-9b7f-d2ed028921cd"),
            },
            shared.C1APITaskV1TaskRef{
                ID: conductoroneapi.String("dc692601-fb57-46b0-95f0-d30c5fbb2587"),
            },
        },
        SortBy: shared.C1APITaskV1TaskSearchRequestSortByTaskSearchSortByUnspecified.ToPointer(),
        SubjectIds: []string{
            "nesciunt",
            "eos",
        },
        TaskStates: []shared.C1APITaskV1TaskSearchRequestTaskStates{
            shared.C1APITaskV1TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.C1APITaskV1TaskType{
            shared.C1APITaskV1TaskType{
                Certify: &shared.C1APITaskV1TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("quam"),
                    AccessReviewSelection: conductoroneapi.String("dolor"),
                    AppEntitlementID: conductoroneapi.String("vero"),
                    AppID: conductoroneapi.String("nostrum"),
                    AppUserID: conductoroneapi.String("hic"),
                    IdentityUserID: conductoroneapi.String("recusandae"),
                    Outcome: shared.C1APITaskV1TaskTypeCertifyOutcomeCertifyOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-10-22T10:35:49.400Z"),
                },
                Grant: &shared.C1APITaskV1TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("voluptatem"),
                    AppID: conductoroneapi.String("porro"),
                    AppUserID: conductoroneapi.String("consequuntur"),
                    GrantDuration: conductoroneapi.String("blanditiis"),
                    IdentityUserID: conductoroneapi.String("error"),
                    Outcome: shared.C1APITaskV1TaskTypeGrantOutcomeGrantOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-08-08T15:48:05.748Z"),
                },
                Revoke: &shared.C1APITaskV1TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("adipisci"),
                    AppID: conductoroneapi.String("asperiores"),
                    AppUserID: conductoroneapi.String("earum"),
                    IdentityUserID: conductoroneapi.String("modi"),
                    Outcome: shared.C1APITaskV1TaskTypeRevokeOutcomeRevokeOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-12-05T23:42:47.745Z"),
                    Source: &shared.C1APITaskV1TaskRevokeSource{
                        Expired: &shared.C1APITaskV1TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-03-26T01:09:56.415Z"),
                        },
                        NonUsage: &shared.C1APITaskV1TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2020-10-23T12:23:35.067Z"),
                            LastLogin: types.MustTimeFromString("2022-01-24T10:05:07.174Z"),
                        },
                        Request: &shared.C1APITaskV1TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("quos"),
                        },
                        Review: &shared.C1APITaskV1TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("aliquid"),
                            CertTicketID: conductoroneapi.String("dolorem"),
                        },
                    },
                },
            },
            shared.C1APITaskV1TaskType{
                Certify: &shared.C1APITaskV1TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("dolorem"),
                    AccessReviewSelection: conductoroneapi.String("dolor"),
                    AppEntitlementID: conductoroneapi.String("qui"),
                    AppID: conductoroneapi.String("ipsum"),
                    AppUserID: conductoroneapi.String("hic"),
                    IdentityUserID: conductoroneapi.String("excepturi"),
                    Outcome: shared.C1APITaskV1TaskTypeCertifyOutcomeCertifyOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-07-05T23:34:50.715Z"),
                },
                Grant: &shared.C1APITaskV1TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("reiciendis"),
                    AppID: conductoroneapi.String("amet"),
                    AppUserID: conductoroneapi.String("dolorum"),
                    GrantDuration: conductoroneapi.String("numquam"),
                    IdentityUserID: conductoroneapi.String("veritatis"),
                    Outcome: shared.C1APITaskV1TaskTypeGrantOutcomeGrantOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-07-26T10:30:36.625Z"),
                },
                Revoke: &shared.C1APITaskV1TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("odio"),
                    AppID: conductoroneapi.String("quaerat"),
                    AppUserID: conductoroneapi.String("accusamus"),
                    IdentityUserID: conductoroneapi.String("quidem"),
                    Outcome: shared.C1APITaskV1TaskTypeRevokeOutcomeRevokeOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-05-20T13:18:59.478Z"),
                    Source: &shared.C1APITaskV1TaskRevokeSource{
                        Expired: &shared.C1APITaskV1TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-06-16T23:42:38.113Z"),
                        },
                        NonUsage: &shared.C1APITaskV1TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-02-23T01:35:05.899Z"),
                            LastLogin: types.MustTimeFromString("2022-04-04T12:00:33.616Z"),
                        },
                        Request: &shared.C1APITaskV1TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("dolorum"),
                        },
                        Review: &shared.C1APITaskV1TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("iusto"),
                            CertTicketID: conductoroneapi.String("voluptate"),
                        },
                    },
                },
            },
            shared.C1APITaskV1TaskType{
                Certify: &shared.C1APITaskV1TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("dolorum"),
                    AccessReviewSelection: conductoroneapi.String("deleniti"),
                    AppEntitlementID: conductoroneapi.String("omnis"),
                    AppID: conductoroneapi.String("necessitatibus"),
                    AppUserID: conductoroneapi.String("distinctio"),
                    IdentityUserID: conductoroneapi.String("asperiores"),
                    Outcome: shared.C1APITaskV1TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-07-18T13:18:42.293Z"),
                },
                Grant: &shared.C1APITaskV1TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("id"),
                    AppID: conductoroneapi.String("saepe"),
                    AppUserID: conductoroneapi.String("eius"),
                    GrantDuration: conductoroneapi.String("aspernatur"),
                    IdentityUserID: conductoroneapi.String("perferendis"),
                    Outcome: shared.C1APITaskV1TaskTypeGrantOutcomeGrantOutcomeGranted.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2020-05-10T15:56:07.975Z"),
                },
                Revoke: &shared.C1APITaskV1TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("ad"),
                    AppID: conductoroneapi.String("saepe"),
                    AppUserID: conductoroneapi.String("suscipit"),
                    IdentityUserID: conductoroneapi.String("deserunt"),
                    Outcome: shared.C1APITaskV1TaskTypeRevokeOutcomeRevokeOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-03-03T16:00:55.670Z"),
                    Source: &shared.C1APITaskV1TaskRevokeSource{
                        Expired: &shared.C1APITaskV1TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-09-28T20:14:16.452Z"),
                        },
                        NonUsage: &shared.C1APITaskV1TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-02-16T11:34:48.736Z"),
                            LastLogin: types.MustTimeFromString("2022-09-23T03:46:17.712Z"),
                        },
                        Request: &shared.C1APITaskV1TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("vel"),
                        },
                        Review: &shared.C1APITaskV1TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("quod"),
                            CertTicketID: conductoroneapi.String("officiis"),
                        },
                    },
                },
            },
            shared.C1APITaskV1TaskType{
                Certify: &shared.C1APITaskV1TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("qui"),
                    AccessReviewSelection: conductoroneapi.String("dolorum"),
                    AppEntitlementID: conductoroneapi.String("a"),
                    AppID: conductoroneapi.String("esse"),
                    AppUserID: conductoroneapi.String("harum"),
                    IdentityUserID: conductoroneapi.String("iusto"),
                    Outcome: shared.C1APITaskV1TaskTypeCertifyOutcomeCertifyOutcomeCertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2020-02-28T15:06:02.733Z"),
                },
                Grant: &shared.C1APITaskV1TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("amet"),
                    AppID: conductoroneapi.String("tempore"),
                    AppUserID: conductoroneapi.String("accusamus"),
                    GrantDuration: conductoroneapi.String("numquam"),
                    IdentityUserID: conductoroneapi.String("enim"),
                    Outcome: shared.C1APITaskV1TaskTypeGrantOutcomeGrantOutcomeGranted.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-06-12T13:41:06.619Z"),
                },
                Revoke: &shared.C1APITaskV1TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("nihil"),
                    AppID: conductoroneapi.String("sit"),
                    AppUserID: conductoroneapi.String("expedita"),
                    IdentityUserID: conductoroneapi.String("neque"),
                    Outcome: shared.C1APITaskV1TaskTypeRevokeOutcomeRevokeOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-04-09T09:19:35.125Z"),
                    Source: &shared.C1APITaskV1TaskRevokeSource{
                        Expired: &shared.C1APITaskV1TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-05-10T02:43:24.423Z"),
                        },
                        NonUsage: &shared.C1APITaskV1TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-10-14T13:38:40.307Z"),
                            LastLogin: types.MustTimeFromString("2022-10-24T22:37:32.805Z"),
                        },
                        Request: &shared.C1APITaskV1TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("cupiditate"),
                        },
                        Review: &shared.C1APITaskV1TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("maxime"),
                            CertTicketID: conductoroneapi.String("pariatur"),
                        },
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.C1APITaskV1TaskSearchRequest](../../models/shared/c1apitaskv1tasksearchrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.C1APITaskV1TaskSearchServiceSearchResponse](../../models/operations/c1apitaskv1tasksearchservicesearchresponse.md), error**


## C1APITaskV1TaskServiceCreateGrantTask

Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Task.C1APITaskV1TaskServiceCreateGrantTask(ctx, shared.C1APITaskV1TaskServiceCreateGrantRequest{
        AppEntitlementID: conductoroneapi.String("soluta"),
        AppID: conductoroneapi.String("dicta"),
        AppUserID: conductoroneapi.String("laborum"),
        Description: conductoroneapi.String("totam"),
        ExpandMask: &shared.C1APITaskV1TaskExpandMask{
            Paths: []string{
                "aspernatur",
                "dolores",
            },
        },
        GrantDuration: conductoroneapi.String("distinctio"),
        IdentityUserID: conductoroneapi.String("facilis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskServiceCreateGrantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [shared.C1APITaskV1TaskServiceCreateGrantRequest](../../models/shared/c1apitaskv1taskservicecreategrantrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.C1APITaskV1TaskServiceCreateGrantTaskResponse](../../models/operations/c1apitaskv1taskservicecreategranttaskresponse.md), error**


## C1APITaskV1TaskServiceCreateRevokeTask

Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Task.C1APITaskV1TaskServiceCreateRevokeTask(ctx, shared.C1APITaskV1TaskServiceCreateRevokeRequest{
        AppEntitlementID: conductoroneapi.String("aliquid"),
        AppID: conductoroneapi.String("quam"),
        AppUserID: conductoroneapi.String("molestias"),
        Description: conductoroneapi.String("temporibus"),
        ExpandMask: &shared.C1APITaskV1TaskExpandMask{
            Paths: []string{
                "neque",
            },
        },
        IdentityUserID: conductoroneapi.String("fugit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskServiceCreateRevokeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [shared.C1APITaskV1TaskServiceCreateRevokeRequest](../../models/shared/c1apitaskv1taskservicecreaterevokerequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.C1APITaskV1TaskServiceCreateRevokeTaskResponse](../../models/operations/c1apitaskv1taskservicecreaterevoketaskresponse.md), error**


## C1APITaskV1TaskServiceGet

Invokes the c1.api.task.v1.TaskService.Get method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/ConductorOne/conductorone-sdk-go"
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Task.C1APITaskV1TaskServiceGet(ctx, operations.C1APITaskV1TaskServiceGetRequest{
        ID: "2715bf0c-bb1e-431b-8b90-f3443a1108e0",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.C1APITaskV1TaskServiceGetRequest](../../models/operations/c1apitaskv1taskservicegetrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.C1APITaskV1TaskServiceGetResponse](../../models/operations/c1apitaskv1taskservicegetresponse.md), error**

