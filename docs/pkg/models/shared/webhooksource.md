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
| `WebhookSourceApprovalStep`                                                                          | [*shared.WebhookSourceApprovalStep](../../../pkg/models/shared/webhooksourceapprovalstep.md)         | :heavy_minus_sign:                                                                                   | The WebhookSourceApprovalStep message.                                                               |
| `WebhookSourcePolicyPostAction`                                                                      | [*shared.WebhookSourcePolicyPostAction](../../../pkg/models/shared/webhooksourcepolicypostaction.md) | :heavy_minus_sign:                                                                                   | The WebhookSourcePolicyPostAction message.                                                           |
| `WebhookSourceProvisionStep`                                                                         | [*shared.WebhookSourceProvisionStep](../../../pkg/models/shared/webhooksourceprovisionstep.md)       | :heavy_minus_sign:                                                                                   | The WebhookSourceProvisionStep message.                                                              |
| `WebhookSourceTest`                                                                                  | [*shared.WebhookSourceTest](../../../pkg/models/shared/webhooksourcetest.md)                         | :heavy_minus_sign:                                                                                   | The WebhookSourceTest message.                                                                       |
| `WebhookSourceWorkflowStep`                                                                          | [*shared.WebhookSourceWorkflowStep](../../../pkg/models/shared/webhooksourceworkflowstep.md)         | :heavy_minus_sign:                                                                                   | The WebhookSourceWorkflowStep message.                                                               |