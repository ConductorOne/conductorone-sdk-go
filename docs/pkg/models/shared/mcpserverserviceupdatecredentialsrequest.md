# MCPServerServiceUpdateCredentialsRequest

MCPServerServiceUpdateCredentialsRequest updates the auth credentials and config fields
 for an existing MCP server. Secrets are sealed before storage.


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ExternalConfig`                                                                         | [*shared.MCPServerExternalConfig](../../../pkg/models/shared/mcpserverexternalconfig.md) | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `HostedConfig`                                                                           | [*shared.MCPServerHostedConfig](../../../pkg/models/shared/mcpserverhostedconfig.md)     | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `UpdateMask`                                                                             | `*string`                                                                                | :heavy_minus_sign:                                                                       | N/A                                                                                      |