# MCPToolServiceListResponse

MCPToolServiceListResponse returns a list of MCP tools.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `NextPageToken`                                           | `*string`                                                 | :heavy_minus_sign:                                        | Token for next page.                                      |
| `Tools`                                                   | [][shared.MCPTool](../../../pkg/models/shared/mcptool.md) | :heavy_minus_sign:                                        | List of MCP tools.                                        |