# PolicyStep

A single step in a policy workflow. Exactly one step type is set.

This message contains a oneof named step. Only a single field of the following list may be set at a time:
  - approval
  - provision
  - accept
  - reject
  - wait
  - form
  - action



## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Accept`                                                     | [*shared.Accept](../../../pkg/models/shared/accept.md)       | :heavy_minus_sign:                                           | N/A                                                          |
| `Action`                                                     | [*shared.Action](../../../pkg/models/shared/action.md)       | :heavy_minus_sign:                                           | N/A                                                          |
| `Approval`                                                   | [*shared.Approval](../../../pkg/models/shared/approval.md)   | :heavy_minus_sign:                                           | N/A                                                          |
| `Form`                                                       | [*shared.Form](../../../pkg/models/shared/form.md)           | :heavy_minus_sign:                                           | N/A                                                          |
| `Provision`                                                  | [*shared.Provision](../../../pkg/models/shared/provision.md) | :heavy_minus_sign:                                           | N/A                                                          |
| `Reject`                                                     | [*shared.Reject](../../../pkg/models/shared/reject.md)       | :heavy_minus_sign:                                           | N/A                                                          |
| `Wait`                                                       | [*shared.Wait](../../../pkg/models/shared/wait.md)           | :heavy_minus_sign:                                           | N/A                                                          |