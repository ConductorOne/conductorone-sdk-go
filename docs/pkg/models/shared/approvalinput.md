# ApprovalInput

The Approval message.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - users
  - manager
  - appOwners
  - group
  - self
  - entitlementOwners
  - expression
  - webhook
  - resourceOwners



## Fields

| Field                                                                                                                                           | Type                                                                                                                                            | Required                                                                                                                                        | Description                                                                                                                                     |
| ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `AppGroupApproval`                                                                                                                              | [*shared.AppGroupApproval](../../../pkg/models/shared/appgroupapproval.md)                                                                      | :heavy_minus_sign:                                                                                                                              | The AppGroupApproval object provides the configuration for setting a group as the approvers of an approval policy step.                         |
| `AppOwnerApproval`                                                                                                                              | [*shared.AppOwnerApproval](../../../pkg/models/shared/appownerapproval.md)                                                                      | :heavy_minus_sign:                                                                                                                              | App owner approval provides the configuration for an approval step when the app owner is the target.                                            |
| `EntitlementOwnerApproval`                                                                                                                      | [*shared.EntitlementOwnerApproval](../../../pkg/models/shared/entitlementownerapproval.md)                                                      | :heavy_minus_sign:                                                                                                                              | The entitlement owner approval allows configuration of the approval step when the target approvers are the entitlement owners.                  |
| `ExpressionApproval`                                                                                                                            | [*shared.ExpressionApproval](../../../pkg/models/shared/expressionapproval.md)                                                                  | :heavy_minus_sign:                                                                                                                              | The ExpressionApproval message.                                                                                                                 |
| `ManagerApproval`                                                                                                                               | [*shared.ManagerApprovalInput](../../../pkg/models/shared/managerapprovalinput.md)                                                              | :heavy_minus_sign:                                                                                                                              | The manager approval object provides configuration options for approval when the target of the approval is the manager of the user in the task. |
| `ResourceOwnerApproval`                                                                                                                         | [*shared.ResourceOwnerApproval](../../../pkg/models/shared/resourceownerapproval.md)                                                            | :heavy_minus_sign:                                                                                                                              | The resource owner approval allows configuration of the approval step when the target approvers are the resource owners.                        |
| `SelfApproval`                                                                                                                                  | [*shared.SelfApproval](../../../pkg/models/shared/selfapproval.md)                                                                              | :heavy_minus_sign:                                                                                                                              | The self approval object describes the configuration of a policy step that needs to be approved by the target of the request.                   |
| `UserApproval`                                                                                                                                  | [*shared.UserApproval](../../../pkg/models/shared/userapproval.md)                                                                              | :heavy_minus_sign:                                                                                                                              | The user approval object describes the approval configuration of a policy step that needs to be approved by a specific list of users.           |
| `WebhookApproval`                                                                                                                               | [*shared.WebhookApproval](../../../pkg/models/shared/webhookapproval.md)                                                                        | :heavy_minus_sign:                                                                                                                              | The WebhookApproval message.                                                                                                                    |
| `AllowReassignment`                                                                                                                             | **bool*                                                                                                                                         | :heavy_minus_sign:                                                                                                                              | Configuration to allow reassignment by reviewers during this step.                                                                              |
| `Assigned`                                                                                                                                      | **bool*                                                                                                                                         | :heavy_minus_sign:                                                                                                                              | A field indicating whether this step is assigned.                                                                                               |
| `RequireApprovalReason`                                                                                                                         | **bool*                                                                                                                                         | :heavy_minus_sign:                                                                                                                              | Configuration to require a reason when approving this step.                                                                                     |
| `RequireDenialReason`                                                                                                                           | **bool*                                                                                                                                         | :heavy_minus_sign:                                                                                                                              | Configuration to require a reason when denying this step.                                                                                       |
| `RequireReassignmentReason`                                                                                                                     | **bool*                                                                                                                                         | :heavy_minus_sign:                                                                                                                              | Configuration to require a reason when reassigning this step.                                                                                   |