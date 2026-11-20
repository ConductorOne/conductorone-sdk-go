# MyOAuthAuthorization

MyOAuthAuthorization exposes the information needed to connect an authorization.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AllowedRoleIds`                                                              | []`string`                                                                    | :heavy_minus_sign:                                                            | The allowedRoleIds field.                                                     |
| `ClientID`                                                                    | `*string`                                                                     | :heavy_minus_sign:                                                            | The clientId field.                                                           |
| `DisplayName`                                                                 | `*string`                                                                     | :heavy_minus_sign:                                                            | The displayName field.                                                        |
| `ID`                                                                          | `*string`                                                                     | :heavy_minus_sign:                                                            | The id field.                                                                 |
| `UseEntitlementAppID`                                                         | `*string`                                                                     | :heavy_minus_sign:                                                            | The C1 app containing the entitlement that governs use of this authorization. |
| `UseEntitlementID`                                                            | `*string`                                                                     | :heavy_minus_sign:                                                            | The useEntitlementId field.                                                   |