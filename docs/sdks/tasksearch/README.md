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
                "beatae",
                "vero",
            },
        },
        AccessReviewIds: []string{
            "eum",
            "velit",
            "ut",
        },
        AccountOwnerIds: []string{
            "earum",
            "dicta",
            "impedit",
        },
        ActorID: conductoroneapi.String("voluptatibus"),
        AppEntitlementIds: []string{
            "itaque",
            "alias",
            "nisi",
        },
        AppResourceIds: []string{
            "velit",
            "laborum",
            "non",
            "dolor",
        },
        AppResourceTypeIds: []string{
            "sit",
            "doloremque",
        },
        AppUserSubjectIds: []string{
            "officia",
        },
        ApplicationIds: []string{
            "ea",
            "quidem",
            "voluptas",
            "facilis",
        },
        AssigneesInIds: []string{
            "perspiciatis",
            "expedita",
            "deleniti",
            "a",
        },
        CreatedAfter: types.MustTimeFromString("2022-08-25T13:02:18.147Z"),
        CreatedBefore: types.MustTimeFromString("2021-03-16T19:03:02.908Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepApproval.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusEmergency.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "corporis",
            "est",
        },
        ExcludeIds: []string{
            "esse",
            "labore",
            "veritatis",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "consectetur",
            "vitae",
            "inventore",
            "dolorem",
        },
        OpenerIds: []string{
            "qui",
            "iste",
        },
        PageSize: conductoroneapi.Float64(4030.26),
        PageToken: conductoroneapi.String("nemo"),
        PreviouslyActedOnIds: []string{
            "libero",
            "rem",
            "dolorum",
        },
        Query: conductoroneapi.String("odio"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("02611435-e139-4dbc-a259-b1abda8c070e"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByUnspecified.ToPointer(),
        SubjectIds: []string{
            "totam",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-07-18T07:24:59.490Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-03-05T22:46:19.288Z"),
                            LastLogin: types.MustTimeFromString("2022-05-01T23:00:45.808Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("facere"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("corrupti"),
                            CertTicketID: conductoroneapi.String("molestiae"),
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

