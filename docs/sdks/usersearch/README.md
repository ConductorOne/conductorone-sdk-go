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
                "molestias",
                "magnam",
                "saepe",
                "consequuntur",
            },
        },
        ExcludeIds: []string{
            "officiis",
            "perspiciatis",
            "in",
        },
        Ids: []string{
            "eveniet",
        },
        PageSize: conductoroneapi.Float64(5808.87),
        PageToken: conductoroneapi.String("consequuntur"),
        Query: conductoroneapi.String("fugit"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("57a15be3-e060-4807-a2b6-e3ab8845f059"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("7a60ff2a-54a3-41e9-8764-a3e865e7956f"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("9251a5a9-da66-40ff-97bf-aad4f9efc1b4"),
            },
        },
        RoleIds: []string{
            "inventore",
            "fugit",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesUnknown,
            shared.SearchUsersRequestUserStatusesUnknown,
            shared.SearchUsersRequestUserStatusesUnknown,
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

