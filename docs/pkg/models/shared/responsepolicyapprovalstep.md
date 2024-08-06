# ResponsePolicyApprovalStep

The ResponsePolicyApprovalStep message.

This message contains a oneof named action. Only a single field of the following list may be set at a time:
  - approve
  - deny
  - reassign
  - replacePolicy



## Fields

| Field                                                                                                            | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ResponsePolicyApprovalReplacePolicy`                                                                            | [*shared.ResponsePolicyApprovalReplacePolicy](../../../pkg/models/shared/responsepolicyapprovalreplacepolicy.md) | :heavy_minus_sign:                                                                                               | The ResponsePolicyApprovalReplacePolicy message.                                                                 |
| `ResponsePolicyApprovalStepApprove`                                                                              | [*shared.ResponsePolicyApprovalStepApprove](../../../pkg/models/shared/responsepolicyapprovalstepapprove.md)     | :heavy_minus_sign:                                                                                               | The ResponsePolicyApprovalStepApprove message.                                                                   |
| `ResponsePolicyApprovalStepDeny`                                                                                 | [*shared.ResponsePolicyApprovalStepDeny](../../../pkg/models/shared/responsepolicyapprovalstepdeny.md)           | :heavy_minus_sign:                                                                                               | The ResponsePolicyApprovalStepDeny message.                                                                      |
| `ResponsePolicyApprovalStepReassign`                                                                             | [*shared.ResponsePolicyApprovalStepReassign](../../../pkg/models/shared/responsepolicyapprovalstepreassign.md)   | :heavy_minus_sign:                                                                                               | The ResponsePolicyApprovalStepReassign message.                                                                  |
| `Version`                                                                                                        | **string*                                                                                                        | :heavy_minus_sign:                                                                                               | version contains the constant value "v1". Future versions of the Webhook Response<br/> will use a different string. |