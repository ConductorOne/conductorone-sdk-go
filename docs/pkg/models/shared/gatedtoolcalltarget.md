# GatedToolCallTarget

The GatedToolCallTarget message.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `AppEntitlementID`                                             | `*string`                                                      | :heavy_minus_sign:                                             | The appEntitlementId field.                                    |
| `AppID`                                                        | `*string`                                                      | :heavy_minus_sign:                                             | The appId field.                                               |
| `CallerKind`                                                   | `*string`                                                      | :heavy_minus_sign:                                             | The callerKind field.                                          |
| `ConnectorID`                                                  | `*string`                                                      | :heavy_minus_sign:                                             | The connectorId field.                                         |
| `GateID`                                                       | `*string`                                                      | :heavy_minus_sign:                                             | The gateId field.                                              |
| `InputSizeBytes`                                               | `*int`                                                         | :heavy_minus_sign:                                             | The inputSizeBytes field.                                      |
| `ToolError`                                                    | `*string`                                                      | :heavy_minus_sign:                                             | The toolError field.                                           |
| `ToolID`                                                       | `*string`                                                      | :heavy_minus_sign:                                             | The toolId field.                                              |
| `ToolInput`                                                    | map[string]`any`                                               | :heavy_minus_sign:                                             | N/A                                                            |
| `ToolKind`                                                     | `*string`                                                      | :heavy_minus_sign:                                             | The toolKind field.                                            |
| `ToolName`                                                     | `*string`                                                      | :heavy_minus_sign:                                             | The toolName field.                                            |
| `ToolOutput`                                                   | [*shared.ToolOutput](../../../pkg/models/shared/tooloutput.md) | :heavy_minus_sign:                                             | N/A                                                            |