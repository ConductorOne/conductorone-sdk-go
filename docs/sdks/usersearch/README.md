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
                "fugiat",
                "ad",
            },
        },
        Email: conductoroneapi.String("Emil.Wintheiser@gmail.com"),
        ExcludeIds: []string{
            "illo",
            "ab",
            "incidunt",
        },
        Ids: []string{
            "saepe",
            "tempore",
            "veniam",
            "eos",
        },
        PageSize: conductoroneapi.Float64(9700.79),
        PageToken: conductoroneapi.String("earum"),
        Query: conductoroneapi.String("reprehenderit"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("5fc37814-d4c9-48e0-82bb-89eb75dad636"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("c600503d-8bb3-4118-8f73-9ae9e057eb80"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("9e281033-1f39-481d-8c70-0b607f3c93c7"),
            },
        },
        RoleIds: []string{
            "soluta",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesDeleted,
            shared.SearchUsersRequestUserStatusesDisabled,
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

