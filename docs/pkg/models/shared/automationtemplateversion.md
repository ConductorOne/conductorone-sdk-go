# AutomationTemplateVersion

The AutomationTemplateVersion message.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AutomationSteps`                                                             | [][shared.AutomationStep](../../../pkg/models/shared/automationstep.md)       | :heavy_minus_sign:                                                            | The automationSteps field.                                                    |
| `AutomationTemplateID`                                                        | **string*                                                                     | :heavy_minus_sign:                                                            | The automationTemplateId field.                                               |
| `CreatedAt`                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                    | :heavy_minus_sign:                                                            | N/A                                                                           |
| `DeletedAt`                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                    | :heavy_minus_sign:                                                            | N/A                                                                           |
| `Triggers`                                                                    | [][shared.AutomationTrigger](../../../pkg/models/shared/automationtrigger.md) | :heavy_minus_sign:                                                            | The triggers field.                                                           |
| `UpdatedAt`                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                    | :heavy_minus_sign:                                                            | N/A                                                                           |
| `Version`                                                                     | **int64*                                                                      | :heavy_minus_sign:                                                            | The version field.                                                            |