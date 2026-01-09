# DirectoryServiceUpdateRequest

Update a directory by app_id.

This message contains a oneof named account_filter. Only a single field of the following list may be set at a time:
  - all
  - celExpression



## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `DirectoryAccountFilterAll`                                                                  | [*shared.DirectoryAccountFilterAll](../../../pkg/models/shared/directoryaccountfilterall.md) | :heavy_minus_sign:                                                                           | The DirectoryAccountFilterAll message.                                                       |
| `DirectoryAccountFilterCel`                                                                  | [*shared.DirectoryAccountFilterCel](../../../pkg/models/shared/directoryaccountfiltercel.md) | :heavy_minus_sign:                                                                           | The DirectoryAccountFilterCel message.                                                       |
| `DirectoryExpandMask`                                                                        | [*shared.DirectoryExpandMask](../../../pkg/models/shared/directoryexpandmask.md)             | :heavy_minus_sign:                                                                           | The fields to be included in the directory response.                                         |