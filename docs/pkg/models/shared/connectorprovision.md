# ConnectorProvision

Indicates that a connector should perform the provisioning. This object has no fields.

This message contains a oneof named provision_type. Only a single field of the following list may be set at a time:
  - defaultBehavior
  - account
  - deleteAccount



## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Account`                                                                  | [*shared.AccountProvision](../../../pkg/models/shared/accountprovision.md) | :heavy_minus_sign:                                                         | N/A                                                                        |
| `DefaultBehavior`                                                          | [*shared.DefaultBehavior](../../../pkg/models/shared/defaultbehavior.md)   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `DeleteAccount`                                                            | [*shared.DeleteAccount](../../../pkg/models/shared/deleteaccount.md)       | :heavy_minus_sign:                                                         | N/A                                                                        |