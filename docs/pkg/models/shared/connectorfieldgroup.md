# ConnectorFieldGroup

The FieldGroup message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Default`                                                                            | `*bool`                                                                              | :heavy_minus_sign:                                                                   | The default field.                                                                   |
| `DisplayName`                                                                        | `*string`                                                                            | :heavy_minus_sign:                                                                   | Nice name this group (e.g. renders as a Tab label)                                   |
| `FieldNames`                                                                         | []`string`                                                                           | :heavy_minus_sign:                                                                   | Field names are "guaranteed" to be unique, but can be repeated in and between lists. |
| `HelpText`                                                                           | `*string`                                                                            | :heavy_minus_sign:                                                                   | Optional. User-facing help text.                                                     |
| `Name`                                                                               | `*string`                                                                            | :heavy_minus_sign:                                                                   | Unique ID.                                                                           |