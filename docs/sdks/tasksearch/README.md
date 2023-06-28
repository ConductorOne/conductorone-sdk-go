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
        AccessReviewIds: []string{
            "quaerat",
            "tempora",
            "vel",
            "quod",
        },
        AccountOwnerIds: []string{
            "qui",
            "dolorum",
            "a",
            "esse",
        },
        ActorID: conductoroneapi.String("harum"),
        AppEntitlementIds: []string{
            "ipsum",
            "quisquam",
        },
        AppResourceIds: []string{
            "amet",
            "tempore",
            "accusamus",
            "numquam",
        },
        AppResourceTypeIds: []string{
            "dolorem",
            "sapiente",
        },
        AppUserSubjectIds: []string{
            "nihil",
            "sit",
            "expedita",
        },
        ApplicationIds: []string{
            "sed",
        },
        AssigneesInIds: []string{
            "libero",
            "voluptas",
        },
        CreatedAfter: types.MustTimeFromString("2022-01-27T14:09:30.399Z"),
        CreatedBefore: types.MustTimeFromString("2022-09-21T15:58:20.330Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepUnspecified.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "maxime",
            "pariatur",
            "soluta",
        },
        ExcludeIds: []string{
            "laborum",
        },
        ExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "incidunt",
                "aspernatur",
                "dolores",
            },
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "facilis",
            "aliquid",
            "quam",
        },
        OpenerIds: []string{
            "temporibus",
            "qui",
            "neque",
        },
        PageSize: conductoroneapi.Float64(1448.47),
        PageToken: conductoroneapi.String("magni"),
        PreviouslyActedOnIds: []string{
            "sunt",
            "ullam",
        },
        Query: conductoroneapi.String("nam"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("0cbb1e31-b8b9-40f3-843a-1108e0adcf4b"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("921879fc-e953-4f73-af7f-bc7abd74dd39"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("c0f5d2cf-f7c7-40a4-9626-d436813f16d9"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("f5fce6c5-5614-46c3-a250-fb008c42e141"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByResource.ToPointer(),
        SubjectIds: []string{
            "placeat",
            "velit",
            "eum",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
        },
        TaskTypes: []shared.TaskType{
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("nulla"),
                    AccessReviewSelection: conductoroneapi.String("voluptas"),
                    AppEntitlementID: conductoroneapi.String("libero"),
                    AppID: conductoroneapi.String("quasi"),
                    AppUserID: conductoroneapi.String("tempora"),
                    IdentityUserID: conductoroneapi.String("numquam"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-11-21T13:50:30.142Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("molestiae"),
                    AppID: conductoroneapi.String("magnam"),
                    AppUserID: conductoroneapi.String("odio"),
                    GrantDuration: conductoroneapi.String("eius"),
                    IdentityUserID: conductoroneapi.String("esse"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-08-20T04:54:39.392Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("reprehenderit"),
                    AppID: conductoroneapi.String("quidem"),
                    AppUserID: conductoroneapi.String("fugiat"),
                    IdentityUserID: conductoroneapi.String("ut"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-03-05T04:36:55.913Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-06-29T02:09:48.123Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-09-28T10:00:45.702Z"),
                            LastLogin: types.MustTimeFromString("2022-05-05T02:03:02.582Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("quidem"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("neque"),
                            CertTicketID: conductoroneapi.String("quo"),
                        },
                    },
                },
            },
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("illum"),
                    AccessReviewSelection: conductoroneapi.String("quo"),
                    AppEntitlementID: conductoroneapi.String("fuga"),
                    AppID: conductoroneapi.String("eius"),
                    AppUserID: conductoroneapi.String("eos"),
                    IdentityUserID: conductoroneapi.String("voluptas"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-12-24T22:15:46.522Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("tempora"),
                    AppID: conductoroneapi.String("debitis"),
                    AppUserID: conductoroneapi.String("ipsam"),
                    GrantDuration: conductoroneapi.String("aspernatur"),
                    IdentityUserID: conductoroneapi.String("sequi"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-01-28T07:34:05.998Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("aperiam"),
                    AppID: conductoroneapi.String("distinctio"),
                    AppUserID: conductoroneapi.String("quod"),
                    IdentityUserID: conductoroneapi.String("dignissimos"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-06-25T15:00:14.304Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-01-30T01:01:49.335Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-06-04T04:43:25.138Z"),
                            LastLogin: types.MustTimeFromString("2022-01-15T19:21:50.024Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("dolores"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("deserunt"),
                            CertTicketID: conductoroneapi.String("molestiae"),
                        },
                    },
                },
            },
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("accusantium"),
                    AccessReviewSelection: conductoroneapi.String("porro"),
                    AppEntitlementID: conductoroneapi.String("eum"),
                    AppID: conductoroneapi.String("quas"),
                    AppUserID: conductoroneapi.String("praesentium"),
                    IdentityUserID: conductoroneapi.String("consequuntur"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-04-27T07:00:05.421Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("mollitia"),
                    AppID: conductoroneapi.String("incidunt"),
                    AppUserID: conductoroneapi.String("atque"),
                    GrantDuration: conductoroneapi.String("explicabo"),
                    IdentityUserID: conductoroneapi.String("minima"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeGranted.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-01-16T21:53:08.951Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("consequuntur"),
                    AppID: conductoroneapi.String("ratione"),
                    AppUserID: conductoroneapi.String("explicabo"),
                    IdentityUserID: conductoroneapi.String("saepe"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-10-25T15:36:05.649Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-02-02T15:37:41.657Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-10-02T03:30:08.260Z"),
                            LastLogin: types.MustTimeFromString("2022-03-14T20:41:04.723Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("nam"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("vero"),
                            CertTicketID: conductoroneapi.String("aliquid"),
                        },
                    },
                },
            },
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("quasi"),
                    AccessReviewSelection: conductoroneapi.String("saepe"),
                    AppEntitlementID: conductoroneapi.String("vel"),
                    AppID: conductoroneapi.String("harum"),
                    AppUserID: conductoroneapi.String("molestiae"),
                    IdentityUserID: conductoroneapi.String("rerum"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-04-14T13:42:03.513Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("eligendi"),
                    AppID: conductoroneapi.String("sit"),
                    AppUserID: conductoroneapi.String("culpa"),
                    GrantDuration: conductoroneapi.String("tempore"),
                    IdentityUserID: conductoroneapi.String("adipisci"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-12-28T10:26:44.811Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("minus"),
                    AppID: conductoroneapi.String("quaerat"),
                    AppUserID: conductoroneapi.String("sapiente"),
                    IdentityUserID: conductoroneapi.String("consectetur"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-10-26T13:57:26.455Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2020-06-05T19:02:55.998Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-02-01T05:26:41.765Z"),
                            LastLogin: types.MustTimeFromString("2022-01-18T13:34:46.515Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("error"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("sint"),
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

