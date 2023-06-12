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
            "maxime",
            "ea",
            "excepturi",
            "odit",
        },
        AccountOwnerIds: []string{
            "accusantium",
            "ab",
        },
        ActorID: conductoroneapi.String("maiores"),
        AppEntitlementIds: []string{
            "ipsam",
            "voluptate",
            "autem",
        },
        AppResourceIds: []string{
            "eaque",
            "pariatur",
            "nemo",
        },
        AppResourceTypeIds: []string{
            "perferendis",
            "fugiat",
            "amet",
            "aut",
        },
        AppUserSubjectIds: []string{
            "corporis",
            "hic",
            "libero",
            "nobis",
        },
        ApplicationIds: []string{
            "quis",
        },
        AssigneesInIds: []string{
            "dignissimos",
            "eaque",
            "quis",
        },
        CreatedAfter: types.MustTimeFromString("2022-10-27T11:39:54.088Z"),
        CreatedBefore: types.MustTimeFromString("2022-10-30T14:09:25.982Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "dolor",
            "vero",
        },
        ExcludeIds: []string{
            "hic",
            "recusandae",
        },
        ExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "facilis",
                "perspiciatis",
                "voluptatem",
            },
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "consequuntur",
            "blanditiis",
            "error",
            "eaque",
        },
        OpenerIds: []string{
            "rerum",
            "adipisci",
            "asperiores",
        },
        PageSize: conductoroneapi.Float64(9342.14),
        PageToken: conductoroneapi.String("modi"),
        PreviouslyActedOnIds: []string{
            "dolorum",
            "deleniti",
            "pariatur",
        },
        Query: conductoroneapi.String("provident"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("bf486333-23f9-4b77-b3a4-100674ebf692"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("80d1ba77-a89e-4bf7-b7ae-4203ce5e6a95"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("d8a0d446-ce2a-4f7a-b3cf-3be453f870b3"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("26b5a734-29cd-4b1a-8422-bb679d232271"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccount.ToPointer(),
        SubjectIds: []string{
            "hic",
            "voluptatem",
            "cumque",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
        },
        TaskTypes: []shared.TaskType{
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("veritatis"),
                    AccessReviewSelection: conductoroneapi.String("nobis"),
                    AppEntitlementID: conductoroneapi.String("quos"),
                    AppID: conductoroneapi.String("tempore"),
                    AppUserID: conductoroneapi.String("cupiditate"),
                    IdentityUserID: conductoroneapi.String("aperiam"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-09-16T08:47:14.764Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("labore"),
                    AppID: conductoroneapi.String("adipisci"),
                    AppUserID: conductoroneapi.String("dolorum"),
                    GrantDuration: conductoroneapi.String("architecto"),
                    IdentityUserID: conductoroneapi.String("quae"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-02-21T10:38:46.777Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("consequatur"),
                    AppID: conductoroneapi.String("est"),
                    AppUserID: conductoroneapi.String("repellendus"),
                    IdentityUserID: conductoroneapi.String("porro"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-04-19T05:23:08.152Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-08-21T09:49:15.782Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-06-27T23:25:59.184Z"),
                            LastLogin: types.MustTimeFromString("2022-06-03T03:16:58.870Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("voluptatibus"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("quisquam"),
                            CertTicketID: conductoroneapi.String("vero"),
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

