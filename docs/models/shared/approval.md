# Approval

The Approval message.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - users
  - manager
  - appOwners
  - group
  - self
  - entitlementOwners



## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `AllowReassignment`                                | **bool*                                            | :heavy_minus_sign:                                 | The allowReassignment field.                       |
| `Assigned`                                         | **bool*                                            | :heavy_minus_sign:                                 | The assigned field.                                |
| `RequireApprovalReason`                            | **bool*                                            | :heavy_minus_sign:                                 | The requireApprovalReason field.                   |
| `RequireReassignmentReason`                        | **bool*                                            | :heavy_minus_sign:                                 | The requireReassignmentReason field.               |
| `Typ`                                              | [*ApprovalTyp](../../models/shared/approvaltyp.md) | :heavy_minus_sign:                                 | N/A                                                |