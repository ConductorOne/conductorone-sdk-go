# UserSearch

### Available Operations

* [Search](#search) - Search

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
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.UserSearch.Search(ctx, shared.SearchUsersRequest{
        UserExpandMask: &shared.UserExpandMask{
            Paths: []string{
                "minus",
                "esse",
            },
        },
        Email: conductoroneapi.String("Al.Price@gmail.com"),
        ExcludeIds: []string{
            "repellat",
            "velit",
        },
        Ids: []string{
            "provident",
            "consectetur",
            "eligendi",
            "dignissimos",
        },
        PageSize: conductoroneapi.Float64(2358.34),
        PageToken: conductoroneapi.String("soluta"),
        Query: conductoroneapi.String("natus"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("a3f2ceda-7e23-4f22-9741-1faf4b7544e4"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("72e80285-7a5b-4404-a3a7-d575f1400e76"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("4ad7334e-c1b7-481b-b6a0-8088d100efad"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("a200ef04-22eb-4216-8cf9-ab8366c723ff"),
            },
        },
        RoleIds: []string{
            "laborum",
            "natus",
            "accusamus",
            "doloremque",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesDisabled,
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

