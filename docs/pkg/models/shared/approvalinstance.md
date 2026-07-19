# ApprovalInstance

The approval instance object describes the way a policy step should be approved as well as its outcomes and state.

This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
  - approved
  - denied
  - reassigned
  - restarted
  - reassignedByError
  - skipped



## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `Approval`                                                                               | [*shared.Approval](../../../pkg/models/shared/approval.md)                               | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Approved`                                                                               | [*shared.ApprovedAction](../../../pkg/models/shared/approvedaction.md)                   | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `AssignedAt`                                                                             | [*time.Time](https://pkg.go.dev/time#Time)                                               | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Denied`                                                                                 | [*shared.DeniedAction](../../../pkg/models/shared/deniedaction.md)                       | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `EscalationInstance`                                                                     | [*shared.EscalationInstance](../../../pkg/models/shared/escalationinstance.md)           | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Reassigned`                                                                             | [*shared.ReassignedAction](../../../pkg/models/shared/reassignedaction.md)               | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `ReassignedByError`                                                                      | [*shared.ReassignedByErrorAction](../../../pkg/models/shared/reassignedbyerroraction.md) | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Restarted`                                                                              | [*shared.RestartAction](../../../pkg/models/shared/restartaction.md)                     | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Skipped`                                                                                | [*shared.SkippedAction](../../../pkg/models/shared/skippedaction.md)                     | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `State`                                                                                  | [*shared.ApprovalInstanceState](../../../pkg/models/shared/approvalinstancestate.md)     | :heavy_minus_sign:                                                                       | The state of the approval instance                                                       |