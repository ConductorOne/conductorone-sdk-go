# SyncResourceTypeDelta

SyncResourceTypeDelta compares resource, entitlement and grant counts for one resource type.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Entitlements`                                                                   | [*shared.SyncCountDelta](../../../pkg/models/shared/synccountdelta.md)           | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Grants`                                                                         | [*shared.SyncCountDelta](../../../pkg/models/shared/synccountdelta.md)           | :heavy_minus_sign:                                                               | N/A                                                                              |
| `ResourceTypeID`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | Connector-defined resource type identifier, not an individual record identifier. |
| `Resources`                                                                      | [*shared.SyncCountDelta](../../../pkg/models/shared/synccountdelta.md)           | :heavy_minus_sign:                                                               | N/A                                                                              |