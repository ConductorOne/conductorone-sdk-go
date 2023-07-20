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
                "maiores",
                "tempore",
                "aperiam",
                "libero",
            },
        },
        AccessReviewIds: []string{
            "labore",
        },
        AccountOwnerIds: []string{
            "occaecati",
            "voluptas",
            "quo",
        },
        ActorID: conductoroneapi.String("velit"),
        AppEntitlementIds: []string{
            "fuga",
            "nostrum",
            "est",
            "impedit",
        },
        AppResourceIds: []string{
            "tempore",
            "vero",
            "odit",
            "repellat",
        },
        AppResourceTypeIds: []string{
            "nemo",
            "reprehenderit",
            "aperiam",
            "odio",
        },
        AppUserSubjectIds: []string{
            "in",
            "ducimus",
        },
        ApplicationIds: []string{
            "dolores",
            "error",
            "veritatis",
        },
        AssigneesInIds: []string{
            "voluptate",
            "pariatur",
        },
        CreatedAfter: types.MustTimeFromString("2021-02-12T15:07:49.977Z"),
        CreatedBefore: types.MustTimeFromString("2021-10-18T10:28:23.665Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepUnspecified.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusAll.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "placeat",
            "quidem",
            "exercitationem",
            "quam",
        },
        ExcludeIds: []string{
            "modi",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "sint",
        },
        OpenerIds: []string{
            "sequi",
            "repudiandae",
            "cum",
            "dicta",
        },
        PageSize: conductoroneapi.Float64(9369.28),
        PageToken: conductoroneapi.String("veniam"),
        PreviouslyActedOnIds: []string{
            "dolores",
            "nam",
            "dicta",
        },
        Query: conductoroneapi.String("consequuntur"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("b07f116d-b995-445f-895f-a88970e189db"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("b30fcb33-ea05-45b1-97cd-44e2f52d82d3"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("513bb6f4-8b65-46bc-9b35-ff2e4b27537a"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("8cd9e731-9c17-47d5-a5f7-7b114eeb52ff"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccount.ToPointer(),
        SubjectIds: []string{
            "nemo",
            "repellat",
            "quisquam",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
        },
        TaskTypes: []shared.TaskType{
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("illo"),
                    AccessReviewSelection: conductoroneapi.String("labore"),
                    AppEntitlementID: conductoroneapi.String("assumenda"),
                    AppID: conductoroneapi.String("aliquam"),
                    AppUserID: conductoroneapi.String("quisquam"),
                    IdentityUserID: conductoroneapi.String("provident"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-12-21T09:31:19.686Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("maxime"),
                    AppID: conductoroneapi.String("aspernatur"),
                    AppUserID: conductoroneapi.String("nam"),
                    GrantDuration: conductoroneapi.String("expedita"),
                    IdentityUserID: conductoroneapi.String("quas"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2020-11-25T10:49:49.319Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-08-22T06:12:50.001Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-02-10T19:58:48.033Z"),
                            LastLogin: types.MustTimeFromString("2021-09-11T20:45:20.947Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("dolorem"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("commodi"),
                            CertTicketID: conductoroneapi.String("impedit"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("commodi"),
                    AppID: conductoroneapi.String("aut"),
                    AppUserID: conductoroneapi.String("voluptatem"),
                    IdentityUserID: conductoroneapi.String("ad"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-02-24T23:57:50.783Z"),
                },
            },
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("praesentium"),
                    AccessReviewSelection: conductoroneapi.String("quidem"),
                    AppEntitlementID: conductoroneapi.String("cum"),
                    AppID: conductoroneapi.String("amet"),
                    AppUserID: conductoroneapi.String("quasi"),
                    IdentityUserID: conductoroneapi.String("dicta"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-01-23T11:30:38.329Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("iusto"),
                    AppID: conductoroneapi.String("amet"),
                    AppUserID: conductoroneapi.String("provident"),
                    GrantDuration: conductoroneapi.String("dolorum"),
                    IdentityUserID: conductoroneapi.String("necessitatibus"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-12-20T23:31:33.018Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-07-11T07:45:44.709Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2020-11-15T21:30:50.859Z"),
                            LastLogin: types.MustTimeFromString("2022-11-29T21:33:00.572Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("sint"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("accusamus"),
                            CertTicketID: conductoroneapi.String("eos"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("totam"),
                    AppID: conductoroneapi.String("dicta"),
                    AppUserID: conductoroneapi.String("voluptatem"),
                    IdentityUserID: conductoroneapi.String("velit"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeRevoked.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-01-17T21:48:03.045Z"),
                },
            },
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("dolor"),
                    AccessReviewSelection: conductoroneapi.String("occaecati"),
                    AppEntitlementID: conductoroneapi.String("atque"),
                    AppID: conductoroneapi.String("beatae"),
                    AppUserID: conductoroneapi.String("at"),
                    IdentityUserID: conductoroneapi.String("labore"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-12-20T04:33:32.836Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("perferendis"),
                    AppID: conductoroneapi.String("rerum"),
                    AppUserID: conductoroneapi.String("ea"),
                    GrantDuration: conductoroneapi.String("aperiam"),
                    IdentityUserID: conductoroneapi.String("dignissimos"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-03-20T19:09:29.668Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-07-16T02:11:27.559Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-07-13T07:54:07.156Z"),
                            LastLogin: types.MustTimeFromString("2022-04-04T09:40:23.441Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("natus"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("temporibus"),
                            CertTicketID: conductoroneapi.String("officia"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("amet"),
                    AppID: conductoroneapi.String("tenetur"),
                    AppUserID: conductoroneapi.String("aspernatur"),
                    IdentityUserID: conductoroneapi.String("quo"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2020-12-27T12:23:00.742Z"),
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

