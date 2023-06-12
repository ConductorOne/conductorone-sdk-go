# UserSearch

### Available Operations

* [Search](#search) - Invokes the c1.api.user.v1.UserSearch.Search method.

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
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.UserSearch.Search(ctx, shared.SearchUsersRequest{
        ExcludeIds: []string{
            "eligendi",
            "ducimus",
        },
        ExpandMask: &shared.UserExpandMask{
            Paths: []string{
                "officia",
            },
        },
        Ids: []string{
            "ipsam",
            "ea",
        },
        PageSize: conductoroneapi.Float64(1369),
        PageToken: conductoroneapi.String("vel"),
        Query: conductoroneapi.String("possimus"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("36813f16-d9f5-4fce-ac55-6146c3e250fb"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("008c42e1-41aa-4c36-ac8d-d6b144290747"),
            },
        },
        RoleIds: []string{
            "esse",
            "esse",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesDisabled,
            shared.SearchUsersRequestUserStatusesEnabled,
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

