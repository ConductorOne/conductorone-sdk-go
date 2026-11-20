# MCPToolEnablementRules

MCPToolEnablementRules sets the enablement for each tool classification.
 Tools without a read, write, or destructive classification always require enablement.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `DestructiveTools`                                                         | [*shared.DestructiveTools](../../../pkg/models/shared/destructivetools.md) | :heavy_minus_sign:                                                         | Enablement for tools that delete data or cannot be undone.                 |
| `ReadTools`                                                                | [*shared.ReadTools](../../../pkg/models/shared/readtools.md)               | :heavy_minus_sign:                                                         | Enablement for tools that only query data.                                 |
| `WriteTools`                                                               | [*shared.WriteTools](../../../pkg/models/shared/writetools.md)             | :heavy_minus_sign:                                                         | Enablement for tools that create or modify records.                        |