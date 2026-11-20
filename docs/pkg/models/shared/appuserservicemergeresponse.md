# AppUserServiceMergeResponse

The AppUserServiceMergeResponse returns the target app user after the merge.


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `AppUserView`                                                                            | [*shared.AppUserView](../../../pkg/models/shared/appuserview.md)                         | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `MergeID`                                                                                | `*string`                                                                                | :heavy_minus_sign:                                                                       | The ID of the merge record written for this merge.                                       |
| `RecordsMoved`                                                                           | `*int`                                                                                   | :heavy_minus_sign:                                                                       | The number of connector and external source records moved from the source to the target. |