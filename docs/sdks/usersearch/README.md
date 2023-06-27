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
            "suscipit",
            "assumenda",
        },
        ExpandMask: &shared.UserExpandMask{
            Paths: []string{
                "praesentium",
            },
        },
        Ids: []string{
            "veritatis",
            "ipsa",
            "id",
            "quidem",
        },
        PageSize: conductoroneapi.Float64(2065.94),
        PageToken: conductoroneapi.String("quo"),
        Query: conductoroneapi.String("illum"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("a4251904-e523-4c7e-8bc7-178e4796f2a7"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("0c688282-aa48-4256-af22-2e9817ee17cb"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("e61e6b7b-95bc-40ab-bc20-c4f3789fd871"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("f99dd2ef-d121-4aa6-b1e6-74bdb04f1575"),
            },
        },
        RoleIds: []string{
            "aut",
            "voluptatum",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
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

