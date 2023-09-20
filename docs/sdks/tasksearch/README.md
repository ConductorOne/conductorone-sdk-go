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
                "repellat",
            },
        },
        AccessReviewIds: []string{
            "doloribus",
        },
        AccountOwnerIds: []string{
            "ullam",
        },
        ActorID: conductoroneapi.String("in"),
        AppEntitlementIds: []string{
            "nam",
        },
        AppResourceIds: []string{
            "earum",
        },
        AppResourceTypeIds: []string{
            "officia",
        },
        AppUserSubjectIds: []string{
            "laborum",
        },
        ApplicationIds: []string{
            "placeat",
        },
        AssigneesInIds: []string{
            "modi",
        },
        CreatedAfter: types.MustTimeFromString("2021-04-23T08:23:19.189Z"),
        CreatedBefore: types.MustTimeFromString("2020-02-18T03:48:05.478Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusUnspecified.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "rerum",
        },
        ExcludeIds: []string{
            "tempora",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "quis",
        },
        OpenerIds: []string{
            "inventore",
        },
        PageSize: conductoroneapi.Float64(1476.85),
        PageToken: conductoroneapi.String("cumque"),
        PreviouslyActedOnIds: []string{
            "quae",
        },
        Query: conductoroneapi.String("perferendis"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("32648dc2-f615-4199-abfd-0e9fe6c632ca"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByUnspecified.ToPointer(),
        SubjectIds: []string{
            "animi",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("nulla"),
                        IntegrationID: conductoroneapi.String("consequatur"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-11-29T01:33:31.768Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-05-19T23:57:30.950Z"),
                            LastLogin: types.MustTimeFromString("2022-03-27T19:38:57.457Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("adipisci"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("quasi"),
                            CertTicketID: conductoroneapi.String("magni"),
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

