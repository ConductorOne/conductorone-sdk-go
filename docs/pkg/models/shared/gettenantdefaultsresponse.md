# GetTenantDefaultsResponse

GetTenantDefaultsResponse contains tenant defaults used by
 legacy MCP servers.


## Fields

| Field                                                                                                                                                   | Type                                                                                                                                                    | Required                                                                                                                                                | Description                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `RequireToolApproval`                                                                                                                                   | `*bool`                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                      | Whether newly discovered tools require approval on MCP servers without a<br/> saved registration access level, unless overridden by the per-server setting. |