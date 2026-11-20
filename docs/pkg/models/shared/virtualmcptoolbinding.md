# VirtualMCPToolBinding

VirtualMCPToolBinding represents one individual tool assigned to a Virtual MCP Server.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `CreatedAt`                                                    | [*time.Time](https://pkg.go.dev/time#Time)                     | :heavy_minus_sign:                                             | N/A                                                            |
| `McpTool`                                                      | [*shared.MCPToolRef](../../../pkg/models/shared/mcptoolref.md) | :heavy_minus_sign:                                             | N/A                                                            |
| `UpdatedAt`                                                    | [*time.Time](https://pkg.go.dev/time#Time)                     | :heavy_minus_sign:                                             | N/A                                                            |
| `VirtualMcpServerID`                                           | `*string`                                                      | :heavy_minus_sign:                                             | Immutable identifier of the Virtual MCP Server.                |