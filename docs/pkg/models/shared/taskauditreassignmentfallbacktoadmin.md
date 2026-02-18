# TaskAuditReassignmentFallbackToAdmin

TaskAuditReassignmentFallbackToAdmin is used when no eligible reviewers are found
 from the policy configuration and the task falls back to system administrators
 without creating a new policy step. This prevents reassignment loops.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AdminUserIds`                                                               | []*string*                                                                   | :heavy_minus_sign:                                                           | The IDs of the system administrator users that the task is being assigned to |
| `AdminUsers`                                                                 | [][shared.User](../../../pkg/models/shared/user.md)                          | :heavy_minus_sign:                                                           | The system administrator users (populated for display)                       |