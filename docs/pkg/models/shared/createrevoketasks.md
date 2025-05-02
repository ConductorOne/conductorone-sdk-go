# CreateRevokeTasks

The CreateRevokeTasks message.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `UserRef`                                                                     | [*shared.UserRef](../../../pkg/models/shared/userref.md)                      | :heavy_minus_sign:                                                            | A reference to a user.                                                        |
| `AppEntitlementRefs`                                                          | [][shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md) | :heavy_minus_sign:                                                            | The appEntitlementRefs field.                                                 |
| `AppEntitlementRefsCel`                                                       | **string*                                                                     | :heavy_minus_sign:                                                            | The appEntitlementRefsCel field.                                              |
| `RevokeAll`                                                                   | **bool*                                                                       | :heavy_minus_sign:                                                            | The revokeAll field.                                                          |
| `UserIDCel`                                                                   | **string*                                                                     | :heavy_minus_sign:                                                            | The userIdCel field.                                                          |