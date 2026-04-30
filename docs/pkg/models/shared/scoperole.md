# ScopeRole

Scope-role variant of TaskTypeAction.target_object. The UI uses the
 embedded identifiers to build links and title strings without a separate
 Action fetch.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `AppID`                                                  | `*string`                                                | :heavy_minus_sign:                                       | The IaaS/sparse-ACL app the (scope, role) pair lives on. |
| `RoleResourceID`                                         | `*string`                                                | :heavy_minus_sign:                                       | The roleResourceId field.                                |
| `RoleResourceTypeID`                                     | `*string`                                                | :heavy_minus_sign:                                       | The roleResourceTypeId field.                            |
| `ScopeResourceID`                                        | `*string`                                                | :heavy_minus_sign:                                       | The scopeResourceId field.                               |
| `ScopeResourceTypeID`                                    | `*string`                                                | :heavy_minus_sign:                                       | The scopeResourceTypeId field.                           |