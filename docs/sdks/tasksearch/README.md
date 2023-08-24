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
                "nulla",
                "distinctio",
                "maxime",
            },
        },
        AccessReviewIds: []string{
            "quia",
        },
        AccountOwnerIds: []string{
            "omnis",
            "libero",
        },
        ActorID: conductoroneapi.String("dicta"),
        AppEntitlementIds: []string{
            "libero",
            "fugiat",
            "officia",
        },
        AppResourceIds: []string{
            "placeat",
            "sit",
            "iusto",
        },
        AppResourceTypeIds: []string{
            "voluptates",
        },
        AppUserSubjectIds: []string{
            "aperiam",
        },
        ApplicationIds: []string{
            "dolore",
            "eligendi",
            "distinctio",
        },
        AssigneesInIds: []string{
            "autem",
        },
        CreatedAfter: types.MustTimeFromString("2022-10-29T20:56:45.429Z"),
        CreatedBefore: types.MustTimeFromString("2022-09-05T14:30:04.979Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusEmergency.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "molestiae",
            "provident",
            "accusamus",
        },
        ExcludeIds: []string{
            "tempore",
            "sint",
            "ea",
            "autem",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "rerum",
            "laudantium",
        },
        OpenerIds: []string{
            "officiis",
            "voluptatibus",
        },
        PageSize: conductoroneapi.Float64(7372.79),
        PageToken: conductoroneapi.String("at"),
        PreviouslyActedOnIds: []string{
            "quia",
        },
        Query: conductoroneapi.String("quidem"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("e0be2d78-2259-4e3e-a4b5-197f92443da7"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("ce52b895-c537-4c64-94ef-b0b34896c3ca"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("5acfbe2f-d570-4757-b929-177deac646ec"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccountOwner.ToPointer(),
        SubjectIds: []string{
            "quam",
            "dolorem",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("sequi"),
                        IntegrationID: conductoroneapi.String("repudiandae"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-10-05T07:29:39.358Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-12-30T18:01:47.888Z"),
                            LastLogin: types.MustTimeFromString("2022-08-25T23:28:51.951Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("nam"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("dicta"),
                            CertTicketID: conductoroneapi.String("consequuntur"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("necessitatibus"),
                        IntegrationID: conductoroneapi.String("nobis"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-07-03T11:09:14.391Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-09-27T07:34:57.107Z"),
                            LastLogin: types.MustTimeFromString("2022-08-12T22:43:14.075Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("pariatur"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("libero"),
                            CertTicketID: conductoroneapi.String("excepturi"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("occaecati"),
                        IntegrationID: conductoroneapi.String("nemo"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-08-29T01:04:50.334Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2020-09-23T16:28:13.144Z"),
                            LastLogin: types.MustTimeFromString("2022-05-16T08:40:32.544Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("hic"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("animi"),
                            CertTicketID: conductoroneapi.String("quas"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("totam"),
                        IntegrationID: conductoroneapi.String("molestias"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-12-13T07:44:48.227Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-09-09T00:48:35.479Z"),
                            LastLogin: types.MustTimeFromString("2021-10-09T11:02:42.252Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("assumenda"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("tempore"),
                            CertTicketID: conductoroneapi.String("libero"),
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

