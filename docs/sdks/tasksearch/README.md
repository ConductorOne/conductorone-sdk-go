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
                "impedit",
                "hic",
                "necessitatibus",
                "asperiores",
            },
        },
        AccessReviewIds: []string{
            "voluptas",
            "debitis",
        },
        AccountOwnerIds: []string{
            "quae",
            "minus",
            "fuga",
            "laborum",
        },
        ActorID: conductoroneapi.String("consectetur"),
        AppEntitlementIds: []string{
            "atque",
        },
        AppResourceIds: []string{
            "impedit",
        },
        AppResourceTypeIds: []string{
            "soluta",
        },
        AppUserSubjectIds: []string{
            "nam",
            "dolore",
            "iusto",
            "voluptate",
        },
        ApplicationIds: []string{
            "dignissimos",
        },
        AssigneesInIds: []string{
            "quo",
        },
        CreatedAfter: types.MustTimeFromString("2021-04-24T17:00:12.334Z"),
        CreatedBefore: types.MustTimeFromString("2022-11-11T05:43:22.650Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusAll.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "quibusdam",
            "inventore",
        },
        ExcludeIds: []string{
            "libero",
            "architecto",
            "voluptatibus",
            "quia",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "aliquam",
            "velit",
            "illo",
            "accusantium",
        },
        OpenerIds: []string{
            "ea",
            "beatae",
        },
        PageSize: conductoroneapi.Float64(8777.51),
        PageToken: conductoroneapi.String("excepturi"),
        PreviouslyActedOnIds: []string{
            "velit",
            "ut",
        },
        Query: conductoroneapi.String("perspiciatis"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("1cf9e06e-3a43-4700-8ae6-b6bc9b8f759e"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("ac55a974-1d31-4135-a965-bb8a72026114"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("35e139db-c225-49b1-abda-8c070e1084cb"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("0672d1ad-879e-4eb9-a65b-85efbd02bae0"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccountOwner.ToPointer(),
        SubjectIds: []string{
            "eos",
            "quibusdam",
            "odio",
            "praesentium",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("error"),
                        IntegrationID: conductoroneapi.String("earum"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-01-28T23:50:43.870Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-06-08T12:41:53.360Z"),
                            LastLogin: types.MustTimeFromString("2022-04-28T01:39:49.849Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("beatae"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("unde"),
                            CertTicketID: conductoroneapi.String("molestiae"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("delectus"),
                        IntegrationID: conductoroneapi.String("cupiditate"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-09-30T10:14:18.921Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-10-19T15:08:58.907Z"),
                            LastLogin: types.MustTimeFromString("2021-02-01T13:09:02.595Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("dignissimos"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("optio"),
                            CertTicketID: conductoroneapi.String("necessitatibus"),
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

