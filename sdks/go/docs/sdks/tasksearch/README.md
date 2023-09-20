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
        AccessReviewIds: []string{
            "odio",
            "praesentium",
            "odit",
            "explicabo",
        },
        AccountOwnerIds: []string{
            "error",
            "earum",
        },
        ActorID: conductoroneapi.String("adipisci"),
        AppEntitlementIds: []string{
            "similique",
            "ut",
            "quidem",
            "quis",
        },
        AppResourceIds: []string{
            "unde",
        },
        AppResourceTypeIds: []string{
            "delectus",
            "cupiditate",
        },
        AppUserSubjectIds: []string{
            "numquam",
        },
        ApplicationIds: []string{
            "nesciunt",
            "at",
        },
        AssigneesInIds: []string{
            "dignissimos",
            "optio",
            "necessitatibus",
        },
        CreatedAfter: types.MustTimeFromString("2022-10-25T06:51:46.936Z"),
        CreatedBefore: types.MustTimeFromString("2021-12-08T03:38:46.630Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepApproval.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusAll.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "enim",
            "neque",
            "in",
            "minus",
        },
        ExcludeIds: []string{
            "modi",
            "corporis",
        },
        ExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "voluptates",
                "maiores",
            },
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "aperiam",
            "libero",
            "ratione",
        },
        OpenerIds: []string{
            "totam",
            "occaecati",
        },
        PageSize: conductoroneapi.Float64(3759.94),
        PageToken: conductoroneapi.String("quo"),
        PreviouslyActedOnIds: []string{
            "minus",
        },
        Query: conductoroneapi.String("fuga"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("acfbe2fd-5707-4577-9291-77deac646ecb"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("573409e3-eb1e-45a2-b12e-b07f116db995"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccount.ToPointer(),
        SubjectIds: []string{
            "doloribus",
            "eligendi",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                Certify: &shared.TaskTypeCertify{},
                Grant: &shared.TaskTypeGrantInput{
                    Source: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("totam"),
                        IntegrationID: conductoroneapi.String("molestias"),
                    },
                },
                Revoke: &shared.TaskTypeRevokeInput{
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-12-13T07:44:48.227Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-09-09T00:48:35.479Z"),
                            LastLogin: types.MustTimeFromString("2021-10-09T11:02:42.252Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("assumenda"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("tempore"),
                            CertTicketID: conductoroneapi.String("libero"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                Certify: &shared.TaskTypeCertify{},
                Grant: &shared.TaskTypeGrantInput{
                    Source: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("velit"),
                        IntegrationID: conductoroneapi.String("doloremque"),
                    },
                },
                Revoke: &shared.TaskTypeRevokeInput{
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2020-09-06T10:40:03.787Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-07-26T16:25:58.578Z"),
                            LastLogin: types.MustTimeFromString("2022-02-05T08:24:32.703Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("deserunt"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("doloremque"),
                            CertTicketID: conductoroneapi.String("quis"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                Certify: &shared.TaskTypeCertify{},
                Grant: &shared.TaskTypeGrantInput{
                    Source: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("veniam"),
                        IntegrationID: conductoroneapi.String("libero"),
                    },
                },
                Revoke: &shared.TaskTypeRevokeInput{
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-06-01T14:58:08.177Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-03-30T19:42:48.151Z"),
                            LastLogin: types.MustTimeFromString("2022-03-31T15:31:53.121Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("magnam"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("itaque"),
                            CertTicketID: conductoroneapi.String("sed"),
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

