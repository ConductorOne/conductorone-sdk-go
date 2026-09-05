# MCPResourceServiceListHistoryResponse

MCPResourceServiceListHistoryResponse returns MCP resource history entries.


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `List`                                                                                    | [][shared.MCPResourceHistoryEntry](../../../pkg/models/shared/mcpresourcehistoryentry.md) | :heavy_minus_sign:                                                                        | The page of history entries, newest first.                                                |
| `NextPageToken`                                                                           | `*string`                                                                                 | :heavy_minus_sign:                                                                        | Pagination token for the next page, or empty if there are no more results.                |