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
                "totam",
            },
        },
        AccessReviewIds: []string{
            "eligendi",
            "distinctio",
        },
        AccountOwnerIds: []string{
            "autem",
        },
        ActorID: conductoroneapi.String("esse"),
        AppEntitlementIds: []string{
            "assumenda",
        },
        AppResourceIds: []string{
            "est",
        },
        AppResourceTypeIds: []string{
            "corrupti",
            "molestiae",
            "provident",
            "accusamus",
        },
        AppUserSubjectIds: []string{
            "tempore",
            "sint",
            "ea",
            "autem",
        },
        ApplicationIds: []string{
            "rerum",
            "laudantium",
        },
        AssigneesInIds: []string{
            "officiis",
            "voluptatibus",
        },
        CreatedAfter: types.MustTimeFromString("2021-04-04T05:15:35.912Z"),
        CreatedBefore: types.MustTimeFromString("2022-11-04T19:42:54.772Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusNonEmergency.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "accusantium",
            "expedita",
            "officiis",
            "eos",
        },
        ExcludeIds: []string{
            "odio",
            "praesentium",
            "odit",
            "explicabo",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "error",
            "earum",
        },
        OpenerIds: []string{
            "recusandae",
        },
        PageSize: conductoroneapi.Float64(6308.71),
        PageToken: conductoroneapi.String("ut"),
        PreviouslyActedOnIds: []string{
            "quis",
            "beatae",
            "unde",
        },
        Query: conductoroneapi.String("molestiae"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("92443da7-ce52-4b89-9c53-7c6454efb0b3"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("4896c3ca-5acf-4be2-bd57-07577929177d"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("eac646ec-b573-4409-a3eb-1e5a2b12eb07"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("f116db99-545f-4c95-ba88-970e189dbb30"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByReverseTicketID.ToPointer(),
        SubjectIds: []string{
            "cum",
            "ipsum",
            "adipisci",
            "saepe",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("architecto"),
                        IntegrationID: conductoroneapi.String("cupiditate"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-03-30T19:42:48.151Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-03-31T15:31:53.121Z"),
                            LastLogin: types.MustTimeFromString("2022-01-25T17:05:34.945Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("sed"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("asperiores"),
                            CertTicketID: conductoroneapi.String("veniam"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("consequuntur"),
                        IntegrationID: conductoroneapi.String("facere"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-09-20T12:28:22.531Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-04-27T00:54:52.450Z"),
                            LastLogin: types.MustTimeFromString("2022-12-06T20:51:23.545Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("velit"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("facilis"),
                            CertTicketID: conductoroneapi.String("tempore"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoroneapi.String("nisi"),
                        IntegrationID: conductoroneapi.String("voluptatibus"),
                    },
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-07-01T03:10:08.438Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-03-20T12:36:58.674Z"),
                            LastLogin: types.MustTimeFromString("2022-08-10T22:40:04.682Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("libero"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("minus"),
                            CertTicketID: conductoroneapi.String("facere"),
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

