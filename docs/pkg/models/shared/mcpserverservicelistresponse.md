# MCPServerServiceListResponse

MCPServerServiceListResponse returns a paginated list of MCP servers.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `List`                                                                | [][shared.MCPServerView](../../../pkg/models/shared/mcpserverview.md) | :heavy_minus_sign:                                                    | List of MCP servers.                                                  |
| `NextPageToken`                                                       | `*string`                                                             | :heavy_minus_sign:                                                    | Token for next page.                                                  |