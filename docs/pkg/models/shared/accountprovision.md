# AccountProvision

The AccountProvision message.

This message contains a oneof named storage_type. Only a single field of the following list may be set at a time:
  - saveToVault
  - doNotSave



## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `DoNotSave`                                                      | [*shared.DoNotSave](../../../pkg/models/shared/donotsave.md)     | :heavy_minus_sign:                                               | The DoNotSave message.                                           |
| `SaveToVault`                                                    | [*shared.SaveToVault](../../../pkg/models/shared/savetovault.md) | :heavy_minus_sign:                                               | The SaveToVault message.                                         |
| `Config`                                                         | map[string]*any*                                                 | :heavy_minus_sign:                                               | N/A                                                              |
| `ConnectorID`                                                    | **string*                                                        | :heavy_minus_sign:                                               | The connectorId field.                                           |
| `SchemaID`                                                       | **string*                                                        | :heavy_minus_sign:                                               | The schemaId field.                                              |