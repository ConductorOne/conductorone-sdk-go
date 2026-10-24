# VirtualMCPToolBindingServiceCreateBindingsRequest

VirtualMCPToolBindingServiceCreateBindingsRequest assigns individual tools to one Virtual MCP Server.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `McpTools`                                                            | [][shared.MCPToolRef](../../../pkg/models/shared/mcptoolref.md)       | :heavy_minus_sign:                                                    | Complete MCP tool references to assign. Duplicate tuples are ignored. |