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
                "aliquid",
            },
        },
        AccessReviewIds: []string{
            "accusantium",
        },
        AccountOwnerIds: []string{
            "repellat",
        },
        ActorID: conductoroneapi.String("doloribus"),
        AppEntitlementIds: []string{
            "ullam",
        },
        AppResourceIds: []string{
            "in",
        },
        AppResourceTypeIds: []string{
            "nam",
        },
        AppUserSubjectIds: []string{
            "earum",
        },
        ApplicationIds: []string{
            "officia",
        },
        AssigneesInIds: []string{
            "laborum",
        },
        CreatedAfter: types.MustTimeFromString("2022-03-14T23:12:21.252Z"),
        CreatedBefore: types.MustTimeFromString("2021-04-23T08:23:19.189Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusEmergency.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "cumque",
        },
        ExcludeIds: []string{
            "vitae",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "rerum",
        },
        OpenerIds: []string{
            "tempora",
        },
        PageSize: conductoroneapi.Float64(3354.98),
        PageToken: conductoroneapi.String("inventore"),
        PreviouslyActedOnIds: []string{
            "fugit",
        },
        Query: conductoroneapi.String("cumque"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("1032648d-c2f6-4151-99eb-fd0e9fe6c632"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccountOwner.ToPointer(),
        SubjectIds: []string{
            "fuga",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("animi"),
                        IntegrationID: conductoroneapi.String("necessitatibus"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-12-22T05:17:09.936Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-11-29T01:33:31.768Z"),
                            LastLogin: types.MustTimeFromString("2022-05-19T23:57:30.950Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("occaecati"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("suscipit"),
                            CertTicketID: conductoroneapi.String("adipisci"),
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

