# WorkflowTemplateVersion

The WorkflowTemplateVersion message.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `CreatedAt`                                                               | [*time.Time](https://pkg.go.dev/time#Time)                                | :heavy_minus_sign:                                                        | N/A                                                                       |
| `DeletedAt`                                                               | [*time.Time](https://pkg.go.dev/time#Time)                                | :heavy_minus_sign:                                                        | N/A                                                                       |
| `Triggers`                                                                | [][shared.WorkflowTrigger](../../../pkg/models/shared/workflowtrigger.md) | :heavy_minus_sign:                                                        | The triggers field.                                                       |
| `UpdatedAt`                                                               | [*time.Time](https://pkg.go.dev/time#Time)                                | :heavy_minus_sign:                                                        | N/A                                                                       |
| `Version`                                                                 | **int64*                                                                  | :heavy_minus_sign:                                                        | The version field.                                                        |
| `WorkflowSteps`                                                           | [][shared.WorkflowStep](../../../pkg/models/shared/workflowstep.md)       | :heavy_minus_sign:                                                        | The workflowSteps field.                                                  |
| `WorkflowTemplateID`                                                      | **string*                                                                 | :heavy_minus_sign:                                                        | The workflowTemplateId field.                                             |