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
                "vero",
                "voluptatem",
                "ipsam",
            },
        },
        AccessReviewIds: []string{
            "alias",
            "quasi",
        },
        AccountOwnerIds: []string{
            "maiores",
        },
        ActorID: conductoroneapi.String("enim"),
        AppEntitlementIds: []string{
            "nulla",
            "deserunt",
            "esse",
        },
        AppResourceIds: []string{
            "reprehenderit",
            "est",
        },
        AppResourceTypeIds: []string{
            "sint",
            "accusamus",
        },
        AppUserSubjectIds: []string{
            "hic",
            "necessitatibus",
            "asperiores",
            "ex",
        },
        ApplicationIds: []string{
            "debitis",
            "delectus",
        },
        AssigneesInIds: []string{
            "minus",
        },
        CreatedAfter: types.MustTimeFromString("2021-08-25T17:55:33.849Z"),
        CreatedBefore: types.MustTimeFromString("2022-10-03T14:46:04.624Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepApproval.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusUnspecified.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "magni",
            "soluta",
            "repudiandae",
            "nam",
        },
        ExcludeIds: []string{
            "iusto",
            "voluptate",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "dignissimos",
        },
        OpenerIds: []string{
            "quo",
        },
        PageSize: conductoroneapi.Float64(5354.68),
        PageToken: conductoroneapi.String("quibusdam"),
        PreviouslyActedOnIds: []string{
            "odit",
            "voluptatibus",
        },
        Query: conductoroneapi.String("vel"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("d1db1f2c-4310-4661-a963-49e1cf9e06e3"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("a437000a-e6b6-4bc9-b8f7-59eac55a9741"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccountOwner.ToPointer(),
        SubjectIds: []string{
            "vitae",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertify{},
                TaskTypeGrant: &shared.TaskTypeGrant{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-05-22T20:46:18.043Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-08-19T19:35:40.341Z"),
                            LastLogin: types.MustTimeFromString("2021-07-19T20:29:58.626Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("rem"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("dolorum"),
                            CertTicketID: conductoroneapi.String("odio"),
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
                            ExpiredAt: types.MustTimeFromString("2022-12-31T19:13:01.264Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-07-29T17:28:50.439Z"),
                            LastLogin: types.MustTimeFromString("2022-12-08T19:16:07.063Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("modi"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("neque"),
                            CertTicketID: conductoroneapi.String("exercitationem"),
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

