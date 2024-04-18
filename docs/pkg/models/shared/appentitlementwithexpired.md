# AppEntitlementWithExpired

The AppEntitlementWithExpired message.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `AppUser`                                                                               | [*shared.AppUser](../../../pkg/models/shared/appuser.md)                                | :heavy_minus_sign:                                                                      | Application User that represents an account in the application.                         |
| `Discovered`                                                                            | [*time.Time](https://pkg.go.dev/time#Time)                                              | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `Expired`                                                                               | [*time.Time](https://pkg.go.dev/time#Time)                                              | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `User`                                                                                  | [*shared.User](../../../pkg/models/shared/user.md)                                      | :heavy_minus_sign:                                                                      | The User object provides all of the details for an user, as well as some configuration. |