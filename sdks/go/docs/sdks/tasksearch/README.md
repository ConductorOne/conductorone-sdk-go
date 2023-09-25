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
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/types"
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TaskSearch.Search(ctx, shared.TaskSearchRequestInput{
        TaskExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "officia",
            },
        },
        AccessReviewIds: []string{
            "laborum",
        },
        AccountOwnerIds: []string{
            "placeat",
        },
        ActorID: conductoronesdkgo.String("modi"),
        AppEntitlementIds: []string{
            "voluptatibus",
        },
        AppResourceIds: []string{
            "molestias",
        },
        AppResourceTypeIds: []string{
            "officiis",
        },
        AppUserSubjectIds: []string{
            "sapiente",
        },
        ApplicationIds: []string{
            "cumque",
        },
        AssigneesInIds: []string{
            "vitae",
        },
        CreatedAfter: types.MustTimeFromString("2022-06-16T06:32:28.803Z"),
        CreatedBefore: types.MustTimeFromString("2022-12-02T01:10:26.015Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepUnspecified.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusEmergency.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "quae",
        },
        ExcludeIds: []string{
            "perferendis",
        },
        IncludeDeleted: conductoronesdkgo.Bool(false),
        MyWorkUserIds: []string{
            "velit",
        },
        OpenerIds: []string{
            "aspernatur",
        },
        PageSize: conductoronesdkgo.Float64(4322.81),
        PageToken: conductoronesdkgo.String("eius"),
        PreviouslyActedOnIds: []string{
            "rem",
        },
        Query: conductoronesdkgo.String("at"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoronesdkgo.String("c2f61519-9ebf-4d0e-9fe6-c632ca3aed01"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByUnspecified.ToPointer(),
        SubjectIds: []string{
            "ducimus",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoronesdkgo.String("occaecati"),
                        IntegrationID: conductoronesdkgo.String("suscipit"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-11-26T18:06:32.022Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-01-06T11:59:03.269Z"),
                            LastLogin: types.MustTimeFromString("2020-04-24T05:50:40.136Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoronesdkgo.String("ipsa"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoronesdkgo.String("tempora"),
                            CertTicketID: conductoronesdkgo.String("nihil"),
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

