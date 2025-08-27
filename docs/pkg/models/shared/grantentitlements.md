# GrantEntitlements

The GrantEntitlements message.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `UserRef`                                                                     | [*shared.UserRef](../../../pkg/models/shared/userref.md)                      | :heavy_minus_sign:                                                            | A reference to a user.                                                        |
| `AppEntitlementRefs`                                                          | [][shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md) | :heavy_minus_sign:                                                            | The appEntitlementRefs field.                                                 |
| `AppEntitlementRefsCel`                                                       | **string*                                                                     | :heavy_minus_sign:                                                            | The appEntitlementRefsCel field.                                              |
| `UseSubjectUser`                                                              | **bool*                                                                       | :heavy_minus_sign:                                                            | If true, the step will use the subject user of the automation as the subject. |
| `UserIDCel`                                                                   | **string*                                                                     | :heavy_minus_sign:                                                            | The userIdCel field.                                                          |