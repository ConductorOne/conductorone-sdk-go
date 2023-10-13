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
)

func main() {
    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TaskSearch.Search(ctx, &shared.TaskSearchRequestInput{
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
        AppEntitlementIds: []string{
            "Future",
        },
        AppResourceIds: []string{
            "Southwest",
        },
        AppResourceTypeIds: []string{
            "broach",
        },
        AppUserSubjectIds: []string{
            "dependent",
        },
        ApplicationIds: []string{
            "Mozambique",
        },
        AssigneesInIds: []string{
            "Global",
        },
        ExcludeAppEntitlementIds: []string{
            "watt",
        },
        ExcludeIds: []string{
            "Gasoline",
        },
        MyWorkUserIds: []string{
            "Protactinium",
        },
        OpenerIds: []string{
            "Arsenic",
        },
        PreviouslyActedOnIds: []string{
            "Gasoline",
        },
        Refs: []shared.TaskRef{
            shared.TaskRef{},
        },
        SubjectIds: []string{
            "Oklahoma",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertifyInput{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{},
                },
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{},
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{},
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{},
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{},
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

