# Workflow

The Workflow message.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `WorkflowContext`                                                         | [*shared.WorkflowContext](../../../pkg/models/shared/workflowcontext.md)  | :heavy_minus_sign:                                                        | The WorkflowContext message.                                              |
| `Description`                                                             | **string*                                                                 | :heavy_minus_sign:                                                        | The description field.                                                    |
| `DisplayName`                                                             | **string*                                                                 | :heavy_minus_sign:                                                        | The displayName field.                                                    |
| `Enabled`                                                                 | **bool*                                                                   | :heavy_minus_sign:                                                        | The enabled field.                                                        |
| `ID`                                                                      | **string*                                                                 | :heavy_minus_sign:                                                        | The id field.                                                             |
| `Triggers`                                                                | [][shared.WorkflowTrigger](../../../pkg/models/shared/workflowtrigger.md) | :heavy_minus_sign:                                                        | The triggers field.                                                       |
| `WorkflowSteps`                                                           | [][shared.WorkflowStep](../../../pkg/models/shared/workflowstep.md)       | :heavy_minus_sign:                                                        | The workflowSteps field.                                                  |