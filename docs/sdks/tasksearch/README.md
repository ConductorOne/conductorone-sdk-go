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
                "eligendi",
                "distinctio",
            },
        },
        AccessReviewIds: []string{
            "autem",
        },
        AccountOwnerIds: []string{
            "dolores",
            "assumenda",
        },
        ActorID: conductoroneapi.String("beatae"),
        AppEntitlementIds: []string{
            "facere",
            "corrupti",
            "molestiae",
        },
        AppResourceIds: []string{
            "accusamus",
            "necessitatibus",
            "tempore",
        },
        AppResourceTypeIds: []string{
            "ea",
            "autem",
            "ipsam",
        },
        AppUserSubjectIds: []string{
            "laudantium",
            "corporis",
            "officiis",
        },
        ApplicationIds: []string{
            "cum",
            "at",
            "alias",
            "quia",
        },
        AssigneesInIds: []string{
            "fuga",
            "repudiandae",
            "accusantium",
        },
        CreatedAfter: types.MustTimeFromString("2021-03-25T19:10:06.582Z"),
        CreatedBefore: types.MustTimeFromString("2022-02-26T15:04:57.243Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepApproval.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusNonEmergency.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "explicabo",
        },
        ExcludeIds: []string{
            "error",
            "earum",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "recusandae",
        },
        OpenerIds: []string{
            "ut",
            "quidem",
            "quis",
        },
        PageSize: conductoroneapi.Float64(1062.55),
        PageToken: conductoroneapi.String("unde"),
        PreviouslyActedOnIds: []string{
            "delectus",
            "cupiditate",
        },
        Query: conductoroneapi.String("fugit"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("43da7ce5-2b89-45c5-b7c6-454efb0b3489"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("6c3ca5ac-fbe2-4fd5-b075-77929177deac"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByResource.ToPointer(),
        SubjectIds: []string{
            "commodi",
            "officiis",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-06-04T22:19:22.955Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-06-01T14:00:05.540Z"),
                            LastLogin: types.MustTimeFromString("2020-10-11T08:25:54.157Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("dicta"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("earum"),
                            CertTicketID: conductoroneapi.String("veniam"),
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
                            ExpiredAt: types.MustTimeFromString("2022-08-25T23:28:51.951Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-10-08T09:27:30.689Z"),
                            LastLogin: types.MustTimeFromString("2022-02-06T13:10:02.431Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("nobis"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("ipsa"),
                            CertTicketID: conductoroneapi.String("ducimus"),
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

