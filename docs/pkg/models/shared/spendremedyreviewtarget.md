# SpendRemedyReviewTarget

A missing row is not a deleted version. Current state may differ from capture.

This message contains a oneof named state. Only a single field of the following list may be set at a time:
  - absent
  - live
  - deleted



## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `Absent`                                                                                         | [*shared.SpendRemedyAbsentTarget](../../../pkg/models/shared/spendremedyabsenttarget.md)         | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `Deleted`                                                                                        | [*shared.SpendRemedyControlsSnapshot](../../../pkg/models/shared/spendremedycontrolssnapshot.md) | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `Live`                                                                                           | [*shared.SpendRemedyControlsSnapshot](../../../pkg/models/shared/spendremedycontrolssnapshot.md) | :heavy_minus_sign:                                                                               | N/A                                                                                              |