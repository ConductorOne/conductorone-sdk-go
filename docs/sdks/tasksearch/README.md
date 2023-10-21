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
                "string",
            },
        },
        AccessReviewIds: []string{
            "string",
        },
        AccountOwnerIds: []string{
            "string",
        },
        AppEntitlementIds: []string{
            "string",
        },
        AppResourceIds: []string{
            "string",
        },
        AppResourceTypeIds: []string{
            "string",
        },
        AppUserSubjectIds: []string{
            "string",
        },
        ApplicationIds: []string{
            "string",
        },
        AssigneesInIds: []string{
            "string",
        },
        ExcludeAppEntitlementIds: []string{
            "string",
        },
        ExcludeIds: []string{
            "string",
        },
        MyWorkUserIds: []string{
            "string",
        },
        OpenerIds: []string{
            "string",
        },
        PreviouslyActedOnIds: []string{
            "string",
        },
        Refs: []shared.TaskRef{
            shared.TaskRef{},
        },
        SubjectIds: []string{
            "string",
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

