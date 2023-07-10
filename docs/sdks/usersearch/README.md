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
        UserExpandMask: &shared.UserExpandMask{
            Paths: []string{
                "dolores",
            },
        },
        ExcludeIds: []string{
            "laboriosam",
            "velit",
        },
        Ids: []string{
            "molestias",
            "magnam",
            "saepe",
            "consequuntur",
        },
        PageSize: conductoroneapi.Float64(5801.07),
        PageToken: conductoroneapi.String("officiis"),
        Query: conductoroneapi.String("perspiciatis"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("3e922a57-a15b-4e3e-8608-07e2b6e3ab88"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("45f0597a-60ff-42a5-8a31-e94764a3e865"),
            },
        },
        RoleIds: []string{
            "esse",
            "provident",
            "quis",
            "eum",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesDisabled,
            shared.SearchUsersRequestUserStatusesUnknown,
            shared.SearchUsersRequestUserStatusesEnabled,
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

