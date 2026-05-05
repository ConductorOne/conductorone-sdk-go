# SetAppEntitlementOwnersV2Request

SetAppEntitlementOwnersV2Request is the request for setting the owners of an app entitlement for a given role.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AppEntitlementRefs`                                                          | [][shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md) | :heavy_minus_sign:                                                            | The appEntitlementRefs field.                                                 |
| `RoleSlug`                                                                    | `*string`                                                                     | :heavy_minus_sign:                                                            | Empty defaults to the "primary" role on the server side.                      |
| `UserRefs`                                                                    | [][shared.UserRef](../../../pkg/models/shared/userref.md)                     | :heavy_minus_sign:                                                            | The userRefs field.                                                           |