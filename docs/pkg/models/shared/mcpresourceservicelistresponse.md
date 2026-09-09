# MCPResourceServiceListResponse

MCPResourceServiceListResponse returns a list of MCP resources.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `NextPageToken`                                                   | `*string`                                                         | :heavy_minus_sign:                                                | Token for next page.                                              |
| `Resources`                                                       | [][shared.MCPResource](../../../pkg/models/shared/mcpresource.md) | :heavy_minus_sign:                                                | List of MCP resources.                                            |