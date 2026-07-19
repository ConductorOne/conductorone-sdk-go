# FormInstance

The FormInstance message.

This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
  - completed
  - restarted
  - reassigned
  - skipped



## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Completed`                                                                      | [*shared.FormCompletedAction](../../../pkg/models/shared/formcompletedaction.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Data`                                                                           | map[string]`any`                                                                 | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Form`                                                                           | [*shared.RequestSchemaForm](../../../pkg/models/shared/requestschemaform.md)     | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Reassigned`                                                                     | [*shared.ReassignedAction](../../../pkg/models/shared/reassignedaction.md)       | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Restarted`                                                                      | [*shared.RestartAction](../../../pkg/models/shared/restartaction.md)             | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Skipped`                                                                        | [*shared.SkippedAction](../../../pkg/models/shared/skippedaction.md)             | :heavy_minus_sign:                                                               | N/A                                                                              |
| `State`                                                                          | [*shared.FormInstanceState](../../../pkg/models/shared/forminstancestate.md)     | :heavy_minus_sign:                                                               | The state field.                                                                 |