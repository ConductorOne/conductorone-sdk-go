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
        Email: conductoroneapi.String("Ike48@gmail.com"),
        ExcludeIds: []string{
            "rem",
            "eligendi",
            "fugiat",
        },
        ExpandMask: &shared.UserExpandMask{
            Paths: []string{
                "officiis",
                "ducimus",
                "dolor",
            },
        },
        Ids: []string{
            "error",
        },
        PageSize: conductoroneapi.Float64(7847.31),
        PageToken: conductoroneapi.String("vitae"),
        Query: conductoroneapi.String("dignissimos"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("d525f77b-114e-4eb5-aff7-85fc37814d4c"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("98e0c2bb-89eb-475d-ad63-6c600503d8bb"),
            },
        },
        RoleIds: []string{
            "quasi",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
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

