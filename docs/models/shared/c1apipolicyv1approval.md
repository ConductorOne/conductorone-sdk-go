# C1APIPolicyV1Approval

The Approval message.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - users
  - manager
  - appOwners
  - group
  - self
  - entitlementOwners



## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `AllowReassignment`                                                                                    | **bool*                                                                                                | :heavy_minus_sign:                                                                                     | The allowReassignment field.                                                                           |
| `AppOwners`                                                                                            | [*C1APIPolicyV1AppOwnerApproval](../../models/shared/c1apipolicyv1appownerapproval.md)                 | :heavy_minus_sign:                                                                                     | The AppOwnerApproval message.                                                                          |
| `Assigned`                                                                                             | **bool*                                                                                                | :heavy_minus_sign:                                                                                     | The assigned field.                                                                                    |
| `EntitlementOwners`                                                                                    | [*C1APIPolicyV1EntitlementOwnerApproval](../../models/shared/c1apipolicyv1entitlementownerapproval.md) | :heavy_minus_sign:                                                                                     | The EntitlementOwnerApproval message.                                                                  |
| `Group`                                                                                                | [*C1APIPolicyV1AppGroupApproval](../../models/shared/c1apipolicyv1appgroupapproval.md)                 | :heavy_minus_sign:                                                                                     | The AppGroupApproval message.                                                                          |
| `Manager`                                                                                              | [*C1APIPolicyV1ManagerApproval](../../models/shared/c1apipolicyv1managerapproval.md)                   | :heavy_minus_sign:                                                                                     | The ManagerApproval message.                                                                           |
| `RequireApprovalReason`                                                                                | **bool*                                                                                                | :heavy_minus_sign:                                                                                     | The requireApprovalReason field.                                                                       |
| `RequireReassignmentReason`                                                                            | **bool*                                                                                                | :heavy_minus_sign:                                                                                     | The requireReassignmentReason field.                                                                   |
| `Self`                                                                                                 | [*C1APIPolicyV1SelfApproval](../../models/shared/c1apipolicyv1selfapproval.md)                         | :heavy_minus_sign:                                                                                     | The SelfApproval message.                                                                              |
| `Users`                                                                                                | [*C1APIPolicyV1UserApproval](../../models/shared/c1apipolicyv1userapproval.md)                         | :heavy_minus_sign:                                                                                     | The UserApproval message.                                                                              |