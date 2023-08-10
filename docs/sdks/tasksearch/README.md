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
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TaskSearch.Search(ctx, shared.TaskSearchRequestInput{
        TaskExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "commodi",
                "officiis",
            },
        },
        AccessReviewIds: []string{
            "quidem",
            "exercitationem",
            "quam",
            "dolorem",
        },
        AccountOwnerIds: []string{
            "ipsa",
            "sint",
        },
        ActorID: conductoroneapi.String("vero"),
        AppEntitlementIds: []string{
            "repudiandae",
        },
        AppResourceIds: []string{
            "dicta",
            "earum",
            "veniam",
        },
        AppResourceTypeIds: []string{
            "dolores",
            "nam",
            "dicta",
        },
        AppUserSubjectIds: []string{
            "necessitatibus",
        },
        ApplicationIds: []string{
            "ipsa",
            "ducimus",
            "maiores",
        },
        AssigneesInIds: []string{
            "quasi",
        },
        CreatedAfter: types.MustTimeFromString("2022-02-19T19:59:49.568Z"),
        CreatedBefore: types.MustTimeFromString("2021-11-13T10:48:46.467Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepApproval.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusAll.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "nostrum",
            "doloribus",
        },
        ExcludeIds: []string{
            "sint",
            "enim",
            "hic",
            "animi",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "totam",
            "molestias",
            "odio",
        },
        OpenerIds: []string{
            "saepe",
        },
        PageSize: conductoroneapi.Float64(1040.78),
        PageToken: conductoroneapi.String("quos"),
        PreviouslyActedOnIds: []string{
            "assumenda",
            "tempore",
            "libero",
        },
        Query: conductoroneapi.String("velit"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("fcb33ea0-55b1-497c-944e-2f52d82d3513"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccountOwner.ToPointer(),
        SubjectIds: []string{
            "nisi",
            "voluptatibus",
            "quaerat",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-05-31T00:14:48.748Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2020-11-18T20:43:00.067Z"),
                            LastLogin: types.MustTimeFromString("2022-09-06T14:17:02.690Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("voluptatibus"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("voluptatibus"),
                            CertTicketID: conductoroneapi.String("consequuntur"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-02-16T12:42:02.892Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-08-22T17:59:10.804Z"),
                            LastLogin: types.MustTimeFromString("2022-08-27T21:22:12.876Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("neque"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("iusto"),
                            CertTicketID: conductoroneapi.String("est"),
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

