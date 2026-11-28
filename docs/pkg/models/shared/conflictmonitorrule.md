# ConflictMonitorRule

The rule a conflict monitor detects. A rule's entitlement lists are set
 through the monitor's entitlement bindings (groups A and B).

This message contains a oneof named type. Only a single field of the following list may be set at a time:
  - holdsBoth
  - holdsAWithoutB



## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `HoldsAWithoutB`                                                                                             | [*shared.ConflictMonitorHoldsAWithoutBRule](../../../pkg/models/shared/conflictmonitorholdsawithoutbrule.md) | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `HoldsBoth`                                                                                                  | [*shared.ConflictMonitorHoldsBothRule](../../../pkg/models/shared/conflictmonitorholdsbothrule.md)           | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |