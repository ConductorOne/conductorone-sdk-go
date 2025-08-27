# UsageBasedRevocationTrigger

The UsageBasedRevocationTrigger message.

This message contains a oneof named cold_start_schedule. Only a single field of the following list may be set at a time:
  - runImmediately
  - runDelayed



## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `RunDelayed`                                                                        | [*shared.RunDelayed](../../../pkg/models/shared/rundelayed.md)                      | :heavy_minus_sign:                                                                  | The RunDelayed message.                                                             |
| `RunImmediately`                                                                    | [*shared.RunImmediately](../../../pkg/models/shared/runimmediately.md)              | :heavy_minus_sign:                                                                  | No fields needed; this just indicates the trigger should run immediately            |
| `AppID`                                                                             | **string*                                                                           | :heavy_minus_sign:                                                                  | The appId field.                                                                    |
| `EnabledAt`                                                                         | [*time.Time](https://pkg.go.dev/time#Time)                                          | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `ExcludedGroupRefs`                                                                 | [][shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md)       | :heavy_minus_sign:                                                                  | The excludedGroupRefs field.                                                        |
| `ExcludedUserRefs`                                                                  | [][shared.UserRef](../../../pkg/models/shared/userref.md)                           | :heavy_minus_sign:                                                                  | The excludedUserRefs field.                                                         |
| `IncludeUsersWithNoActivity`                                                        | **bool*                                                                             | :heavy_minus_sign:                                                                  | The includeUsersWithNoActivity field.                                               |
| `TargetedAppUserTypes`                                                              | [][shared.TargetedAppUserTypes](../../../pkg/models/shared/targetedappusertypes.md) | :heavy_minus_sign:                                                                  | The targetedAppUserTypes field.                                                     |
| `TargetedEntitlementRefs`                                                           | [][shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md)       | :heavy_minus_sign:                                                                  | The targetedEntitlementRefs field.                                                  |
| `UnusedForDays`                                                                     | **int64*                                                                            | :heavy_minus_sign:                                                                  | The unusedForDays field.                                                            |