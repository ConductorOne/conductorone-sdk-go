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
                "vero",
                "odit",
                "repellat",
            },
        },
        Email: conductoroneapi.String("Francisca48@hotmail.com"),
        ExcludeIds: []string{
            "in",
            "ducimus",
        },
        Ids: []string{
            "dolores",
            "error",
            "veritatis",
        },
        PageSize: conductoroneapi.Float64(4981.8),
        PageToken: conductoroneapi.String("voluptate"),
        Query: conductoroneapi.String("pariatur"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("ac646ecb-5734-409e-beb1-e5a2b12eb07f"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("116db995-45fc-495f-a889-70e189dbb30f"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("cb33ea05-5b19-47cd-84e2-f52d82d3513b"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("b6f48b65-6bcd-4b35-bf2e-4b27537a8cd9"),
            },
        },
        RoleIds: []string{
            "ducimus",
            "dolor",
            "dicta",
            "error",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesUnknown,
            shared.SearchUsersRequestUserStatusesEnabled,
            shared.SearchUsersRequestUserStatusesEnabled,
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

