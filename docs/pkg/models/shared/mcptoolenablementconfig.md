# MCPToolEnablementConfig

MCPToolEnablementConfig controls what happens when a new tool is discovered on an MCP server.
 Tools that already exist are not affected.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ClassificationRules`                                                                            | [*shared.MCPToolEnablementRules](../../../pkg/models/shared/mcptoolenablementrules.md)           | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `Mode`                                                                                           | [*shared.MCPToolEnablementConfigMode](../../../pkg/models/shared/mcptoolenablementconfigmode.md) | :heavy_minus_sign:                                                                               | How newly discovered tools are enabled.                                                          |