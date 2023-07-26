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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.UserSearch.Search(ctx, shared.SearchUsersRequest{
        UserExpandMask: &shared.UserExpandMask{
            Paths: []string{
                "impedit",
                "delectus",
                "tempore",
            },
        },
        Email: conductoroneapi.String("Bruce.Zieme44@hotmail.com"),
        ExcludeIds: []string{
            "odio",
        },
        Ids: []string{
            "in",
            "ducimus",
        },
        PageSize: conductoroneapi.Float64(5678.46),
        PageToken: conductoroneapi.String("dolores"),
        Query: conductoroneapi.String("error"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("77deac64-6ecb-4573-809e-3eb1e5a2b12e"),
            },
        },
        RoleIds: []string{
            "ipsa",
            "ducimus",
            "maiores",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
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

