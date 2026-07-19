# DirectoryServiceUpdateRequest

Update a directory by app_id.

This message contains a oneof named account_filter. Only a single field of the following list may be set at a time:
  - all
  - celExpression



## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `All`                                                                                        | [*shared.DirectoryAccountFilterAll](../../../pkg/models/shared/directoryaccountfilterall.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `CelExpression`                                                                              | [*shared.DirectoryAccountFilterCel](../../../pkg/models/shared/directoryaccountfiltercel.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `ExpandMask`                                                                                 | [*shared.DirectoryExpandMask](../../../pkg/models/shared/directoryexpandmask.md)             | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `MergeConfig`                                                                                | [*shared.DirectoryMergeConfig](../../../pkg/models/shared/directorymergeconfig.md)           | :heavy_minus_sign:                                                                           | N/A                                                                                          |