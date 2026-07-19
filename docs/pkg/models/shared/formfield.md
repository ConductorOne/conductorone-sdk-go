# FormField

A field is a single input meant to collect a piece of data from a user

This message contains a oneof named type. Only a single field of the following list may be set at a time:
  - stringField
  - boolField
  - stringSliceField
  - int64Field
  - fileField
  - oauth2Field
  - stringMapField


This message contains a oneof named provider_config. Only a single field of the following list may be set at a time:
  - userConfig
  - adminConfig
  - sharedConfig



## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `AdminConfig`                                                                      | [*shared.AdminProviderConfig](../../../pkg/models/shared/adminproviderconfig.md)   | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `BoolField`                                                                        | [*shared.BoolField](../../../pkg/models/shared/boolfield.md)                       | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `Description`                                                                      | `*string`                                                                          | :heavy_minus_sign:                                                                 | The description field.                                                             |
| `DisplayName`                                                                      | `*string`                                                                          | :heavy_minus_sign:                                                                 | The displayName field.                                                             |
| `FileField`                                                                        | [*shared.FileField](../../../pkg/models/shared/filefield.md)                       | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `Int64Field`                                                                       | [*shared.Int64Field](../../../pkg/models/shared/int64field.md)                     | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `Name`                                                                             | `*string`                                                                          | :heavy_minus_sign:                                                                 | The name field.                                                                    |
| `Oauth2Field`                                                                      | [*shared.Oauth2Field1](../../../pkg/models/shared/oauth2field1.md)                 | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `ReadOnly`                                                                         | `*bool`                                                                            | :heavy_minus_sign:                                                                 | When true, this field is displayed to the user but cannot be edited.               |
| `Required`                                                                         | `*bool`                                                                            | :heavy_minus_sign:                                                                 | The required field.                                                                |
| `SharedConfig`                                                                     | [*shared.SharedProviderConfig](../../../pkg/models/shared/sharedproviderconfig.md) | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `StringField`                                                                      | [*shared.FormStringField](../../../pkg/models/shared/formstringfield.md)           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `StringMapField`                                                                   | [*shared.FormStringMapField](../../../pkg/models/shared/formstringmapfield.md)     | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `StringSliceField`                                                                 | [*shared.StringSliceField](../../../pkg/models/shared/stringslicefield.md)         | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `UserConfig`                                                                       | [*shared.UserProviderConfig](../../../pkg/models/shared/userproviderconfig.md)     | :heavy_minus_sign:                                                                 | N/A                                                                                |