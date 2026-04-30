# UserWithAppEntitlementUserBindingView

The UserWithAppEntitlementUserBindingView message.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `User`                                                                                  | [*shared.User](../../../pkg/models/shared/user.md)                                      | :heavy_minus_sign:                                                                      | The User object provides all of the details for an user, as well as some configuration. |
| `AppEntitlementID`                                                                      | `*string`                                                                               | :heavy_minus_sign:                                                                      | The ID of the app entitlement.                                                          |
| `AppID`                                                                                 | `*string`                                                                               | :heavy_minus_sign:                                                                      | The ID of the app that contains the entitlement.                                        |
| `AppUserID`                                                                             | `*string`                                                                               | :heavy_minus_sign:                                                                      | The ID of the app user associated with this binding.                                    |