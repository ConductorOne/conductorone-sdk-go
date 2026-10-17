# ScopeRoleBindingRef

ScopeRoleBindingRef uniquely identifies a scope's binding state within an access profile and app.


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `AppID`                                          | `*string`                                        | :heavy_minus_sign:                               | The unique identifier of the app.                |
| `CatalogID`                                      | `*string`                                        | :heavy_minus_sign:                               | The unique identifier of the access profile.     |
| `ScopeID`                                        | `*string`                                        | :heavy_minus_sign:                               | The unique identifier of the scope app resource. |