# CreateAccessRequestDefaultsRequest

The CreateAccessRequestDefaultsRequest message is used to create a new app request defaults.

This message contains a oneof named max_grant_duration. Only a single field of the following list may be set at a time:
  - durationUnset
  - durationGrant



## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `CatalogIds`                                                         | []*string*                                                           | :heavy_minus_sign:                                                   | The catalogIds field.                                                |
| `DurationGrant`                                                      | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `DurationUnset`                                                      | [*shared.DurationUnset](../../../pkg/models/shared/durationunset.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `EmergencyGrantEnabled`                                              | **bool*                                                              | :heavy_minus_sign:                                                   | The emergencyGrantEnabled field.                                     |
| `EmergencyGrantPolicyID`                                             | **string*                                                            | :heavy_minus_sign:                                                   | The emergencyGrantPolicyId field.                                    |