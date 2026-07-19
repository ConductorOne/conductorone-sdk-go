# MCPServerServiceListCatalogResponse

MCPServerServiceListCatalogResponse returns a paginated list of catalog entries.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `List`                                                                                | [][shared.MCPServerCatalogEntry](../../../pkg/models/shared/mcpservercatalogentry.md) | :heavy_minus_sign:                                                                    | List of catalog entries.                                                              |
| `NextPageToken`                                                                       | `*string`                                                                             | :heavy_minus_sign:                                                                    | Token for next page.                                                                  |