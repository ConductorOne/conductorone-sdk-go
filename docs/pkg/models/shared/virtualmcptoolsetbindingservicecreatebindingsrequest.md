# VirtualMCPToolsetBindingServiceCreateBindingsRequest

VirtualMCPToolsetBindingServiceCreateBindingsRequest assigns Toolsets to one Virtual MCP Server.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `McpAccessProfiles`                                                               | [][shared.MCPAccessProfileRef](../../../pkg/models/shared/mcpaccessprofileref.md) | :heavy_minus_sign:                                                                | Complete Toolset references to assign. Duplicate tuples are ignored.              |