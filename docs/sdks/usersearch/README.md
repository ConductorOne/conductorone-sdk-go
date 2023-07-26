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
                "dolores",
                "assumenda",
            },
        },
        Email: conductoroneapi.String("Marcella.Schumm@gmail.com"),
        ExcludeIds: []string{
            "accusamus",
            "necessitatibus",
            "tempore",
        },
        Ids: []string{
            "ea",
            "autem",
            "ipsam",
        },
        PageSize: conductoroneapi.Float64(7029.52),
        PageToken: conductoroneapi.String("laudantium"),
        Query: conductoroneapi.String("corporis"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("fbd02bae-0be2-4d78-a259-e3ea4b5197f9"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("2443da7c-e52b-4895-8537-c6454efb0b34"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("896c3ca5-acfb-4e2f-9570-7577929177de"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("ac646ecb-5734-409e-beb1-e5a2b12eb07f"),
            },
        },
        RoleIds: []string{
            "quasi",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesDeleted,
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

