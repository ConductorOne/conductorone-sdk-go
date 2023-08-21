# AppUser

### Available Operations

* [Update](#update) - Update

## Update

Update an app user by ID. Only the fields specified in the update mask are updated.
 Currently, only the appUserType, and identityUserId fields can be updated.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppUser.Update(ctx, operations.C1APIAppV1AppUserServiceUpdateRequest{
        AppUserServiceUpdateRequestInput: &shared.AppUserServiceUpdateRequestInput{
            AppUser: &shared.AppUserInput{
                AppUserStatus: &shared.AppUserStatusInput{},
                AppUserType: shared.AppUserAppUserTypeAppUserTypeUnspecified.ToPointer(),
            },
            AppUserExpandMask: &shared.AppUserExpandMask{
                Paths: []string{
                    "praesentium",
                },
            },
            UpdateMask: conductoroneapi.String("natus"),
        },
        AppUserAppID: "magni",
        AppUserID: "sunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppUserServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.C1APIAppV1AppUserServiceUpdateRequest](../../models/operations/c1apiappv1appuserserviceupdaterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.C1APIAppV1AppUserServiceUpdateResponse](../../models/operations/c1apiappv1appuserserviceupdateresponse.md), error**
