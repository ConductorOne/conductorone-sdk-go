# MCPServerServiceResolveMCPIdentifierResponse

MCPServerServiceResolveMcpIdentifierResponse is exactly one of three
 outcomes. registered_server takes precedence over catalog_entry when both
 would match: a tenant that already registered a server for this product
 should never also be told to register it again from the catalog.

This message contains a oneof named match. Only a single field of the following list may be set at a time:
  - registeredServer
  - catalogEntry
  - unmatched



## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `CatalogEntry`                                                                       | [*shared.MCPServerCatalogEntry](../../../pkg/models/shared/mcpservercatalogentry.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `RegisteredServer`                                                                   | [*shared.MCPServerView](../../../pkg/models/shared/mcpserverview.md)                 | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Unmatched`                                                                          | [*shared.Unmatched](../../../pkg/models/shared/unmatched.md)                         | :heavy_minus_sign:                                                                   | N/A                                                                                  |