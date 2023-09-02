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
                "dignissimos",
                "libero",
            },
        },
        Email: conductoroneapi.String("Annie_Goyette73@gmail.com"),
        ExcludeIds: []string{
            "eos",
            "reiciendis",
        },
        Ids: []string{
            "reprehenderit",
            "praesentium",
            "nemo",
            "repellat",
        },
        PageSize: conductoroneapi.Float64(7897.7),
        PageToken: conductoroneapi.String("sequi"),
        Query: conductoroneapi.String("nihil"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("14d4c98e-0c2b-4b89-ab75-dad636c60050"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("3d8bb311-80f7-439a-a9e0-57eb809e2810"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("331f3981-d4c7-400b-a07f-3c93c73b9da3"),
            },
        },
        RoleIds: []string{
            "aspernatur",
            "quo",
            "itaque",
            "illum",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesEnabled,
            shared.SearchUsersRequestUserStatusesDeleted,
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

