# VirtualMCPToolsetBinding

VirtualMCPToolsetBinding represents one Toolset assigned to a Virtual MCP Server.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `CreatedAt`                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                       | :heavy_minus_sign:                                                               | N/A                                                                              |
| `McpAccessProfile`                                                               | [*shared.MCPAccessProfileRef](../../../pkg/models/shared/mcpaccessprofileref.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `UpdatedAt`                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                       | :heavy_minus_sign:                                                               | N/A                                                                              |
| `VirtualMcpServerID`                                                             | `*string`                                                                        | :heavy_minus_sign:                                                               | Immutable identifier of the Virtual MCP Server.                                  |