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
| `AppOwnerProvisioner`                                                                            | [*shared.AppOwnerProvisioner](../../../pkg/models/shared/appownerprovisioner.md)                 | :heavy_minus_sign:                                                                               | AppOwnerProvisioner resolves to app owners.                                                      |
| `EntitlementOwnerProvisioner`                                                                    | [*shared.EntitlementOwnerProvisioner](../../../pkg/models/shared/entitlementownerprovisioner.md) | :heavy_minus_sign:                                                                               | EntitlementOwnerProvisioner resolves to entitlement owners.                                      |
| `ExpressionProvisioner`                                                                          | [*shared.ExpressionProvisioner](../../../pkg/models/shared/expressionprovisioner.md)             | :heavy_minus_sign:                                                                               | ExpressionProvisioner evaluates CEL expressions to determine provisioners.                       |
| `GroupProvisioner`                                                                               | [*shared.GroupProvisioner](../../../pkg/models/shared/groupprovisioner.md)                       | :heavy_minus_sign:                                                                               | GroupProvisioner resolves to members of a specific group.                                        |
| `ManagerProvisioner`                                                                             | [*shared.ManagerProvisioner](../../../pkg/models/shared/managerprovisioner.md)                   | :heavy_minus_sign:                                                                               | ManagerProvisioner resolves to the user's manager.                                               |
| `UserProvisioner`                                                                                | [*shared.UserProvisioner](../../../pkg/models/shared/userprovisioner.md)                         | :heavy_minus_sign:                                                                               | UserProvisioner assigns specific users as provisioners.                                          |