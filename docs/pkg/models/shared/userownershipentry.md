# UserOwnershipEntry

A single ownership entry. Fields are populated based on ownership_type:
 APP — only app_id and app_display_name are set.
 RESOURCE — app_id, app_display_name, resource_type_id, resource_id, and resource_display_name are set.
 ENTITLEMENT — app_id, app_display_name, resource_type_id, entitlement_id, and entitlement_display_name are set.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `AppDisplayName`                                                     | `*string`                                                            | :heavy_minus_sign:                                                   | The app display name.                                                |
| `AppID`                                                              | `*string`                                                            | :heavy_minus_sign:                                                   | The app ID.                                                          |
| `EntitlementDisplayName`                                             | `*string`                                                            | :heavy_minus_sign:                                                   | The entitlement display name, if applicable.                         |
| `EntitlementID`                                                      | `*string`                                                            | :heavy_minus_sign:                                                   | The entitlement ID, if applicable.                                   |
| `OwnershipType`                                                      | [*shared.OwnershipType](../../../pkg/models/shared/ownershiptype.md) | :heavy_minus_sign:                                                   | The type of ownership.                                               |
| `ResourceDisplayName`                                                | `*string`                                                            | :heavy_minus_sign:                                                   | The resource display name, if applicable.                            |
| `ResourceID`                                                         | `*string`                                                            | :heavy_minus_sign:                                                   | The resource ID, if applicable.                                      |
| `ResourceTypeID`                                                     | `*string`                                                            | :heavy_minus_sign:                                                   | The resource type ID, if applicable.                                 |