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
                "adipisci",
            },
        },
        AccessReviewIds: []string{
            "id",
            "suscipit",
            "velit",
        },
        AccountOwnerIds: []string{
            "est",
            "recusandae",
            "totam",
        },
        ActorID: conductoroneapi.String("fugiat"),
        AppEntitlementIds: []string{
            "ducimus",
            "quos",
        },
        AppResourceIds: []string{
            "labore",
            "possimus",
        },
        AppResourceTypeIds: []string{
            "cum",
            "commodi",
            "in",
        },
        AppUserSubjectIds: []string{
            "reiciendis",
            "assumenda",
        },
        ApplicationIds: []string{
            "recusandae",
            "aliquid",
        },
        AssigneesInIds: []string{
            "cum",
        },
        CreatedAfter: types.MustTimeFromString("2022-07-21T02:01:44.496Z"),
        CreatedBefore: types.MustTimeFromString("2022-01-23T21:23:20.237Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusAll.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "suscipit",
            "reiciendis",
            "quidem",
            "saepe",
        },
        ExcludeIds: []string{
            "dolore",
            "sunt",
            "asperiores",
            "adipisci",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "amet",
        },
        OpenerIds: []string{
            "dignissimos",
        },
        PageSize: conductoroneapi.Float64(9509.53),
        PageToken: conductoroneapi.String("debitis"),
        PreviouslyActedOnIds: []string{
            "corporis",
        },
        Query: conductoroneapi.String("harum"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("0eb1ea42-6555-4ba3-8287-44ed53b88f3a"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("8d8f5c0b-2f2f-4b7b-994a-276b26916fe1"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccountOwner.ToPointer(),
        SubjectIds: []string{
            "corrupti",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.TaskType{
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("ipsum"),
                    AccessReviewSelection: conductoroneapi.String("ea"),
                    AppEntitlementID: conductoroneapi.String("occaecati"),
                    AppID: conductoroneapi.String("quos"),
                    AppUserID: conductoroneapi.String("voluptatibus"),
                    IdentityUserID: conductoroneapi.String("tempora"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeCertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-01-11T22:08:32.388Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("ex"),
                    AppID: conductoroneapi.String("sit"),
                    AppUserID: conductoroneapi.String("non"),
                    GrantDuration: conductoroneapi.String("officiis"),
                    IdentityUserID: conductoroneapi.String("praesentium"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-09-21T14:42:08.721Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-02-08T08:59:54.184Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-12-12T15:19:28.678Z"),
                            LastLogin: types.MustTimeFromString("2021-02-14T22:16:10.503Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("veniam"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("minima"),
                            CertTicketID: conductoroneapi.String("recusandae"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("reiciendis"),
                    AppID: conductoroneapi.String("nulla"),
                    AppUserID: conductoroneapi.String("magni"),
                    IdentityUserID: conductoroneapi.String("aperiam"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-09-02T13:45:48.565Z"),
                },
            },
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("in"),
                    AccessReviewSelection: conductoroneapi.String("officiis"),
                    AppEntitlementID: conductoroneapi.String("beatae"),
                    AppID: conductoroneapi.String("laudantium"),
                    AppUserID: conductoroneapi.String("exercitationem"),
                    IdentityUserID: conductoroneapi.String("praesentium"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-04-27T14:41:34.966Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("voluptatum"),
                    AppID: conductoroneapi.String("error"),
                    AppUserID: conductoroneapi.String("hic"),
                    GrantDuration: conductoroneapi.String("expedita"),
                    IdentityUserID: conductoroneapi.String("debitis"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeGranted.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-04-26T13:26:55.921Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-08-25T08:13:54.077Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-03-30T03:48:24.857Z"),
                            LastLogin: types.MustTimeFromString("2022-06-16T13:22:48.429Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("fugit"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("ut"),
                            CertTicketID: conductoroneapi.String("fugiat"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("voluptatem"),
                    AppID: conductoroneapi.String("culpa"),
                    AppUserID: conductoroneapi.String("expedita"),
                    IdentityUserID: conductoroneapi.String("magnam"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-08-17T23:19:52.079Z"),
                },
            },
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("sit"),
                    AccessReviewSelection: conductoroneapi.String("voluptatum"),
                    AppEntitlementID: conductoroneapi.String("quas"),
                    AppID: conductoroneapi.String("repudiandae"),
                    AppUserID: conductoroneapi.String("corporis"),
                    IdentityUserID: conductoroneapi.String("et"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-11-05T22:13:21.002Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("sit"),
                    AppID: conductoroneapi.String("vel"),
                    AppUserID: conductoroneapi.String("nostrum"),
                    GrantDuration: conductoroneapi.String("saepe"),
                    IdentityUserID: conductoroneapi.String("error"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-01-12T08:45:01.658Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-04-23T19:46:15.105Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-11-25T17:06:43.735Z"),
                            LastLogin: types.MustTimeFromString("2022-06-03T11:07:35.218Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("quidem"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("atque"),
                            CertTicketID: conductoroneapi.String("laborum"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("nam"),
                    AppID: conductoroneapi.String("tenetur"),
                    AppUserID: conductoroneapi.String("laboriosam"),
                    IdentityUserID: conductoroneapi.String("alias"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeRevoked.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-02-02T22:51:17.351Z"),
                },
            },
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("unde"),
                    AccessReviewSelection: conductoroneapi.String("reiciendis"),
                    AppEntitlementID: conductoroneapi.String("provident"),
                    AppID: conductoroneapi.String("repellendus"),
                    AppUserID: conductoroneapi.String("delectus"),
                    IdentityUserID: conductoroneapi.String("voluptates"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-08-10T13:37:39.961Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("reprehenderit"),
                    AppID: conductoroneapi.String("facere"),
                    AppUserID: conductoroneapi.String("fuga"),
                    GrantDuration: conductoroneapi.String("praesentium"),
                    IdentityUserID: conductoroneapi.String("mollitia"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeGranted.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-03-18T08:14:24.399Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-09-16T12:27:20.507Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-02-12T07:32:08.165Z"),
                            LastLogin: types.MustTimeFromString("2021-06-10T15:39:04.964Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("suscipit"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("quidem"),
                            CertTicketID: conductoroneapi.String("maxime"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("et"),
                    AppID: conductoroneapi.String("esse"),
                    AppUserID: conductoroneapi.String("amet"),
                    IdentityUserID: conductoroneapi.String("assumenda"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-10-02T23:52:38.012Z"),
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

