# TaskSearch
(*TaskSearch*)

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
                "transition",
            },
        },
        AccessReviewIds: []string{
            "turquoise",
        },
        AccountOwnerIds: []string{
            "Hyundai",
        },
        ActorID: conductoronesdkgo.String("Practical yearningly Southeast"),
        AppEntitlementIds: []string{
            "Books",
        },
        AppResourceIds: []string{
            "Other",
        },
        AppResourceTypeIds: []string{
            "eligendi",
        },
        AppUserSubjectIds: []string{
            "Tobago",
        },
        ApplicationIds: []string{
            "Protactinium",
        },
        AssigneesInIds: []string{
            "Arsenic",
        },
        CreatedAfter: types.MustTimeFromString("2023-06-10T20:22:45.491Z"),
        CreatedBefore: types.MustTimeFromString("2022-09-06T19:12:57.415Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepUnspecified.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusNonEmergency.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "National",
        },
        ExcludeIds: []string{
            "Wagon",
        },
        IncludeDeleted: conductoronesdkgo.Bool(false),
        MyWorkUserIds: []string{
            "Estate",
        },
        OpenerIds: []string{
            "female",
        },
        PageSize: conductoronesdkgo.Float64(2339.71),
        PageToken: conductoronesdkgo.String("Soft Bacon"),
        PreviouslyActedOnIds: []string{
            "mindshare",
        },
        Query: conductoronesdkgo.String("frizzy"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoronesdkgo.String("<ID>"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccountOwner.ToPointer(),
        SubjectIds: []string{
            "boo",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoronesdkgo.String("Frozen"),
                        IntegrationID: conductoronesdkgo.String("discrete HDD chargrill"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2023-03-10T15:48:59.869Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2023-01-14T01:21:25.379Z"),
                            LastLogin: types.MustTimeFromString("2021-01-22T06:17:27.525Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoronesdkgo.String("overriding Focused"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoronesdkgo.String("plum Astatine"),
                            CertTicketID: conductoronesdkgo.String("soliloquize instead"),
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

