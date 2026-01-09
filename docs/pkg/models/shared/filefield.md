# FileField

The FileField message.

This message contains a oneof named view. Only a single field of the following list may be set at a time:
  - fileInputField


This message contains a oneof named _max_file_size. Only a single field of the following list may be set at a time:
  - maxFileSize



## Fields

| Field                                                                                                                                           | Type                                                                                                                                            | Required                                                                                                                                        | Description                                                                                                                                     |
| ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `FileInputField`                                                                                                                                | [*shared.FileInputField](../../../pkg/models/shared/fileinputfield.md)                                                                          | :heavy_minus_sign:                                                                                                                              | The FileInputField message.                                                                                                                     |
| `AcceptedFileTypes`                                                                                                                             | []*string*                                                                                                                                      | :heavy_minus_sign:                                                                                                                              | The acceptedFileTypes field.                                                                                                                    |
| `MaxFileSize`                                                                                                                                   | **int64*                                                                                                                                        | :heavy_minus_sign:                                                                                                                              | The maxFileSize field.<br/>This field is part of the `_max_file_size` oneof.<br/>See the documentation for `c1.api.form.v1.FileField` for more details. |