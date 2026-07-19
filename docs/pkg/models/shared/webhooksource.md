# WebhookSource

The WebhookSource message.

This message contains a oneof named source. Only a single field of the following list may be set at a time:
  - test
  - policyPostAction
  - approvalStep
  - provisionStep
  - workflowStep



## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ApprovalStep`                                                                                       | [*shared.WebhookSourceApprovalStep](../../../pkg/models/shared/webhooksourceapprovalstep.md)         | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `PolicyPostAction`                                                                                   | [*shared.WebhookSourcePolicyPostAction](../../../pkg/models/shared/webhooksourcepolicypostaction.md) | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `ProvisionStep`                                                                                      | [*shared.WebhookSourceProvisionStep](../../../pkg/models/shared/webhooksourceprovisionstep.md)       | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `Test`                                                                                               | [*shared.WebhookSourceTest](../../../pkg/models/shared/webhooksourcetest.md)                         | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `WorkflowStep`                                                                                       | [*shared.WebhookSourceWorkflowStep](../../../pkg/models/shared/webhooksourceworkflowstep.md)         | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |