# FileField

The FileField message.

This message contains a oneof named view. Only a single field of the following list may be set at a time:
  - fileInputField



## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `FileInputField`                                                       | [*shared.FileInputField](../../../pkg/models/shared/fileinputfield.md) | :heavy_minus_sign:                                                     | The FileInputField message.                                            |
| `AcceptedFileTypes`                                                    | []`string`                                                             | :heavy_minus_sign:                                                     | The acceptedFileTypes field.                                           |
| `MaxFileSize`                                                          | `*int64`                                                               | :heavy_minus_sign:                                                     | The maxFileSize field.                                                 |