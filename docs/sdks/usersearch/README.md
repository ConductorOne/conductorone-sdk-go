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
                "officia",
            },
        },
        Email: conductoroneapi.String("Green.Bahringer@yahoo.com"),
        ExcludeIds: []string{
            "iste",
        },
        Ids: []string{
            "id",
        },
        PageSize: conductoroneapi.Float64(700.42),
        PageToken: conductoroneapi.String("error"),
        Query: conductoroneapi.String("possimus"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("eaab5851-d6c6-445b-88b6-1891baa0fe1a"),
            },
        },
        RoleIds: []string{
            "pariatur",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesDeleted,
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

