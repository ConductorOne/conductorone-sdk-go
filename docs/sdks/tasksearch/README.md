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
            "voluptate",
            "expedita",
            "ab",
        },
        AccountOwnerIds: []string{
            "dolore",
            "laborum",
            "sed",
        },
        ActorID: conductoroneapi.String("in"),
        AppEntitlementIds: []string{
            "quidem",
            "explicabo",
        },
        AppResourceIds: []string{
            "unde",
            "architecto",
        },
        AppResourceTypeIds: []string{
            "sapiente",
            "debitis",
        },
        AppUserSubjectIds: []string{
            "reiciendis",
        },
        ApplicationIds: []string{
            "corrupti",
        },
        AssigneesInIds: []string{
            "incidunt",
            "sed",
            "provident",
            "eius",
        },
        CreatedAfter: types.MustTimeFromString("2022-05-09T23:54:09.125Z"),
        CreatedBefore: types.MustTimeFromString("2022-06-03T07:58:02.123Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepApproval.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "tempora",
            "tempora",
            "voluptate",
            "reiciendis",
        },
        ExcludeIds: []string{
            "sit",
            "non",
        },
        ExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "praesentium",
                "facilis",
                "quaerat",
                "incidunt",
            },
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "debitis",
            "rem",
        },
        OpenerIds: []string{
            "nobis",
        },
        PageSize: conductoroneapi.Float64(6256.37),
        PageToken: conductoroneapi.String("veniam"),
        PreviouslyActedOnIds: []string{
            "recusandae",
            "reiciendis",
        },
        Query: conductoroneapi.String("nulla"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("0e457e18-58b6-4a89-bbe3-a5aa8e4824d0"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByResource.ToPointer(),
        SubjectIds: []string{
            "magnam",
            "consequatur",
            "esse",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
        },
        TaskTypes: []shared.TaskType{
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("repudiandae"),
                    AccessReviewSelection: conductoroneapi.String("corporis"),
                    AppEntitlementID: conductoroneapi.String("et"),
                    AppID: conductoroneapi.String("blanditiis"),
                    AppUserID: conductoroneapi.String("ex"),
                    IdentityUserID: conductoroneapi.String("sed"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-08-28T22:43:39.407Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("saepe"),
                    AppID: conductoroneapi.String("error"),
                    AppUserID: conductoroneapi.String("consequatur"),
                    GrantDuration: conductoroneapi.String("incidunt"),
                    IdentityUserID: conductoroneapi.String("reiciendis"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeGranted.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-10-08T12:52:44.610Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("architecto"),
                    AppID: conductoroneapi.String("occaecati"),
                    AppUserID: conductoroneapi.String("labore"),
                    IdentityUserID: conductoroneapi.String("quidem"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-07-21T08:55:46.635Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-11-01T04:33:49.270Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-10-10T02:44:39.117Z"),
                            LastLogin: types.MustTimeFromString("2022-02-02T22:51:17.351Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("unde"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("reiciendis"),
                            CertTicketID: conductoroneapi.String("provident"),
                        },
                    },
                },
            },
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("repellendus"),
                    AccessReviewSelection: conductoroneapi.String("delectus"),
                    AppEntitlementID: conductoroneapi.String("voluptates"),
                    AppID: conductoroneapi.String("perferendis"),
                    AppUserID: conductoroneapi.String("est"),
                    IdentityUserID: conductoroneapi.String("quidem"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2020-12-11T19:46:56.272Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("praesentium"),
                    AppID: conductoroneapi.String("mollitia"),
                    AppUserID: conductoroneapi.String("veniam"),
                    GrantDuration: conductoroneapi.String("voluptatem"),
                    IdentityUserID: conductoroneapi.String("quisquam"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-06-17T00:04:35.396Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("reprehenderit"),
                    AppID: conductoroneapi.String("asperiores"),
                    AppUserID: conductoroneapi.String("totam"),
                    IdentityUserID: conductoroneapi.String("suscipit"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-09-23T11:31:21.970Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-10-09T20:49:35.642Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-10-08T17:16:29.459Z"),
                            LastLogin: types.MustTimeFromString("2021-10-02T23:52:38.012Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("officiis"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("officiis"),
                            CertTicketID: conductoroneapi.String("accusamus"),
                        },
                    },
                },
            },
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("natus"),
                    AccessReviewSelection: conductoroneapi.String("minima"),
                    AppEntitlementID: conductoroneapi.String("aspernatur"),
                    AppID: conductoroneapi.String("ex"),
                    AppUserID: conductoroneapi.String("maiores"),
                    IdentityUserID: conductoroneapi.String("corrupti"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-12-30T00:19:55.496Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("suscipit"),
                    AppID: conductoroneapi.String("repudiandae"),
                    AppUserID: conductoroneapi.String("atque"),
                    GrantDuration: conductoroneapi.String("atque"),
                    IdentityUserID: conductoroneapi.String("sunt"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-05-05T04:11:52.897Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("labore"),
                    AppID: conductoroneapi.String("reiciendis"),
                    AppUserID: conductoroneapi.String("doloremque"),
                    IdentityUserID: conductoroneapi.String("repudiandae"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-11-23T03:40:41.154Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-09-07T11:33:56.286Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-10-03T00:21:18.046Z"),
                            LastLogin: types.MustTimeFromString("2021-04-24T18:03:33.752Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("magnam"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("saepe"),
                            CertTicketID: conductoroneapi.String("consequuntur"),
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

