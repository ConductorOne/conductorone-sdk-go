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
                "repellat",
                "quisquam",
            },
        },
        Email: conductoroneapi.String("Jamison29@gmail.com"),
        ExcludeIds: []string{
            "aliquam",
            "quisquam",
            "provident",
            "laudantium",
        },
        Ids: []string{
            "consequatur",
            "maxime",
            "aspernatur",
            "nam",
        },
        PageSize: conductoroneapi.Float64(7119.91),
        PageToken: conductoroneapi.String("quas"),
        Query: conductoroneapi.String("provident"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("b75dad63-6c60-4050-bd8b-b31180f739ae"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("9e057eb8-09e2-4810-b31f-3981d4c700b6"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("07f3c93c-73b9-4da3-b2ce-da7e23f22574"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("11faf4b7-544e-4472-a802-857a5b40463a"),
            },
        },
        RoleIds: []string{
            "fugiat",
            "nostrum",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
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

