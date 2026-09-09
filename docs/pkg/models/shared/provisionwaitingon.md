# ProvisionWaitingOn

Describes why a provision step is paused in the WAITING state.

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - entitlementMerge
  - devicePlacement



## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `DevicePlacement`                                                                              | [*shared.WaitingForDevicePlacement](../../../pkg/models/shared/waitingfordeviceplacement.md)   | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `EntitlementMerge`                                                                             | [*shared.WaitingForEntitlementMerge](../../../pkg/models/shared/waitingforentitlementmerge.md) | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `FallbackAt`                                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                                     | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `StartedWaitingAt`                                                                             | [*time.Time](https://pkg.go.dev/time#Time)                                                     | :heavy_minus_sign:                                                                             | N/A                                                                                            |