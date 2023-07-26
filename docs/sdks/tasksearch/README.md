# TaskSearch

### Available Operations

* [Search](#search) - Search

## Search

 Search tasks based on filters specified in the request body.


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
    res, err := s.TaskSearch.Search(ctx, shared.TaskSearchRequestInput{
        TaskExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "voluptate",
                "ullam",
                "unde",
                "necessitatibus",
            },
        },
        AccessReviewIds: []string{
            "impedit",
            "ipsam",
            "corporis",
        },
        AccountOwnerIds: []string{
            "error",
            "esse",
            "labore",
        },
        ActorID: conductoroneapi.String("veritatis"),
        AppEntitlementIds: []string{
            "consectetur",
            "vitae",
            "inventore",
            "dolorem",
        },
        AppResourceIds: []string{
            "qui",
            "iste",
        },
        AppResourceTypeIds: []string{
            "nemo",
            "soluta",
        },
        AppUserSubjectIds: []string{
            "rem",
            "dolorum",
            "odio",
        },
        ApplicationIds: []string{
            "alias",
        },
        AssigneesInIds: []string{
            "vel",
        },
        CreatedAfter: types.MustTimeFromString("2022-12-08T19:16:07.063Z"),
        CreatedBefore: types.MustTimeFromString("2022-10-16T23:42:04.526Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepApproval.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusEmergency.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "ipsum",
        },
        ExcludeIds: []string{
            "nulla",
            "distinctio",
            "maxime",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "quia",
        },
        OpenerIds: []string{
            "omnis",
            "libero",
        },
        PageSize: conductoroneapi.Float64(1156.61),
        PageToken: conductoroneapi.String("id"),
        PreviouslyActedOnIds: []string{
            "fugiat",
            "officia",
            "quos",
        },
        Query: conductoroneapi.String("placeat"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("70e1084c-b067-42d1-ad87-9eeb9665b85e"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccountOwner.ToPointer(),
        SubjectIds: []string{
            "at",
            "alias",
            "quia",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertify{},
                TaskTypeGrant: &shared.TaskTypeGrant{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-06-21T04:17:16.724Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-07-20T13:08:36.205Z"),
                            LastLogin: types.MustTimeFromString("2022-09-20T13:39:46.907Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("explicabo"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("corporis"),
                            CertTicketID: conductoroneapi.String("error"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertify{},
                TaskTypeGrant: &shared.TaskTypeGrant{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-04-13T22:13:24.007Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-02-09T04:42:29.895Z"),
                            LastLogin: types.MustTimeFromString("2022-04-22T18:47:14.845Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("quis"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("beatae"),
                            CertTicketID: conductoroneapi.String("unde"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertify{},
                TaskTypeGrant: &shared.TaskTypeGrant{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-01-14T10:23:30.043Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-09-15T02:31:13.378Z"),
                            LastLogin: types.MustTimeFromString("2022-09-29T05:24:35.816Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("nesciunt"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("at"),
                            CertTicketID: conductoroneapi.String("officia"),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.TaskSearchRequestInput](../../models/shared/tasksearchrequestinput.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.C1APITaskV1TaskSearchServiceSearchResponse](../../models/operations/c1apitaskv1tasksearchservicesearchresponse.md), error**

