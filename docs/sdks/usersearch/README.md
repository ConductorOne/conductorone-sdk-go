# UserSearch

### Available Operations

* [Search](#search) - Invokes the c1.api.user.v1.UserSearch.Search method.

## Search

Invokes the c1.api.user.v1.UserSearch.Search method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.UserSearch.Search(ctx, shared.SearchUsersRequest{
        ExcludeIds: []string{
            "distinctio",
            "voluptatum",
            "rem",
        },
        ExpandMask: &shared.UserExpandMask{
            Paths: []string{
                "ad",
                "repellat",
            },
        },
        Ids: []string{
            "corporis",
        },
        PageSize: conductoroneapi.Float64(5973.03),
        PageToken: conductoroneapi.String("nihil"),
        Query: conductoroneapi.String("mollitia"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("0ff2a54a-31e9-4476-8a3e-865e7956f925"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("1a5a9da6-60ff-457b-baad-4f9efc1b4512"),
            },
        },
        RoleIds: []string{
            "quae",
            "perferendis",
            "velit",
            "aspernatur",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesEnabled,
            shared.SearchUsersRequestUserStatusesDisabled,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchUsersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.SearchUsersRequest](../../models/shared/searchusersrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.C1APIUserV1UserSearchSearchResponse](../../models/operations/c1apiuserv1usersearchsearchresponse.md), error**

