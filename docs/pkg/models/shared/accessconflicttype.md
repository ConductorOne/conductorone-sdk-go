# AccessConflictType

AccessConflictType identifies an access-conflict finding: one conflict
 monitor and one piece of conflicting access held by the target user.

This message contains a oneof named match. Only a single field of the following list may be set at a time:
  - entitlementPair
  - aWithoutB



## Fields

| Field                                                                                                          | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `AWithoutB`                                                                                                    | [*shared.AccessConflictAWithoutBMatch](../../../pkg/models/shared/accessconflictawithoutbmatch.md)             | :heavy_minus_sign:                                                                                             | N/A                                                                                                            |
| `EntitlementPair`                                                                                              | [*shared.AccessConflictEntitlementPairMatch](../../../pkg/models/shared/accessconflictentitlementpairmatch.md) | :heavy_minus_sign:                                                                                             | N/A                                                                                                            |
| `MonitorID`                                                                                                    | `*string`                                                                                                      | :heavy_minus_sign:                                                                                             | The conflict monitor that detected the conflict.                                                               |