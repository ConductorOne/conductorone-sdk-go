# UserSearch

### Available Operations

* [Search](#search) - Search

## Search

Search users based on filters specified in the request body.

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
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.UserSearch.Search(ctx, shared.SearchUsersRequest{
        UserExpandMask: &shared.UserExpandMask{
            Paths: []string{
                "tempore",
                "nisi",
                "voluptatibus",
            },
        },
        Email: conductoroneapi.String("Joanie.Raynor@gmail.com"),
        ExcludeIds: []string{
            "libero",
            "minus",
        },
        Ids: []string{
            "facilis",
            "ipsum",
            "ad",
            "voluptatibus",
        },
        PageSize: conductoroneapi.Float64(9745.89),
        PageToken: conductoroneapi.String("consequuntur"),
        Query: conductoroneapi.String("debitis"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("b27537a8-cd9e-4731-9c17-7d525f77b114"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("eeb52ff7-85fc-4378-94d4-c98e0c2bb89e"),
            },
        },
        RoleIds: []string{
            "dignissimos",
            "corporis",
            "vero",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesDeleted,
            shared.SearchUsersRequestUserStatusesEnabled,
            shared.SearchUsersRequestUserStatusesUnknown,
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

