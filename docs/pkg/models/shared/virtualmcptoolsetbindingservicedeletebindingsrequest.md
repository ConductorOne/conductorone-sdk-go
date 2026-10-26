# VirtualMCPToolsetBindingServiceDeleteBindingsRequest

VirtualMCPToolsetBindingServiceDeleteBindingsRequest removes Toolsets from one Virtual MCP Server.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `McpAccessProfiles`                                                               | [][shared.MCPAccessProfileRef](../../../pkg/models/shared/mcpaccessprofileref.md) | :heavy_minus_sign:                                                                | Complete Toolset references to unassign. Duplicate tuples are ignored.            |