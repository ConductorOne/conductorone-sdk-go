# FileField

The FileField message.

This message contains a oneof named view. Only a single field of the following list may be set at a time:
  - fileInputField



## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AcceptedFileTypes`                                                    | []`string`                                                             | :heavy_minus_sign:                                                     | The acceptedFileTypes field.                                           |
| `FileInputField`                                                       | [*shared.FileInputField](../../../pkg/models/shared/fileinputfield.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `MaxFileSize`                                                          | `*int64`                                                               | :heavy_minus_sign:                                                     | The maxFileSize field.                                                 |