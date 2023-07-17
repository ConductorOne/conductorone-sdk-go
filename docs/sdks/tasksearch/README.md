# TaskSearch

### Available Operations

* [Search](#search) - Invokes the c1.api.task.v1.TaskSearchService.Search method.

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
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.TaskSearch.Search(ctx, shared.TaskSearchRequest{
        TaskExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "fuga",
            },
        },
        AccessReviewIds: []string{
            "voluptatibus",
            "distinctio",
            "omnis",
            "delectus",
        },
        AccountOwnerIds: []string{
            "praesentium",
            "maxime",
        },
        ActorID: conductoroneapi.String("magnam"),
        AppEntitlementIds: []string{
            "quos",
            "commodi",
            "itaque",
            "commodi",
        },
        AppResourceIds: []string{
            "earum",
            "modi",
            "nam",
        },
        AppResourceTypeIds: []string{
            "voluptatem",
            "ipsam",
            "vel",
            "alias",
        },
        AppUserSubjectIds: []string{
            "non",
        },
        ApplicationIds: []string{
            "enim",
            "sint",
            "nulla",
            "deserunt",
        },
        AssigneesInIds: []string{
            "nemo",
            "reprehenderit",
        },
        CreatedAfter: types.MustTimeFromString("2022-04-30T15:29:17.062Z"),
        CreatedBefore: types.MustTimeFromString("2021-03-29T02:31:09.447Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusEmergency.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "asperiores",
            "ex",
            "voluptas",
            "debitis",
        },
        ExcludeIds: []string{
            "quae",
            "minus",
            "fuga",
            "laborum",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "velit",
        },
        OpenerIds: []string{
            "ipsum",
            "impedit",
            "magni",
        },
        PageSize: conductoroneapi.Float64(7465.85),
        PageToken: conductoroneapi.String("repudiandae"),
        PreviouslyActedOnIds: []string{
            "dolore",
            "iusto",
            "voluptate",
        },
        Query: conductoroneapi.String("sequi"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("3c8d72f6-4d1d-4b1f-ac43-10661e96349e"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("1cf9e06e-3a43-4700-8ae6-b6bc9b8f759e"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByResource.ToPointer(),
        SubjectIds: []string{
            "ipsam",
            "corporis",
            "est",
            "error",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.TaskType{
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("consectetur"),
                    AccessReviewSelection: conductoroneapi.String("vitae"),
                    AppEntitlementID: conductoroneapi.String("inventore"),
                    AppID: conductoroneapi.String("dolorem"),
                    AppUserID: conductoroneapi.String("ad"),
                    IdentityUserID: conductoroneapi.String("qui"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-08-19T19:35:40.341Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("soluta"),
                    AppID: conductoroneapi.String("libero"),
                    AppUserID: conductoroneapi.String("rem"),
                    GrantDuration: conductoroneapi.String("dolorum"),
                    IdentityUserID: conductoroneapi.String("odio"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-10-31T15:56:50.300Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-12-08T11:32:42.651Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-09-26T12:26:41.334Z"),
                            LastLogin: types.MustTimeFromString("2022-08-26T20:23:40.892Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("itaque"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("et"),
                            CertTicketID: conductoroneapi.String("ipsum"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("unde"),
                    AppID: conductoroneapi.String("nulla"),
                    AppUserID: conductoroneapi.String("distinctio"),
                    IdentityUserID: conductoroneapi.String("maxime"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-08-29T02:52:30.806Z"),
                },
            },
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("omnis"),
                    AccessReviewSelection: conductoroneapi.String("libero"),
                    AppEntitlementID: conductoroneapi.String("dicta"),
                    AppID: conductoroneapi.String("id"),
                    AppUserID: conductoroneapi.String("libero"),
                    IdentityUserID: conductoroneapi.String("fugiat"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-05-18T06:50:10.787Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("sit"),
                    AppID: conductoroneapi.String("iusto"),
                    AppUserID: conductoroneapi.String("ipsa"),
                    GrantDuration: conductoroneapi.String("voluptates"),
                    IdentityUserID: conductoroneapi.String("inventore"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-06-01T04:35:31.707Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2020-11-09T00:25:03.486Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-08-01T09:21:47.394Z"),
                            LastLogin: types.MustTimeFromString("2022-10-29T20:56:45.429Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("assumenda"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("beatae"),
                            CertTicketID: conductoroneapi.String("est"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("facere"),
                    AppID: conductoroneapi.String("corrupti"),
                    AppUserID: conductoroneapi.String("molestiae"),
                    IdentityUserID: conductoroneapi.String("provident"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2020-10-20T01:09:41.345Z"),
                },
            },
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("sint"),
                    AccessReviewSelection: conductoroneapi.String("ea"),
                    AppEntitlementID: conductoroneapi.String("autem"),
                    AppID: conductoroneapi.String("ipsam"),
                    AppUserID: conductoroneapi.String("rerum"),
                    IdentityUserID: conductoroneapi.String("laudantium"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeCertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2020-01-31T15:52:45.660Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("cum"),
                    AppID: conductoroneapi.String("at"),
                    AppUserID: conductoroneapi.String("alias"),
                    GrantDuration: conductoroneapi.String("quia"),
                    IdentityUserID: conductoroneapi.String("quidem"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-11-24T16:37:12.153Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-03-25T19:10:06.582Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-02-26T15:04:57.243Z"),
                            LastLogin: types.MustTimeFromString("2022-06-28T19:16:42.798Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("odit"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("explicabo"),
                            CertTicketID: conductoroneapi.String("corporis"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("error"),
                    AppID: conductoroneapi.String("earum"),
                    AppUserID: conductoroneapi.String("adipisci"),
                    IdentityUserID: conductoroneapi.String("recusandae"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-04-22T18:47:14.845Z"),
                },
            },
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("quis"),
                    AccessReviewSelection: conductoroneapi.String("beatae"),
                    AppEntitlementID: conductoroneapi.String("unde"),
                    AppID: conductoroneapi.String("molestiae"),
                    AppUserID: conductoroneapi.String("delectus"),
                    IdentityUserID: conductoroneapi.String("cupiditate"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-09-29T05:24:35.816Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("nesciunt"),
                    AppID: conductoroneapi.String("at"),
                    AppUserID: conductoroneapi.String("officia"),
                    GrantDuration: conductoroneapi.String("dignissimos"),
                    IdentityUserID: conductoroneapi.String("optio"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-10-25T06:51:46.936Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-12-08T03:38:46.630Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-05-07T19:45:21.246Z"),
                            LastLogin: types.MustTimeFromString("2022-01-19T09:15:40.360Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("neque"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("in"),
                            CertTicketID: conductoroneapi.String("minus"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("eum"),
                    AppID: conductoroneapi.String("modi"),
                    AppUserID: conductoroneapi.String("corporis"),
                    IdentityUserID: conductoroneapi.String("magnam"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2020-10-21T17:45:45.582Z"),
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

