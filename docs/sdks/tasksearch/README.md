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
            "id",
            "saepe",
        },
        AccountOwnerIds: []string{
            "aspernatur",
            "perferendis",
        },
        ActorID: conductoroneapi.String("amet"),
        AppEntitlementIds: []string{
            "accusamus",
            "ad",
            "saepe",
            "suscipit",
        },
        AppResourceIds: []string{
            "provident",
            "minima",
            "repellendus",
        },
        AppResourceTypeIds: []string{
            "similique",
            "alias",
            "at",
        },
        AppUserSubjectIds: []string{
            "tempora",
            "vel",
        },
        ApplicationIds: []string{
            "officiis",
            "qui",
            "dolorum",
            "a",
        },
        AssigneesInIds: []string{
            "harum",
            "iusto",
        },
        CreatedAfter: types.MustTimeFromString("2022-03-19T02:38:33.395Z"),
        CreatedBefore: types.MustTimeFromString("2022-04-24T18:16:06.669Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "numquam",
            "enim",
            "dolorem",
            "sapiente",
        },
        ExcludeIds: []string{
            "nihil",
            "sit",
            "expedita",
        },
        ExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "sed",
            },
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "libero",
            "voluptas",
        },
        OpenerIds: []string{
            "quam",
            "ipsum",
            "incidunt",
        },
        PageSize: conductoroneapi.Float64(1864.58),
        PageToken: conductoroneapi.String("cupiditate"),
        PreviouslyActedOnIds: []string{
            "pariatur",
            "soluta",
            "dicta",
            "laborum",
        },
        Query: conductoroneapi.String("totam"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("22bb679d-2322-4715-bf0c-bb1e31b8b90f"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("3443a110-8e0a-4dcf-8b92-1879fce953f7"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByUnspecified.ToPointer(),
        SubjectIds: []string{
            "tenetur",
            "dignissimos",
            "hic",
            "distinctio",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
        },
        TaskTypes: []shared.TaskType{
            shared.TaskType{
                TaskType: &shared.TaskTypeTaskType{
                    Certify: &shared.TaskTypeCertify{
                        AccessReviewID: conductoroneapi.String("dolore"),
                        AccessReviewSelection: conductoroneapi.String("quibusdam"),
                        AppEntitlementID: conductoroneapi.String("illum"),
                        AppID: conductoroneapi.String("sequi"),
                        AppUserID: conductoroneapi.String("natus"),
                        IdentityUserID: conductoroneapi.String("impedit"),
                        Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeUnspecified.ToPointer(),
                        OutcomeTime: types.MustTimeFromString("2021-12-16T18:42:11.269Z"),
                    },
                    Grant: &shared.TaskTypeGrant{
                        AppEntitlementID: conductoroneapi.String("nulla"),
                        AppID: conductoroneapi.String("fugit"),
                        AppUserID: conductoroneapi.String("porro"),
                        GrantDuration: conductoroneapi.String("maiores"),
                        IdentityUserID: conductoroneapi.String("doloribus"),
                        Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeDenied.ToPointer(),
                        OutcomeTime: types.MustTimeFromString("2021-07-05T08:33:16.995Z"),
                    },
                    Revoke: &shared.TaskTypeRevoke{
                        AppEntitlementID: conductoroneapi.String("alias"),
                        AppID: conductoroneapi.String("officia"),
                        AppUserID: conductoroneapi.String("tempora"),
                        IdentityUserID: conductoroneapi.String("ipsam"),
                        Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeDenied.ToPointer(),
                        OutcomeTime: types.MustTimeFromString("2022-07-28T16:45:32.822Z"),
                        Source: &shared.TaskRevokeSource{
                            Origin: &shared.TaskRevokeSourceOrigin{
                                Expired: &shared.TaskRevokeSourceExpired{
                                    ExpiredAt: types.MustTimeFromString("2022-02-08T20:43:00.221Z"),
                                },
                                NonUsage: &shared.TaskRevokeSourceNonUsage{
                                    ExpiresAt: types.MustTimeFromString("2022-08-07T14:04:49.880Z"),
                                    LastLogin: types.MustTimeFromString("2022-10-04T22:05:21.038Z"),
                                },
                                Request: &shared.TaskRevokeSourceRequest{
                                    RequestUserID: conductoroneapi.String("dolor"),
                                },
                                Review: &shared.TaskRevokeSourceReview{
                                    AccessReviewID: conductoroneapi.String("maiores"),
                                    CertTicketID: conductoroneapi.String("quasi"),
                                },
                            },
                        },
                    },
                },
            },
            shared.TaskType{
                TaskType: &shared.TaskTypeTaskType{
                    Certify: &shared.TaskTypeCertify{
                        AccessReviewID: conductoroneapi.String("ex"),
                        AccessReviewSelection: conductoroneapi.String("nulla"),
                        AppEntitlementID: conductoroneapi.String("excepturi"),
                        AppID: conductoroneapi.String("voluptatibus"),
                        AppUserID: conductoroneapi.String("nostrum"),
                        IdentityUserID: conductoroneapi.String("sapiente"),
                        Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeError.ToPointer(),
                        OutcomeTime: types.MustTimeFromString("2021-10-07T13:07:57.949Z"),
                    },
                    Grant: &shared.TaskTypeGrant{
                        AppEntitlementID: conductoroneapi.String("impedit"),
                        AppID: conductoroneapi.String("corporis"),
                        AppUserID: conductoroneapi.String("veniam"),
                        GrantDuration: conductoroneapi.String("aliquid"),
                        IdentityUserID: conductoroneapi.String("inventore"),
                        Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeGranted.ToPointer(),
                        OutcomeTime: types.MustTimeFromString("2022-03-24T01:04:28.722Z"),
                    },
                    Revoke: &shared.TaskTypeRevoke{
                        AppEntitlementID: conductoroneapi.String("consectetur"),
                        AppID: conductoroneapi.String("recusandae"),
                        AppUserID: conductoroneapi.String("aspernatur"),
                        IdentityUserID: conductoroneapi.String("minima"),
                        Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeUnspecified.ToPointer(),
                        OutcomeTime: types.MustTimeFromString("2020-10-28T11:22:47.751Z"),
                        Source: &shared.TaskRevokeSource{
                            Origin: &shared.TaskRevokeSourceOrigin{
                                Expired: &shared.TaskRevokeSourceExpired{
                                    ExpiredAt: types.MustTimeFromString("2022-12-27T19:53:42.787Z"),
                                },
                                NonUsage: &shared.TaskRevokeSourceNonUsage{
                                    ExpiresAt: types.MustTimeFromString("2021-06-17T11:25:30.782Z"),
                                    LastLogin: types.MustTimeFromString("2022-11-08T08:44:49.857Z"),
                                },
                                Request: &shared.TaskRevokeSourceRequest{
                                    RequestUserID: conductoroneapi.String("accusamus"),
                                },
                                Review: &shared.TaskRevokeSourceReview{
                                    AccessReviewID: conductoroneapi.String("inventore"),
                                    CertTicketID: conductoroneapi.String("non"),
                                },
                            },
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

