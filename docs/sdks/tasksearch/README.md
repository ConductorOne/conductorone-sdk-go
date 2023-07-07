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
        ExcludeAppEntitlementIds: []string{
            "doloribus",
            "suscipit",
        },
        ExcludeIds: []string{
            "quidem",
            "saepe",
            "necessitatibus",
            "dolore",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "asperiores",
        },
        OpenerIds: []string{
            "non",
        },
        PageSize: conductoroneapi.Float64(2282.63),
        PageToken: conductoroneapi.String("beatae"),
        PreviouslyActedOnIds: []string{
            "a",
            "debitis",
        },
        Query: conductoroneapi.String("consectetur"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("b60eb1ea-4265-455b-a3c2-8744ed53b88f"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("3a8d8f5c-0b2f-42fb-bb19-4a276b26916f"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccountOwner.ToPointer(),
        SubjectIds: []string{
            "reiciendis",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
        },
        TaskTypes: []shared.TaskType{
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("incidunt"),
                    AccessReviewSelection: conductoroneapi.String("sed"),
                    AppEntitlementID: conductoroneapi.String("provident"),
                    AppID: conductoroneapi.String("eius"),
                    AppUserID: conductoroneapi.String("necessitatibus"),
                    IdentityUserID: conductoroneapi.String("ipsum"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-11-23T23:35:18.899Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("voluptatibus"),
                    AppID: conductoroneapi.String("tempora"),
                    AppUserID: conductoroneapi.String("tempora"),
                    GrantDuration: conductoroneapi.String("voluptate"),
                    IdentityUserID: conductoroneapi.String("reiciendis"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-10-02T07:53:52.364Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-06-26T01:49:52.614Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-05-19T10:07:39.882Z"),
                            LastLogin: types.MustTimeFromString("2022-08-17T20:12:51.918Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("debitis"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("rem"),
                            CertTicketID: conductoroneapi.String("sit"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("nobis"),
                    AppID: conductoroneapi.String("error"),
                    AppUserID: conductoroneapi.String("veniam"),
                    IdentityUserID: conductoroneapi.String("minima"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2020-05-31T18:16:06.681Z"),
                },
            },
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("magni"),
                    AccessReviewSelection: conductoroneapi.String("aperiam"),
                    AppEntitlementID: conductoroneapi.String("saepe"),
                    AppID: conductoroneapi.String("numquam"),
                    AppUserID: conductoroneapi.String("veniam"),
                    IdentityUserID: conductoroneapi.String("in"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-06-27T22:55:23.952Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("exercitationem"),
                    AppID: conductoroneapi.String("praesentium"),
                    AppUserID: conductoroneapi.String("cum"),
                    GrantDuration: conductoroneapi.String("laboriosam"),
                    IdentityUserID: conductoroneapi.String("dolorum"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-02-10T08:43:29.852Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-03-20T05:03:12.319Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-04-28T20:28:39.956Z"),
                            LastLogin: types.MustTimeFromString("2022-05-12T18:07:12.039Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("dolorum"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("corrupti"),
                            CertTicketID: conductoroneapi.String("accusamus"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("tempora"),
                    AppID: conductoroneapi.String("atque"),
                    AppUserID: conductoroneapi.String("fugit"),
                    IdentityUserID: conductoroneapi.String("ut"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-05-14T04:54:08.545Z"),
                },
            },
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("expedita"),
                    AccessReviewSelection: conductoroneapi.String("magnam"),
                    AppEntitlementID: conductoroneapi.String("consequatur"),
                    AppID: conductoroneapi.String("esse"),
                    AppUserID: conductoroneapi.String("ipsam"),
                    IdentityUserID: conductoroneapi.String("sit"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-02-26T20:36:25.696Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("corporis"),
                    AppID: conductoroneapi.String("et"),
                    AppUserID: conductoroneapi.String("blanditiis"),
                    GrantDuration: conductoroneapi.String("ex"),
                    IdentityUserID: conductoroneapi.String("sed"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-08-28T22:43:39.407Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-02-18T15:46:03.371Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-09-21T03:21:52.611Z"),
                            LastLogin: types.MustTimeFromString("2022-05-16T07:46:15.799Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("harum"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("dicta"),
                            CertTicketID: conductoroneapi.String("architecto"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("occaecati"),
                    AppID: conductoroneapi.String("labore"),
                    AppUserID: conductoroneapi.String("quidem"),
                    IdentityUserID: conductoroneapi.String("atque"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-02-07T07:57:17.421Z"),
                },
            },
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("laboriosam"),
                    AccessReviewSelection: conductoroneapi.String("alias"),
                    AppEntitlementID: conductoroneapi.String("amet"),
                    AppID: conductoroneapi.String("deserunt"),
                    AppUserID: conductoroneapi.String("voluptate"),
                    IdentityUserID: conductoroneapi.String("unde"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-05-02T07:30:16.176Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("delectus"),
                    AppID: conductoroneapi.String("voluptates"),
                    AppUserID: conductoroneapi.String("perferendis"),
                    GrantDuration: conductoroneapi.String("est"),
                    IdentityUserID: conductoroneapi.String("quidem"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2020-12-11T19:46:56.272Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-09-14T12:33:27.065Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-12-21T09:05:01.168Z"),
                            LastLogin: types.MustTimeFromString("2020-03-30T02:42:49.718Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("quasi"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("atque"),
                            CertTicketID: conductoroneapi.String("reprehenderit"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("asperiores"),
                    AppID: conductoroneapi.String("totam"),
                    AppUserID: conductoroneapi.String("suscipit"),
                    IdentityUserID: conductoroneapi.String("quidem"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-07-16T17:34:48.115Z"),
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

