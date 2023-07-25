# TaskSearch

### Available Operations

* [Search](#search) - Search

## Search

Invokes the c1.api.task.v1.TaskSearchService.Search method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/types"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TaskSearch.Search(ctx, shared.TaskSearchRequest{
        TaskExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "numquam",
            },
        },
        AccessReviewIds: []string{
            "nesciunt",
            "at",
        },
        AccountOwnerIds: []string{
            "dignissimos",
            "optio",
            "necessitatibus",
        },
        ActorID: conductoroneapi.String("corporis"),
        AppEntitlementIds: []string{
            "expedita",
        },
        AppResourceIds: []string{
            "cupiditate",
            "minima",
            "placeat",
        },
        AppResourceTypeIds: []string{
            "neque",
            "in",
        },
        AppUserSubjectIds: []string{
            "eum",
            "modi",
            "corporis",
            "magnam",
        },
        ApplicationIds: []string{
            "maiores",
            "tempore",
            "aperiam",
            "libero",
        },
        AssigneesInIds: []string{
            "labore",
        },
        CreatedAfter: types.MustTimeFromString("2021-11-05T06:31:50.944Z"),
        CreatedBefore: types.MustTimeFromString("2022-03-22T14:23:17.198Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepUnspecified.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusEmergency.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "nostrum",
            "est",
            "impedit",
        },
        ExcludeIds: []string{
            "tempore",
            "vero",
            "odit",
            "repellat",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "nemo",
            "reprehenderit",
            "aperiam",
            "odio",
        },
        OpenerIds: []string{
            "in",
            "ducimus",
        },
        PageSize: conductoroneapi.Float64(5678.46),
        PageToken: conductoroneapi.String("dolores"),
        PreviouslyActedOnIds: []string{
            "veritatis",
            "ducimus",
            "voluptate",
        },
        Query: conductoroneapi.String("pariatur"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("ac646ecb-5734-409e-beb1-e5a2b12eb07f"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("116db995-45fc-495f-a889-70e189dbb30f"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("cb33ea05-5b19-47cd-84e2-f52d82d3513b"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("b6f48b65-6bcd-4b35-bf2e-4b27537a8cd9"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccountOwner.ToPointer(),
        SubjectIds: []string{
            "dolor",
            "dicta",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
        },
        TaskTypes: []shared.TaskType{
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("fugiat"),
                    AccessReviewSelection: conductoroneapi.String("ad"),
                    AppEntitlementID: conductoroneapi.String("aspernatur"),
                    AppID: conductoroneapi.String("enim"),
                    AppUserID: conductoroneapi.String("delectus"),
                    IdentityUserID: conductoroneapi.String("iusto"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-11-09T04:25:16.566Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("ab"),
                    AppID: conductoroneapi.String("incidunt"),
                    AppUserID: conductoroneapi.String("accusamus"),
                    GrantDuration: conductoroneapi.String("saepe"),
                    IdentityUserID: conductoroneapi.String("tempore"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeGranted.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-01-11T22:06:57.023Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-09-01T16:29:21.118Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-04-08T01:20:58.799Z"),
                            LastLogin: types.MustTimeFromString("2020-08-19T04:51:15.983Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("sequi"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("nihil"),
                            CertTicketID: conductoroneapi.String("deleniti"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("illo"),
                    AppID: conductoroneapi.String("labore"),
                    AppUserID: conductoroneapi.String("assumenda"),
                    IdentityUserID: conductoroneapi.String("aliquam"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-12-21T01:19:26.741Z"),
                },
            },
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("repudiandae"),
                    AccessReviewSelection: conductoroneapi.String("consequatur"),
                    AppEntitlementID: conductoroneapi.String("maxime"),
                    AppID: conductoroneapi.String("aspernatur"),
                    AppUserID: conductoroneapi.String("nam"),
                    IdentityUserID: conductoroneapi.String("expedita"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-02-26T17:19:48.572Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("rerum"),
                    AppID: conductoroneapi.String("dignissimos"),
                    AppUserID: conductoroneapi.String("corporis"),
                    GrantDuration: conductoroneapi.String("vero"),
                    IdentityUserID: conductoroneapi.String("similique"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-10-14T22:48:07.675Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-03-25T12:03:53.573Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-12-27T13:22:49.195Z"),
                            LastLogin: types.MustTimeFromString("2022-09-05T04:00:47.034Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("quae"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("amet"),
                            CertTicketID: conductoroneapi.String("illum"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("praesentium"),
                    AppID: conductoroneapi.String("quidem"),
                    AppUserID: conductoroneapi.String("cum"),
                    IdentityUserID: conductoroneapi.String("amet"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-06-27T01:16:52.839Z"),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.TaskSearchRequest](../../models/shared/tasksearchrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.C1APITaskV1TaskSearchServiceSearchResponse](../../models/operations/c1apitaskv1tasksearchservicesearchresponse.md), error**

