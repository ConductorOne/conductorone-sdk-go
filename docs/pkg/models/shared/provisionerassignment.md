# ProvisionerAssignment

ProvisionerAssignment defines how a provisioner is dynamically assigned.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - users
  - appOwners
  - group
  - manager
  - expression
  - entitlementOwners



## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `AppOwners`                                                                                      | [*shared.AppOwnerProvisioner](../../../pkg/models/shared/appownerprovisioner.md)                 | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `EntitlementOwners`                                                                              | [*shared.EntitlementOwnerProvisioner](../../../pkg/models/shared/entitlementownerprovisioner.md) | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `Expression`                                                                                     | [*shared.ExpressionProvisioner](../../../pkg/models/shared/expressionprovisioner.md)             | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `Group`                                                                                          | [*shared.GroupProvisioner](../../../pkg/models/shared/groupprovisioner.md)                       | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `Manager`                                                                                        | [*shared.ManagerProvisioner](../../../pkg/models/shared/managerprovisioner.md)                   | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `Users`                                                                                          | [*shared.UserProvisioner](../../../pkg/models/shared/userprovisioner.md)                         | :heavy_minus_sign:                                                                               | N/A                                                                                              |