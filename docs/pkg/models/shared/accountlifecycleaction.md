# AccountLifecycleAction

The AccountLifecycleAction message.

This message contains a oneof named account_identifier. Only a single field of the following list may be set at a time:
  - accountRef
  - accountInContext



## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `AccountInContext`                                                         | [*shared.AccountInContext](../../../pkg/models/shared/accountincontext.md) | :heavy_minus_sign:                                                         | N/A                                                                        |
| `AccountRef`                                                               | [*shared.AccountRef](../../../pkg/models/shared/accountref.md)             | :heavy_minus_sign:                                                         | N/A                                                                        |
| `ActionName`                                                               | `*string`                                                                  | :heavy_minus_sign:                                                         | The actionName field.                                                      |
| `ConnectorRef`                                                             | [*shared.ConnectorRef](../../../pkg/models/shared/connectorref.md)         | :heavy_minus_sign:                                                         | N/A                                                                        |