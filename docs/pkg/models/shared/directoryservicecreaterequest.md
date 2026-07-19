# DirectoryServiceCreateRequest

Uplevel an app into a full directory.

This message contains a oneof named account_filter. Only a single field of the following list may be set at a time:
  - all
  - celExpression



## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `All`                                                                                        | [*shared.DirectoryAccountFilterAll](../../../pkg/models/shared/directoryaccountfilterall.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `AppID`                                                                                      | `*string`                                                                                    | :heavy_minus_sign:                                                                           | The AppID to make into a directory, providing identities and more for the C1 app.            |
| `CelExpression`                                                                              | [*shared.DirectoryAccountFilterCel](../../../pkg/models/shared/directoryaccountfiltercel.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `ExpandMask`                                                                                 | [*shared.DirectoryExpandMask](../../../pkg/models/shared/directoryexpandmask.md)             | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `MergeConfig`                                                                                | [*shared.DirectoryMergeConfig](../../../pkg/models/shared/directorymergeconfig.md)           | :heavy_minus_sign:                                                                           | N/A                                                                                          |