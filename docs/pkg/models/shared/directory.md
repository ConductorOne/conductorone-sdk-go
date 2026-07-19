# Directory

This object indicates that an app is also a directory.

This message contains a oneof named account_filter. Only a single field of the following list may be set at a time:
  - all
  - celExpression



## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `All`                                                                                        | [*shared.DirectoryAccountFilterAll](../../../pkg/models/shared/directoryaccountfilterall.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `AppID`                                                                                      | `*string`                                                                                    | :heavy_minus_sign:                                                                           | The ID of the app associated with the directory.                                             |
| `CelExpression`                                                                              | [*shared.DirectoryAccountFilterCel](../../../pkg/models/shared/directoryaccountfiltercel.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `CreatedAt`                                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `DeletedAt`                                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `MergeConfig`                                                                                | [*shared.DirectoryMergeConfig](../../../pkg/models/shared/directorymergeconfig.md)           | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `UpdatedAt`                                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |