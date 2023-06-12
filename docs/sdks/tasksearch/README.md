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
            "sint",
            "aliquid",
            "provident",
            "necessitatibus",
        },
        AccountOwnerIds: []string{
            "officia",
            "dolor",
            "debitis",
        },
        ActorID: conductoroneapi.String("a"),
        AppEntitlementIds: []string{
            "in",
            "in",
            "illum",
        },
        AppResourceIds: []string{
            "rerum",
            "dicta",
            "magnam",
            "cumque",
        },
        AppResourceTypeIds: []string{
            "ea",
            "aliquid",
            "laborum",
            "accusamus",
        },
        AppUserSubjectIds: []string{
            "occaecati",
        },
        ApplicationIds: []string{
            "accusamus",
            "delectus",
        },
        AssigneesInIds: []string{
            "provident",
            "nam",
            "id",
        },
        CreatedAfter: types.MustTimeFromString("2021-12-07T18:13:34.827Z"),
        CreatedBefore: types.MustTimeFromString("2022-04-23T13:35:30.978Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepApproval.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "vel",
            "natus",
        },
        ExcludeIds: []string{
            "molestiae",
            "perferendis",
            "nihil",
        },
        ExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "distinctio",
                "id",
            },
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "labore",
            "suscipit",
        },
        OpenerIds: []string{
            "nobis",
            "eum",
            "vero",
        },
        PageSize: conductoroneapi.Float64(1354.74),
        PageToken: conductoroneapi.String("architecto"),
        PreviouslyActedOnIds: []string{
            "et",
            "excepturi",
        },
        Query: conductoroneapi.String("ullam"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("890afa56-3e25-416f-a4c8-b711e5b7fd2e"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("d028921c-ddc6-4926-81fb-576b0d5f0d30"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("c5fbb258-7053-4202-873d-5fe9b90c2890"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByResource.ToPointer(),
        SubjectIds: []string{
            "adipisci",
            "asperiores",
            "earum",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
        },
        TaskTypes: []shared.TaskType{
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("pariatur"),
                    AccessReviewSelection: conductoroneapi.String("provident"),
                    AppEntitlementID: conductoroneapi.String("nobis"),
                    AppID: conductoroneapi.String("libero"),
                    AppUserID: conductoroneapi.String("delectus"),
                    IdentityUserID: conductoroneapi.String("quaerat"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-10-15T11:27:32.342Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("dolorem"),
                    AppID: conductoroneapi.String("dolor"),
                    AppUserID: conductoroneapi.String("qui"),
                    GrantDuration: conductoroneapi.String("ipsum"),
                    IdentityUserID: conductoroneapi.String("hic"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-02-04T23:02:58.199Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("dignissimos"),
                    AppID: conductoroneapi.String("reiciendis"),
                    AppUserID: conductoroneapi.String("amet"),
                    IdentityUserID: conductoroneapi.String("dolorum"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeRevoked.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-12-10T19:39:51.391Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-07-26T10:30:36.625Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-09-09T04:40:04.540Z"),
                            LastLogin: types.MustTimeFromString("2020-11-29T12:05:35.198Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("voluptatibus"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("voluptas"),
                            CertTicketID: conductoroneapi.String("natus"),
                        },
                    },
                },
            },
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("eos"),
                    AccessReviewSelection: conductoroneapi.String("atque"),
                    AppEntitlementID: conductoroneapi.String("sit"),
                    AppID: conductoroneapi.String("fugiat"),
                    AppUserID: conductoroneapi.String("ab"),
                    IdentityUserID: conductoroneapi.String("soluta"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-07-19T09:36:55.923Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("dolorum"),
                    AppID: conductoroneapi.String("deleniti"),
                    AppUserID: conductoroneapi.String("omnis"),
                    GrantDuration: conductoroneapi.String("necessitatibus"),
                    IdentityUserID: conductoroneapi.String("distinctio"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-10-13T19:58:56.024Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("voluptate"),
                    AppID: conductoroneapi.String("id"),
                    AppUserID: conductoroneapi.String("saepe"),
                    IdentityUserID: conductoroneapi.String("eius"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-10-09T08:02:18.659Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2020-05-10T15:56:07.975Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-02-04T21:14:13.821Z"),
                            LastLogin: types.MustTimeFromString("2022-05-10T06:55:36.509Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("provident"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("minima"),
                            CertTicketID: conductoroneapi.String("repellendus"),
                        },
                    },
                },
            },
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("totam"),
                    AccessReviewSelection: conductoroneapi.String("similique"),
                    AppEntitlementID: conductoroneapi.String("alias"),
                    AppID: conductoroneapi.String("at"),
                    AppUserID: conductoroneapi.String("quaerat"),
                    IdentityUserID: conductoroneapi.String("tempora"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2020-05-06T13:19:34.143Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("qui"),
                    AppID: conductoroneapi.String("dolorum"),
                    AppUserID: conductoroneapi.String("a"),
                    GrantDuration: conductoroneapi.String("esse"),
                    IdentityUserID: conductoroneapi.String("harum"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-03-19T02:38:33.395Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("tenetur"),
                    AppID: conductoroneapi.String("amet"),
                    AppUserID: conductoroneapi.String("tempore"),
                    IdentityUserID: conductoroneapi.String("accusamus"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeRevoked.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-10-15T03:23:12.773Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-06-12T13:41:06.619Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-12-22T15:11:39.695Z"),
                            LastLogin: types.MustTimeFromString("2022-08-02T13:07:26.403Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("sed"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("vel"),
                            CertTicketID: conductoroneapi.String("libero"),
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

