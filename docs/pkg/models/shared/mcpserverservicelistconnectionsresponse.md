# MCPServerServiceListConnectionsResponse

MCPServerServiceListConnectionsResponse returns a list of passthrough-mode
 MCP servers with per-user connection status.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `List`                                                                        | [][shared.MCPConnectionView](../../../pkg/models/shared/mcpconnectionview.md) | :heavy_minus_sign:                                                            | List of passthrough MCP servers with connection status.                       |
| `NextPageToken`                                                               | `*string`                                                                     | :heavy_minus_sign:                                                            | Token for next page.                                                          |