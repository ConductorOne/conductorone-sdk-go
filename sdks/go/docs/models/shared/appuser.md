# AppUser

Application User that represents an account in the application.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `AppID`                                                          | **string*                                                        | :heavy_minus_sign:                                               | The ID of the application.                                       |
| `AppUserType`                                                    | [*AppUserAppUserType](../../models/shared/appuserappusertype.md) | :heavy_minus_sign:                                               | The appplication user type. Type can be user, system or service. |
| `CreatedAt`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_minus_sign:                                               | N/A                                                              |
| `DeletedAt`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_minus_sign:                                               | N/A                                                              |
| `DisplayName`                                                    | **string*                                                        | :heavy_minus_sign:                                               | The display name of the application user.                        |
| `Email`                                                          | **string*                                                        | :heavy_minus_sign:                                               | The email field of the application user.                         |
| `ID`                                                             | **string*                                                        | :heavy_minus_sign:                                               | A unique idenditfier of the application user.                    |
| `IdentityUserID`                                                 | **string*                                                        | :heavy_minus_sign:                                               | The conductor one user ID of the account owner.                  |
| `Profile`                                                        | map[string]*interface{}*                                         | :heavy_minus_sign:                                               | N/A                                                              |
| `Status`                                                         | [*AppUserStatus](../../models/shared/appuserstatus.md)           | :heavy_minus_sign:                                               | The satus of the applicaiton user.                               |
| `UpdatedAt`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_minus_sign:                                               | N/A                                                              |