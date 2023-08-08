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
                "adipisci",
                "recusandae",
                "similique",
                "ut",
            },
        },
        AccessReviewIds: []string{
            "quis",
            "beatae",
            "unde",
        },
        AccountOwnerIds: []string{
            "delectus",
            "cupiditate",
        },
        ActorID: conductoroneapi.String("fugit"),
        AppEntitlementIds: []string{
            "numquam",
            "nesciunt",
        },
        AppResourceIds: []string{
            "officia",
            "dignissimos",
            "optio",
            "necessitatibus",
        },
        AppResourceTypeIds: []string{
            "qui",
            "expedita",
        },
        AppUserSubjectIds: []string{
            "cupiditate",
            "minima",
            "placeat",
        },
        ApplicationIds: []string{
            "neque",
            "in",
        },
        AssigneesInIds: []string{
            "eum",
            "modi",
            "corporis",
            "magnam",
        },
        CreatedAfter: types.MustTimeFromString("2020-01-25T21:38:07.032Z"),
        CreatedBefore: types.MustTimeFromString("2022-11-28T09:50:30.826Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusUnspecified.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "totam",
            "occaecati",
        },
        ExcludeIds: []string{
            "quo",
            "velit",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "fuga",
            "nostrum",
            "est",
            "impedit",
        },
        OpenerIds: []string{
            "tempore",
            "vero",
            "odit",
            "repellat",
        },
        PageSize: conductoroneapi.Float64(8659.46),
        PageToken: conductoroneapi.String("nemo"),
        PreviouslyActedOnIds: []string{
            "aperiam",
            "odio",
        },
        Query: conductoroneapi.String("minima"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("7929177d-eac6-446e-8b57-3409e3eb1e5a"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("2b12eb07-f116-4db9-9545-fc95fa88970e"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByUnspecified.ToPointer(),
        SubjectIds: []string{
            "iste",
            "assumenda",
            "tempore",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-07-26T16:25:58.578Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-02-05T08:24:32.703Z"),
                            LastLogin: types.MustTimeFromString("2022-11-30T15:58:04.315Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("quis"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("veniam"),
                            CertTicketID: conductoroneapi.String("libero"),
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
                            ExpiredAt: types.MustTimeFromString("2022-06-01T14:58:08.177Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-03-30T19:42:48.151Z"),
                            LastLogin: types.MustTimeFromString("2022-03-31T15:31:53.121Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("magnam"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("itaque"),
                            CertTicketID: conductoroneapi.String("sed"),
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
                            ExpiredAt: types.MustTimeFromString("2022-01-04T08:34:55.183Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-03-09T22:24:35.771Z"),
                            LastLogin: types.MustTimeFromString("2022-09-20T12:28:22.531Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("pariatur"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("amet"),
                            CertTicketID: conductoroneapi.String("exercitationem"),
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
                            ExpiredAt: types.MustTimeFromString("2022-10-04T10:29:48.523Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-07-16T07:44:14.632Z"),
                            LastLogin: types.MustTimeFromString("2022-01-09T04:56:37.970Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("quaerat"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("blanditiis"),
                            CertTicketID: conductoroneapi.String("distinctio"),
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

