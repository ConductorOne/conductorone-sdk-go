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
                "similique",
                "ut",
                "quidem",
                "quis",
            },
        },
        Email: conductoroneapi.String("Layla_Kreiger14@gmail.com"),
        ExcludeIds: []string{
            "numquam",
            "nesciunt",
        },
        Ids: []string{
            "officia",
            "dignissimos",
            "optio",
            "necessitatibus",
        },
        PageSize: conductoroneapi.Float64(3591.11),
        PageToken: conductoroneapi.String("qui"),
        Query: conductoroneapi.String("expedita"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("95c537c6-454e-4fb0-b348-96c3ca5acfbe"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("2fd57075-7792-4917-bdea-c646ecb57340"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("9e3eb1e5-a2b1-42eb-87f1-16db99545fc9"),
            },
        },
        RoleIds: []string{
            "hic",
            "animi",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesDisabled,
            shared.SearchUsersRequestUserStatusesDisabled,
            shared.SearchUsersRequestUserStatusesEnabled,
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

